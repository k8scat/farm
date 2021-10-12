package exchange

import (
	"errors"
	"testing"

	"github.com/reactivex/rxgo/v2"
)

var _ Subscriber = (*TestSubscriber)(nil)

type TestSubscriber struct {
	handleOK bool
}

func (t *TestSubscriber) IsEnable() bool { return true }

func (t *TestSubscriber) Label() string { return "test" }

func (t *TestSubscriber) Actions() []Action { return []Action{ActionCreate} }

func (t *TestSubscriber) Handle(event *Event) error {
	if t.handleOK {
		return nil
	} else {
		return errors.New("not ok")
	}
}

func (t *TestSubscriber) LastOffset() uint64 { return 0 }

func (t *TestSubscriber) SetOffset(u uint64) error { return nil }

func TestPipeEvent_Wait(t *testing.T) {
	type fields struct {
		event              *Event
		affectedSubscriber Subscriber
		obs                rxgo.Observable
	}

	type args struct {
		shouldFunc func(*Event) error
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				event: &Event{
					Action:      ActionCreate,
					Context:     nil,
					Offset:      0,
					Users:       nil,
					Departments: nil,
				},
				affectedSubscriber: &TestSubscriber{handleOK: true},
				obs:                rxgo.Just("1", "2")(),
			},
			args: args{func(event *Event) error {
				return errors.New("handle err")
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PipeEvent{
				event:              tt.fields.event,
				affectedSubscriber: tt.fields.affectedSubscriber,
				obs:                tt.fields.obs,
			}
			if err := p.Wait(tt.args.shouldFunc); (err != nil) != tt.wantErr {
				t.Errorf("Wait() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
