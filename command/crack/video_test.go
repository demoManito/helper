package crack

import "testing"

func TestOpenBrowser(t *testing.T) {
	v := new(Video)
	v.run(nil, []string{"po", "https://v.qq.com/x/cover/m441e3rjq9kwpsc/k00417vlpjq.html"})
}
