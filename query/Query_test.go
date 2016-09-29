package query

import (
	"testing"
)

func TestCompareEq(t *testing.T) {
	query := Eq("key", 10)
	if query.String() != `{"key":10}` {
		t.Error("query is wrong:" + query.String())
	}
}
