package globa

// Systems 系统命令
var Systems = map[string]string{
	"windows": "handler /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}
