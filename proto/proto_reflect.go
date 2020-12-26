// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoreflect build tag disables use of fast-path methods.
// +build protoreflect

package proto

import (
	"github.com/xen0n/protobuf-gogogo/reflect/protoreflect"
	"github.com/xen0n/protobuf-gogogo/runtime/protoiface"
)

const hasProtoMethods = false

func protoMethods(m protoreflect.Message) *protoiface.Methods {
	return nil
}
