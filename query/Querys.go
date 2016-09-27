package query

type Query struct {
	order string
	query string
	limit int
	skip  int
}

func And(query ...Query) Query {

}

func Or(query ...Query) Query {

}

func Eq(key, value string) Query {

}

func Lt(key, value string) Query {

}

func Gt(key, value string) Query {

}
