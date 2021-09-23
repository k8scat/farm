package farm

import (
	"github.com/molizz/farm/exchange"
	"github.com/molizz/farm/puller"
	"github.com/molizz/farm/thirdparty"
)

type Puller interface {
	Count() int
	Register(label string, puller thirdparty.ThirdPartyPuller) error
	RegisterEvent(fn puller.EventCallbck)
	Start()
}

// Processor 处理Puller的数据
type Processor interface {
}

var _ Puller = (*puller.Puller)(nil)

type Synchronizer struct {
	puller Puller

	// processor 当从puller中拉取到数据后，对数据进行处理
	processes []Processor

	// exchange 当processes处理完成后，将数据推送到exchange，Exchange将负责将数据推送出去
	exchange *exchange.Exchange
}

func NewSynchronizer() *Synchronizer {
	p := &Synchronizer{}
	p.puller = puller.New()
	p.puller.RegisterEvent(p.onEvent)
	return p
}

func (p *Synchronizer) RegisterPuller(label string, puller thirdparty.ThirdPartyPuller) error {
	if err := p.puller.Register(label, puller); err != nil {
		return err
	}
	return nil
}

func (p *Synchronizer) RegisterProcessor(processes ...Processor) error {
	return nil
}

func (p *Synchronizer) RegisterSubscriber(subscriber exchange.Subscriber) error {
	return nil
}

func (p *Synchronizer) Do() error {
	p.puller.Start()
	return nil
}

func (p *Synchronizer) onEvent(event *puller.Event) error {
	// TODO 将event数据进行filter操作
	// TODO 将event数据进行merge到数据库
	// TODO 根据merge得结果产生event通知到订阅者

	return nil
}
