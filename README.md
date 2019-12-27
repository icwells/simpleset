[![Build Status](https://travis-ci.org/icwells/simpleset.svg?branch=master)](https://travis-ci.org/icwells/simpleset)
[![GoDoc](https://godoc.org/github.com/icwells/simpleset?status.svg)](https://godoc.org/github.com/icwells/simpleset)

# simpleset  

## Simple python-style set for string, int, and float64 types.  

Copyright 2019 by Shawn Rupp

1. [Description](#Description)
2. [Installation](#Installation)  
3. [Usage](#Usage)  

## Description  
simpleset provides a straightforward set similar to a python set. It supports string, integer, and float64 
based sets. It does not support mixed-type sets, so, for example, you cannot add an integer to a string set.  

## Installation  

	go get github.com/icwells/simpleset   

### Usage  
To make a new empty set, either call the appropriate constructor for the type you want the set to hold. To convert a slice of 
strings, integers, or float64s, pass the slice to ToSet which will return a set of the appropriate type.  

```go
letters := []string{"a", "b", "c", "d", "e"}
set := simpleset.NewStringSet()
for _, i := range letters {
	// Can also pass the whole slice to Extend
	set.Add(i)
}
```

For documentation on additional methods, follow the GoDocs link above.  
