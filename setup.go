package main

import (
	"vugu-sample-project/error"
	"vugu-sample-project/page"
	"vugu-sample-project/static"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {
	router := vgrouter.New(eventEnv)

	root := &Root{}
	toastContainer := &static.ToastContainer{ToastDuration: 2000}
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
	})

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
