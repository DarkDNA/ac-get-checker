package checker

import (
	"strings"

	"bufio"

	"net/http"
)

var recmmended = []string{"Description"}
var deprecated = []string{"Version"}

func init() {
	DefineCheck(func(url string, r Results) {
		exists := r.BeginTest("Checking for .../packages.list")

		res, err := MustHttp200(url + "/packages.list")
		if err != nil {
			exists.Fail("Failed to get package list: %s", err.Error())

			return
		}

		defer res.Body.Close()

		exists.Success()

		valid := r.BeginTest("Checking Packages...")

		pkgs := 0

		reader := bufio.NewReader(res.Body)

		line, _, err := reader.ReadLine()

		for err == nil {
			pkgs++
			parts := strings.Split(string(line), "::")

			var pkg Results

			if len(parts) == 3 {
				pkg = valid.BeginTest("Checking package %s", parts[0])

				check_package(url+"/"+parts[0], pkg)
			} else {
				valid.Fail("Invalid package on line %d", pkgs)

				return
			}

			line, _, err = reader.ReadLine()
		}

		valid.Done()
	})
}

func check_package(url string, r Results) {
	res, err := MustHttp200(url)
	if err != nil {
		r.Fail("Failed to get details.pkg: %s", err.Error())

		return
	}

	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)

	line, _, err := reader.ReadLine()

	warned_blank := false
	warned_invalid := false

	directives := make(map[string][]string)

	for err == nil {
		str_line := string(line)

		if str_line == "" && !warned_blank {
			r.Warn("Previous verisons of ac-get didn't work with blank lines in the details.pkg")
			warned_blank = true
		}

		parts := strings.SplitN(str_line, ": ", 2)

		if len(parts) != 2 && !warned_invalid {
			r.Warn("Previous versions of ac-get will fail to parse this.")
			warned_invalid = true
		} else {
			if _, ok := directives[parts[0]]; !ok {
				directives[parts[0]] = []string{}
			}

			directives[parts[0]] = append(directives[parts[0]], parts[1])
		}

		line, _, err = reader.ReadLine()
	}

	for _, rec := range recmmended {
		if _, ok := directives[rec]; !ok {
			r.Warn("Missing recommended directive: %s", rec)
		}
	}

	for _, dep := range deprecated {
		if _, ok := directives[dep]; ok {
			r.Warn("Contains deprecated directive %s", dep)
		}
	}

	if check_files(url+"/lib", directives["Library"], ".lua", r) &&
		check_files(url+"/bin", directives["Executable"], ".lua", r) &&
		check_files(url+"/cfg", directives["Config"], "", r) &&
		check_files(url+"/startup", directives["Startup"], ".lua", r) &&
		check_files(url+"/docs", directives["Documentation"], ".txt", r) {
		r.Success()
	}
}
