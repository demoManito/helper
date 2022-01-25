package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var commandMap = make(map[string]*cobra.Command)

type CmdFunc func(cmd *cobra.Command, args []string)

type ICommand interface {
	Use() string
	Short() string              // 是 'help' 输出中显示的简短说明
	Long() string               // 是 'help <this-command>' 显示的长说明
	Args() cobra.PositionalArgs // 返回一个错误，如果至少没有 n 个参数
	Run() CmdFunc
}

// Register 注册
func Register(commands ...ICommand) {
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

type ParentCommand struct{}

func (*ParentCommand) Short() string {
	return ""
}

func (*ParentCommand) Long() string {
	return ""
}

// Args example: cobra.MinimumNArgs(1)
func (*ParentCommand) Args() cobra.PositionalArgs {
	return nil
}
