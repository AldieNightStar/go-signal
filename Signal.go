package signal

import (
	"reflect"
)

type CallBack func(item interface{})

type Signal struct {
	cbs     []CallBack
	cbsOnce []CallBack
}

func (s *Signal) Connect(fun CallBack) {
	s.cbs = append(s.cbs, fun)
}

func (s *Signal) ConnectOnce(fun CallBack) {
	s.cbsOnce = append(s.cbs, fun)
}

func (s *Signal) Disconnect(fun CallBack) {
	toDelete := -1
	for i, c := range s.cbs {
		cInd := reflect.ValueOf(c)
		fInd := reflect.ValueOf(fun)
		if cInd.Pointer() == fInd.Pointer() {
			toDelete = i
		}
	}
	if toDelete > -1 {
		s.cbs = withRemovedCallBack(s.cbs, toDelete)
	}
}

func (s *Signal) Emit(data interface{}) {
	for _, f := range s.cbs {
		f(data)
	}
	// Call's CBSOnce and clear it
	for _, f := range s.cbsOnce {
		f(data)
	}
	s.cbsOnce = make([]CallBack, 0, 8)
}

func (s *Signal) Wait() interface{} {
	c := make(chan interface{}, 1)

	// Add in parrallel func that will wait on data
	//    and will write to a data channel one record and close it
	go func() {
		s.ConnectOnce(func(data interface{}) {
			c <- data
			close(c)
		})
	}()
	return <-c
}

// Creates new signal
func NewSignal() *Signal {
	return &Signal{
		cbs:     make([]CallBack, 0, 8),
		cbsOnce: make([]CallBack, 0, 8),
	}
}

func withRemovedCallBack(arr []CallBack, index int) []CallBack {
	arrNew := make([]CallBack, 0, len(arr))
	for i, cb := range arr {
		if i == index {
			continue
		}
		arrNew = append(arrNew, cb)
	}
	return arrNew
}
