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
func NewCluster(points Points) Cluster {
	c := Cluster{
		points: points.Copy(),
		size:   len(points),
	}

	c.points.Sort()
	c.Update()
	return c
}

// Append ...
func (c *Cluster) Append(point Point) {
	c.points = append(c.points, point)
	c.points.Sort()
	c.size++
	c.Update()
}

// Clusters ...
type Clusters []Cluster

// Coalesce ensures clusters containing equivalent points are placed in one
// cluster. For each cluster i, the points in each cluster j are compared and
// equivalent points are moved to cluster i.
func (cs Clusters) Coalesce() {
	cs.SortAll(SortByVariance)

	m := len(cs)
	for i := 0; i < m; i++ {
		n := len(cs[i])
		for j := 0; j < n; j++ {
			if j+1 < n && cs[i][j].Compare(cs[i][j+1]) == 0 {
				// Points j and j+1 are equal, so keep iterating until the last
				// equal point is found.
				continue
			}

			for k := i + 1; k < m; k++ {
				for b := 0; b < len(cs[k]); b++ {
					c := cs[i][j].Compare(cs[k][b])
					if c == 0 {
						cs[i], cs[k] = Transfer(b, cs[i], cs[k])
						continue
					}

					if 0 < c {
						break
					}
				}
			}
		}

		cs[i].Sort(SortByVariance)
	}
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

// Update ...
func (c *Cluster) Update() {
	// Update mean
	switch c.size {
	case 0:
		c.mean = Zero()
	case 1:
		c.mean = c.points[0].Copy()
	default:
		c.mean = c.points[0]
		c.variance = c.points.Variance(c.points[0])
		for _, p := range c.points[1:] {
			if v := c.points.Variance(p); v < c.variance {
				c.mean = p.Copy()
				c.variance = v
			}
		}
	}

	c.mean.label = ""

	// Update variance
	c.variance = c.points.Variance(c.mean)
}

// Sort a cluster by a sorting option.
func (c *Cluster) Sort(sortOpt SortOption) {
	switch sortOpt {
	case SortByVariance:
		if c.mean.Compare(Zero()) != 0 {
			sort.Slice(c.points, func(i, j int) bool { return c.mean.Dist(c.points[i]) < c.mean.Dist(c.points[j]) })
		}
	case SortByDimension:
		c.points.Sort()
	}
}

// Points returns a set of points from a cluster.
func (c *Cluster) Points() Points {
	return c.points.Copy()
}

// Remove ...
func (c *Cluster) Remove(i int) Point {
	p := c.points[i]
	if i+1 < c.size {
		c.points = append(c.points[:i], c.points[i+1:]...)
	} else {
		c.points = c.points[:i]
	}

	c.size--
	c.points.Sort()
	c.Update()
	return p
}

// transfer ...
func (c *Cluster) transfer(i int, cluster *Cluster) {
	cluster.Append(c.Remove(i))
}

// Variance ...
func (c *Cluster) Variance() float64 {
	return c.variance
}

// VarianceTo returns the Variance of the cluster to a point (usually the mean).
func (c *Cluster) VarianceTo(point Point) float64 {
	return c.points.Variance(point)
}
