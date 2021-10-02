package farm

import (
	"log"
	"time"

	"github.com/molizz/farm/exchange"
	"github.com/molizz/farm/processor"
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
	Name() string
	Prepare(event *puller.Event) error
	Process() error
}

var _ Puller = (*puller.Puller)(nil)

type Synchronizer struct {
	puller Puller

	// processor 当从puller中拉取到数据后，对数据进行处理
	processes []Processor

	// exchange 当processes处理完成后，将数据推送到exchange，Exchange将负责将数据推送出去
	exchange *exchange.Exchange
}

func NewSynchronizer() thirdparty.Synchronizer {
	syncer := &Synchronizer{}
	syncer.puller = puller.New()
	syncer.puller.RegisterEvent(syncer.onEvent)

	syncer.processes = syncer.defaultProcessors()
	return syncer
}

func (p *Synchronizer) RegisterPuller(label string, puller thirdparty.ThirdPartyPuller) error {
	if err := p.puller.Register(label, puller); err != nil {
		return err
	}
	return nil
}

func (p *Synchronizer) RegisterProcessor(processes ...Processor) error {
	p.processes = append(p.processes, processes...)
	return nil
}

func (p *Synchronizer) PullerCount() int {
	return p.puller.Count()
}

func (p *Synchronizer) RegisterSubscriber(subscriberes ...exchange.Subscriber) error {
	for _, sub := range subscriberes {
		p.exchange.AddSubscriber(sub)
	}
	return nil
}

func (p *Synchronizer) Do() error {
	p.puller.Start()
	return nil
}

func (p *Synchronizer) defaultProcessors() []Processor {
	ret := []Processor{
		processor.NewPrimaryProcessor(),
	}
	return ret
}

func (p *Synchronizer) onEvent(event *puller.Event) (err error) {
	// TODO 预备&检查数据
	// 		例 thirdparty.ThirdPartyPulledPack 中 Users 与 depts hash是否与上一次相同，相同则跳过所有的处理器）
	// 		例 thirdparty.ThirdPartyPulledPack 中 Users 的属性是否与 ThirdPartyUserPuller.UserPrimaryAttrs() 中提供的字段匹配）
	// 		例 thirdparty.ThirdPartyPulledPack 中 Users 的属性是否与 ThirdPartyUserPuller.DepartmentPrimaryAttr() 中提供的字段匹配）
	// 		例 thirdparty.ThirdPartyPulledPack 中 Depts 的属性是否与 ThirdPartyUserPuller.DepartmentPrimaryAttr() 中提供的字段匹配）
	// TODO 清洗不合法的数据（例 db.metadata 是否初始化，没有初始化则根据 thirdparty.ThirdPartyPulledPack.Users 字段初始化该三方数据库中 columns metadata ）
	// TODO 清洗不合法的数据（例 返回的属性与db.metadata匹配不上）
	// TODO 清洗不合法的数据（例 primary属性值为空）
	// TODO 清洗不合法的数据（例 存在重复的primary属性值）
	// TODO 清洗不合法的数据（例 部门中的父子层级对不上的将放在根节点）
	// TODO 将event数据进行merge到数据库
	// TODO 根据merge得结果产生event通知信息
	// TODO 根据拿到通知信息，根据filter过滤，将通知信息推送到exchange

	for _, process := range p.processes {
		now := time.Now()
		log.Printf("process '%s'\n", process.Name())
		if err = process.Prepare(event); err != nil {
			return err
		}
		if err = process.Process(); err != nil {
			return err
		}
		log.Printf("process '%s' total: %v", process.Name(), time.Now().Sub(now))
	}
	return nil
}
