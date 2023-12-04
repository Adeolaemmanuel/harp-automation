package hasura

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var baseUrl = "https://harp-neon.onrender.com/v1/graphql"

func prepareHeader(authorization string, res *http.Request) Hasura {
	config := Hasura{
		BaseURL: baseUrl,
		Headers: Headers{
			X_Hasura_Admin_Secret: "harp_secret",
			Content_Type: "application/json",
			Authorization: authorization,
			X_Hasura_Role: "admin",
		},
	}

	if len(config.Headers.Authorization) > 0 {
		res.Header.Add("x-hasura-admin-secret", config.Headers.X_Hasura_Admin_Secret)
		res.Header.Add("x-hasura-role", config.Headers.X_Hasura_Role)
	} else {
		res.Header.Add("Authorization", config.Headers.Authorization)
	}

	return config
}

func (h Hasura) Execute(query string, variables any, authorization string) (HasuraResponse, error) {

	var data =  HasuraResponse{}

	body, err1 := json.Marshal(map[string] any {
		"query": query,
		"variables": variables,
	})

	if err1 != nil {
		return HasuraResponse{}, err1
	}

	req, err2 := http.NewRequest("POST", baseUrl, bytes.NewBuffer(body))

	prepareHeader(authorization, req)

	if err2 != nil {
		return HasuraResponse{}, err2
	}

	defer req.Body.Close()

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil {
		return HasuraResponse{}, err
	}

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return HasuraResponse{}, nil
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		json.Unmarshal(responseBody, &data)
	}

	json.Unmarshal(responseBody, &data)

	return data, nil
}

func (h Hasura) TransformData (res HasuraResponse) []map[string]any {

	data := []map[string]any{}

	for _, d := range res.Data {
		for _, val := range d {
			data = append(d, val)
		}
	}

	return data
}

func (h Hasura) ScheduleEvent(P map[string]any) (HasuraResponse, error) {
	config := HasuraEvent{}
	config.Arg.Webhook = baseUrl + "/webhook/event"

	data := HasuraResponse{}

	body, err := json.Marshal(config)

	if err != nil {
		return HasuraResponse{}, err
	}

	req, err1 := http.NewRequest("POST", baseUrl, bytes.NewBuffer(body))

	prepareHeader("", req)

	if err1 != nil {
		return HasuraResponse{}, err1
	}

	defer req.Body.Close()

	httpClient := http.Client{}

	resp, _ := httpClient.Do(req)

	responseBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		json.Unmarshal(responseBody, &data)
	}

	return data, err
}