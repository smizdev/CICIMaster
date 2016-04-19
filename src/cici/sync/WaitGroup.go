package sync

import(
    "sync"
)

// WaitGroup
type WaitGroup interface {
    // Add
    Add(int)
    
    // Done
    Done()
    
    // Wait
    Wait()
}

// NewWaitGroup
func NewWaitGroup() WaitGroup {
    return &sync.WaitGroup{}
}