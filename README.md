# Examples of testing golang

1. Run the math package simple unit tests:

```bash
go test
PASS
ok      github.com/heyjdp/go-testing-examples/math      0.002s
```

More verbose (custom messages):

```bash
go test -v
=== RUN   TestAdd
    math_test.go:11: Add(1, 3) PASSED. Expected 4, got 4
--- PASS: TestAdd (0.00s)
=== RUN   TestDivide
    math_test.go:21: Divide(1, 3) PASSED. Expected 0.000000, got 0.000000
--- PASS: TestDivide (0.00s)
PASS
ok      github.com/heyjdp/go-testing-examples/math      0.002s
```

Add coverage:

```bash
go test -v --cover
=== RUN   TestAdd
    math_test.go:11: Add(1, 3) PASSED. Expected 4, got 4
--- PASS: TestAdd (0.00s)
=== RUN   TestDivide
    math_test.go:21: Divide(1, 3) PASSED. Expected 0.000000, got 0.000000
--- PASS: TestDivide (0.00s)
PASS
coverage: 50.0% of statements
ok      github.com/heyjdp/go-testing-examples/math      0.002s
```

2. Run the server package mocked unit tests:

```bash
go test -v --cover
=== RUN   TestGetWeather
=== RUN   TestGetWeather/basic-request
    server_test.go:44: PASSED: expected &{Larnaca, CY scorchio}, got &{Larnaca, CY scorchio}
    server_test.go:50: PASSED: expected error <nil>, got <nil>
--- PASS: TestGetWeather (0.00s)
    --- PASS: TestGetWeather/basic-request (0.00s)
PASS
coverage: 77.8% of statements
ok      github.com/heyjdp/go-testing-examples/server    0.005s
```