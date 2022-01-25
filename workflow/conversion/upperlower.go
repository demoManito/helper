package conversion

import (
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"

	"helper/handler/command"
	"helper/handler/workflow"
)

const (
	modelToUpper = "tu"
	modelToLower = "tl"
)

// specialSymbol 大小写转换间隔特殊符号
var specialSymbol = []string{
	"-",
	"_",
}

var (
	upperToLower workflow.IWorkflow = new(UpperToLower)
	lowerToUpper workflow.IWorkflow = new(LowerToUpper)

	upperlowerwf = lowerToUpper.GetWF()
)

func init() {
	workflow.Register(upperToLower, lowerToUpper)
}

// UpperToLower 转小写
type UpperToLower struct {
	workflow.ParentWorkflow
}

func (*UpperToLower) Use() string {
	return modelToLower
}

func (*UpperToLower) Run() command.CmdFunc {
	return upperLowerRun
}

// LowerToUpper 转大写
type LowerToUpper struct {
	workflow.ParentWorkflow
}

func (*LowerToUpper) Use() string {
	return modelToUpper
}

func (*LowerToUpper) Run() command.CmdFunc {
	return upperLowerRun
}

func upperLowerRun(cmd *cobra.Command, args []string) {
	switch cmd.Use {
	case modelToUpper:
		toUpper(args)
	case modelToLower:
		toLower(args)
	default:
		return
	}
}

func toUpper(args []string) {
	upperlowerwf.Run(func() {
		defer func() { upperlowerwf.SendFeedback() }()
		var uppers string
		for index, arg := range args {
			if index == 0 {
				uppers = uppers + upper(arg)
				continue
			}
			uppers = uppers + "," + upper(arg)
		}
		upperlowerwf.NewItem(uppers).Valid(true).Copytext(uppers).Arg(uppers).Subtitle("[CMD+C] copy")
	})
}

func upper(arg string) string {
	reg := regexp.MustCompile(`^([a-z]?)|_([a-z]?)|-([a-z]?)`) // user_id/user-id
	return reg.ReplaceAllStringFunc(arg, func(w string) string {
		for _, s := range specialSymbol {
			w = strings.Replace(w, s, "", -1)
		}
		return strings.ToUpper(w)
	})
}

func toLower(args []string) {
	upperlowerwf.Run(func() {
		defer func() { upperlowerwf.SendFeedback() }()
		var lowers string
		for index, arg := range args {
			if index == 0 {
				lowers = lowers + gorm.ToColumnName(arg)
				continue
			}
			lowers = lowers + ", " + gorm.ToColumnName(arg)
		}
		upperlowerwf.NewItem(lowers).Valid(true).Copytext(lowers).Arg(lowers).Subtitle("[CMD+C] copy")
	})
}
