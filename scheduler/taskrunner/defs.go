package taskrunner

const (
	READY_TO_DISPATCH = "d" // 准备好了接收 -> controlChan
	READY_TO_EXCUTE   = "e" // 准备好了执行 -> controlChan
	CLOSE             = "c" // 准备关闭通道 -> controlChan
	VIDEO_PATH        = "./videos"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error // 闭包函数
