package main

// SortOpt indicates how a cluster is sorted.
type SortOpt int

const (
	// VarSort dictates points will be compared by variance to the cluster mean.
	VarSort SortOpt = 1 << iota

	// LexiSort dictates points will be compared by the default comparer, which is lexicographic.
	LexiSort
)
