package farm

import "errors"

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

func (f *Farm) GetSynchronizer() *Synchronizer {
	return f.synchronizer
}
