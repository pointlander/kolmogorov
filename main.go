// Copyright 2024 The Kolmogorov Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/pointlander/datum/iris"
)

func main() {
	datum, err := iris.Load()
	if err != nil {
		panic(err)
	}
	_ = datum
}
