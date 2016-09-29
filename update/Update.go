package update

import (
	"encoding/json"
)

type Update struct {
	update map[string]interface{}
}

func (this *Update) String() string {
	if str, err := json.Marshal(this.update); nil != err {
		return ""
	} else {
		return string(str)
	}
}

func (this *Update) And(u *Update) *Update {
	for key, value := range u.update {
		this.update[key] = value
	}
	return this
}

func Increment(key string, value interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":   "Increment",
			"amount": value,
		},
	}
	return &Update{ret}
}

func BitOr(key string, value int64) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":  "BitOr",
			"value": value,
		},
	}
	return &Update{ret}
}

func BitAnd(key string, value int64) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":  "BitAnd",
			"value": value,
		},
	}
	return &Update{ret}
}

func BitXor(key string, value int64) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":  "BitXor",
			"value": value,
		},
	}
	return &Update{ret}
}

func AddToArray(key string, value ...interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":    "Add",
			"objects": value,
		},
	}
	return &Update{ret}
}

func AddUniqueToArray(key string, value ...interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":    "AddUnique",
			"objects": value,
		},
	}
	return &Update{ret}
}

func RemoveFromArray(key string, value ...interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":    "Remove",
			"objects": value,
		},
	}
	return &Update{ret}
}
