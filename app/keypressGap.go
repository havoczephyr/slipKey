package app

type keypressGap struct {
	//prefix keypress is the last successful keypress before the gap
	PrefixKeypress *item
	//suffix keypress is the first successful keypress after the gap
	SuffixKeypress *item
	//the cluster of keypresses that are missing after curation.
	MissingKeypresses []item
	//the time between prefixKeypress and MissingKeypresses[0]
	TriggerTime float64
	//the time between the first and last MissingKeypresses
	Duration float64
}

func (g *keypressGap) Size() int {
	return len(g.MissingKeypresses)
}
