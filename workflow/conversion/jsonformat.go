package conversion

import (
	"bytes"
	encodingjson "encoding/json"
	"log"

	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

// TODO: 目前有问题，暂时无法正常使用

var (
	json workflow.IWorkflow = new(JsonFormat)

	jsonwf = json.GetWF()
)

//func init() {
//	workflow.Register(json)
//}

type JsonFormat struct {
	workflow.ParentWorkflow
}

func (*JsonFormat) Use() string {
	return "json"
}

func (j *JsonFormat) Run() command.CmdFunc {
	return j.run
}

func (j *JsonFormat) run(cmd *cobra.Command, args []string) {
	jsonwf.Run(func() {
		defer func() { jsonwf.SendFeedback() }()
		var jsonFormat bytes.Buffer
		if err := encodingjson.Indent(&jsonFormat, []byte(args[0]), "", "  "); err != nil {
			jsonwf.NewItem("JSON 格式异常").Valid(true).Copytext(jsonFormat.String())
			log.Printf("err json %s %s", err, jsonFormat.String())
		}
		// "回车自动复制"
		jsonwf.NewItem(jsonFormat.String()).Valid(true).Copytext(jsonFormat.String())
	})
}
