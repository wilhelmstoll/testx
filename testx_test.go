// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package testx

import (
	"testing"
)

func TestEqualFloat(t *testing.T) {
	resultFromJsProject := 33.997467251559286
	resultFromGoProject := 33.997467248417244

	EqualFloat(t, resultFromJsProject, resultFromGoProject, "Float")
}
