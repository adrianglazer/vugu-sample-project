package testutils

import (
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/js"
)

type MockDOMEvent struct{}

func (e *MockDOMEvent) Prop(keys ...string) interface{} {
	var ret interface{}
	return ret
}

func (e *MockDOMEvent) PropString(keys ...string) string {
	return "empty"
}

func (e *MockDOMEvent) PropFloat64(keys ...string) float64 {
	return 0
}

func (e *MockDOMEvent) PropBool(keys ...string) bool {
	return false
}

func (e *MockDOMEvent) EventSummary() map[string]interface{} {
	var ret map[string]interface{}
	return ret
}

func (e *MockDOMEvent) JSEvent() js.Value {
	return js.Value{}
}

func (e *MockDOMEvent) JSEventTarget() js.Value {
	return js.Value{}
}

func (e *MockDOMEvent) JSEventCurrentTarget() js.Value {
	return js.Value{}
}

func (e *MockDOMEvent) EventEnv() vugu.EventEnv {
	return &MockEventEnv{}
}

func (e *MockDOMEvent) PreventDefault()  {}
func (e *MockDOMEvent) StopPropagation() {}
