package conversion

import "testing"

func TestJsonRun(t *testing.T) {
	json := new(JsonFormat)
	json.run(nil, []string{"{:1}"}) // print error log
	json.run(nil, []string{"{\"a\":1}"})
}
