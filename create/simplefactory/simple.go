package simplefactory

import "fmt"

var _ API = (*hiAPI)(nil)
var _ API = (*helloAPI)(nil)

type API interface {
	i() //防止API接口被其他的struct实现
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (h *hiAPI) i() {
}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (h *helloAPI) i() {
}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
