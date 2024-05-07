package router

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vugu/vugu"
	"github.com/vugu/vugu/js"
)

type Router struct {
	ShowWasm bool `vugu:"data"`
	ShowGo   bool `vugu:"data"`
	ShowVugu bool `vugu:"data"`
	Cnt      int  `vugu:"data"`
}

func (c *Router) PresentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (c *Router) LocalStorageItemGet() string {
	return js.Global().Get("localStorage").Call("getItem", "mykey").String()
}

func (c *Router) AddCnt() {
	var mykey = js.Global().Get("localStorage").Call("getItem", "mykey").String()
	cnt, err := strconv.Atoi(mykey)

	if err != nil {
		cnt = 0
	}

	c.Cnt = cnt
	fmt.Println(cnt)

	js.Global().Get("localStorage").Call("setItem", "mykey", cnt+1)
	fmt.Println("local storage key set")
}

func (c *Router) Init(ctx vugu.InitCtx) {
	c.AddCnt()
}
