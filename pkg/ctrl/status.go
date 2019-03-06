package ctrl

import (
	"net/http"
)

type StatusCtrl interface {
	BaseCtrl
	Show(rw http.ResponseWriter, r *http.Request)
}

type statusCtrl struct {
}

func NewStatusCtrl() StatusCtrl {
	c := &statusCtrl{}
	return c
}

func (c *statusCtrl) Name() string {
	return "status controller"
}

func (c *statusCtrl) Show(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Status Ok."))
}
