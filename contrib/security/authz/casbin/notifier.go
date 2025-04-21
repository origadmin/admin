package casbin

import (
	"sync"

	"github.com/origadmin/runtime/log"
)

type PolicyNotifier interface {
	AddObserver(name string, callback func())
	RemoveObserver(name string)
	NotifyAll()
}

type policyNotifier struct {
	observers map[string]func()
	mu        sync.RWMutex
}

func NewPolicyNotifier() PolicyNotifier {
	return &policyNotifier{
		observers: make(map[string]func()),
	}
}

func (n *policyNotifier) AddObserver(name string, callback func()) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.observers[name] = callback
	log.Debugf("Added policy observer: %s", name)
}

func (n *policyNotifier) RemoveObserver(name string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.observers, name)
	log.Debugf("Removed policy observer: %s", name)
}

func (n *policyNotifier) NotifyAll() {
	n.mu.RLock()
	defer n.mu.RUnlock()

	log.Info("Notifying all policy observers")
	for name, cb := range n.observers {
		log.Debugf("Triggering policy update for: %s", name)
		go func(name string, callback func()) {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("Policy observer %s panic: %v", name, err)
				}
			}()
			callback()
		}(name, cb)
	}
}
