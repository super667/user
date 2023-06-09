package ldap

import (
	"context"
	"fmt"
	"testing"
)

var pool = NewLdapPool(Config{
	Url:       "ldap://121.5.70.40:388",
	BaseDn:    "dc=superboys,dc=top",
	AdminDN:   "cn=admin,dc=superboys,dc=top",
	UserDn:    "ou=people,dc=superboys,dc=top",
	AdminPass: "123456",
})

func TestPool_GetAllUsers(t *testing.T) {
	users, err := pool.GetAllUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, u := range users {
		fmt.Println(u)
	}
}

func TestPool_GetAllDepts(t *testing.T) {
	depts, err := pool.GetAllDepts(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range depts {
		fmt.Println(d)
	}
}
