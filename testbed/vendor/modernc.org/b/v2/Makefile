# Copyright 2021 The B Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	edit editor

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi


editor:
	gofmt -l -s -w *.go
	go test -c -o /dev/null
