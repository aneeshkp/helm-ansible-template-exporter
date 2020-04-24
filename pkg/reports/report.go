package reports

type Report struct {
	FileName string
	LineNumber int64
	OldLine string
	NewLine string
	pos Pos
}


// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

func (p Pos) Position() Pos {
	return p
}


