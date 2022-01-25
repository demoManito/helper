package crack

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

var (
	video workflow.IWorkflow = new(Video)

	videowf = video.GetWF()
)

func init() {
	workflow.Register(video)
}

type Url struct {
	Name      string
	Png       string
	ParentURL string
	URL       string
}

type Video struct {
	workflow.ParentWorkflow
}

func (*Video) Use() string {
	return "crack-video"
}

func (v *Video) Run() command.CmdFunc {
	return v.run
}

func (v *Video) run(cmd *cobra.Command, args []string) {
	videowf.Run(func() {
		defer func() { videowf.SendFeedback() }()

		for _, url := range v.loadURLs() {
			url.URL = fmt.Sprintf(url.URL, args[0])
			videowf.NewItem(url.Name).Valid(true).Copytext(url.URL).Arg(url.URL).
				Icon(&aw.Icon{Value: url.Png}).Subtitle("[" + url.ParentURL + "]" + "  [CMD+C] copy / " + "[" + url.URL + "]" + "[Enter] direct")
		}
	})
}

// add more crack website
func (v *Video) loadURLs() []*Url {
	return []*Url{
		{
			Name:      "CC视频破解",
			Png:       "video/cc.png",
			ParentURL: "https://thinkibm.vercel.app",
			URL:       "https://thinkibm.vercel.app/?url=%s",
		},
		{
			Name:      "17kyun",
			Png:       "video/17kyun.png",
			ParentURL: "https://17kyun.com/api.php",
			URL:       "https://17kyun.com/api.php?url=%s",
		},
	}
}
