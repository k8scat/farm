package puller

import (
	"errors"
	"log"

	"github.com/molizz/farm/thridparty"
	"github.com/robfig/cron"
)

var (
	ErrPullerExist = errors.New("ErrPullerExist")
)

type Event struct {
	Puller thridparty.ThridPartyPuller
	// 是否增量信息（例如新增用户、部门）
	IsIncrement bool
	// 用户、部门信息
	Users []*thridparty.ThridPartyUser
	Depts []*thridparty.ThridPartyDepartment
}

type Puller struct {
	thridPartyHub map[string]thridparty.ThridPartyPuller
	eventFuncs    []func(*Event)

	cron *cron.Cron
}

func New() *Puller {
	return &Puller{
		thridPartyHub: map[string]thridparty.ThridPartyPuller{},
		eventFuncs:    make([]func(*Event), 0, 2),
		cron:          cron.New(),
	}
}

func (p *Puller) Count() int {
	return len(p.thridPartyHub)
}

func (p *Puller) Register(puller thridparty.ThridPartyPuller) error {
	if _, exist := p.thridPartyHub[puller.Label()]; exist {
		return ErrPullerExist
	}

	p.thridPartyHub[puller.Label()] = puller

	userPuller := puller.GetPuller()
	if userPuller.HasIncrement() {
		userPuller.InjectPullIncrementCallback(p.onInjectIncrementCallback)
	}

	userPuller.InjectPullActionFunc(p.onInjectPullCallback)

	if hasCron, spec := puller.Cron(); hasCron {
		err := p.cron.AddFunc(spec, func() {
			err := p.pull(userPuller)
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

func (p *Puller) onInjectIncrementCallback(puller thridparty.ThridPartyPuller, users []*thridparty.ThridPartyUser, depts []*thridparty.ThridPartyDepartment) error {
	event := &Event{
		Puller:      puller,
		IsIncrement: true,
		Users:       users,
		Depts:       depts,
	}
	return p.onEvent(event)
}

func (p *Puller) onInjectPullCallback(puller thridparty.ThridPartyPuller) error {
	return p.pull(puller.GetPuller())
}

func (p *Puller) pull(puller thridparty.ThridPartyUserPuller) error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic: pull users and depts was err: %+v\n", e)
		}
	}()

	users, err := puller.PullUsers()
	if err != nil {
		return err
	}
	depts, err := puller.PullDepts()
	if err != nil {
		return err
	}

	// TODO puller.GetFilter()
	p.onEvent(event * Event)
	return nil
}

func (p *Puller) onEvent(event *Event) error {
	for _, fn := range p.eventFuncs {

	}
	return nil
}

func (p *Puller) RegisterEvent(fn func(event *Event)) {
	p.eventFuncs = append(p.eventFuncs, fn)
}
