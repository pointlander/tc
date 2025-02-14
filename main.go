// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	type T struct {
		T []*T
	}
	var prnt func(int, *T)
	prnt = func(depth int, a *T) {
		for i := 0; i < depth; i++ {
			fmt.Printf(" ")
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
	fmt.Println("not false = true")
	prnt(0, apply(_not, _false))
	fmt.Println("not true = false")
	prnt(0, apply(_not, _true))
	fmt.Println("not not = identity")
	_identity := apply(_not, _not)
	prnt(0, _identity)
	fmt.Println("identity false = false")
	prnt(0, apply(_identity, _false))
	fmt.Println("identity true = true")
	prnt(0, apply(_identity, _true))
}
