package main

import (
	"helper/handler/command"
	"log"

	"github.com/spf13/cobra"

	_ "helper/command/conversion" // 转换相关
	_ "helper/command/crack"      // 破解相关
	_ "helper/command/example"    // 测试
	_ "helper/command/search"     // 搜索相关
	"helper/globa"
)

var rootCmd = &cobra.Command{
	Use:     "helper",
	Version: globa.Release,
}

func main() {
	if err := command.Execute(rootCmd); err != nil {
		log.Fatal("init ", err)
	}
}
