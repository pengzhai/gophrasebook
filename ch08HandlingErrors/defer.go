package main
import "sync"

func callLocked(lock *sync.Mutex, f func()) {
	lock.Lock()
	defer lock.Unlock()
	f()
}
