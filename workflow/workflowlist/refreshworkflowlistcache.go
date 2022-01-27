package workflowlist

import (
	"log"

	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

var (
	refreshCache workflow.IWorkflow = new(RefreshCacheWorkflows)

	rcwf = refreshCache.GetWF()
)

func init() {
	workflow.Register(workflowList)
}

type RefreshCacheWorkflows struct {
	workflow.ParentWorkflow
}

func (*RefreshCacheWorkflows) Use() string {
	return "refresh-workflow-list-cache"
}

func (w *RefreshCacheWorkflows) Run() command.CmdFunc {
	return func(cmd *cobra.Command, args []string) {
		wl := new(WorkflowList)
		if err := wl.initData(); err != nil {
			log.Fatalf("【]")
		}
		wl.cacheAllImage()
	}
}
