package main

import (
	"helper/handler/workflow"
	"log"

	"github.com/spf13/cobra"

	"helper/globa"
	_ "helper/workflow/conversion" // 转换相关
	_ "helper/workflow/crack"      // 破解相关
)

var rootCmd = &cobra.Command{
	Use:     "helper",
	Version: globa.Release,
}

func main() {
	if err := workflow.Execute(rootCmd); err != nil {
		log.Fatal("init ", err)
	}
}
