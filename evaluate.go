package evaluate

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Evaluator is a string that can be evaluated
type Evaluator string

var allowedExtra = map[byte]bool{
	'=': true,
	'>': true,
	'<': true,
	'!': true,
}

// Test will run evaluation on the value according to the rules in the Evaluator
func (t Evaluator) Test(value float64) Status {
	fields := strings.Fields(string(t))
	resp := StatusUnknown

	if string(t) == "" {
		return StatusOK
	}

	for _, f := range fields {
		if resp == StatusCritical {
			return resp
		}

		reader := bufio.NewReader(strings.NewReader(f))
		_, _ = reader.Discard(1)
		typB, _ := reader.ReadByte()
		typ := string(typB)

		op1 := readOp(reader)

		readMore := true
		numString, err := reader.ReadString(',')
		cutPoint := len(numString) - 1
		if err == io.EOF {
			readMore = false
			cutPoint = len(numString)
		}
		num, _ := strconv.ParseFloat(numString[:cutPoint], 64)

		ev1 := evaluate(op1, num, value)

		if readMore {
			op2 := readOp(reader)
			numString, _ = reader.ReadString(',')
			num2, _ := strconv.ParseFloat(numString, 64)
			ev2 := evaluate(op2, num2, value)

			if (op1 == ">" || op1 == ">=" || op1 == "=>") &&
				(op2 == "<" || op2 == "<=" || op2 == "=<") {
				if ev1 == true && ev2 == true {
					resp = getStatusFromType(typ)
				} else {
					resp = StatusOK
				}
			} else if (op1 == "<" || op2 == "<=" || op2 == "=<") &&
				(op2 == ">" || op1 == ">=" || op1 == "=>") {
				if ev1 || ev2 {
					resp = getStatusFromType(typ)
				} else {
					resp = StatusOK
				}
			} else {
				resp = StatusUnknown
			}
		} else {
			if !ev1 {
				if resp == StatusWarning {
					return resp
				}
				resp = StatusOK
			} else {
				resp = getStatusFromType(typ)
			}
		}
	}

	return resp
}

func getStatusFromType(typ string) Status {
	switch typ {
	case "w":
		return StatusWarning
	case "c":
		return StatusCritical
	default:
		return StatusUnknown
	}
}

func readOp(r *bufio.Reader) string {
	opFirstB, _ := r.ReadByte()
	opFirst := string(opFirstB)
	peek, _ := r.Peek(1)

	opExtra := ""
	if len(peek) > 0 && allowedExtra[peek[0]] {
		opExtra = string(peek[0])
		_, _ = r.Discard(1)
	}

	return opFirst + opExtra
}

func evaluate(op string, num float64, value float64) bool {
	switch op {
	case "<":
		return value < num
	case ">":
		return value > num
	case "=":
		return value == num
	case "!":
		fallthrough
	case "!=":
		return value != num
	case "==":
		return value == num
	case ">=":
		fallthrough
	case "=>":
		return value >= num
	case "<=":
		fallthrough
	case "=<":
		return value <= num
	default:
		return true
	}
}
