package kmeans

// triMatrix is a two dimensional array of values populating only the
// lower triangular region. Given n data points, The nxn matrix
// 	    [ d00 d01 d02 ... ]
// 	A = [ d10 d11 d12 ... ]
// 	    [ d20 d21 d22 ... ]
// 	    [ ...         ... ]
// wastes space with dij = dji for all i,j in [0,n) and dii = 0 for all
// i in [0,n), so this triangular matrix is represented as the
// flattened vector v = [ d10 d20 d21 d30 ... ]. The kth index of v,
// read left-to-right, top-to-bottom of A for the lower triangle
// entries only, is k = i + (j^2-j) / 2.
type triMatrix []float64

// newTriMatrix returns a new lower-triangular matrix ready to be
// updated for n data points.
func newTriMatrix(n int) triMatrix {
	return make(triMatrix, (n*n-n)>>1)
}

// dist returns the distance between two points.
func (mtx triMatrix) dist(i, j int) float64 {
	switch {
	case i == j:
		return 0.0
	case j < i:
		return mtx[j+(i-1)*i>>1]
	default:
		return mtx[i+(j-1)*j>>1]
	}
}

// update a triangular matrix. Assumes the size of the data is the
// same as when the triangular matrix was initialized.
func (mtx triMatrix) update(data []Point) {
	for i, k := 0, 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			mtx[k] = data[i].Dist(data[j])
			k++
		}
	}
}
