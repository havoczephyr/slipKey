package app

type KeypressGap struct {
	// Start             *Item
	// End               *Item
	MissingKeypresses []Item
	TriggerTime       float64
	Duration          float64
}

func (g *KeypressGap) Size() int {
	return len(g.MissingKeypresses)
}
