package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptName struct {
	*config.Config
}

func (p *PromptName) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Application Name",
		Validate: RegexValidator{`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`, &p.Name}.validate,
		Default:  p.Name,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
