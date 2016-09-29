package update

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	u := Increment("number", 2)
	if u.String() != `{"number":{"__op":"Increment","amount":2}}` {
		t.Error("value is not right :" + u.String())
	}
}
