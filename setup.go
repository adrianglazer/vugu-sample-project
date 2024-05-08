package main

import (
	"vugu-sample-project/auth"
	"vugu-sample-project/error"
	"vugu-sample-project/page"
	"vugu-sample-project/static"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {
	router := vgrouter.New(eventEnv)

	root := &Root{}
	toastContainer := &static.ToastContainer{ToastDuration: 5000}
	authService := &auth.Auth{}
	root.ToastContainer = toastContainer
	buildEnv.WireComponent(root)

	// loop on components and inject dependencies if implemented
	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
		if c, ok := b.(static.ToastContainerSetter); ok {
			c.SetToastContainer(toastContainer)
		}
		if c, ok := b.(auth.AuthSetter); ok {
			c.SetAuth(authService)
		}
	})

	router.MustAddRouteExact("/",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {

			if authService.IsAuthenticated() {
				root.Body = &page.DashboardPage{}
			} else {
				router.Navigate("/login", nil)
			}
		}))

	router.MustAddRouteExact("/login",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &page.LoginPage{}
		}))

	router.SetNotFound(vgrouter.RouteHandlerFunc(
		func(rm *vgrouter.RouteMatch) {
			root.Body = &error.Error404{}
		}))

	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}
