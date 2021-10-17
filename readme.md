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

## Example

```go
package main

import (
    "log"

    "github.com/nathangreene3/kmeans"
    "github.com/nathangreene3/kmeans/lpoint"
)

func main() {
    labeledData, err := lpoint.ReadJSONFile("./data.json")
    if err != nil {
        log.Fatal(err)
    }

    mdl := kmeans.New(
        len(lpoint.Labels(labeledData...)), // k
        lpoint.Points(labeledData...),
        kmeans.SetTrainRounds(3),              // Train three models keeping the highest scoring model
        kmeans.SetInitMethod(kmeans.PlusPlus), // Use the k-means++ method
    )

    log.Println(mdl)
}
```
