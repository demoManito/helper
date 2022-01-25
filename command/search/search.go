package search

import (
	"fmt"
	"helper/handler/command"
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"

	"helper/globa"
)

const (
	UseBaidu  = "baidu"
	UseGoogle = "google"
)

var (
	baidu  command.ICommand = new(Baidu)
	google command.ICommand = new(Google)
)

func init() {
	command.Register(baidu, google)
}

type Baidu struct {
	command.ParentCommand
}

func (*Baidu) Use() string {
	return UseBaidu
}

func (b *Baidu) Run() command.CmdFunc {
	return searchRun
}

type Google struct {
	command.ParentCommand
}

func (*Google) Use() string {
	return UseGoogle
}

func (b *Google) Run() command.CmdFunc {
	return searchRun
}

func searchRun(cmd *cobra.Command, args []string) {
	switch cmd.Use {
	case UseBaidu:
		search(fmt.Sprintf("https://www.baidu.com/s?wd=%s", args[0]))
	case UseGoogle:
		search(fmt.Sprintf("https://www.google.com.hk/search?q=%s", args[0]))
	default:
		return
	}
}

func search(url string) {
	sysCmd, ok := globa.Systems[runtime.GOOS]
	if !ok {
		return
	}
	openCmd := exec.Command(sysCmd, url)
	err := openCmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
