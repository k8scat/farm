package farm

import (
	"database/sql"
	"errors"

	dbModel "github.com/molizz/farm/model/db"
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
	synchronizer thirdparty.Synchronizer
}

// SetMysqlDB driverName 大概是
func (f *Farm) SetMysqlDB(db *sql.DB) {
	dbModel.SetDB(db, "mysql")
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
	dbModel.MustInit()

	if len(f.thirdparties) == 0 {
		return errors.New("must init thirdParties")
	}

	if f.synchronizer.PullerCount() == 0 {
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
		err := f.synchronizer.RegisterPuller(tp.Label(), tp.GetThirdPartyPuller())
		if err != nil {
			return err
		}
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

func (f *Farm) GetSynchronizer() thirdparty.Synchronizer {
	return f.synchronizer
}
