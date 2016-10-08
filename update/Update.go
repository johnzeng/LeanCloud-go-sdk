package update

import (
	"encoding/json"
	"errors"
)

type Update struct {
	update map[string]interface{}
}

func (t Update) MarshalJSON() ([]byte, error) {
	str := t.String()
	if "" == str {
		return nil, errors.New("can not marshal json")
	}
	return []byte(str), nil
}

func (t *Update) UnmarshalJSON(i []byte) error {
	//do your serializing here
	if err := json.Unmarshal(i, &t.update); err != nil {
		return err
	} else {
		return nil
	}
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

func Decrement(key string, value interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":   "Decrement",
			"amount": value,
		},
	}
	return &Update{ret}
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

//accept only LeanPointer as targets
func AddRelation(key string, targets ...interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":    "AddRelation",
			"objects": targets,
		},
	}
	return &Update{ret}
}

//accept only LeanPointer as targets
func RemoveRelation(key string, targets ...interface{}) *Update {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			"__op":    "RemoveRelation",
			"objects": targets,
		},
	}
	return &Update{ret}
}
