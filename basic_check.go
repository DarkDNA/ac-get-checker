package checker

import (
	"net/http"
)

func init() {
	DefineCheck(func(url string, r Results) {
		exists := r.BeginTest("Checking for .../desc.txt")

		res, err := MustHttp200(url + "/desc.txt")
		if err != nil {
			exists.Fail("Error getting desc.txt: %s", err.Error())

			return
		}

		res.Body.Close()

		exists.Success()
	})
}
