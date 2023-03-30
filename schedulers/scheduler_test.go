package schedulers

import (
	"testing"
	"time"
)

func TestScheduler_Attach(t *testing.T) {
	scheduler := NewScheduler()
	_, err := scheduler.Attach(func() {

		t.Error("hello world")
	},
		time.Second)
	if err != nil {
		t.Errorf("expected nil, got nil %v", err)
	}
	scheduler.Start()
	time.Sleep(10 * time.Second)
}

func TestScheduler_Detach(t *testing.T) {
	scheduler := NewScheduler()
	_, err := scheduler.Attach(func() {

		t.Log("hello world")
	},
		time.Second)
	if err != nil {
		t.Errorf("expected nil, got nil %v", err)
	}
	scheduler.Start()
	time.Sleep(10 * time.Second)
}
