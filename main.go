// Copyright 2024 The Kolmogorov Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"math"

	"github.com/pointlander/compress"
	"github.com/pointlander/datum/iris"
)

func main() {
	datum, err := iris.Load()
	if err != nil {
		panic(err)
	}
	type Min struct {
		Label string
		Min   float64
	}
	mins := make([]Min, 150)
	for i := range mins {
		mins[i].Min = 1
	}
	vector := make([]byte, 8*64)
	for i := range datum.Bezdek {
		index := 0
		for _, value := range datum.Bezdek[i].Measures {
			bits := math.Float64bits(value)
			for i := 0; i < 64; i++ {
				vector[index] = byte(bits & 1)
				bits >>= 1
				index++
			}
		}
		for j, entry := range datum.Bezdek {
			if i == j {
				continue
			}
			index := 4 * 64
			for _, value := range entry.Measures {
				bits := math.Float64bits(value)
				for i := 0; i < 64; i++ {
					vector[index] = byte(bits & 1)
					bits >>= 1
					index++
				}
			}
			output := bytes.Buffer{}
			compress.Mark1Compress1(vector, &output)
			factor := float64(output.Len()) / float64(len(vector))
			if factor < mins[i].Min {
				mins[i].Min, mins[i].Label = factor, entry.Label
			}
		}
	}
	for i := range mins {
		fmt.Println(mins[i].Label, mins[i].Min)
	}
}
