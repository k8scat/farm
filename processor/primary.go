package processor

import "github.com/molizz/farm/puller"

func NewPrimaryProcessor() *PrimaryProcessor {
	return &PrimaryProcessor{}
}

type PrimaryProcessor struct {
}

func (p *PrimaryProcessor) Name() string {
	return "primary_processor"
}

func (p *PrimaryProcessor) Prepare(event *puller.Event) error {
	return nil
}

func (p *PrimaryProcessor) Process() error {
	return nil
}
