package casbinx

import (
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/jmoiron/sqlx"
	"sync"
)

func New(db *sqlx.DB) *Casbin {
	c := &Casbin{
		db: db,
	}
	c.Init()
	return c
}

type CasbinConf struct {
	Table string `json:"table"`
}

type Casbin struct {
	conf CasbinConf
	db   *sqlx.DB
}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (c *Casbin) Init() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := sqlxadapter.NewAdapter(c.db, "casbin_rule")
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
