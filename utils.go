package checker

import (
	"fmt"

	"net/http"
)

func MustHttp200(url string) (http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		res.Body.Close()
		return nil, fmt.Errorf("HTTP Status %d", res.StatusCode)
	}

	return res, nil
}
