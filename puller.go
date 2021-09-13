package farm

import (
	"github.com/molizz/farm/puller"
	"github.com/molizz/farm/thridparty"
)

type Puller interface {
	Count() int
	Register(puller thridparty.ThridPartyPuller) error
	RegisterEvent(fn func(event *puller.Event))
	Start()
	Stop()
}

var _ Puller = (*puller.Puller)(nil)

type PullManager struct {
	puller Puller
}

func NewPullManager() *PullManager {
	p := &PullManager{}
	p.puller = puller.New()
	return p
}
