package application

import "sync"

type StopSignal struct {
	stoplock      sync.Mutex
	stopGap       int
	stopThreshold int
}

func NewStopSignal(stop_Threshold int) *StopSignal {
	return &StopSignal{
		stopGap:       0,
		stopThreshold: stop_Threshold,
	}
}

func (ss *StopSignal) StopGap_Inc() {
	ss.stoplock.Lock()
	ss.stopGap++
	ss.stoplock.Unlock()
}

func (ss *StopSignal) StopGap_Reset() {
	ss.stoplock.Lock()
	ss.stopGap = 0
	ss.stoplock.Unlock()
}
func (ss *StopSignal) GapEnough() bool {
	ss.stoplock.Lock()
	defer ss.stoplock.Unlock()
	return ss.stopGap >= ss.stopThreshold

}
