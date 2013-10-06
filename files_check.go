package checker

import (
	"strings"

	"net/http"
)

func check_files(url string, files []string, extension string, r Results) bool {
	for _, file := range files {
		if strings.HasSuffix(file, "/") {
			continue
		}

		file = file_name(file) + extension

		res, err := MustHttp200(url + "/" + file)
		if err != nil {
			r.Fail("Could not get %s/%s: %s", url, file, err.Error())

			return false
		}

		// TODO: Lua syntax checking.

		res.Body.Close()
	}

	return true
}

func file_name(spec string) string {
	if strings.Contains(spec, " => ") {
		return strings.Split(spec, " => ")[0]
	} else {
		return spec
	}
}
