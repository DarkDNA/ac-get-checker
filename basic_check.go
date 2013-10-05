package checker

import (
	"net/http"
)

func init() {
	DefineCheck(func(url string, r Results) {
		exists := r.BeginTest("Checking for .../desc.txt")

		res, err := http.Get(url + "/desc.txt")
		if err != nil {
			exists.Fail("Error getting desc.txt: %s", err.Error())

			return
		}

		if res.StatusCode != 200 {
			exists.Fail("Error getting desc.txt: HTTP Status %d", res.StatusCode)

			return
		}

		res.Body.Close()

		exists.Success()
	})
}
