package farm

import "errors"

var (
// synchronizer
// message push center
//
)

func New() *Farm {
	f := &Farm{}
	f.pm = NewPullManager()

	return f
}

type Farm struct {
	pm *PullManager
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
	if f.pm.puller.Count() == 0 {
		return errors.New("must init pullers")
	}
	return nil
}

func (f *Farm) run() error {
	f.pm.Run()
	return nil
}

func (f *Farm) GetPuller() *PullManager {
	return f.pm
}
