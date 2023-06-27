package structs

import "log"

type Match struct {
	UUID   string
	Horses []*Horse
	Winner *Horse
}

func (m *Match) Start() {
	stop := make(chan *Horse, 1)
	ch := make(chan *Horse, 1)
	for _, h := range m.Horses {
		go h.Running(ch, stop)
	}
	m.Winner = <-ch
	log.Printf("Match#%s, Winner is %s", m.UUID, m.Winner.Name)
}

func NewMatch(
	matchName string,
	numberOfContestor int,
	trackLength int,
) *Match {

	m := &Match{
		UUID:   matchName,
		Horses: nil,
		Winner: nil,
	}

	for i := 0; i < numberOfContestor; i++ {
		contestor := NewHorse(i, trackLength, nil, nil)
		m.Horses = append(m.Horses, contestor)
	}

	return m
}
