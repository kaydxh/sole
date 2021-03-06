// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "go-atomicvalue -type tracer<*github.com/searKing/sole/app/core/tracing.Tracer>"; DO NOT EDIT.

package provider

import (
	"sync/atomic"

	"github.com/searKing/sole/pkg/micro/tracing"
)

func _() {
	// An "cannot convert tracer literal (type tracer) to type atomic.Value" compiler error signifies that the base type have changed.
	// Re-run the go-atomicvalue command to generate them again.
	_ = (atomic.Value)(tracer{})
}
func (m *tracer) Store(value *tracing.Tracer) {
	(*atomic.Value)(m).Store(value)
}

func (m *tracer) Load() *tracing.Tracer {
	value := (*atomic.Value)(m).Load()
	if value == nil {
		return nil
	}
	return value.(*tracing.Tracer)
}
