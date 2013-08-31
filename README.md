# Benford's Law

This is a simple implementation of
[Benford's Law](http://en.wikipedia.org/wiki/Benford's_law).

### Dataset

Source: http://www.infochimps.com/datasets/flow-of-funds-accounts-financial-assets-and-liabilities-of-forei/

The dataset is modified to only corporate total financial assets in billion
dollars. The commas are removed, so the dataset is barely the same.

Results of this dataset (example output):

    2013/08/31 19:45:14 {num: 1, estimate: 30.103000, dataset: 25.454545, diff: -4.648454}
    2013/08/31 19:45:14 {num: 2, estimate: 17.609126, dataset: 23.636364, diff: 6.027238}
    2013/08/31 19:45:14 {num: 3, estimate: 12.493874, dataset: 9.090909, diff: -3.402965}
    2013/08/31 19:45:14 {num: 4, estimate: 9.691001, dataset: 12.727273, diff: 3.036271}
    2013/08/31 19:45:14 {num: 5, estimate: 7.918125, dataset: 9.090909, diff: 1.172784}
    2013/08/31 19:45:14 {num: 6, estimate: 6.694679, dataset: 7.272727, diff: 0.578048}
    2013/08/31 19:45:14 {num: 7, estimate: 5.799195, dataset: 5.454545, diff: -0.344649}
    2013/08/31 19:45:14 {num: 8, estimate: 5.115252, dataset: 3.636364, diff: -1.478889}
    2013/08/31 19:45:14 {num: 9, estimate: 4.575749, dataset: 3.636364, diff: -0.939385}

## Build

    mkdir bin
    go build -o bin/benford benford.go
