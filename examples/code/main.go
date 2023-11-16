package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"

	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
)

func main() {

	//print("GET", "/user/list", "userList", 3)
	//print("OPTION", "/user/kkk", "userList", 3)
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		if err := getUser(name); err != nil {
			core.WriteResponse(c, err, nil)
			return
		}

		core.WriteResponse(c, nil, map[string]string{"email": name + "@foxmail.com"})
	})

	//r.Run(":7070")

	ea := Entity{
		name: "zhang",
		attr: map[string]interface{}{
			"li": "abk",
		},
	}

	eb := ea.clone()

	if &ea == eb {
		fmt.Println("point equal")
	} else {
		fmt.Println("point not equal")
	}

}

func getUser(name string) error {
	if err := queryDataBase(name); err != nil {
		return errors.Wrap(err, "get user error")
	}
	return nil
}

func queryDataBase(name string) error {
	return errors.WithCode(code.ErrDatabase, "user '%s' not found")
}

func print(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
}

type Entity struct {
	name string
	attr map[string]interface{}
}

func (e *Entity) clone() *Entity {
	copy := *e

	return &copy
}
