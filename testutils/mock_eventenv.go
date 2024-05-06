package testutils

type MockEventEnv struct{}

func (ee *MockEventEnv) Lock()         {}
func (ee *MockEventEnv) UnlockOnly()   {}
func (ee *MockEventEnv) UnlockRender() {}
func (ee *MockEventEnv) RLock()        {}
func (ee *MockEventEnv) RUnlock()      {}
