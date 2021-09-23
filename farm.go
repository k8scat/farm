package farm

import (
	"errors"

	"github.com/molizz/farm/thirdparty"
)

var (
// synchronizer
// message push center
//
)

func New() *Farm {
	f := &Farm{}
	f.synchronizer = NewSynchronizer()

	return f
}

type Farm struct {
	thirdparties map[string]thirdparty.ThirdParty
	synchronizer *Synchronizer
}

func (f *Farm) Start() error {
	if err := f.verify(); err != nil {
		return err
	}
	if err := f.run(); err != nil {
		return err
	}
	return nil
}

func (f *Farm) verify() error {
	if f.synchronizer.puller.Count() == 0 {
		return errors.New("must init pullers")
	}
	return nil
}

func (f *Farm) run() error {
	if err := f.synchronizer.Do(); err != nil {
		return err
	}
	return nil
}

func (f *Farm) Register(tp thirdparty.ThirdParty) error {
	f.thirdparties[tp.Label()] = tp

	if tp.GetThirdPartyPuller() != nil {
		f.synchronizer.RegisterPuller(tp.Label(), tp.GetThirdPartyPuller())
	}
	if tp.GetUserManager() != nil {
		// TODO
	}
	if tp.GetMessager() != nil {
		// TODO
	}
	if tp.GetOAuth2() != nil {
		// TODO
	}
	if tp.GetAuthorizer() != nil {
		// TODO
	}
	if tp.GetCaller() != nil {
		// TODO
	}
	if tp.GetConfiger() != nil {
		// TODO
	}
	return nil
}

func (f *Farm) GetSynchronizer() *Synchronizer {
	return f.synchronizer
}
