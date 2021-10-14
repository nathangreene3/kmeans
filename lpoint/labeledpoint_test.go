package lpoint

import (
	"testing"
)

func TestJSON(t *testing.T) {
	tests := [][]LPoint{
		{
			NewLPoint(1, "one", 1.0, 1.0, 1.0),
			NewLPoint(2, "two", 2.0, 2.0, 2.0),
			NewLPoint(3, "three", 3.0, 3.0, 3.0),
		},
	}

	for _, ps0 := range tests {
		s, err := JSON(ps0...)
		if err != nil {
			t.Error(err)
			continue
		}

		ps1, err := ParseJSON(s)
		if err != nil {
			t.Error(err)
			continue
		}

		if len(ps0) != len(ps1) {
			t.Errorf("\nexpected length %d\nreceived length %d\n", len(ps0), len(ps1))
			continue
		}

		for i := 0; i < len(ps0); i++ {
			if !ps0[i].Equals(ps1[i]) {
				t.Errorf("\nexpected %v\nreceived %v\n", ps0[i], ps1[i])
			}
		}
	}
}

func TestLabels(t *testing.T) {
	tests := []struct {
		filePath  string
		expLabels []string
	}{
		{
			filePath:  "../data/iris",
			expLabels: []string{"setosa", "versicolor", "virginica"},
		},
		{
			filePath:  "../data/radial_k4",
			expLabels: []string{"1", "2", "3", "4"},
		},
	}

	for _, test := range tests {
		data, err := ReadJSONFile(test.filePath)
		if err != nil {
			t.Fatal(err)
		}

		recLabels := Labels(data...)
		if len(test.expLabels) != len(recLabels) {
			t.Errorf("\nexpected %q\nreceived %q\n", test.expLabels, recLabels)
		}

		for i, expLabel := range test.expLabels {
			if expLabel != recLabels[i] {
				t.Errorf("\nexpected %q\nreceived %q\n", expLabel, recLabels[i])
			}
		}
	}
}
