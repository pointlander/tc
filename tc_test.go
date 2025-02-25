// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte("t t (t (t t) (t (t t) (t t (t t (t t (t (t t) (t (t t) (t (t t) t))))))))")
	_, tt := Parse(input)
	if tt.String() != string(input) {
		t.Fatal("parsing failed")
	}
}

func TestLabel(t *testing.T) {
	tt := &T{
		N: 5,
		T: []*T{
			&T{
				N: 2,
				T: []*T{
					&T{N: 1},
					&T{
						N: 3,
						T: []*T{
							&T{N: 4},
						},
					},
				},
			},
			&T{
				N: 6,
				T: []*T{
					&T{N: 7},
				},
			},
		},
	}
	ttt := &T{
		T: []*T{
			&T{
				T: []*T{
					&T{},
					&T{
						T: []*T{
							&T{},
						},
					},
				},
			},
			&T{
				T: []*T{
					&T{},
				},
			},
		},
	}
	n := ttt.Label()
	if n != 8 {
		t.Fatal("wrong number of nodes")
	}
	var labels func(int, *T)
	labels = func(depth int, tt *T) {
		prefix := ""
		for i := 0; i < depth; i++ {
			prefix += " "
		}
		t.Logf("%s%d\n", prefix, tt.N)
		for _, v := range tt.T {
			labels(depth+1, v)
		}
	}
	labels(0, ttt)
	var compare func(a, b *T)
	compare = func(a, b *T) {
		if a.N != b.N {
			t.Fatal("node is mislabeled")
		}
		for i, v := range a.T {
			compare(v, b.T[i])
		}
	}
	compare(tt, ttt)
}

func TestTriangulation(t *testing.T) {
	tt := &T{
		T: []*T{
			&T{
				T: []*T{
					&T{},
					&T{
						T: []*T{
							&T{},
						},
					},
				},
			},
			&T{
				T: []*T{
					&T{},
				},
			},
		},
	}
	test := func(tt *T) {
		n := tt.Label()
		polygon := tt.Triangulation(n)
		for i, v := range polygon {
			t.Log(i, v)
		}
		ttt := ITriangulation(polygon)
		var compare func(a, b *T)
		compare = func(a, b *T) {
			if len(a.T) != len(b.T) {
				t.Fatal("node is mismatched")
			}
			for i, v := range a.T {
				compare(v, b.T[i])
			}
		}
		compare(tt, ttt)
	}
	test(tt)

	input := []byte("t t (t (t t) (t (t t) (t t (t t (t t (t (t t) (t (t t) (t (t t) t))))))))")
	_, tt = Parse(input)
	test(tt)
}
