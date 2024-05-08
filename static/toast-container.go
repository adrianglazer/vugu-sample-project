package static

import (
	"github.com/vugu/vugu"
)

type ToastContainer struct {
	ToastDuration int            `vugu:"data"` // millis
	Toasts        map[uint]Toast `vugu:"data"`
	vugu.EventEnv
}

func (tc *ToastContainer) Init(ctx vugu.InitCtx) {
	tc.EventEnv = ctx.EventEnv()

	if tc.Toasts == nil {
		// init only once
		tc.Toasts = make(map[uint]Toast)
	}
}

func (tc *ToastContainer) AddToast(tostType uint8, toastMessage string) {

	go func() {
		tc.Lock()
		defer tc.UnlockRender()
		key := uint(len(tc.Toasts) + 1)
		tc.Toasts[key] = Toast{ToastKey: key, ToastType: tostType, ToastMessage: toastMessage, ToastContainer: tc}
	}()
}

func (tc *ToastContainer) RemoveToast(t *Toast) {
	go func() {
		tc.Lock()
		defer tc.UnlockRender()
		delete(tc.Toasts, t.ToastKey)
	}()
}

type ToastContainerRef struct{ *ToastContainer }

type ToastContainerSetter interface {
	SetToastContainer(tc *ToastContainer)
}

func (tcr *ToastContainerRef) SetToastContainer(tc *ToastContainer) { tcr.ToastContainer = tc }
