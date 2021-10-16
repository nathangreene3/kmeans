# *k*-Means

```go
go get github.com/nathangreene3/kmeans
```

## Description

The *k*-means algorithm is an unsupervised learning technique for grouping data into *k* classes. The method is useful in determining how many classes naturally appear in a data set and in determining which dimensions are most important in determining known classifications.

## Types

| Type | Description |
| :- | :- |
| **Model** | A model is a list of points representing the mean (center) of a class (cluster). Each mean is not necessarily a member of the data set it was trained upon. A model may be initialized-only and trained after initialization (updated). |
| **Labeled point** | A labeled point (l-point) extends a point adding an id and label. This may be used for training purposes or comparing labeled data to new, unlabeled data. |
| **Point** | A point is an *n*-tuple of real numbers. It is the basic type used to define and interact with a model. |

## Options

| Option | Description |
| :- | :- |
| **Training rounds** | The number of training rounds dictates how many initialization and training attempts are made. *k*-Means is inherently random and multiple initialization and training attempts is sometimes necessary. The model with the highest score will be returned. By default, one training round is applied. |
| **Initialization method** | The initialization method dictates how a model is initialized *before* training. |

| Method | Description |
| :- | :- |
| **Random** | The classic (naive, Lloyd's algorithm) method is random initialization. For small data sets, this is faster than plus-plus, but in some cases, a model will be returned that does not represent the data it was trained upon due to severe overlap, dimension bias, or other reasons beyond the scope or responsibility of *k*-means, which is an unsupervised method. That is, *k*-means does not train to match data to labels, it discovers labels. |
| **Plus-plus** | This improves upon random initialization by selecting representatives of the training data set that have the maximum distance from *any* mean. This attempts to prevent means from being initialized that are already close to each other. |
| **First-*k*** | The first *k* data points will be used as the means of the model. This method is fast, but exists only to allow the caller to initialize the model with means they know to be close to the expected means representing their data. Since there is no random behavior in this method, training more than once is not necessary. |

## Examples

Below is a taylored list of points that naturally form three clusters with known means.

```go
data := []Point{
    {1.0, 1.0},
    {2.0, 2.0},
    {3.0, 1.0},
    // Exp mean: (2, 1.333...)

    {1.0, 4.0},
    {1.0, 5.0},
    {2.0, 5.0},
    // Exp mean: (1.333..., 4.666...)

    {4.0, 3.0},
    {5.0, 2.0},
    {5.0, 3.0},
    {5.0, 4.0},
    // Exp mean: (4.75, 3.0)
}

mdl := New(3, data, SetTrainRounds(3), SetInitMethod(PlusPlus)) // [(1.333..., 4.666...) (2, 1.333...) (4.75, 3)]
```
