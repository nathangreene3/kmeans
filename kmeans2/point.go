package kmeans2

// Point ...
type Point struct {
	label string
	value Interface
}

// NewPoint ...
func NewPoint(label string, value Interface) Point {
	return Point{
		label: label,
		value: value,
	}
}

// Zero ...
func Zero() Point {
	return Point{}
}

// At ...
func (p *Point) At(i int) float64 {
	return p.At(i)
}

// Compare ...
func (p *Point) Compare(point Point) int {
	return p.value.Compare(point.value)
}

// Copy ...
func (p *Point) Copy() Point {
	return NewPoint(p.label, p.value)
}

// Dist ...
func (p *Point) Dist(q Point) float64 {
	return p.value.Dist(q.value)
}

// Label ...
func (p *Point) Label() string {
	return p.label
}

// Len ...
func (p *Point) Len() int {
	return p.value.Len()
}

// Relabel ...
func (p *Point) Relabel(label string) {
	p.label = label
}
