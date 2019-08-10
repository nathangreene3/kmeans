# k-Means

```go
go get github.com/nathangreene3/kmeans
```

## Description

The k-means algorithm is an unsuperised learning technique for grouping data into k groups. The method is useful in determining how many groups naturally appear in a population (classification) and in determining which fields are most important in determining known classifications (feature extraction).

## Point

Point is an interface to implement k-means upon data. The provided implementation is FPoint, a slice of float64s.

```go
type Point interface {
    At(i int) float64
    CompareTo(p Point) int
    Copy() Point
    Dist(p Point) float64
    Len() int
    SqDist(p Point) float64
}
```
