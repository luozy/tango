// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tango

import (
	"strings"
)

func Prefix(prefix string, handler Handler) HandlerFunc {
	return func(ctx *Context) {
		if strings.HasPrefix(ctx.Req().URL.Path, prefix) {
			handler.Handle(ctx)
			return
		}

		ctx.Next()
	}
}
