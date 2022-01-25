package crack

import (
	"fmt"
	"helper/handler/command"
	"log"
	"os/exec"
	"runtime"
	"sync"

	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	"helper/globa"
)

var (
	wg = &sync.WaitGroup{}

	video command.ICommand = new(Video)
)

func init() {
	command.Register(video)
}

type Video struct {
	command.ParentCommand
}

func (*Video) Use() string {
	return "crack-video"
}

// Run 支持同时打开多个视频窗口
// helper crack-video https://v.qq.com/x/cover/m441e3rjq9kwpsc/k00417vlpjq.html
func (v *Video) Run() command.CmdFunc {
	return v.run
}

func (v *Video) run(cmd *cobra.Command, args []string) {
	sysCmd, ok := globa.Systems[runtime.GOOS]
	if !ok {
		return
	}

	for _, arg := range funk.UniqString(args) {
		wg.Add(1)
		go func() {
			openCmd := exec.Command(sysCmd, fmt.Sprintf("https://thinkibm.vercel.app/?url=%s", arg))
			err := openCmd.Start()
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
