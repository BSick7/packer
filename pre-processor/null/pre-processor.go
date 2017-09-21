package null

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	ctx interpolate.Context
}

type PreProcessor struct {
	config Config
}

func (p *PreProcessor) Configure(raws ...interface{}) error {
	return nil
}

func (p *PreProcessor) PreProcess(ui packer.Ui) error {
	return nil
}

func (p *PreProcessor) Cancel() {
}
