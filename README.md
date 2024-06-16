# Learn Go with tests
This repository is my journey of learning Golang by reading the book [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests).

## Scripts
### Run all test
#### Custom script
```shell
./test.sh # Run all test
```
```shell
./test.sh -w # Run all test and keep watching changes
```

#### Native way
```shell
go test ./... -v # Run all test in project
```

### Run all benchmarks
#### Custom script
```shell
./benchmark.sh # Run all benchmarks in project
```

#### Native way
```shell
go test -bench=. ./...  # Run all benchmarks in project
```

### Check for missing error handling
You will need to have installed `errcheck`. Install it with:
```shell
go install github.com/kisielk/errcheck@latest
```
Make sure you have this in your `.zshenv`:

```shell
# If you are using Go modules, the binaries are installed in the default location ($HOME/go/bin)
export PATH=$PATH:$HOME/go/bin
```

Then you can run this for checking unhandled errors in your whole project:
```shell
errcheck ./...
```
