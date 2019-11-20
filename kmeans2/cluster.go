package kmeans2

import "sort"

// Cluster is a set of points.
type Cluster struct {
	points   Points
	mean     Point
	variance float64
	size     int
}

// NewCluster ...
func NewCluster(points ...Point) Cluster {
	n := len(points)
	c := Cluster{points: make(Points, n), size: n}

	copy(c.points, points)
	c.points.Sort(false)
	c.Update()
	return c
}

// Append a point to a cluster.
func (c *Cluster) Append(point Point) {
	c.points = append(c.points, point)
	c.points.Sort(true)
	c.size++
	c.Update()
}

// Compare returns -1, 0, or 1 indicating cluster c precedes, is equal to, or
// follows another cluster.
func (c *Cluster) Compare(cluster Cluster) int {
	comparison := c.mean.Compare(cluster.mean)
	switch {
	case comparison != 0:
		return comparison
	case c.variance < cluster.variance:
		return -1
	case cluster.variance < c.variance:
		return 1
	case c.size < cluster.size:
		return -1
	case cluster.size < c.size:
		return 1
	default:
		return 0
	}
}

// Copy a cluster.
func (c *Cluster) Copy() Cluster {
	return Cluster{
		points:   c.points.Copy(),
		mean:     c.mean.Copy(),
		variance: c.variance,
		size:     c.size,
	}
}

// Update cluster size, mean, and variance.
func (c *Cluster) Update() {
	c.size = len(c.points)
	switch c.size {
	case 0:
		c.mean = NewPoint("", nil)
	case 1:
		c.mean = c.points[0].Copy()
	default:
		c.mean = c.points[0].Copy()
		c.variance = c.points.Variance(c.points[0])
		for _, p := range c.points[1:] {
			if v := c.points.Variance(p); v < c.variance {
				c.mean = p.Copy()
				c.variance = v
			}
		}
	}

	c.mean.label = ""
	c.variance = c.points.Variance(c.mean)
}

// Sort a cluster by a sorting option.
func (c *Cluster) Sort(sortOpt SortOption, stable bool) {
	switch sortOpt {
	case SortByVariance:
		if !c.mean.IsZero() {
			if stable {
				sort.SliceStable(c.points, func(i, j int) bool { return c.mean.Dist(c.points[i]) < c.mean.Dist(c.points[j]) })
			} else {
				sort.Slice(c.points, func(i, j int) bool { return c.mean.Dist(c.points[i]) < c.mean.Dist(c.points[j]) })
			}
		}
	case SortByDimension:
		c.points.Sort(stable)
	}
}

// Points returns a set of points from a cluster.
func (c *Cluster) Points() Points {
	return c.points.Copy()
}

// Remove the ith point from a cluster.
func (c *Cluster) Remove(i int) Point {
	p := c.points[i]
	if i+1 < c.size {
		c.points = append(c.points[:i], c.points[i+1:]...)
	} else {
		c.points = c.points[:i]
	}

	c.size--
	c.Update()
	return p
}

// transfer ith point to another cluster.
func (c *Cluster) transfer(i int, destination *Cluster) {
	destination.Append(c.Remove(i))
}

// Variance of the cluster with respect to the mean.
func (c *Cluster) Variance() float64 {
	return c.variance
}

// VarianceTo returns the Variance of the cluster to a point (usually the mean).
func (c *Cluster) VarianceTo(point Point) float64 {
	return c.points.Variance(point)
}
