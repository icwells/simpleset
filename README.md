[![Build Status](https://travis-ci.com/icwells/go-tools.svg?branch=master)](https://travis-ci.com/icwells/go-tools)

# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

Copyright 2019 by Shawn Rupp

1. [Description](#Description)
2. [Installation](#Installation)  
3. [Set](#Set)  

## Description  

## Installation  
	go get github.com/icwells/go-tools/dataframe  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

### Set  
The set struct is a simple python-style set for strings.  

#### strarray.NewSet() Set
Initializes new set.  

#### strarray.ToSet(s []string) Set  
Converts slice of strings to set.  

#### set.Length()  
Returns length of set.  

#### set.Add(value string)  
Adds string value to set.  

#### set.Extend(v []string)  
Adds all elements of slice to set.  

#### set.Pop(v string)  
Removes v from set.  

#### set.InSet(value string)  
Reurns true if value is in the set. Returns false if it is not.  

#### set.ToSlice() []string  
Returns set as a sorted string slice.
