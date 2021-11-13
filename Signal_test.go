package signal

import (
	"testing"
	"time"
)

func Test_Signal_Connect(t *testing.T) {
	s := NewSignal()

	dt := make(chan interface{}, 1)

	s.Connect(func(dat interface{}) {
		dt <- dat
		close(dt)
	})

	s.Emit("Heek")

	d := <-dt

	if d != "Heek" {
		t.Fatal("Connected signal returned wrong data. Need 'Heek' instead we have: ", d)
	}
}

func Test_Signal_Disconnect(t *testing.T) {
	s := NewSignal()

	num := 0

	adder1 := func(d interface{}) {
		num += 1
	}
	adder2 := func(d interface{}) {
		num += 2
	}
	adder3 := func(d interface{}) {
		num += 3
	}

	s.Connect(adder1)
	s.Connect(adder2)
	s.Connect(adder3)

	s.Disconnect(adder2)

	s.Emit(0)

	if num != 4 {
		t.Fatal("Function is actually not disconnected! Num should be 4, but", num)
	}

}

func Test_Signal_Wait(t *testing.T) {
	s := NewSignal()

	// Emit after 1 sec
	go func() {
		time.Sleep(time.Millisecond * 100)
		s.Emit(333)
	}()

	// Wait for data
	dat := s.Wait().(int)

	// Try to compare
	if dat != 333 {
		t.Fatal("Value isn't waited or not eq to 333: ", dat)
	}
}
