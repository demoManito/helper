package main

import (
	"log"

	"github.com/spf13/cobra"

	"helper/globa"
	"helper/handler/workflow"

	_ "helper/workflow/conversion"   // 转换相关
	_ "helper/workflow/crack"        // 破解相关
	_ "helper/workflow/workflowlist" // workflow list
)

var rootCmd = &cobra.Command{
	Use:     "helper",
	Version: globa.Release,
}

func main() {
	if err := workflow.Execute(rootCmd); err != nil {
		log.Fatal("workflow init ", err)
	}
}
