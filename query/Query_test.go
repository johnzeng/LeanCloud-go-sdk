package query

import (
	"testing"
)

func TestCompare(t *testing.T) {
	query := Eq("key", 10)
	if query.String() != `{"key":{"$eq":10}}` {
		t.Error("query is wrong:" + query.String())
	}
}
