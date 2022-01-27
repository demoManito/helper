package workflowlist

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLoadData(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/jesse/GolandProjects/helper/static/workflows.json")
	fmt.Println(data, err)
}
