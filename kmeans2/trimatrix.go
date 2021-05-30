package kmeans2

// triMatrix is a two dimensional array of values populating only the lower triangular region. The nxn matrix
// 	    [ d00         ]
// 	A = [ d10 d11     ]
// 	    [ d20 d21 d22 ]
// 	    [ ...         ]
// flattens as the vector
// 	v = [ d00 d10 d11 d20 d21 d22 ... ].
// The kth index of v, read left-to-right, top-to-bottom of A for the lower triangle entries only, is
// 	k = i + j*(j+1)/2.
type triMatrix []float64

// newDistMtx returns a matrix holding the distances from each point to every other point.
func newDistMtx(ps ...Point) triMatrix {
	return make(triMatrix, len(ps)*(len(ps)+1)>>1).update(ps...)
}

// dist returns the distance between points i and j, where i <= j;
func (mtx triMatrix) dist(i, j int) float64 {
	if j < i {
		i, j = j, i
	}

	return mtx[i+j*(j+1)>>1]
}

// update a triangular matrix.
func (mtx triMatrix) update(ps ...Point) triMatrix {
	if len(mtx) != len(ps)*(len(ps)+1)>>1 {
		panic("dimension mismatch")
	}

	for i, k := 0, 0; i < len(ps); i++ {
		for j := 0; j <= i; j++ {
			mtx[k] = ps[i].Dist(ps[j])
			k++
		}
	}

	return mtx
}
