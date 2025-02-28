// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package experiments

import (
	"fmt"
)

// Experiment1
func Experiment1() {
	type T struct {
		T []*T
	}
	var prnt func(int, *T)
	prnt = func(depth int, a *T) {
		for i := 0; i < depth; i++ {
			fmt.Printf("_")
		}
		fmt.Println("T")
		for _, v := range a.T {
			prnt(depth+1, v)
		}
	}
	var cp func(*T) *T
	cp = func(a *T) *T {
		if a == nil {
			return nil
		}
		b := &T{}
		for _, v := range a.T {
			b.T = append(b.T, cp(v))
		}
		return b
	}
	pop := func(a *T) *T {
		if len(a.T) == 0 {
			return nil
		}
		b := a.T[len(a.T)-1]
		a.T = a.T[:len(a.T)-1]
		return b
	}
	push := func(a, b *T) {
		a.T = append(a.T, b)
	}
	apply := func(a, b *T) *T {
		a, b = cp(a), cp(b)
		expression := &T{
			T: []*T{b},
		}
		for _, v := range a.T {
			expression.T = append(expression.T, v)
		}
		todo := &T{
			T: []*T{expression},
		}
		for len(todo.T) > 0 {
			f := pop(todo)
			if len(f.T) < 3 {
				continue
			}
			push(todo, f)
			a, b, c := pop(f), pop(f), pop(f)
			if length := len(a.T); length == 0 {
				for _, v := range b.T {
					push(f, v)
				}
			} else if length == 1 {
				newPotRedex := &T{
					T: []*T{c},
				}
				for _, v := range b.T {
					newPotRedex.T = append(newPotRedex.T, v)
				}
				push(f, newPotRedex)
				push(f, c)
				for _, v := range a.T[0].T {
					push(f, v)
				}
				push(todo, newPotRedex)
			} else if length == 2 {
				if length := len(c.T); length == 0 {
					for _, v := range a.T[1].T {
						push(f, v)
					}
				} else if length == 1 {
					push(f, c.T[0])
					for _, v := range a.T[0].T {
						push(f, v)
					}
				} else if length == 2 {
					push(f, c.T[0])
					push(f, c.T[1])
					for _, v := range b.T {
						push(f, v)
					}
				}
			}
		}
		return expression
	}
	apply2 := func(a, b, c *T) *T {
		a, b, c = cp(a), cp(b), cp(c)
		expression := &T{
			T: []*T{c, b},
		}
		for _, v := range a.T {
			expression.T = append(expression.T, v)
		}
		todo := &T{
			T: []*T{expression},
		}
		for len(todo.T) > 0 {
			f := pop(todo)
			if len(f.T) < 3 {
				continue
			}
			push(todo, f)
			a, b, c := pop(f), pop(f), pop(f)
			if length := len(a.T); length == 0 {
				for _, v := range b.T {
					push(f, v)
				}
			} else if length == 1 {
				newPotRedex := &T{
					T: []*T{c},
				}
				for _, v := range b.T {
					newPotRedex.T = append(newPotRedex.T, v)
				}
				push(f, newPotRedex)
				push(f, c)
				for _, v := range a.T[0].T {
					push(f, v)
				}
				push(todo, newPotRedex)
			} else if length == 2 {
				if length := len(c.T); length == 0 {
					for _, v := range a.T[1].T {
						push(f, v)
					}
				} else if length == 1 {
					push(f, c.T[0])
					for _, v := range a.T[0].T {
						push(f, v)
					}
				} else if length == 2 {
					push(f, c.T[0])
					push(f, c.T[1])
					for _, v := range b.T {
						push(f, v)
					}
				}
			}
		}
		return expression
	}
	multi := func(x ...*T) *T {
		length := len(x)
		v := apply(x[length-2], x[length-1])
		for i := length - 3; i >= 0; i-- {
			v = apply(x[i], v)
		}
		return v
	}

	K := func() *T {
		return &T{
			T: []*T{
				&T{
					T: []*T{
						&T{},
					},
				},
			},
		}
	}
	I := func() *T {
		return &T{
			T: []*T{
				&T{
					T: []*T{
						&T{
							T: []*T{
								&T{},
							},
						},
					},
				},
				&T{
					T: []*T{
						&T{
							T: []*T{
								&T{},
							},
						},
					},
				},
			},
		}
	}
	d := func(x *T) *T {
		return &T{
			T: []*T{
				&T{
					T: []*T{
						x,
					},
				},
			},
		}
	}
	and := func() *T {
		i := I()
		k := K()
		k.T = append(k.T, i)
		kk := K()
		kk.T = append(kk.T, k)
		return d(kk)
	}

	_false := &T{}
	_true := &T{
		T: []*T{
			&T{},
		},
	}
	_not := &T{
		T: []*T{
			&T{},
			&T{
				T: []*T{
					&T{
						T: []*T{
							_false,
							&T{},
						},
					},
					_true,
				},
			},
		},
	}

	prnt(0, and())
	fmt.Println()
	prnt(0, apply2(and(), _false, _false))

	fmt.Println("not false = true")
	prnt(0, apply(_not, _false))
	fmt.Println("not true = false")
	prnt(0, apply(_not, _true))
	fmt.Println("not not false = false")
	prnt(0, multi(_not, _not, _false))
	fmt.Println("not not true = true")
	prnt(0, multi(_not, _not, _true))
	fmt.Println("identity false = false")
	prnt(0, apply(I(), _false))
	fmt.Println("identity true = true")
	prnt(0, apply(I(), _true))
}

// Experiment2
func Experiment2() {
	type T struct {
		L string
		T []*T
	}
	var number func(*T, *[]byte)
	number = func(a *T, results *[]byte) {
		if a != nil {
			*results = append(*results, 1)
			t := make([]*T, len(a.T))
			copy(t, a.T)
			if len(t) == 1 {
				t = append(t, nil)
			}
			for _, v := range t {
				number(v, results)
			}
		} else {
			*results = append(*results, 0)
		}
	}
	var prnt func(int, *T)
	prnt = func(depth int, a *T) {
		for i := 0; i < depth; i++ {
			fmt.Printf("_")
		}
		if a.L == "" {
			fmt.Println("T")
		} else {
			fmt.Println(a.L)
		}
		for _, v := range a.T {
			prnt(depth+1, v)
		}
	}
	var cp func(*T) *T
	cp = func(a *T) *T {
		if a == nil {
			return nil
		}
		b := &T{
			L: a.L,
		}
		for _, v := range a.T {
			b.T = append(b.T, cp(v))
		}
		return b
	}
	var equal func(a, b *T) bool
	equal = func(a, b *T) bool {
		if len(a.T) != len(b.T) {
			return false
		}
		e := true
		for i, v := range a.T {
			e = e && equal(v, b.T[i])
		}
		return e
	}
	hoist := func(a ...*T) *T {
		return &T{
			L: "h",
			T: a,
		}
	}
	const (
		X = 0
		Y = 1
		Z = 2
	)
	debug := false
	var apply func(*T) *T
	apply = func(a *T) *T {
		if debug {
			fmt.Println("apply")
		}
		for len(a.T) > 2 {
			if len(a.T) > 3 {
				panic(fmt.Errorf("len(a.T),%d > 3", len(a.T)))
			}
			switch len(a.T[X].T) {
			case 0:
				if debug {
					fmt.Println(0)
				}
				a = apply(a.T[Y])
			case 1:
				if debug {
					fmt.Println(1)
				}
				x := apply(a.T[X].T[0])
				y := apply(a.T[Y])
				z := apply(a.T[Z])
				if len(x.T) == 0 {
					x = hoist(x, z)
				} else if len(x.T) == 1 {
					x = hoist(x.T[0], z)
				} else if len(x.T) == 2 {
					x = hoist(x.T[0], x.T[1], z)
				} else {
					panic(fmt.Errorf("len(x.T),%d > 2", len(x.T)))
				}
				if len(y.T) == 0 {
					a = hoist(y, z, x)
				} else if len(y.T) == 1 {
					a = hoist(y.T[0], z, x)
				} else if len(y.T) == 2 {
					a = hoist(y.T[0], y.T[1], hoist(z, x))
					/*y = apply(hoist(y.T[0], y.T[1], z))
					if len(y.T) == 0 {
						a = hoist(y, x)
					} else if len(y.T) == 1 {
						a = hoist(y.T[0], x)
					} else if len(y.T) == 2 {
						a = hoist(y.T[0], y.T[1], x)
					} else {
						panic(fmt.Errorf("len(y.T),%d > 2", len(y.T)))
					}*/
				} else {
					panic(fmt.Errorf("len(y.T),%d > 2", len(y.T)))
				}
			case 2:
				if debug {
					fmt.Println(2)
				}
				w := apply(a.T[X].T[0])
				x := apply(a.T[X].T[1])
				z := apply(a.T[Z])
				if len(z.T) == 0 {
					a = hoist(z, w, x)
				} else if len(z.T) == 1 {
					a = hoist(z.T[0], w, x)
				} else if len(z.T) == 2 {
					a = hoist(z.T[0], z.T[1], hoist(w, x))
					/*z = apply(hoist(z.T[0], z.T[1], w))
					if len(z.T) == 0 {
						a = hoist(z, x)
					} else if len(z.T) == 1 {
						a = hoist(z.T[0], x)
					} else if len(z.T) == 2 {
						a = hoist(z.T[0], z.T[1], x)
					} else {
						panic(fmt.Errorf("len(z.T),%d > 2", len(z.T)))
					}*/
				} else {
					panic(fmt.Errorf("len(z.T),%d > 2", len(z.T)))
				}
			}
		}

		if debug {
			fmt.Println("return")
		}
		return a
	}

	K := func() *T {
		return &T{
			L: "K",
			T: []*T{
				&T{
					L: "K",
				},
			},
		}
	}
	I := func() *T {
		return &T{
			L: "I",
			T: []*T{
				&T{
					L: "I",
					T: []*T{
						&T{},
					},
				},
				&T{
					L: "I",
					T: []*T{
						&T{},
					},
				},
			},
		}
	}
	d := func(x *T) *T {
		return &T{
			L: "d",
			T: []*T{
				&T{
					L: "d",
					T: []*T{x},
				},
			},
		}
	}
	and := func() *T {
		return d(hoist(hoist(I(), K().T[0]), K().T[0]))
	}
	or := func() *T {
		x := d(hoist(K(), K()))
		x.T[0].T = []*T{I(), x.T[0].T[0]}
		return x
		//return hoist(hoist(I(), d(hoist(K(), K().T[0])).T[0]))
	}
	prnt(0, and())
	_true := K()
	_false := hoist(I(), K().T[0])
	/*_or := &T{
		L: "X",
		T: []*T{
			&T{
				L: "X",
				T: []*T{
					&T{
						L: "X",
					},
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
				},
			},
			&T{
				L: "X",
				T: []*T{
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
				},
			},
		},
	}*/
	_or := &T{
		L: "X",
		T: []*T{
			&T{
				L: "X",
				T: []*T{
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
				},
			},
			&T{
				L: "X",
				T: []*T{
					&T{
						L: "X",
						T: []*T{
							&T{},
						},
					},
					&T{
						L: "X",
					},
				},
			},
		},
	}
	fmt.Println("or")
	prnt(0, _or)

	fmt.Println("Kxy")
	a := apply(hoist(K().T[0], &T{L: "X"}, &T{L: "Y"}))
	prnt(0, a)

	fmt.Println("KIxy")
	a = apply(hoist(I(), K().T[0],
		hoist(
			&T{L: "X"},
			&T{L: "Y"},
		),
	))
	prnt(0, a)

	fmt.Println("and")
	cases := [][]*T{
		{_true, _true, _true},
		{_false, _true, _false},
		{_true, _false, _false},
		{_false, _false, _false},
	}
	outs := []string{
		"true and true != true",
		"false and true != false",
		"true and false != false",
		"false and false != false",
	}
	for i, v := range cases {
		a := apply(hoist(and(), v[0], v[1]))
		if !equal(a, v[2]) {
			fmt.Println(outs[i])
			prnt(0, a)
		}
	}

	prnt(0, or())
	fmt.Println("or")
	cases = [][]*T{
		{_true, _true, _true},
		{_false, _true, _true},
		{_true, _false, _true},
		{_false, _false, _false},
	}
	outs = []string{
		"true or true != true",
		"false or true != true",
		"true or false != true",
		"false or false != false",
	}
	for i, v := range cases {
		if i == 1 {
			debug = true
		} else {
			debug = false
		}
		a := apply(hoist(or(), v[0], v[1]))
		if !equal(a, v[2]) {
			fmt.Println(outs[i])
			prnt(0, a)
		}
	}

	fmt.Println("I")
	results := []byte{}
	number(I(), &results)
	fmt.Println(results)
	value := 0
	scale := 1
	for _, v := range results {
		if v == 1 {
			value += scale
		}
		scale *= 2
	}
	fmt.Println(value)
	fmt.Println("K")
	results = []byte{}
	number(K(), &results)
	fmt.Println(results)
	value = 0
	scale = 1
	for _, v := range results {
		if v == 1 {
			value += scale
		}
		scale *= 2
	}
	fmt.Println(value)
}
