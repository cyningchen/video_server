package taskrunner

import (
	"github.com/pkg/errors"
	"log"
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %d", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received: %v", d)
			default:
				break forloop
			}
		}
		return errors.New("Executor")
	}

	runner := NewRunner(30, false, d, e)
	go runner.startAll()
	time.Sleep(3 * time.Second)
}
