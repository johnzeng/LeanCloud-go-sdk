package lean

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func (this leanClient) CloudQuery(cql string, pvalues ...interface{}) *Agent {
	if cql == "" {
		return nil
	}

	url := UrlBase + "/cloudQuery"
	request := gorequest.New()
	superAgent := request.Get(url)
	superAgent.QueryData.Add("cql", cql)

	if len(pvalues) != 0 {
		superAgent.QueryData.Add("pvalues", fmt.Sprintf("%x", pvalues))
	}

	return &Agent{
		superAgent: superAgent,
		client:     this,
	}
}
