// +build unit

package imunit_test

import "testing"

func Test(t *testing.T) {
	t.Error("found error in 'Unit Test'")
}
