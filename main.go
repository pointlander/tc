// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
)

var (
	// FlagOriginal original mode
	FlagOriginal = flag.Bool("original", false, "original mode")
)

// Original mode
func Original() {
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

func main() {
	flag.Parse()

	if *FlagOriginal {
		Original()
		return
	}

	type T struct {
		Label string
		T     []*T
	}
	var prnt func(int, *T)
	prnt = func(depth int, a *T) {
		for i := 0; i < depth; i++ {
			fmt.Printf("_")
		}
		if a.Label == "" {
			fmt.Println("T")
		} else {
			fmt.Println(a.Label)
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
			Label: a.Label,
		}
		for _, v := range a.T {
			b.T = append(b.T, cp(v))
		}
		return b
	}
	hoist := func(a ...*T) *T {
		return &T{
			T: a,
		}
	}
	var apply func(*T) *T
	apply = func(a *T) *T {
		for len(a.T) > 2 {
			if len(a.T[0].T) == 0 {
				a = a.T[1]
			} else if len(a.T[0].T) == 1 {
				x := a.T[0].T[0]
				y := a.T[1]
				z := a.T[2]
				x = hoist(x, z)
				if len(y.T) == 1 {
					a = hoist(y.T[0], z, x)
				} else {
					a = hoist(y.T[0], y.T[1], hoist(z, x))
				}
			} else if len(a.T[0].T[0].T) == 2 {
				w := a.T[0].T[0].T[0]
				x := a.T[0].T[0].T[1]
				z := a.T[2]
				if len(z.T) == 1 {
					a = hoist(z.T[0], w, x)
				} else {
					a = hoist(z.T[0], z.T[1], hoist(w, x))
				}
			}
		}
		x := &T{
			Label: a.Label,
		}
		for _, v := range a.T {
			x.T = append(x.T, apply(v))
		}
		return x
	}

	K := func() *T {
		return &T{
			T: []*T{
				&T{},
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
		x = cp(x)
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
		k = hoist(k, i)
		kk := K()
		kk = hoist(kk, k)
		return d(kk)
	}
	prnt(0, and())
	_true := K()
	_false := K()
	_false.T = append(_false.T, I())

	fmt.Println("KIxy")
	k := K()
	k = hoist(k, I().T[0],
		hoist(
			&T{Label: "X"},
			&T{Label: "Y"},
		),
	)
	k = apply(k)
	prnt(0, k)

	fmt.Println("and")
	a := and()
	_ = _true
	a.T = append(a.T, _true, _true)
	a = apply(a)
	prnt(0, a)
}
