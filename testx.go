// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Pure go implementation of a test framework.

package testx

import (
	"fmt"
	"math"
	"testing"
)

// EqualFloat tests the expected with the actual value.
func EqualFloat(t *testing.T, expected, actual float64, msg string) {
	maxmul := math.Max(1.0, math.Max(expected, actual))
	if math.Abs(expected-actual) > (math.SmallestNonzeroFloat64 * maxmul) {
		t.Errorf(fmt.Sprintf("%s not equal.\n"+
			"expected: %g\n"+
			"actual  : %g", msg, expected, actual))
	}
}
