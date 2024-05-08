package static

import (
	"time"

	"github.com/vugu/vugu"
)

const (
	ToastSuccess uint8 = 1
	ToastWarning uint8 = 2
	ToastDanger  uint8 = 3
)

type Toast struct {
	ToastKey       uint   `vugu:"data"`
	ToastType      uint8  `vugu:"data"`
	ToastMessage   string `vugu:"data"`
	ToastContainer *ToastContainer
}

func (t *Toast) Init(ctx vugu.InitCtx) {
	go func() {
		defer t.ToastContainer.RemoveToast(t)
		time.Sleep(time.Duration(t.ToastContainer.ToastDuration * int(time.Millisecond)))
	}()
}

func (t *Toast) makeAttrs() (ret []vugu.VGAttribute) {

	tp := "success"
	switch t.ToastType {
	case ToastWarning:
		tp = "warning"
	case ToastDanger:
		tp = "danger"
	}

	ret = append(ret, vugu.VGAttribute{
		Key: "class",
		Val: "toast-header bg-" + tp + "-subtle",
	})
	return
}
