package exchange

type done struct {
	key string
	p   *Process
}

func (d *done) Wait() {
	d.p.output(d.key)
}

type Process struct {
	keys map[string]chan struct{}
}

func NewProcess(size int) *Process {
	return &Process{
		keys: make(map[string]chan struct{}, size),
	}
}

func (p *Process) Start(key string) *done {
	p.keys[key] <- struct{}{}
	return &done{
		key: key,
		p:   p,
	}
}

func (p *Process) input(key string) {
	p.keys[key] <- struct{}{}
}

func (p *Process) output(key string) {
	<-p.keys[key]
}
