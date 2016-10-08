package lean

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

func (this *leanClient) CloudQuery(cql string, pvalues ...interface{}) *Agent {
	if cql == "" {
		return nil
	}

	url := UrlBase + "/cloudQuery"
	request := gorequest.New()
	superAgent := request.Get(url)
	superAgent.QueryData.Add("cql", cql)

	if len(pvalues) != 0 {
		if jsonByte, err := json.Marshal(pvalues); nil != err {
			return nil
		} else {
			superAgent.QueryData.Add("pvalues", string(jsonByte))

		}
	}

	return &Agent{
		superAgent: superAgent,
		client:     this,
	}
}
