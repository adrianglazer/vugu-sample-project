package layout

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/js"
)

type Sidebar struct {
	vgrouter.NavigatorRef
	ctx     vugu.InitCtx
	Toggled bool `vugu:"data"`
}

func (c *Sidebar) Init(ctx vugu.InitCtx) {
	c.ctx = ctx
	c.Toggled = false

	callback := js.FuncOf(func(v1 js.Value, v2 []js.Value) interface{} {
		if js.Global().Get("innerWidth").Int() < 768 {
			ctx.EventEnv().Lock()
			c.Toggled = true
			ctx.EventEnv().UnlockRender()
		}
		return nil
	})

	js.Global().Call("addEventListener", "resize", callback)
}

func (c *Sidebar) makeAttrs() (ret []vugu.VGAttribute) {
	class := "navbar-nav bg-gradient-primary sidebar sidebar-dark accordion"

	if c.Toggled {
		class = class + " toggled"
	}
	ret = append(ret, vugu.VGAttribute{
		Key: "class",
		Val: class,
	})
	return
}

func (c *Sidebar) Toggle() {
	go func() {
		c.ctx.EventEnv().Lock()
		c.Toggled = !c.Toggled
		c.ctx.EventEnv().UnlockRender()
	}()
}
