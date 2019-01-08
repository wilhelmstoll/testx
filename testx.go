// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Pure go implementation of a test framework.

package testx

import (
	"fmt"
	"testing"
)

// Tolerance adjust the accuracy of the function.
const tolerance float64 = 0.00000001

// EqualFloat tests the expected with the actual value.
func EqualFloat(t *testing.T, expected, actual float64, msg string) {
	if !((expected-actual) < tolerance && (actual-expected) < tolerance) {
		t.Errorf(fmt.Sprintf("%s not equal.\n"+
			"expected: %g\n"+
			"actual  : %g", msg, expected, actual))
	}
}
