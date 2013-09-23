# Benford's Law

[![Build Status](https://travis-ci.org/antonlindstrom/benfordslaw.png?branch=master)](https://travis-ci.org/antonlindstrom/benfordslaw)

This is a simple implementation of
[Benford's Law](http://en.wikipedia.org/wiki/Benford's_law).

### Usage

Use by importing this library and use the following to get JSON data as a string:

    benfordslaw.Process(dataset)

Where data set is an `[]int`, slice of integers.

A simple CSV loader is included in this repo.

### Dataset

Source: http://www.infochimps.com/datasets/flow-of-funds-accounts-financial-assets-and-liabilities-of-forei/

The dataset is modified to only corporate total financial assets in billion
dollars. The commas are removed, so the dataset is barely the same.

Results of this dataset (example output):

    {
      "Digits": [
        {"LeadingDigit":0,"Count":0,"Estimate":0,"Dataset":0},
        {"LeadingDigit":1,"Count":14,"Estimate":30.10299956639812,"Dataset":25.454545454545453},
        {"LeadingDigit":2,"Count":13,"Estimate":17.609125905568124,"Dataset":23.636363636363637},
        {"LeadingDigit":3,"Count":5,"Estimate":12.493873660829994,"Dataset":9.090909090909092},
        {"LeadingDigit":4,"Count":7,"Estimate":9.691001300805642,"Dataset":12.727272727272727},
        {"LeadingDigit":5,"Count":5,"Estimate":7.918124604762482,"Dataset":9.090909090909092},
        {"LeadingDigit":6,"Count":4,"Estimate":6.694678963061322,"Dataset":7.2727272727272725},
        {"LeadingDigit":7,"Count":3,"Estimate":5.799194697768673,"Dataset":5.454545454545454},
        {"LeadingDigit":8,"Count":2,"Estimate":5.115252244738129,"Dataset":3.6363636363636362},
        {"LeadingDigit":9,"Count":2,"Estimate":4.575749056067514,"Dataset":3.6363636363636362}
      ]
    }

## Build

Clone this repository into a structure like: `$GOPATH/src/github.com/antonlindstrom/benfordslaw`
and run:

    make get-deps
    make

This will create an example that uses `./example/example.go`.

## Test

    make test
