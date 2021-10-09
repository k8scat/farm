package processor

import (
	"github.com/molizz/farm/exchange"
	"github.com/molizz/farm/puller"
)

type MergedEventFunc func(*exchange.Event)

func NewMergeProcessor(f MergedEventFunc) *MergeProcessor {
	return &MergeProcessor{f: f}
}

type MergeProcessor struct {
	f MergedEventFunc
}

func (p *MergeProcessor) Name() string {
	return "merge_processor"
}

func (p *MergeProcessor) Prepare(event *puller.Event) error {
	return nil
}

func (p *MergeProcessor) Process() error {
	return nil
}
