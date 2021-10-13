package exchange

import (
	"errors"
	"fmt"
	"testing"

	"github.com/reactivex/rxgo/v2"
)

var _ Subscriber = (*TestSubscriber)(nil)

type TestSubscriber struct {
	handleMustOK   bool
	handleExecuted bool
}

func (t *TestSubscriber) IsEnable() bool { return true }

func (t *TestSubscriber) Label() string { return "test" }

func (t *TestSubscriber) Actions() []Action {
	return []Action{ActionCreate, ActionDelete, ActionUpdate}
}

func (t *TestSubscriber) Handle(event *Event) error {
	t.handleExecuted = true
	fmt.Println("handle...", event.ToJSON())
	if t.handleMustOK {
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
		observable         rxgo.Observable
	}

	type args struct {
		shouldFunc func(*Event) error
	}

	var items = []interface{}{"1", "2"}

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
				affectedSubscriber: &TestSubscriber{handleMustOK: true},
				observable:         rxgo.Just(items...)(),
			},
			args: args{func(event *Event) error {
				return errors.New("handle err")
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PipeEvent{
				event:              tt.fields.event,
				affectedSubscriber: tt.fields.affectedSubscriber,
				observable:         tt.fields.observable,
			}
			if err := p.Wait(tt.args.shouldFunc); (err != nil) != tt.wantErr {
				t.Errorf("Wait() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				// err
				i := 0
				for val := range p.observable.Observe() {
					if items[i] != val.V.(string) {
						t.Errorf("Observe 返回错误时间，items的数据应该恢复一致，当前值：%s", val.V.(string))
					}
					i++
				}
			}
		})
	}
}
