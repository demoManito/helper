package conversion

import (
	"helper/handler/command"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
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
	upperToLower command.ICommand = new(UpperToLower)
	lowerToUpper command.ICommand = new(LowerToUpper)
)

func init() {
	command.Register(upperToLower, lowerToUpper)
}

// UpperToLower 转小写
type UpperToLower struct {
	command.ParentCommand
}

func (*UpperToLower) Use() string {
	return modelToLower
}

func (*UpperToLower) Run() command.CmdFunc {
	return upperLowerRun
}

// LowerToUpper 转大写
type LowerToUpper struct {
	command.ParentCommand
}

func (*LowerToUpper) Use() string {
	return modelToUpper
}

func (*LowerToUpper) Run() command.CmdFunc {
	return upperLowerRun
}

/**
 * to upper:
 * helper tu user_id => UserId
 * helper tu user_id user_id => UserId, UserId
 *
 * to lower:
 * helper tl UserID => user_id
 * helper tl UserID userID userId => user_id, user_id, user_id
 */
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
	var uppers string
	for index, arg := range args {
		if index == 0 {
			uppers = uppers + upper(arg)
			continue
		}
		uppers = uppers + "," + upper(arg)
	}
	clipboard.WriteAll(uppers)
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
	var lowers string
	for index, arg := range args {
		if index == 0 {
			lowers = lowers + gorm.ToColumnName(arg)
			continue
		}
		lowers = lowers + ", " + gorm.ToColumnName(arg)
	}
	clipboard.WriteAll(lowers)
}
