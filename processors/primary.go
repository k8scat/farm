package processors

import "github.com/molizz/farm/puller"

func NewPrimaryProccessor() *PrimaryProccessor {
	return &PrimaryProccessor{}
}

type PrimaryProccessor struct {
}

func (p *PrimaryProccessor) Prepare(event *puller.Event) error {
	return nil
}

func (p *PrimaryProccessor) Process() error {
	return nil
}
