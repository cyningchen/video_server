package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXCUTE   = "e"
	CLOSE             = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
