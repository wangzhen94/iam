package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
	"os"
	"path/filepath"
)

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
	deletePKFiles()
	//d := dog{"11"}
	//d.change()
	//fmt.Println(d.name)
	//
	//c := &cat{"22"}
	//c.change()
	//fmt.Println(c.name)

	//
	//d := Data{Name: "John"}
	//
	//p := &d
	//
	//// 传递指针
	//processData(p.Name)
	//
	//// 传递值
	//processData(p.Name)

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

func deletePKFiles() {
	// 指定目录
	targetDir := "/Users/wangzhen/go/src/github.com/wangzhen94/iam"

	// 遍历目录
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// 检查是否是目录
		if info.IsDir() {
			// 检查目录下是否包含.pk文件
			pkFileExists, otherFilesExist := checkDirectoryContents(path)

			// 如果包含.pk文件和其他文件，则删除.pk文件
			if pkFileExists && otherFilesExist {
				err := deletePkFile(path)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

// 检查目录是否包含.pk文件以及其他文件
func checkDirectoryContents(dirPath string) (pkFileExists bool, otherFilesExist bool) {
	pkFileExists = false
	otherFilesExist = false

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			// 如果存在子目录，则递归检查
			subdirPath := filepath.Join(dirPath, file.Name())
			subPkFileExists, subOtherFilesExist := checkDirectoryContents(subdirPath)
			pkFileExists = pkFileExists || subPkFileExists
			otherFilesExist = otherFilesExist || subOtherFilesExist
		} else {
			// 检查是否是.pk文件
			if filepath.Ext(file.Name()) == ".pk" {
				pkFileExists = true
			} else {
				otherFilesExist = true
			}
		}
	}

	return pkFileExists, otherFilesExist
}

// 删除目录下的.pk文件
func deletePkFile(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".pk" {
			filePath := filepath.Join(dirPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
			fmt.Printf("Deleted: %s\n", filePath)
		}
	}

	return nil
}
