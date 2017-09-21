package packer

type MockPreProcessor struct {
	Error error

	ConfigureCalled  bool
	ConfigureConfigs []interface{}
	ConfigureError   error

	PreProcessCalled bool
	PreProcessUi     Ui
}

func (t *MockPreProcessor) Configure(configs ...interface{}) error {
	t.ConfigureCalled = true
	t.ConfigureConfigs = configs
	return t.ConfigureError
}

func (t *MockPreProcessor) PostProcess(ui Ui) error {
	t.PreProcessCalled = true
	t.PreProcessUi = ui

	return t.Error
}
