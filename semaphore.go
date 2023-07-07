package semaphore

// Semaphore is a struct defining a counting semaphore.
type Semaphore struct {
	sem chan token
}

// Make creates a new Semaphore with the given capacity.
func Make(capacity int) Semaphore {
	return Semaphore{
		make(chan token, capacity),
	}
}

// Acquire acquires a token from the semaphore.
func (s Semaphore) Acquire() {
	s.sem <- token{}
}

// Release releases a token to the semaphore.
func (s Semaphore) Release() {
	<-s.sem
}

// Capacity returns the capacity of the semaphore.
func (s Semaphore) Capacity() int {
	return cap(s.sem)
}

// Acquired returns the number of tokens currently acquired.
func (s Semaphore) Acquired() int {
	return len(s.sem)
}

// Wait blocks until all tokens have been released.
func (s Semaphore) Wait() {
	for i := 0; i < cap(s.sem); i++ {
		s.Acquire()
	}
}
