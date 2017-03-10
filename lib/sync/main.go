package sync

// Worker defines worker func
type Worker func()

var pipe chan Worker
var done chan bool

func waiter() {
	for {
		select {
		case f := <-pipe:
			f()
			done <- true
		}
	}
}

func init() {
	pipe = make(chan Worker)
	done = make(chan bool)

	go waiter()
}

// Sync takes a woker, then call woker syncronizely
func Sync(f Worker) {
	pipe <- f
	<-done
}
