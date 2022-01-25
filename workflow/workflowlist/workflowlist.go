package workflowlist

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

const (
	cacheKey      = "helper-workflows" // alfred cache 缓存 Key
	localDataFile = "workflows.json"
)

var (
	workflowList workflow.IWorkflow = new(WorkflowList)

	wf = workflowList.GetWF()
)

func init() {
	workflow.Register(workflowList)
}

type Workflow struct {
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	Author  string `json:"author"`
	Version string `json:"version"`
	Website string `json:"website"`
	Query   string `json:"query"`
}

func (w *Workflow) makeQuery() {
	w.Query = strings.ToLower(w.Name + " " + w.Author + " " + w.Desc)
}

type WorkflowList struct {
	workflows []*Workflow

	workflow.ParentWorkflow
}

// Use wl gitlab
func (*WorkflowList) Use() string {
	return "workflow-list"
}

func (w *WorkflowList) Run() command.CmdFunc {
	return w.run
}

func (w *WorkflowList) run(cmd *cobra.Command, args []string) {
	err := wf.Cache.LoadJSON(cacheKey, &w.workflows)
	if err != nil {
		w.initData()
	}
	for _, item := range w.workflows {
		item.makeQuery()
	}
	duration, _ := wf.Cache.Age(cacheKey)
	if duration > 10*time.Minute {
		// TODO
	}
}

func (w *WorkflowList) initData() error {
	data, err := ioutil.ReadFile(localDataFile)
	if err != nil {
		log.Printf("[ERROR] load local backup data failed [%s]", err.Error())
	} else {
		err = json.Unmarshal(data, &w.workflows)
		if err != nil {
			log.Printf("[ERROR] load local backup data failed [%s]", err.Error())
		}
		for _, item := range w.workflows {
			item.makeQuery()
		}

		if err = wf.Cache.Store(cacheKey, data); err != nil {
			return err
		}
	}
	return nil
}
