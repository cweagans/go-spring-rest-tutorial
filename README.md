# go-spring-rest-tutorial

This is a port of the [Spring REST tutorial](https://github.com/spring-guides/tut-rest) to Go. This is intended to be used for demonstration only.

## Get started:

Install go through Homebrew or similar (`brew install go`)

```bash
# Pull down dependencies. These are locked to specific versions in go.sum (including transitive dependencies)
go mod vendor

# Run tests (really minimal stuff just for demo - see employee_test.go)
go test

# Build the binary
go build

# Run the binary
./go-spring-rest-tutorial
```

## Other info:

* Entry point is `func main` in `main.go`.
* All HTTP related functions are in `http.go`
* `employee.go` is roughly analagous to `Employee.java` from the tutorial (https://github.com/spring-guides/tut-rest/blob/main/rest/src/main/java/payroll/Employee.java). You could add getters and setters for the struct fields, but it's not really necessary, so I chose to omit that (struct fields that begin with a capital letter are public; lowercase is private -- this is a convention that Go uses a lot. Functions are only callable from outside of a package if they start with a capital letter, for example).
* All of the normal engineering tools that you'd expect are pre-bundled with Go:
    * `go test .` runs tests for the current package. If we had subpackages, you could run the tests recursively for all packages with `go test ./...`
    * `go vet` is a static analysis tool that checks not only for bad syntax, but for things that might need some attentino like `Printf` calls whose arguments do not align with the format string.
    * `go test -coverprofile=cover.out ./...` will run tests and emit a code coverage profile, which you can then convert into an HTML document with `go tool cover -html=cover.out -o cover.html`.
    * `go test -race ./...` will run tests with race detection on (which is super useful for code that relies on the Go concurrency tools)
    * `go tool pprof` is a profiling tool that's built in. It requires some setup, but it's not difficult and it's well documented. There's some info here if you want to see how to do it: https://www.jajaldoang.com/post/profiling-go-app-with-pprof/
