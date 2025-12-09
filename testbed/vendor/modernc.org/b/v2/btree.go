// Copyright 2021 The B Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package b implements the B+tree flavor of a BTree.
//
// This is a generic-enabled version of modernc.org/b.
//
// Concurrency considerations
//
// Tree.{Clear,Delete,Put,Set} mutate the tree. One can use eg. a
// sync.Mutex.Lock/Unlock (or sync.RWMutex.Lock/Unlock) to wrap those calls if
// they are to be invoked concurrently.
//
// Tree.{First,Get,Last,Len,Seek,SeekFirst,SekLast} read but do not mutate the
// tree.  One can use eg. a sync.RWMutex.RLock/RUnlock to wrap those calls if
// they are to be invoked concurrently with any of the tree mutating methods.
//
// Enumerator.{Next,Prev} mutate the enumerator and read but not mutate the
// tree.  One can use eg. a sync.RWMutex.RLock/RUnlock to wrap those calls if
// they are to be invoked concurrently with any of the tree mutating methods. A
// separate mutex for the enumerator, or the whole tree in a simplified
// variant, is necessary if the enumerator's Next/Prev methods per se are to
// be invoked concurrently.
package b // import "modernc.org/b/v2"

import (
	"io"
	"sync"
)

const (
	kx = 32
	kd = 32
)

type (
	// Cmp compares a and b. Return value is:
	//
	//	< 0 if a <  b
	//	  0 if a == b
	//	> 0 if a >  b
	//
	Cmp[K any] func(a, b K) int

	d[K, V any] struct { // data page
		c int
		d [2*kd + 1]de[K, V]
		n *d[K, V]
		p *d[K, V]
	}

	de[K, V any] struct { // d element
		k K
		v V
	}

	// Enumerator captures the state of enumerating a tree. It is returned
	// from the Seek* methods. The enumerator is aware of any mutations
	// made to the tree in the process of enumerating it and automatically
	// resumes the enumeration at the proper key, if possible.
	//
	// However, once an Enumerator returns io.EOF to signal "no more
	// items", it does no more attempt to "resync" on tree mutation(s).  In
	// other words, io.EOF from an Enumerator is "sticky" (idempotent).
	Enumerator[K, V any] struct {
		err error
		hit bool
		i   int
		k   K
		q   *d[K, V]
		t   *Tree[K, V]
		ver int64
	}

	// Tree is a B+tree.
	Tree[K, V any] struct {
		c     int
		cmp   Cmp[K]
		first *d[K, V]
		last  *d[K, V]
		r     any
		ver   int64
		dPool sync.Pool
		ePool sync.Pool
		xPool sync.Pool
	}

	xe[K any] struct { // x element
		ch any
		k  K
	}

	x[K, V any] struct { // index page
		c int
		x [2*kx + 2]xe[K]
	}
)

func (t *Tree[K, V]) clr(q any) {
	switch px := q.(type) {
	case *x[K, V]:
		for i := 0; i <= px.c; i++ { // Ch0 Sep0 ... Chn-1 Sepn-1 Chn
			t.clr(px.x[i].ch)
		}
		*px = x[K, V]{}
		t.xPool.Put(px)
	case *d[K, V]:
		*px = d[K, V]{}
		t.dPool.Put(px)
	}
}

// -------------------------------------------------------------------------- x

func (t *Tree[K, V]) newX(ch0 any) *x[K, V] {
	r := t.xPool.Get().(*x[K, V])
	r.x[0].ch = ch0
	return r
}

func (q *x[K, V]) extract(i int) {
	q.c--
	if i < q.c {
		copy(q.x[i:], q.x[i+1:q.c+1])
		q.x[q.c].ch = q.x[q.c+1].ch
		var zk K
		q.x[q.c].k = zk      // GC
		q.x[q.c+1] = xe[K]{} // GC
	}
}

func (q *x[K, V]) insert(i int, k K, ch any) *x[K, V] {
	c := q.c
	if i < c {
		q.x[c+1].ch = q.x[c].ch
		copy(q.x[i+2:], q.x[i+1:c])
		q.x[i+1].k = q.x[i].k
	}
	c++
	q.c = c
	q.x[i].k = k
	q.x[i+1].ch = ch
	return q
}

func (q *x[K, V]) siblings(i int) (l, r *d[K, V]) {
	if i >= 0 {
		if i > 0 {
			l = q.x[i-1].ch.(*d[K, V])
		}
		if i < q.c {
			r = q.x[i+1].ch.(*d[K, V])
		}
	}
	return
}

// -------------------------------------------------------------------------- d

func (l *d[K, V]) mvL(r *d[K, V], c int) {
	copy(l.d[l.c:], r.d[:c])
	copy(r.d[:], r.d[c:r.c])
	l.c += c
	r.c -= c
}

func (l *d[K, V]) mvR(r *d[K, V], c int) {
	copy(r.d[c:], r.d[:r.c])
	copy(r.d[:c], l.d[l.c-c:])
	r.c += c
	l.c -= c
}

// ----------------------------------------------------------------------- Tree

// TreeNew returns a newly created, empty Tree. The compare function is used
// for key collation.
func TreeNew[K, V any](cmp Cmp[K]) *Tree[K, V] {
	return &Tree[K, V]{
		cmp:   cmp,
		dPool: sync.Pool{New: func() any { return &d[K, V]{} }},
		ePool: sync.Pool{New: func() any { return &Enumerator[K, V]{} }},
		xPool: sync.Pool{New: func() any { return &x[K, V]{} }},
	}
}

// Clear removes all K/V pairs from the tree.
func (t *Tree[K, V]) Clear() {
	if t.r == nil {
		return
	}

	t.clr(t.r)
	t.c, t.first, t.last, t.r = 0, nil, nil, nil
	t.ver++
}

// Close performs Clear and zeroes *t.
func (t *Tree[K, V]) Close() {
	t.Clear()
	*t = Tree[K, V]{}
}

func (t *Tree[K, V]) cat(p *x[K, V], q, r *d[K, V], pi int) {
	t.ver++
	q.mvL(r, r.c)
	if r.n != nil {
		r.n.p = q
	} else {
		t.last = q
	}
	q.n = r.n
	*r = d[K, V]{}
	t.dPool.Put(r)
	if p.c > 1 {
		p.extract(pi)
		p.x[pi].ch = q
		return
	}

	switch px := t.r.(type) {
	case *x[K, V]:
		*px = x[K, V]{}
		t.xPool.Put(px)
	case *d[K, V]:
		*px = d[K, V]{}
		t.dPool.Put(px)
	}
	t.r = q
}

func (t *Tree[K, V]) catX(p, q, r *x[K, V], pi int) {
	t.ver++
	q.x[q.c].k = p.x[pi].k
	copy(q.x[q.c+1:], r.x[:r.c])
	q.c += r.c + 1
	q.x[q.c].ch = r.x[r.c].ch
	*r = x[K, V]{}
	t.xPool.Put(r)
	if p.c > 1 {
		p.c--
		pc := p.c
		if pi < pc {
			p.x[pi].k = p.x[pi+1].k
			copy(p.x[pi+1:], p.x[pi+2:pc+1])
			p.x[pc].ch = p.x[pc+1].ch
			var k K
			p.x[pc].k = k      // GC
			p.x[pc+1].ch = nil // GC
		}
		return
	}

	switch px := t.r.(type) {
	case *x[K, V]:
		*px = x[K, V]{}
		t.xPool.Put(px)
	case *d[K, V]:
		*px = d[K, V]{}
		t.dPool.Put(px)
	}
	t.r = q
}

// Delete removes the k's KV pair, if it exists, in which case Delete returns
// true.
func (t *Tree[K, V]) Delete(k K) (ok bool) {
	pi := -1
	var p *x[K, V]
	q := t.r
	if q == nil {
		return false
	}

	for {
		var i int
		i, ok = t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x[K, V]:
				if x.c < kx && q != t.r {
					x, i = t.underflowX(p, x, pi, i)
				}
				pi = i + 1
				p = x
				q = x.x[pi].ch
				continue
			case *d[K, V]:
				t.extract(x, i)
				if x.c >= kd {
					return true
				}

				if q != t.r {
					t.underflow(p, x, pi)
				} else if t.c == 0 {
					t.Clear()
				}
				return true
			}
		}

		switch x := q.(type) {
		case *x[K, V]:
			if x.c < kx && q != t.r {
				x, i = t.underflowX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d[K, V]:
			return false
		}
	}
}

func (t *Tree[K, V]) extract(q *d[K, V], i int) {
	t.ver++
	q.c--
	if i < q.c {
		copy(q.d[i:], q.d[i+1:q.c+1])
	}
	q.d[q.c] = de[K, V]{} // GC
	t.c--
}

func (t *Tree[K, V]) find(q any, k K) (i int, ok bool) {
	var mk K
	l := 0
	switch x := q.(type) {
	case *x[K, V]:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			mk = x.x[m].k
			switch cmp := t.cmp(k, mk); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	case *d[K, V]:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			mk = x.d[m].k
			switch cmp := t.cmp(k, mk); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	}
	return l, false
}

// First returns the first item of the tree in the key collating order, or
// (zero-value, zero-value) if the tree is empty.
func (t *Tree[K, V]) First() (k K, v V) {
	if q := t.first; q != nil {
		q := &q.d[0]
		k, v = q.k, q.v
	}
	return
}

// Get returns the value associated with k and true if it exists. Otherwise Get
// returns (zero-value, false).
func (t *Tree[K, V]) Get(k K) (v V, ok bool) {
	q := t.r
	if q == nil {
		return
	}

	for {
		var i int
		if i, ok = t.find(q, k); ok {
			switch x := q.(type) {
			case *x[K, V]:
				q = x.x[i+1].ch
				continue
			case *d[K, V]:
				return x.d[i].v, true
			}
		}
		switch x := q.(type) {
		case *x[K, V]:
			q = x.x[i].ch
		default:
			return
		}
	}
}

func (t *Tree[K, V]) insert(q *d[K, V], i int, k K, v V) *d[K, V] {
	t.ver++
	c := q.c
	if i < c {
		copy(q.d[i+1:], q.d[i:c])
	}
	c++
	q.c = c
	q.d[i].k, q.d[i].v = k, v
	t.c++
	return q
}

// Last returns the last item of the tree in the key collating order, or
// (zero-value, zero-value) if the tree is empty.
func (t *Tree[K, V]) Last() (k K, v V) {
	if q := t.last; q != nil {
		q := &q.d[q.c-1]
		k, v = q.k, q.v
	}
	return
}

// Len returns the number of items in the tree.
func (t *Tree[K, V]) Len() int {
	return t.c
}

func (t *Tree[K, V]) overflow(p *x[K, V], q *d[K, V], pi, i int, k K, v V) {
	t.ver++
	l, r := p.siblings(pi)
	if l != nil && l.c < 2*kd && i != 0 {
		s := (2*kd-l.c)/2 + 1 // half plus one
		if i < s {
			s = i
		}
		l.mvL(q, s)
		t.insert(q, i-s, k, v)
		p.x[pi-1].k = q.d[0].k
		return
	}

	if r != nil && r.c < 2*kd {
		if i < 2*kd {
			s := (2*kd-r.c)/2 + 1 // half plus one
			if 2*kd-i < s {
				s = 2*kd - i
			}
			q.mvR(r, s)
			t.insert(q, i, k, v)
			p.x[pi].k = r.d[0].k
			return
		}

		t.insert(r, 0, k, v)
		p.x[pi].k = k
		return
	}

	t.split(p, q, pi, i, k, v)
}

// Seek returns an Enumerator positioned on an item such that k >= item's key.
// ok reports if k == item.key The Enumerator's position is possibly after the
// last item in the tree.
func (t *Tree[K, V]) Seek(k K) (e *Enumerator[K, V], ok bool) {
	q := t.r
	if q == nil {
		e = t.ePoolGet(nil, false, 0, k, nil)
		return
	}

	for {
		var i int
		if i, ok = t.find(q, k); ok {
			switch x := q.(type) {
			case *x[K, V]:
				q = x.x[i+1].ch
				continue
			case *d[K, V]:
				return t.ePoolGet(nil, ok, i, k, x), true
			}
		}

		switch x := q.(type) {
		case *x[K, V]:
			q = x.x[i].ch
		case *d[K, V]:
			return t.ePoolGet(nil, ok, i, k, x), false
		}
	}
}

func (t *Tree[K, V]) ePoolGet(err error, hit bool, i int, k K, q *d[K, V]) *Enumerator[K, V] {
	x := t.ePool.Get().(*Enumerator[K, V])
	x.err, x.hit, x.i, x.k, x.q, x.t, x.ver = err, hit, i, k, q, t, t.ver
	return x
}

// SeekFirst returns an enumerator positioned on the first KV pair in the tree,
// if any. For an empty tree, err == io.EOF is returned and e will be nil.
func (t *Tree[K, V]) SeekFirst() (e *Enumerator[K, V], err error) {
	q := t.first
	if q == nil {
		return nil, io.EOF
	}

	return t.ePoolGet(nil, true, 0, q.d[0].k, q), nil
}

// SeekLast returns an enumerator positioned on the last KV pair in the tree,
// if any. For an empty tree, err == io.EOF is returned and e will be nil.
func (t *Tree[K, V]) SeekLast() (e *Enumerator[K, V], err error) {
	q := t.last
	if q == nil {
		return nil, io.EOF
	}

	return t.ePoolGet(nil, true, q.c-1, q.d[q.c-1].k, q), nil
}

// Set sets the value associated with k.
func (t *Tree[K, V]) Set(k K, v V) {
	pi := -1
	var p *x[K, V]
	q := t.r
	if q == nil {
		z := t.insert(t.dPool.Get().(*d[K, V]), 0, k, v)
		t.r, t.first, t.last = z, z, z
		return
	}

	for {
		i, ok := t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x[K, V]:
				i++
				if x.c > 2*kx {
					x, i = t.splitX(p, x, pi, i)
				}
				pi = i
				p = x
				q = x.x[i].ch
				continue
			case *d[K, V]:
				x.d[i].v = v
			}
			return
		}

		switch x := q.(type) {
		case *x[K, V]:
			if x.c > 2*kx {
				x, i = t.splitX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d[K, V]:
			switch {
			case x.c < 2*kd:
				t.insert(x, i, k, v)
			default:
				t.overflow(p, x, pi, i, k, v)
			}
			return
		}
	}
}

// Updater is the function of the Tree.Put upd argument.
type Updater[V any] func(oldV V, exists bool) (newV V, write bool)

// Put combines Get and Set in a more efficient way where the tree is walked
// only once. The upd(ater) receives (old-value, true) if a KV pair for k
// exists or (zero-value, false) otherwise. It can then return a (new-value,
// true) to create or overwrite the existing value in the KV pair, or
// (whatever, false) if it decides not to create or not to update the value of
// the KV pair.
//
// 	tree.Set(k, v) call conceptually equals calling
//
// 	tree.Put(k, func(oldv V, exists bool){ return v, true })
//
// modulo the differing return values.
func (t *Tree[K, V]) Put(k K, upd Updater[V]) (oldV V, written bool) {
	pi := -1
	var p *x[K, V]
	q := t.r
	var newV V
	if q == nil {
		// new KV pair in empty tree
		newV, written = upd(newV, false)
		if !written {
			return
		}

		z := t.insert(t.dPool.Get().(*d[K, V]), 0, k, newV)
		t.r, t.first, t.last = z, z, z
		return
	}

	for {
		i, ok := t.find(q, k)
		if ok {
			switch x := q.(type) {
			case *x[K, V]:
				i++
				if x.c > 2*kx {
					x, i = t.splitX(p, x, pi, i)
				}
				pi = i
				p = x
				q = x.x[i].ch
				continue
			case *d[K, V]:
				oldV = x.d[i].v
				newV, written = upd(oldV, true)
				if !written {
					return
				}

				x.d[i].v = newV
			}
			return
		}

		switch x := q.(type) {
		case *x[K, V]:
			if x.c > 2*kx {
				x, i = t.splitX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *d[K, V]: // new KV pair
			newV, written = upd(newV, false)
			if !written {
				return
			}

			switch {
			case x.c < 2*kd:
				t.insert(x, i, k, newV)
			default:
				t.overflow(p, x, pi, i, k, newV)
			}
			return
		}
	}
}

func (t *Tree[K, V]) split(p *x[K, V], q *d[K, V], pi, i int, k K, v V) {
	t.ver++
	r := t.dPool.Get().(*d[K, V])
	if q.n != nil {
		r.n = q.n
		r.n.p = r
	} else {
		t.last = r
	}
	q.n = r
	r.p = q

	copy(r.d[:], q.d[kd:2*kd])
	for i := range q.d[kd:] {
		q.d[kd+i] = de[K, V]{}
	}
	q.c = kd
	r.c = kd
	var done bool
	if i > kd {
		done = true
		t.insert(r, i-kd, k, v)
	}
	if pi >= 0 {
		p.insert(pi, r.d[0].k, r)
	} else {
		t.r = t.newX(q).insert(0, r.d[0].k, r)
	}
	if done {
		return
	}

	t.insert(q, i, k, v)
}

func (t *Tree[K, V]) splitX(p, q *x[K, V], pi int, i int) (*x[K, V], int) {
	t.ver++
	r := t.xPool.Get().(*x[K, V])
	copy(r.x[:], q.x[kx+1:])
	q.c = kx
	r.c = kx
	if pi >= 0 {
		p.insert(pi, q.x[kx].k, r)
	} else {
		t.r = t.newX(q).insert(0, q.x[kx].k, r)
	}

	var zk K
	q.x[kx].k = zk
	for i := range q.x[kx+1:] {
		q.x[kx+i+1] = xe[K]{}
	}
	if i > kx {
		q = r
		i -= kx + 1
	}

	return q, i
}

func (t *Tree[K, V]) underflow(p *x[K, V], q *d[K, V], pi int) {
	t.ver++
	l, r := p.siblings(pi)

	if l != nil && l.c+q.c >= 2*kd {
		l.mvR(q, 1)
		p.x[pi-1].k = q.d[0].k
		return
	}

	if r != nil && q.c+r.c >= 2*kd {
		q.mvL(r, 1)
		p.x[pi].k = r.d[0].k
		r.d[r.c] = de[K, V]{} // GC
		return
	}

	if l != nil {
		t.cat(p, l, q, pi-1)
		return
	}

	t.cat(p, q, r, pi)
}

func (t *Tree[K, V]) underflowX(p, q *x[K, V], pi int, i int) (*x[K, V], int) {
	t.ver++
	var l, r *x[K, V]

	if pi >= 0 {
		if pi > 0 {
			l = p.x[pi-1].ch.(*x[K, V])
		}
		if pi < p.c {
			r = p.x[pi+1].ch.(*x[K, V])
		}
	}

	if l != nil && l.c > kx {
		q.x[q.c+1].ch = q.x[q.c].ch
		copy(q.x[1:], q.x[:q.c])
		q.x[0].ch = l.x[l.c].ch
		q.x[0].k = p.x[pi-1].k
		q.c++
		i++
		l.c--
		p.x[pi-1].k = l.x[l.c].k
		return q, i
	}

	if r != nil && r.c > kx {
		q.x[q.c].k = p.x[pi].k
		q.c++
		q.x[q.c].ch = r.x[0].ch
		p.x[pi].k = r.x[0].k
		copy(r.x[:], r.x[1:r.c])
		r.c--
		rc := r.c
		r.x[rc].ch = r.x[rc+1].ch
		var zk K
		r.x[rc].k = zk
		r.x[rc+1].ch = nil
		return q, i
	}

	if l != nil {
		i += l.c + 1
		t.catX(p, l, q, pi-1)
		q = l
		return q, i
	}

	t.catX(p, q, r, pi)
	return q, i
}

// ----------------------------------------------------------------- Enumerator

// Close recycles e to a pool for possible later reuse. No references to e
// should exist or such references must not be used afterwards.
func (e *Enumerator[K, V]) Close() {
	t := e.t
	*e = Enumerator[K, V]{}
	t.ePool.Put(e)
}

// Next returns the currently enumerated item, if it exists and moves to the
// next item in the key collation order. If there is no item to return, err ==
// io.EOF is returned.
func (e *Enumerator[K, V]) Next() (k K, v V, err error) {
	if err = e.err; err != nil {
		return
	}

	if e.ver != e.t.ver {
		f, _ := e.t.Seek(e.k)
		*e = *f
		f.Close()
	}
	if e.q == nil {
		e.err, err = io.EOF, io.EOF
		return
	}

	if e.i >= e.q.c {
		if err = e.next(); err != nil {
			return
		}
	}

	i := e.q.d[e.i]
	k, v = i.k, i.v
	e.k, e.hit = k, true
	e.next()
	return
}

func (e *Enumerator[K, V]) next() error {
	if e.q == nil {
		e.err = io.EOF
		return io.EOF
	}

	switch {
	case e.i < e.q.c-1:
		e.i++
	default:
		if e.q, e.i = e.q.n, 0; e.q == nil {
			e.err = io.EOF
		}
	}
	return e.err
}

// Prev returns the currently enumerated item, if it exists and moves to the
// previous item in the key collation order. If there is no item to return, err
// == io.EOF is returned.
func (e *Enumerator[K, V]) Prev() (k K, v V, err error) {
	if err = e.err; err != nil {
		return
	}

	if e.ver != e.t.ver {
		f, _ := e.t.Seek(e.k)
		*e = *f
		f.Close()
	}
	if e.q == nil {
		e.err, err = io.EOF, io.EOF
		return
	}

	if !e.hit {
		// move to previous because Seek overshoots if there's no hit
		if err = e.prev(); err != nil {
			return
		}
	}

	if e.i >= e.q.c {
		if err = e.prev(); err != nil {
			return
		}
	}

	i := e.q.d[e.i]
	k, v = i.k, i.v
	e.k, e.hit = k, true
	e.prev()
	return
}

func (e *Enumerator[K, V]) prev() error {
	if e.q == nil {
		e.err = io.EOF
		return io.EOF
	}

	switch {
	case e.i > 0:
		e.i--
	default:
		if e.q = e.q.p; e.q == nil {
			e.err = io.EOF
			break
		}

		e.i = e.q.c - 1
	}
	return e.err
}
