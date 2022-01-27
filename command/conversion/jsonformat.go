package conversion

import (
	"bytes"
	encodingjson "encoding/json"
	"log"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"helper/handler/command"
)

var json command.ICommand = new(JsonFormat)

func init() {
	command.Register(json)
}

type JsonFormat struct {
	command.ParentCommand
}

func (*JsonFormat) Use() string {
	return "json"
}

func (j *JsonFormat) Run() command.CmdFunc {
	return j.run
}

func (j *JsonFormat) run(cmd *cobra.Command, args []string) {
	var jsonFormat bytes.Buffer
	if err := encodingjson.Indent(&jsonFormat, []byte(args[0]), "", "  "); err != nil {
		log.Printf("err json %s %s", err, jsonFormat.String())
		return
	}
	if err := clipboard.WriteAll(jsonFormat.String()); err != nil {
		log.Printf("剪切板意外退出 %s", err)
	}
}
