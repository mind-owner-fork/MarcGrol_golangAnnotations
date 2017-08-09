// +build !ci

package structExample

//go:generate golangAnnotations -input-dir .

// @JsonEnum( base = "Color" )
type Color int

const (
	ColorRed Color = iota
	ColorGreen
	ColorBlue
)

// @JsonEnum( )
type Word int

const (
	Aap Word = iota
	Noot
	Mies
)
