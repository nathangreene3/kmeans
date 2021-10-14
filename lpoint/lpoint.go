package lpoint

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/nathangreene3/kmeans"
)

// LPoint labels a point.
type LPoint struct {
	ID    int          `json:"id"`
	Label string       `json:"label"`
	Point kmeans.Point `json:"point"`
}

// NewLPoint ...
func NewLPoint(id int, label string, values ...float64) LPoint {
	lp := LPoint{
		ID:    id,
		Label: label,
		Point: append(make(kmeans.Point, 0, len(values)), values...),
	}

	return lp
}

// Copy a labeled point.
func (lp LPoint) Copy() LPoint {
	cpy := LPoint{
		ID:    lp.ID,
		Label: lp.Label,
		Point: lp.Point.Copy(),
	}

	return cpy
}

// Dims returns the number of dimensions in a list of points. If the
// number of dimensions is inconsistent in the set of points, zero is
// returned.
func Dims(lps ...LPoint) int {
	if err := Validate(lps...); err != nil {
		return 0
	}

	return len(lps[0].Point)
}

// Equals determines if two labeled points are equal.
func (lp LPoint) Equals(lq LPoint) bool {
	return lp.ID == lq.ID && lp.Label == lq.Label && lp.Point.Equals(lq.Point)
}

// JSON returns the points marshalled into a json-encoded string.
func JSON(lps ...LPoint) (string, error) {
	var sb strings.Builder
	if err := json.NewEncoder(&sb).Encode(lps); err != nil {
		return "", err
	}

	return sb.String(), nil
}

// Labels returns a list of the distinct labels found in the given set
// of labeled points. The list will be sorted.
func Labels(lps ...LPoint) []string {
	// var labels []string
	// for i := 0; i < len(lps); i++ {
	// 	var (
	// 		label = lps[i].Label
	// 		index = sort.SearchStrings(labels, label)
	// 	)

	// 	if index == len(labels) {
	// 		labels = append(labels, label)
	// 	} else if labels[index] != label {
	// 		labels = append(labels[:index+1], labels[index:]...)
	// 		labels[index] = label
	// 	}
	// }

	// TODO: Compare these methods
	labelFreq := LabelFreq(lps...)
	labels := make([]string, 0, len(labelFreq))
	for label := range labelFreq {
		labels = append(labels, label)
	}

	sort.Strings(labels)
	return labels
}

// ParseJSON parses labeled points from a json-encoded string.
func ParseJSON(s string) ([]LPoint, error) {
	var lps []LPoint
	if err := json.NewDecoder(strings.NewReader(s)).Decode(&lps); err != nil {
		return nil, err
	}

	return lps, nil
}

// Points ...
func Points(lps ...LPoint) []kmeans.Point {
	ps := make([]kmeans.Point, 0, len(lps))
	for i := 0; i < len(lps); i++ {
		ps = append(ps, append(make(kmeans.Point, 0, len(lps[i].Point)), lps[i].Point...))
	}

	return ps
}

// ReadCSVFile ...
func ReadCSVFile(file string, header bool) ([]LPoint, error) {
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

	ps := make([]LPoint, 0, n)
	for ; i < len(records); i++ {
		p := make(kmeans.Point, 0, len(records[i]))
		for j := 0; j < len(records[i])-1; j++ {
			pj, err := strconv.ParseFloat(records[i][j], 64)
			if err != nil {
				return nil, err
			}

			p = append(p, pj)
		}

		ps = append(ps, LPoint{ID: i, Label: records[i][len(records[i])-1], Point: p})
	}

	return ps, nil
}

// ReadJSONFile returns labeled points parsed from a json file. If
// the file extension .json is not provided in the file name, it will
// be appended before reading the file.
func ReadJSONFile(file string) ([]LPoint, error) {
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

// LabelFreq ...
func LabelFreq(lps ...LPoint) map[string]int {
	labelFreq := make(map[string]int)
	for i := 0; i < len(lps); i++ {
		labelFreq[lps[i].Label]++
	}

	return labelFreq
}

// String ...
func (lp LPoint) String() string {
	return fmt.Sprintf("{%d %s %v}", lp.ID, lp.Label, lp.Point)
}

// Validate ...
func Validate(lps ...LPoint) error {
	for i := 1; i < len(lps); i++ {
		if len(lps[0].Point) != len(lps[i].Point) {
			return errors.New("dimension mismatch")
		}
	}

	return nil
}

// WriteCSVFile writes a list of points to a csv file. The header will
// only be written if provided. The last column will be the label.
func WriteCSVFile(file string, header []string, lps ...LPoint) error {
	if err := Validate(lps...); err != nil {
		return err
	}

	var (
		buf = bytes.NewBuffer(make([]byte, 0))
		w   = csv.NewWriter(buf)
	)

	if len(header) != 0 {
		if len(header) != Dims(lps...) {
			return errors.New("dimension mismatch")
		}

		for j := 0; j < len(header); j++ {
			header[j] = strings.ToTitle(header[j])
		}

		if err := w.Write(append(header, "Label")); err != nil {
			return err
		}
	}

	for i := 0; i < len(lps); i++ {
		record := make([]string, 0, len(lps[i].Point))
		for j := 0; j < len(lps[i].Point); j++ {
			record = append(record, strconv.FormatFloat(lps[i].Point[j], 'f', -1, 64))
		}

		if err := w.Write(append(record, lps[i].Label)); err != nil {
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return os.WriteFile(file, buf.Bytes(), os.ModePerm)
}

// WriteJSONFile writes labeled points to a json file. If the file
// extension .json is not provided in the file name, it will be
// appended before writing the file.
func WriteJSONFile(file string, lps ...LPoint) error {
	s, err := JSON(lps...)
	if err != nil {
		return err
	}

	file = filepath.Clean(file)
	if !strings.EqualFold(filepath.Ext(file), ".json") {
		file += ".json"
	}

	return os.WriteFile(file, []byte(s), os.ModePerm)
}
