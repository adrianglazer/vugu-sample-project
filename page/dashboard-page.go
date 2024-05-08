package page

import (
	"vugu-sample-project/auth"
	"vugu-sample-project/static"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type DashboardPage struct {
	vgrouter.NavigatorRef
	auth.AuthRef
	static.ToastContainerRef
}

func (c *DashboardPage) Logout(event vugu.DOMEvent) {
	event.PreventDefault()
	c.RemoveJWT()
	c.AddToast(static.ToastSuccess, "You have been logged out")
	c.Navigate("/login", nil)
}
