package landing

import (
	"errors"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Quotes  []Quote `json:"quotes"`
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
}

func FetchQuotes(c *http.Client) (*Result, error) {
	req, _ := http.NewRequest("GET", "https://goquotes-api.herokuapp.com/api/v1/random?count=5", nil)

	req.Header.Add("Accept", "application/json")

	resp, err := c.Do(req)

	if err != nil {
		return nil, errors.New("errored when sending request to the server")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var res Result
	json.Unmarshal(body, &res)
	return &res, json.Unmarshal(body, &res)
}

func (q *Quote) FormatText() string {
	if len(q.Text) > 150 {
		description := q.Text[0:150]

		return fmt.Sprint(description + "...")
	}
	return fmt.Sprint(q.Text)
}