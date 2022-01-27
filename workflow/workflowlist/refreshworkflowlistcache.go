package workflowlist

import (
	"log"

	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

const refreshCacheUse = "refresh-workflow-list-cache"

var (
	refreshCache workflow.IWorkflow = new(RefreshCacheWorkflows)

	rcwf = refreshCache.GetWF()
)

func init() {
	workflow.Register(refreshCache)
}

type RefreshCacheWorkflows struct {
	workflow.ParentWorkflow
}

func (*RefreshCacheWorkflows) Use() string {
	return refreshCacheUse
}

func (w *RefreshCacheWorkflows) Run() command.CmdFunc {
	return func(cmd *cobra.Command, args []string) {
		rcwf.Run(func() {
			defer func() { rcwf.SendFeedback() }()

			log.Println("init workflow data")
			wl := new(WorkflowList)
			if err := wl.initData(); err != nil {
				log.Fatalf("[ERROR] init workflow data err %s", err)
			}
			wl.cacheAllImage()
		})
	}
}
