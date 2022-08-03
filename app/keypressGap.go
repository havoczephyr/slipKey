package app

type keypressGap struct {
	// Start             *Item
	// End               *Item
	MissingKeypresses []item
	TriggerTime       float64
	Duration          float64
}

func (g *keypressGap) Size() int {
	return len(g.MissingKeypresses)
}
