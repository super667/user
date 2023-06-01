package ldap

import (
	"context"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

type Dept struct {
	DN       string `json:"dn"`
	Id       string `json:"id"`       // 部门ID
	Name     string `json:"name"`     // 部门名称拼音
	Remark   string `json:"remark"`   // 部门中文名
	ParentId string `json:"parentid"` // 父部门ID
}

type User struct {
	Name          string `json:"name"`
	DN            string `json:"dn"`
	CN            string `json:"cn"`
	SN            string `json:"sn"`
	Mobile        string `json:"mobile"`
	DeptCode      string `json:"deptCode"`      // 部门编号，此处可以存放员工的职位
	Description   string `json:"description"`   // 描述
	DisplayName   string `json:"displayName"`   // 展示名字，可以是中文名字
	Mail          string `json:"mail"`          // 邮箱
	Number        string `json:"Number"`        // 员工工号
	PostalAddress string `json:"postalAddress"` // 家庭住址
	DeptName      string `json:"deptName"`
	Position      string `json:"position"`
	Password      string `json:"password"`
}

func NewLdapPool(conf Config) *Pool {
	return &Pool{
		conf: conf,
	}
}

type Config struct {
	Url              string `json:"Url,default=ldap://121.5.70.40:389"`
	MaxConn          int64  `json:"MaxConn,default=10"`
	BaseDn           string `json:"BaseDn"`
	AdminDN          string `json:"AdminDn"`
	AdminPass        string `json:"AdminPass"`
	UserDn           string `json:"UserDn"`
	UserInitPassword string `json:"UserInitPassword"`
	GroupNameModify  bool   `json:"GroupNameModify"`
	UserNameModify   bool   `json:"UserNameModify"`
}

type Pool struct {
	conf     Config
	mu       sync.Mutex
	conns    []*ldap.Conn
	reqConns map[uint64]chan *ldap.Conn
	openConn int
	maxOpen  int
}

func (p *Pool) initLDAPConn(ctx context.Context) (*ldap.Conn, error) {
	// Dail有两个参数 network,  address, 返回 (*Conn,  error)
	ldapConn, err := ldap.DialURL(p.conf.Url, ldap.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}))
	if err != nil {
		logx.Error(fmt.Sprintf("初始化ldap连接异常: %s", err.Error()))
		panic(fmt.Errorf("初始化ldap连接异常: %v", err))
	}
	err = ldapConn.Bind(p.conf.AdminDN, p.conf.AdminPass)
	if err != nil {
		logx.Error(fmt.Sprintf("绑定admin账号异常: %s", err.Error()))
		panic(fmt.Errorf("绑定admin账号异常: %v", err))
	}
	logx.WithContext(ctx).Info(fmt.Sprintf("创建连接%s:******@tcp(%s)", p.conf.AdminDN, p.conf.Url))
	return ldapConn, nil
}

func (p *Pool) getConnection(ctx context.Context) (*ldap.Conn, error) {
	p.mu.Lock()
	// 判断当前连接池内是否存在连接
	connNum := len(p.conns)
	if connNum > 0 {
		p.openConn++
		conn := p.conns[0]
		copy(p.conns, p.conns[1:])
		p.conns = p.conns[:connNum-1]

		p.mu.Unlock()
		// 发现连接已经 close 重新获取连接
		if conn.IsClosing() {
			return p.initLDAPConn(ctx)
		}
		return conn, nil
	}

	// 当现有连接池为空时，并且当前超过最大连接限制
	if p.maxOpen != 0 && p.openConn > p.maxOpen {
		// 创建一个等待队列
		req := make(chan *ldap.Conn, 1)
		reqKey := p.nextRequestKeyLocked()
		p.reqConns[reqKey] = req
		p.mu.Unlock()

		// 等待请求归还
		return <-req, nil
	} else {
		p.openConn++
		p.mu.Unlock()
		return p.initLDAPConn(ctx)
	}
}

func (p *Pool) putConnection(ctx context.Context, conn *ldap.Conn) {
	logx.WithContext(ctx).Info("返回了一个 LDAP 连接")
	p.mu.Lock()
	defer p.mu.Unlock()

	// 先判断是否存在等待的队列
	if num := len(p.reqConns); num > 0 {
		var req chan *ldap.Conn
		var reqKey uint64
		for reqKey, req = range p.reqConns {
			break
		}
		delete(p.reqConns, reqKey)
		req <- conn
		return
	} else {
		p.openConn--
		if !conn.IsClosing() {
			p.conns = append(p.conns, conn)
		}
	}
}

func (p *Pool) nextRequestKeyLocked() uint64 {
	for {
		reqKey := rand.Uint64()
		if _, ok := p.reqConns[reqKey]; !ok {
			return reqKey
		}
	}
}

func (p *Pool) GetAllDepts(ctx context.Context) (ret []*Dept, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		p.conf.BaseDn,                                               // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=*))", // This is Filter for LDAP query
		[]string{},           // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	// 获取 LDAP 连接
	conn, err := p.getConnection(ctx)
	defer p.putConnection(ctx, conn)
	if err != nil {
		return nil, err
	}

	// Search through ldap built-in search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return ret, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			if v.DN == p.conf.BaseDn || v.DN == p.conf.AdminDN || strings.Contains(v.DN, p.conf.UserDn) {
				continue
			}
			var ele Dept
			ele.DN = v.DN
			ele.Name = strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			ele.Id = strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			ele.Remark = v.GetAttributeValue("description")
			if len(strings.Split(v.DN, ","))-len(strings.Split(p.conf.BaseDn, ",")) == 1 {
				ele.ParentId = "0"
			} else {
				ele.ParentId = strings.Split(strings.Split(v.DN, ",")[1], "=")[1]
			}
			ret = append(ret, &ele)
		}
	}
	return
}

func (p *Pool) GetAllUsers(ctx context.Context) (ret []*User, err error) {
	searchReq := ldap.NewSearchRequest(
		p.conf.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false, "(&(objectClass=*))", []string{},
		nil,
	)
	conn, err := p.getConnection(ctx)
	defer p.putConnection(ctx, conn)
	if err != nil {
		return nil, err
	}
	sr, err := conn.Search(searchReq)
	if err != nil {
		return ret, err
	}
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			if v.DN == p.conf.UserDn || !strings.Contains(v.DN, p.conf.UserDn) {
				continue
			}
			name := strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			_, err := p.getUserDeptIds(ctx, v.DN)
			if err != nil {
				return ret, err
			}
			ret = append(ret, &User{
				Name:          name,
				DN:            v.DN,
				CN:            v.GetAttributeValue("cn"),
				SN:            v.GetAttributeValue("sn"),
				Mobile:        v.GetAttributeValue("mobile"),
				DeptCode:      v.GetAttributeValue("departmentNumber"),
				Description:   v.GetAttributeValue("description"),
				DisplayName:   v.GetAttributeValue("displayName"),
				Mail:          v.GetAttributeValue("mail"),
				Number:        v.GetAttributeValue("employeeNumber"),
				PostalAddress: v.GetAttributeValue("postalAddress"),
				//DeptName:      deptIds[0],
			})
		}
	}
	return
}

// 获取用户所在部门

func (p *Pool) getUserDeptIds(ctx context.Context, udn string) (ret []string, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		p.conf.BaseDn,                                               // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		fmt.Sprintf("(|(Member=%s)(uniqueMember=%s))", udn, udn), // This is Filter for LDAP query
		[]string{}, // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	// 获取 LDAP 连接
	conn, err := p.getConnection(ctx)
	defer p.putConnection(ctx, conn)
	if err != nil {
		return nil, err
	}

	// Search through ldap built-in search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return ret, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			ret = append(ret, strings.Split(strings.Split(v.DN, ",")[0], "=")[1])
		}
	}
	return ret, nil
}
