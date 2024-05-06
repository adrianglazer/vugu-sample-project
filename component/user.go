package component

import (
	"log"
	"time"

	"github.com/vugu/vugu"
)

type User struct {
	ShowWasmVal bool `vugu:"data"`
	ShowGo      bool `vugu:"data"`
	ShowVugu    bool `vugu:"data"`
	DefaultSlot vugu.Builder
	Router      *Router
	Name        string
}

func (c *User) ShowWasm(e vugu.DOMEvent) {
	e.PreventDefault()
	c.ShowWasmVal = !c.ShowWasmVal
	log.Printf("Toggled! Show is now %t", c.ShowWasmVal)
	c.Router.AddCnt()
}

func (c *User) PresentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
