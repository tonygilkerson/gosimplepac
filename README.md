# gosimplepac

This is a sample project with minimal main and one simple package, and of coures a test.

## Steps taken to create this project

Initialize the go module. The name of the module should correspond to your project such as `github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME`.

```sh
# In the root folder
go mod init github.com/tonygilkerson/gosimplepac
```

Create a `main.go` using `package main`
Create a folder to hold your packages, or as go calls them, modules. [Here](https://github.com/golang-standards/project-layout) is a good reference for how to name them. In this project we use `pkg`

```sh
mkdir pkg
touch pkg/simple.go
touch pkg/simple_test.go
```

In `simple.go` use `package pkg`, most of the time the package name will be the same as the parent folder. One exception is the tests.  Append `_test` to the end. Therefore, in `simple_test.go` use `package pkg`

```sh
echo "package pkg" >> pkg/simple.go
echo "package pkg_test" >> pkg/simple_test.go
```

Import the `pkg` package. To use `pkg` from `main.go` you need to first import it using `<module name>/<package name>`, for example:

```go
import "github.com/tonygilkerson/gosimplepac/pkg"
```

Create a test. To crate a test import `testing` in `simle_test.go` and write a test. Note, you will need to import `pkg` in order to test it.  

```go
import (
  "testing"
  "github.com/tonygilkerson/gosimplepac/pkg"
)
```

A simple test would start someting like the following. Note, the function name starts with `Test`, with a big `T`:

```go
func TestGetName(t *testing.T){}
```

Run

```sh
# Examples of running the test
go test ./...
go test -v ./...

# Run main
go run main.go 
```

 <!-- TODO - show build examples -->
