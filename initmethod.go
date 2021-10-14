package kmeans

// InitMethod defines how each mean is initialized before training.
type InitMethod uint

const (
	// Random indicates a model will be initialized with random data
	// points.
	Random InitMethod = 1 + iota

	// PlusPlus indicates a model will be initialized with the
	// k-means++ method.
	PlusPlus

	// FirstK indicates a model will be initialized with the first k
	// data points.
	FirstK
)

// String describes an initialization method.
func (mthd InitMethod) String() string {
	switch mthd {
	case Random:
		return "random"
	case PlusPlus:
		return "k-means++"
	case FirstK:
		return "first-k"
	default:
		return "invalid"
	}
}
