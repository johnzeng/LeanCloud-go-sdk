package query

import (
	"encoding/json"
)

type Query struct {
	query map[string]interface{}
}

func (this Query) String() string {
	if str, err := json.Marshal(this.query); nil != err {
		return ""
	} else {
		return string(str)
	}
}

func And(query ...Query) *Query {
	var ret map[string]interface{}
	for i := range query {
		q := query[i].query
		for key, value := range q {
			ret[key] = value
		}
	}

	return &Query{ret}
}

func Or(query ...Query) *Query {
	ret := map[string](interface{}){
		"$or": query,
	}
	return &Query{ret}

}

func compare(key, cmp string, value interface{}) *Query {
	ret := map[string]interface{}{
		key: map[string]interface{}{
			cmp: value,
		},
	}
	return &Query{ret}
}

func Eq(key string, value interface{}) *Query {
	return &Query{map[string]interface{}{key: value}}
}

func Lt(key string, value interface{}) *Query {
	return compare(key, "$lt", value)
}

func Gt(key string, value interface{}) *Query {
	return compare(key, "$gt", value)
}

func Lte(key string, value interface{}) *Query {
	return compare(key, "$lte", value)
}

func Gte(key string, value interface{}) *Query {
	return compare(key, "$gte", value)
}

func In(key string, value []interface{}) *Query {
	return compare(key, "$in", value)
}

func Exists(key string, value bool) *Query {
	return compare(key, "$exists", value)
}

//????
func Select(key string, value interface{}) *Query {
	return nil
}

//????
func DontSelect(key string, value interface{}) *Query {
	return nil
}

func All(key string, value []interface{}) *Query {
	if len(value) == 1 {
		ret := map[string]interface{}{
			key: value[0],
		}
		return &Query{ret}

	} else {
		return compare(key, "$all", value)
	}
}
