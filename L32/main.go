package main

type errKind int

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidHdrType
	// ...
)

type WaveError struct {
	kind  errKind
	value int
	err   error
}

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header"
	case cantReadHeader:
		return "cant-read-header"
	case invalidHdrType:
		return "invalid header type"
	}
}

func main() {

}
