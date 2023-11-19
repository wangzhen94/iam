package main

import (
	"fmt"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var dir = "/Users/wangzhen/go/src/github.com/wangzhen94/iam"

type Data struct {
	Name string
}

func processData(v interface{}) {
	switch val := v.(type) {
	case *string:
		fmt.Printf("Received pointer: %s\n", val)
	case string:
		fmt.Printf("Received value: %s\n", val)
	default:
		fmt.Println("Unknown type")
	}
}

type dog struct {
	name string
}

func (d *dog) change() {
	d.name = "li"
}

type cat struct {
	name string
}

func (c cat) change() {
	c.name = "zhao"
}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case value := <-ch:
				fmt.Println("一：Received:", value)
				//time.Sleep(time.Second)
				ch <- 1
				fmt.Println("一：send: 1")
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case ch <- 2:
				fmt.Println("二：Sent 2")
				//time.Sleep(time.Second)
				value := <-ch
				fmt.Println("二：Received: ", value)
			}
		}
	}()

	wg.Wait()
	//d := dog{"11"}
	//d.change()
	//fmt.Println(d.name)
	//
	//c := &cat{"22"}
	//c.change()
	//fmt.Println(c.name)
	//
	////
	////d := Data{Name: "John"}
	////
	////p := &d
	////
	////// 传递指针
	////processData(p.Name)
	////
	////// 传递值
	////processData(p.Name)
	//
	////print("GET", "/user/list", "userList", 3)
	////print("OPTION", "/user/kkk", "userList", 3)
	//r := gin.Default()
	//
	//r.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Params.ByName("name")
	//	if err := getUser(name); err != nil {
	//		core.WriteResponse(c, err, nil)
	//		return
	//	}
	//
	//	core.WriteResponse(c, nil, map[string]string{"email": name + "@foxmail.com"})
	//})
	//
	////r.Run(":7070")
	//
	//ea := Entity{
	//	name: "zhang",
	//	attr: map[string]interface{}{
	//		"li": "abk",
	//	},
	//}
	//
	//eb := ea.clone()
	//
	//if &ea == eb {
	//	fmt.Println("point equal")
	//} else {
	//	fmt.Println("point not equal")
	//}

}

func deadLock() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			value := <-ch
			fmt.Println("Received:", value)
			mu.Unlock()
			//time.Sleep(time.Second)
			mu.Lock()
			ch <- 1
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			ch <- 2
			fmt.Println("Sent 2")
			//time.Sleep(time.Second)
			value := <-ch
			fmt.Println("Received after sending 2:", value)
			mu.Unlock()
		}
	}()

	wg.Wait()
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

func deletePKFiles(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否为目录
		if info.IsDir() {
			return nil
		}

		// 判断文件名中是否包含.pk
		if strings.Contains(info.Name(), ".pk") {
			fmt.Printf("Deleting file: %s\n", path)
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
