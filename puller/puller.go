package puller

import (
	"errors"
	"log"

	"github.com/molizz/farm/thirdparty"
	"github.com/robfig/cron/v3"
)

var (
	ErrPullerExist = errors.New("ErrPullerExist")
)

type EventCallbck func(event *Event) error

type Event struct {
	// 是否增量信息（例如新增用户、部门）
	IsIncrement bool
	// 用户、部门信息
	Pack *thirdparty.ThirdPartyPulledPack
}

type Puller struct {
	thirdPartyHub map[string]thirdparty.ThirdPartyPuller
	eventFuncs    []EventCallbck

	cron *cron.Cron
}

func New() *Puller {
	return &Puller{
		thirdPartyHub: map[string]thirdparty.ThirdPartyPuller{},
		eventFuncs:    make([]EventCallbck, 0, 2),
		cron:          cron.New(),
	}
}

func (p *Puller) Count() int {
	return len(p.thirdPartyHub)
}

func (p *Puller) Register(label string, puller thirdparty.ThirdPartyPuller) error {
	if _, exist := p.thirdPartyHub[label]; exist {
		return ErrPullerExist
	}

	p.thirdPartyHub[label] = puller

	userPuller := puller.GetPuller()
	if userPuller.HasIncrement() {
		err := userPuller.InjectPullIncrementCallback(p.onInjectIncrementCallback)
		if err != nil {
			return err
		}
	}

	if err := userPuller.InjectPullActionFunc(p.onInjectPullCallback); err != nil {
		return err
	}

	if hasCron, spec := puller.Cron(); hasCron {
		_, err := p.cron.AddFunc(spec, func() {
			err := p.pull(puller)
			if err != nil {
				log.Printf("pull was error: %+v\n", err)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Puller) Start() {
	p.cron.Start()
}

func (p *Puller) Stop() {
	p.cron.Stop()
}

func (p *Puller) onInjectIncrementCallback(
	puller thirdparty.ThirdPartyPuller,
	pack *thirdparty.ThirdPartyPulledPack) error {

	event := &Event{
		IsIncrement: true,
		Pack:        pack,
	}
	return p.onEvent(event)
}

func (p *Puller) onInjectPullCallback(puller thirdparty.ThirdPartyPuller) error {
	return p.pull(puller)
}

func (p *Puller) pull(puller thirdparty.ThirdPartyPuller) error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic: pull users and depts got err: %+v\n", e)
		}
	}()

	pack, err := puller.GetPuller().Pull()
	if err != nil {
		return err
	}

	return p.onEvent(&Event{
		IsIncrement: false,
		Pack:        pack,
	})
}

func (p *Puller) onEvent(event *Event) error {
	log.Printf("onEvent: users count: %d deps count: %d \n",
		len(event.Pack.Users),
		len(event.Pack.Depts))

	for _, fn := range p.eventFuncs {
		if err := fn(event); err != nil {
			log.Printf("eval eventFuncs() got err: %+v\n", err)
		}
	}
	return nil
}

func (p *Puller) RegisterEvent(fn EventCallbck) {
	p.eventFuncs = append(p.eventFuncs, fn)
}
