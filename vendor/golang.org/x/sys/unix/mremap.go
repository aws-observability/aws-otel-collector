// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:build linux
// +build linux
=======
//go:build linux || netbsd
// +build linux netbsd
>>>>>>> main

package unix

import "unsafe"

type mremapMmapper struct {
	mmapper
	mremap func(oldaddr uintptr, oldlength uintptr, newlength uintptr, flags int, newaddr uintptr) (xaddr uintptr, err error)
}

<<<<<<< HEAD
func (m *mremapMmapper) Mremap(oldData []byte, newLength int, flags int) (data []byte, err error) {
	if newLength <= 0 || len(oldData) == 0 || len(oldData) != cap(oldData) || flags&MREMAP_FIXED != 0 {
=======
var mapper = &mremapMmapper{
	mmapper: mmapper{
		active: make(map[*byte][]byte),
		mmap:   mmap,
		munmap: munmap,
	},
	mremap: mremap,
}

func (m *mremapMmapper) Mremap(oldData []byte, newLength int, flags int) (data []byte, err error) {
	if newLength <= 0 || len(oldData) == 0 || len(oldData) != cap(oldData) || flags&mremapFixed != 0 {
>>>>>>> main
		return nil, EINVAL
	}

	pOld := &oldData[cap(oldData)-1]
	m.Lock()
	defer m.Unlock()
	bOld := m.active[pOld]
	if bOld == nil || &bOld[0] != &oldData[0] {
		return nil, EINVAL
	}
	newAddr, errno := m.mremap(uintptr(unsafe.Pointer(&bOld[0])), uintptr(len(bOld)), uintptr(newLength), flags, 0)
	if errno != nil {
		return nil, errno
	}
	bNew := unsafe.Slice((*byte)(unsafe.Pointer(newAddr)), newLength)
	pNew := &bNew[cap(bNew)-1]
<<<<<<< HEAD
	if flags&MREMAP_DONTUNMAP == 0 {
=======
	if flags&mremapDontunmap == 0 {
>>>>>>> main
		delete(m.active, pOld)
	}
	m.active[pNew] = bNew
	return bNew, nil
}
<<<<<<< HEAD
=======

func Mremap(oldData []byte, newLength int, flags int) (data []byte, err error) {
	return mapper.Mremap(oldData, newLength, flags)
}
>>>>>>> main
