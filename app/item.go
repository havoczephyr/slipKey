package app

//item represents the data structure of the .tsv files generated from SpeakFaster, with each element in item
// representing a column.
// TBegin and TEnd represent when an item started and ended.
// Tier is what type of item
// with Content being the value of the item.
type item struct {
	TBegin  float64
	TEnd    float64
	Tier    string
	Content string
}
