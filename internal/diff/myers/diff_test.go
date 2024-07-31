// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package myers_test

import (
	"github.com/pymeta/go-diff/internal/diff/difftest"
	"github.com/pymeta/go-diff/internal/diff/myers"
	"testing"
)

func TestDiff(t *testing.T) {
	difftest.DiffTest(t, myers.ComputeEdits)
}
