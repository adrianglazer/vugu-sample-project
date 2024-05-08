package error

import (
	"log"

	"github.com/vugu/vugu"
)

type Error404 struct{}

func (c *Error404) Init(ctx vugu.InitCtx) {
	log.Println("404 Not Found")
}
