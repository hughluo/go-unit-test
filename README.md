# Write Unit Test in Go

Examples for https://wexort.medium.com/write-unit-test-in-go-309747c7453

## Generate and illustrate test coverage
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```