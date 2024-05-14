package page

import (
	"vugu-sample-project/auth"
	"vugu-sample-project/generic"
	"vugu-sample-project/layout"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type DashboardPage struct {
	vgrouter.NavigatorRef
	auth.AuthRef
	generic.ToastContainerRef
	Sidebar       *layout.Sidebar       `vugu:"data"`
	SidebarToggle *layout.SidebarToggle `vugu:"data"`
}

func (c *DashboardPage) Init(ctx vugu.InitCtx) {
	c.Sidebar = new(layout.Sidebar)
	c.SidebarToggle = new(layout.SidebarToggle)
	c.SidebarToggle.Toggle = c.Sidebar.Toggle
}

func (c *DashboardPage) Logout(event vugu.DOMEvent) {
	event.PreventDefault()
	c.RemoveJWT()
	c.AddToast(generic.ToastSuccess, "You have been logged out")
	c.Navigate("/login", nil)
}
