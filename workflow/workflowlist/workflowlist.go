package workflowlist

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

const (
	cacheKey      = "helper:workflows" // alfred cache 缓存 Key
	localDataFile = "workflows.json"

	searchLimit = 20
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

func (*WorkflowList) Use() string {
	return "workflow-list"
}

func (w *WorkflowList) Run() command.CmdFunc {
	return w.run
}

func (w *WorkflowList) run(cmd *cobra.Command, args []string) {
	wf.Run(func() {
		defer func() { wf.SendFeedback() }()
		if len(args) > 0 && len(args[0]) > 0 {
			urlPlain := args[0]
			validUrl, err := url.Parse(urlPlain)
			if err == nil && len(validUrl.Scheme) > 0 && len(validUrl.Host) > 0 && len(validUrl.Path) > 0 {
				wf.NewItem("Download & Install ...").Valid(true).Arg(urlPlain).Icon(&aw.Icon{Value: "lock.png"}).Var("cmd", "download")
				return
			}
		}
		w.search(args)
	})
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

func (w *WorkflowList) cacheAllImage() {
	wf.Add(len(w.workflows))
	for _, workflow := range w.workflows {
		go func(icon string) {
			file := fmt.Sprintf("%x", md5.Sum([]byte(icon)))
			localPath := wf.Cache.Dir + "/icons/" + file
			if _, err := os.Stat(localPath); os.IsNotExist(err) {
				_, err = grab.Get(localPath, icon)
				if err != nil {
					log.Printf("[ERROR] cache image [%s] to [%s] failed[%s]\n", icon, localPath, err.Error())
				} else {
					log.Printf("[DEBUG] cache image [%s] to [%s] success\n", icon, localPath)
				}
			}
			wf.Done()
		}(workflow.Icon)
	}
	wf.Wait()
}

func (w *WorkflowList) loadData() error {
	err := wf.Cache.LoadJSON(cacheKey, &w.workflows)
	if err != nil {
		w.initData()
		w.cacheAllImage()
	}
	for _, workflow := range w.workflows {
		workflow.makeQuery()
	}
	duration, _ := wf.Cache.Age(cacheKey)
	if duration > 10*time.Minute {
		go func() {
			w.initData()
			w.cacheAllImage()
		}()
	}
	return nil
}

func (w *WorkflowList) search(keywords []string) {
	err := w.loadData()
	if err != nil {
		wf.NewWarningItem(fmt.Sprintf("Error: %s", err.Error()), "")
		return
	}
	count := 0
	lowers := make([]string, len(keywords))
	for i, keyword := range keywords {
		lowers[i] = strings.ToLower(keyword)
	}
	for _, workflow := range w.workflows {
		notFound := false
		for _, keyword := range lowers {
			if !strings.Contains(workflow.Query, keyword) {
				notFound = true
				break
			}
		}
		if !notFound {
			text := workflow.Url
			if len(text) == 0 {
				text = workflow.Website
			}
			wf.NewItem(workflow.Name+" @"+workflow.Author).Subtitle(workflow.Desc).Icon(&aw.Icon{
				Value: fmt.Sprintf(wf.Cache.Dir+"/icons/%x", md5.Sum([]byte(workflow.Icon))),
			}).Valid(true).
				Var("title", fmt.Sprintf("Downloading [%s]", workflow.Name)).Arg(workflow.Url).
				Var("cmd", "download").
				Var("website", workflow.Website).Copytext(text)
			count += 1
			if count > searchLimit {
				return
			}
		}
	}
}
