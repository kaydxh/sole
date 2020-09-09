// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "go-atomicvalue -type prometheusMetricsManager<*github.com/searKing/sole/app/shared/micro/prometheus.MetricsManager>"; DO NOT EDIT.

package provider

import (
	"sync/atomic"

	"github.com/searKing/sole/pkg/micro/prometheus"
)

func _() {
	// An "cannot convert prometheusMetricsManager literal (type prometheusMetricsManager) to type atomic.Value" compiler error signifies that the base type have changed.
	// Re-run the go-atomicvalue command to generate them again.
	_ = (atomic.Value)(prometheusMetricsManager{})
}
func (m *prometheusMetricsManager) Store(value *prometheus.MetricsManager) {
	(*atomic.Value)(m).Store(value)
}

func (m *prometheusMetricsManager) Load() *prometheus.MetricsManager {
	value := (*atomic.Value)(m).Load()
	if value == nil {
		return nil
	}
	return value.(*prometheus.MetricsManager)
}