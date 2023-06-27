package structs

import (
	"fmt"
	"log"
)

type Horse struct {
	Name            string
	Run             int
	Target          int
	RunningCallback func()
	FinishCallback  func()
}

func (h *Horse) Running(ch chan<- *Horse, stop chan<- *Horse) {
	for {
		if len(stop) == cap(stop) {
			return
		} else {
			h.Run += 1
			log.Printf("%s: %d", h.Name, h.Run)
			if h.Target == h.Run {
				stop <- h
				ch <- h
			}
		}
	}
}

func NewHorse(id int, target int, runningCallback func(), finishCallback func()) *Horse {
	name := fmt.Sprintf("horse#%d", id)
	p := Horse{
		Name:            name,
		Target:          target,
		Run:             0,
		RunningCallback: runningCallback,
		FinishCallback:  finishCallback,
	}
	return &p
}
