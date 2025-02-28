// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))
}
