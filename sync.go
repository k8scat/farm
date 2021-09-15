package farm

import (
	"github.com/molizz/farm/puller"
	"github.com/molizz/farm/thridparty"
)

type Puller interface {
	Count() int
	Register(puller thridparty.ThridPartyPuller) error
	RegisterEvent(fn puller.EventCallbck)
	Start()
}

var _ Puller = (*puller.Puller)(nil)

type PullManager struct {
	puller Puller
}

func NewPullManager() *PullManager {
	p := &PullManager{}
	p.puller = puller.New()
	p.puller.RegisterEvent(p.onEvent)
	return p
}

func (p *PullManager) RegisterPuller(pullers ...thridparty.ThridPartyPuller) {
	for _, puller := range pullers {
		p.puller.Register(puller)
	}
}

func (p *PullManager) Run() {
	p.puller.Start()
}

func (p *PullManager) onEvent(event *puller.Event) error {
	// TODO
	return nil
}
