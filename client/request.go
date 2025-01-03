package client

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func MakeRequest(url *url.URL) (int, *CheckResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data CheckResponse
	err = decoder.Decode(&data)

	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, &data, nil
}
