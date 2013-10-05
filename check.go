package checker

type Results interface {
	BeginTest(string, ...interface{}) Results

	Success()
	Fail(string, ...interface{})
	Warn(string, ...interface{})

	// Called when the task is merely a parent, and doesn't have a
	// success/failure condition on it's own.
	Done()

	BeginCheck()
	EndCheck()
}

type Check func(string, Results)

var checks []Check

func DefineCheck(c Check) {
	checks = append(checks, c)
}

func Run(url string, r Results) {
	for _, check := range checks {
		r.BeginCheck()
		check(url, r)
		r.EndCheck()
	}
}
