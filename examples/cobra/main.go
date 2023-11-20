package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wangzhen94/iam/examples/cobra/cmd"
)

func main() {
	cmd.Execute()

	fmt.Println(viper.GetString("author"))
	fmt.Println(viper.GetString("license"))
	fmt.Println(viper.GetString("projectbase"))

}
