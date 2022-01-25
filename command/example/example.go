package example

import (
	"fmt"
	"helper/handler/command"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	command.Register(new(Example))
}

type Example struct {
	command.ParentCommand
}

func (*Example) Use() string {
	return "example"
}

func (*Example) Short() string {
	return "example the directive was clear"
}

func (*Example) Long() string {
	return "example the directive was clear"
}

func (*Example) Args() cobra.PositionalArgs {
	return nil
}

func (*Example) Run() command.CmdFunc {
	return run
}

func run(cmd *cobra.Command, args []string) {
	fd := exec.Command("echo", "success")
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	if err := fd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err.Error())
		return
	}
}
