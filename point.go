package kmeans

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Point is real-valued n-dimensional number.
type Point []float64

// Add a point q to p. That is, p += q.
func (p Point) Add(q Point) {
	if len(p) != len(q) {
		panic(errDims)
	}

	for i := 0; i < len(p); i++ {
		p[i] += q[i]
	}
}

// Add returns the addition of a list of points.
func Add(ps ...Point) Point {
	p := ps[0].Copy()
	for i := 1; i < len(ps); i++ {
		p.Add(ps[i])
	}

	return p
}

// Compare returns -1 if p < q, 1 if p > q, or 0 if p = q.
func (p Point) Compare(q Point) int {
	if len(p) != len(q) {
		panic(errDims)
	}

	for i := 0; i < len(p); i++ {
		switch {
		case p[i] < q[i]:
			return -1
		case q[i] < p[i]:
			return 1
		}
	}

	return 0
}

// Copy a point.
func (p Point) Copy() Point {
	return append(make(Point, 0, len(p)), p...)
}

// Near determines if two points differ by a given tolerance. That is,
// if |pi-qi|<=tol for each dimension i, then p is near q. If tol < 0,
// then this function behaves the same as if tol = 0. Furthermore,
// tol = 0 determines if p and q are equal.
func (p Point) Near(q Point, tol float64) bool {
	if len(p) != len(q) {
		return false
	}

	for i := 0; i < len(p); i++ {
		if tol < math.Abs(p[i]-q[i]) {
			return false
		}
	}

	return true
}

// Dims returns the dimensions of the set of points.
func Dims(ps ...Point) int {
	if err := Validate(ps...); err != nil {
		panic(err)
	}

	return len(ps[0])
}

// Dist returns the Euclidean distance between two points.
func (p Point) Dist(q Point) float64 {
	return math.Sqrt(p.SqDist(q))
}

// Dot returns the dot product of p and q.
func (p Point) Dot(q Point) float64 {
	var r float64
	for i := 0; i < len(p); i++ {
		r += p[i] * q[i]
	}

	return r
}

// Equals determines if two points are equal.
func (p Point) Equals(q Point) bool {
	return p.Near(q, 0.0)
}

// JSON returns a json-encoded string representation of a list of
// points.
func JSON(ps ...Point) (string, error) {
	var sb strings.Builder
	if err := json.NewEncoder(&sb).Encode(ps); err != nil {
		return "", err
	}

	return sb.String(), nil
}

// Mag returns the magnitude of a point. This is, the Euclidean
// distance from the origin.
func (p Point) Mag() float64 {
	return math.Sqrt(p.Dot(p))
}

// ParseJSON returns a list of points parsed from a json-encoded string.
func ParseJSON(s string) ([]Point, error) {
	var ps []Point
	if err := json.NewDecoder(strings.NewReader(s)).Decode(&ps); err != nil {
		return nil, err
	}

	return ps, nil
}

// ReadCSVFile returns a list of points read from a csv file formatted
// without a header. Each line should be formatted as
// 0.0, 0.0, ..., 0.0, though the choice of delimiter is not limited
// to a comma.
func ReadCSVFile(file string, header bool) ([]Point, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	records, err := csv.NewReader(bytes.NewReader(b)).ReadAll()
	if err != nil {
		return nil, err
	}

	var (
		i int
		n = len(records)
	)

	if header {
		i++
		n--
	}

	ps := make([]Point, 0, n)
	for ; i < len(records); i++ {
		p := make(Point, 0, len(records[i]))
		for j := 0; j < len(records[i]); j++ {
			pj, err := strconv.ParseFloat(records[i][j], 64)
			if err != nil {
				return nil, err
			}

			p = append(p, pj)
		}

		ps = append(ps, p)
	}

	return ps, nil
}

// ReadJSONFile returns a list of points read from a json file.
func ReadJSONFile(file string) ([]Point, error) {
	file = filepath.Clean(file)
	if !strings.EqualFold(filepath.Ext(file), ".json") {
		file += ".json"
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return ParseJSON(string(b))
}

// ScalMult mulitplies a point p by a. That is, p *= a.
func (p Point) ScalMult(a float64) {
	for i := 0; i < len(p); i++ {
		p[i] *= a
	}
}

// ScalMult returns the scalar multiple a of a point p.
func ScalMult(p Point, a float64) Point {
	q := p.Copy()
	q.ScalMult(a)
	return q
}

// Normalize a point. That is, p /= |p|.
func (p Point) Normalize() {
	if r := p.Dot(p); r != 0 {
		p.ScalMult(1.0 / math.Sqrt(r))
	}
}

// Normalize a list of points.
func Normalize(ps ...Point) {
	for i := 0; i < len(ps); i++ {
		ps[i].Normalize()
	}
}

// Norm returns the norm of a point.
func Norm(p Point) Point {
	q := p.Copy()
	q.Normalize()
	return q
}

// Norms returns the norms of a list of points.
func Norms(ps ...Point) []Point {
	norms := make([]Point, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		norms = append(norms, Norm(ps[i]))
	}

	return norms
}

// SqDist returns the squared Euclidean distance between two points.
func (p Point) SqDist(q Point) float64 {
	if len(p) != len(q) {
		panic(errDims)
	}

	var sd float64 // sd = sum((pi-qi)^2)
	for i := 0; i < len(p); i++ {
		d := p[i] - q[i]
		sd += d * d
	}

	return sd
}

// String returns a representation of a point formatted as
// (0.0, ..., 0.0). A nil point is represented as the empty string.
func (p Point) String() string {
	var sb strings.Builder
	if len(p) != 0 {
		sb.WriteString("(" + strconv.FormatFloat(p[0], 'f', -1, 64))
		for i := 1; i < len(p); i++ {
			sb.WriteString(", " + strconv.FormatFloat(p[i], 'f', -1, 64))
		}

		sb.WriteByte(')')
	}

	return sb.String()
}

// Validate a list of points.
func Validate(data ...Point) error {
	for i := 1; i < len(data); i++ {
		if len(data[0]) != len(data[i]) {
			return errors.New(errDims)
		}
	}

	return nil
}

// WriteCSVFile writes a list of points to a csv file. The header will
// only be written if provided.
func WriteCSVFile(file string, header []string, ps ...Point) error {
	var (
		buf = bytes.NewBuffer(make([]byte, 0))
		w   = csv.NewWriter(buf)
	)

	if len(header) != 0 {
		if err := w.Write(header); err != nil {
			return err
		}
	}

	for i := 0; i < len(ps); i++ {
		record := make([]string, 0, len(ps[i]))
		for j := 0; j < len(ps[i]); j++ {
			record = append(record, strconv.FormatFloat(ps[i][j], 'f', -1, 64))
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return os.WriteFile(file, buf.Bytes(), os.ModePerm)
}

// WriteJSONFile writes a list of points to a json file.
func WriteJSONFile(file string, ps ...Point) error {
	s, err := JSON(ps...)
	if err != nil {
		return err
	}

	file = filepath.Clean(file)
	if !strings.EqualFold(filepath.Ext(file), ".json") {
		file += ".json"
	}

	return os.WriteFile(file, []byte(s), os.ModePerm)
}
