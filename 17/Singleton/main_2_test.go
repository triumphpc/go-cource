package Singleton

import (
	"strconv"
	"sync"
	"testing"
)

func TestSecondGetInstance(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	var w sync.WaitGroup

	for i := 0; i < 3000; i++ {
		j := i
		w.Add(1)
		go func() {
			t := "title_" + strconv.Itoa(j)
			s1.SetTitle(t)
			w.Done()
		}()
		w.Add(1)
		go func() {

			t2 := "title_2_" + strconv.Itoa(j)
			s2.SetTitle(t2)
			w.Done()
		}()
	}
	t1 := s1.GetTitle()
	t2 := s2.GetTitle()

	if t1 != t2 {
		t.Error("New instance different")
	}
}
