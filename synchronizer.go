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

// Processor 处理Puller的数据
type Processor interface {
}

// Subscriber 同步后的数据进行整理后拆成单独的event，并推送到订阅者
type Subscriber interface {
}

var _ Puller = (*puller.Puller)(nil)

type Synchronizer struct {
	puller      Puller
	processes   []Processor  // TODO
	subscribers []Subscriber // TODO
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
	// TODO 将event数据进行filter操作
	// TODO 将event数据进行merge到数据库
	// TODO 根据merge得结果产生event通知到订阅者

	return nil
}
