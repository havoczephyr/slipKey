package app

//item represents the data structure of the .tsv files generated from SpeakFaster, with each element in item
// representing a column.
type item struct {
	TBegin  float64 // where an item starts.
	TEnd    float64 // where an item ends.
	Tier    string  // what category of item is it.
	Content string  // the value of an item.
}
