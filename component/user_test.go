package component

import (
	"testing"
	testutils "vugu-sample-project/testutils"
)

func TestShowWasm(t *testing.T) {
	var c = User{ShowWasmVal: false, ShowGo: false, ShowVugu: false}

	d := testutils.MockDOMEvent{}
	c.ShowWasm(&d)

	if c.ShowWasmVal == false {
		t.Fail()
	}
}
