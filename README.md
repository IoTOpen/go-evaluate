# Go-evaluate

This lib provides functions to evaluate values based on rules defined as a
string. The result is either OK, WARNING, CRITICAL or UNKNOWN. This type of
evaluation is typically used in monitoring tools like Nagios or OP5 Monitor.

## Installing

```shell
go get github.com/IoTOpen/go-evaluate
```

There are currently no dependencies for go-evaluate.

## Usage

```go
package main

import (
	"github.com/IoTOpen/go-evaluate"
	"log"
)

func main() {
	ev := evaluate.Evaluator("-c<10")
	status := ev.Test(22.3)
	log.Println(status)
}
```

## Defining rules

Rules can be created and combined using the following operators:

| Operator    | Description                              |
|-------------|------------------------------------------|
| `-w`        | The following is conditions for Warning  |
| `-c`        | The following is conditions for Critical |
| `<`         | Less than                                |
| `>`         | Bigger than                              |
| `!`  / `!=` | Not equal to                             |
| `=` / `==`  | Equal                                    |
| `<=`        | Less than or equal to                    |
| `=>`        | Equal to or more than                    |
| `,`         | Range combinator                         |

All malformed rules results in status Unknown.

## Examples

| Thresholds        | Value | Result   | Description                                            |
|-------------------|-------|----------|--------------------------------------------------------|
| `-w<10`           | 5     | Warning  | Warning if less than 10                                |
| `-w<10`           | 15    | OK       | Warning if less than 10                                |
| `-w!10`           | 15    | Warning  | Warning if not 10                                      |
| `-w!10`           | 10    | OK       | Warning if not 10                                      |
| `-c<10 -w<20`     | 15    | Warning  | Critical if under 10, Warning if below 20              |
| `-c<10 -w<20`     | 5     | Critical | Critical if under 10, Warning if below 20              |
| `-w>15,<20 -c<10` | 8     | Critical | Warning if over 15, but under 20. Critical if under 10 |

## License

See [LICENSE](LICENSE).
