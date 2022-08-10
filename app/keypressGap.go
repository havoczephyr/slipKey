package app

//PrefixKeypress is the last successful keypress before the gap. Uses a *item so the value can be nil if it is the first element.
// SuffixKeypress is the first successful keypress after the gap. Uses a *item so the value can be nil if it is the last element.
// MissingKeypresses is a []item representing the cluster of keypresses that are missing after curation.
// TriggerTime is the time between prefixKeypress and MissingKeypresses[0].
// and Duration is the time between the first and last MissingKeypresses.
type keypressGap struct {
	PrefixKeypress    *item
	SuffixKeypress    *item
	MissingKeypresses []item
	TriggerTime       float64
	Duration          float64
}

// func (g *keypressGap) Size() int {
// 	return len(g.MissingKeypresses)
// }
