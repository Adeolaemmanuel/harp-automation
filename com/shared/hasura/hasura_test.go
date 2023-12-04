package hasura

import (
	"encoding/json"
	"testing"

	"harp-automation.com/com/shared/graphql"
)

var queries = graphql.Queries{}
func TestExecute(t *testing.T) {
	conn := Hasura{}

	req, err := conn.Execute(queries.GetHarpCampaign(), map[string]any{}, "")

	if err != nil {
		t.Error(err)
	}

	con := []Campaign{}

	res, _ := json.Marshal(conn.TransformData(req))

	err1 := json.Unmarshal(res, &con)

	if err1 != nil {
		t.Error(err1)
	}

	t.Log(con)

}