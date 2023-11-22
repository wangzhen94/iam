package main

import (
	"fmt"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
	"os"
	"path/filepath"
	"strings"
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
			if strings.Contains(err.Error(), ".pk") {
				return nil
			}
			//if _, ok := err.(fs.PathError); ok {
			//
			//}
			fmt.Println(err)
			return nil
		}

		// 检查是否是目录
		if info.IsDir() {
			// 检查目录下是否包含.pk文件和其他文件
			canDel := checkDirectoryContents(path)

			// 如果同时包含.pk文件和其他文件，则删除.pk文件
			if canDel {
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

// 检查目录是否同时包含.pk文件和其他文件
func checkDirectoryContents(dirPath string) (canDel bool) {
	var pkFileExists = false
	var otherFileExist = false
	var otherDirExit = false

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		subdirPath := filepath.Join(dirPath, file.Name())
		subFile, err := os.Stat(subdirPath)
		if err != nil {
			return false
		}
		if subFile.IsDir() {
			otherDirExit = true
		} else {
			if filepath.Ext(subdirPath) == ".pk" {
				pkFileExists = true
			} else {
				otherFileExist = true
			}
		}
		if pkFileExists && (otherFileExist || otherDirExit) {
			return true
		}
	}

	return false
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
