package domain

type EvidenceSignal struct {
	Name  string
	Score float64
}

type Evidence struct {
	FactID  FactID
	Score   float64
	Rank    int
	Signals []EvidenceSignal
}

func (e Evidence) Signal(name string) (float64, bool) {
	for _, signal := range e.Signals {
		if signal.Name == name {
			return signal.Score, true
		}
	}

	return 0, false
}
