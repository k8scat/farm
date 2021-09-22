package farm

import (
	"github.com/molizz/farm/puller"
	"github.com/molizz/farm/thirdparty"
)

type Puller interface {
	Count() int
	Register(puller thirdparty.ThirdPartyPuller) error
	RegisterEvent(fn puller.EventCallbck)
	Start()
}

type Processor interface {
}

// EventConsumer 将Puller拉下来的数据进行消费
type EventConsumer interface {
}

var _ Puller = (*puller.Puller)(nil)

type Synchronizer struct {
	puller Puller
}

func NewSynchronizer() *Synchronizer {
	p := &Synchronizer{}
	p.puller = puller.New()
	p.puller.RegisterEvent(p.onEvent)
	return p
}

func (p *Synchronizer) RegisterPuller(pullers ...thirdparty.ThirdPartyPuller) error {
	for _, pl := range pullers {
		if err := p.puller.Register(pl); err != nil {
			return err
		}
	}
	return nil
}

func (p *Synchronizer) RegisterProcessor(processes ...Processor) error {
	return nil
}

func (p *Synchronizer) Do() error {
	p.puller.Start()
	return nil
}

func (p *Synchronizer) onEvent(event *puller.Event) error {
	// TODO
	return nil
}
