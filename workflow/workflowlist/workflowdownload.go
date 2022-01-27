package workflowlist

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cavaliergopher/grab/v3"
	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

var (
	workflowDownload workflow.IWorkflow = new(WorkflowDownload)

	wdwf = workflowDownload.GetWF()
)

func init() {
	workflow.Register(workflowDownload)
}

type WorkflowDownload struct {
	workflow.ParentWorkflow
}

func (w *WorkflowDownload) Use() string {
	return "workflow-download"
}

func (w *WorkflowDownload) Run() command.CmdFunc {
	return func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(args[0]) > 0 {
			log.Printf("[DEBUG] start to download [%s]", args[0])
			dir, _ := os.UserHomeDir()
			resp, err := grab.Get(dir+"/Downloads/", args[0])
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
			cmd := exec.Command("open", resp.Filename)
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
			log.Printf("[DEBUG] Download success. Wait Install...")
		}
	}
}
