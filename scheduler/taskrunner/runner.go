package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	Longlived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		Longlived:  longlived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatcher() {
	defer func() {
		if !r.Longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXCUTE
				}
			}
			if c == READY_TO_EXCUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:
		}
	}
}

func (r *Runner) startAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatcher()
}
