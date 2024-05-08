package login

import (
	"log"
	"time"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type Login struct {
	vgrouter.NavigatorRef
	Username  string `vugu:"data"`
	Password  string `vugu:"data"`
	IsLoading bool   `vugu:"data"`
}

func (c *Login) Login(event vugu.DOMEvent) {
	ee := event.EventEnv()
	log.Println("login...")

	go func() {
		// lock and unlock allows go routine to safely access component
		// Lock will block the component or put routine in a queue if component is already locked
		ee.Lock()
		log.Println("loading...")
		c.IsLoading = true
		// below unlocks the component and renders it again to apply all the changes made in between (in this example IsLoading was affected)
		ee.UnlockRender()

		time.Sleep(2 * time.Second)

		ee.Lock()
		defer ee.UnlockRender()
		c.IsLoading = false
		log.Println("login finished")
	}()
}
