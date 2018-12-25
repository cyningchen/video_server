package taskrunner

type Runner struct {
	Controller controlChan // 控制通道， dispatcher和executor交换信息
	Error      controlChan // 错误信息通道
	Data       dataChan    // 数据通道
	dataSize   int
	Longlived  bool
	Dispatcher fn // 将数据写入Data通道
	Executor   fn // 从Data通道中获取数据处理
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
