package workflow

import (
	"fmt"
	"helper/handler/command"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var (
	commandMap = make(map[string]*cobra.Command)

	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

type IWorkflow interface {
	GetWF() *aw.Workflow

	Use() string
	Short() string              // 是 'help' 输出中显示的简短说明
	Long() string               // 是 'help <this-command>' 显示的长说明
	Args() cobra.PositionalArgs // 返回一个错误，如果至少没有 n 个参数
	Run() command.CmdFunc
}

func Register(commands ...IWorkflow) {
	for _, command := range commands {
		if _, ok := commandMap[command.Use()]; ok {
			log.Println(fmt.Sprintf("%s: is duplicate", command.Use()))
		}

		commandMap[command.Use()] = &cobra.Command{
			Use:   command.Use(),
			Short: command.Short(),
			Long:  command.Long(),
			Args:  command.Args(),
			Run:   command.Run(),
		}
	}
}

// Execute 执行
func Execute(cmd *cobra.Command) error {
	for _, command := range commandMap {
		cmd.AddCommand(command)
	}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return nil
}

type ParentWorkflow struct{}

func (*ParentWorkflow) Short() string {
	return ""
}

func (*ParentWorkflow) Long() string {
	return ""
}

// Args example: cobra.MinimumNArgs(1)
func (*ParentWorkflow) Args() cobra.PositionalArgs {
	return nil
}

func (*ParentWorkflow) GetWF() *aw.Workflow {
	return wf
}
