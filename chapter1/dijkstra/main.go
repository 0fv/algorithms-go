package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	cal := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"

	log.Println(calulate(cal))

}
func calulate(cal string) float64 {
	c := strings.Split(cal, " ")
	calContainer := calContainer{}
	for _, char := range c {
		switch char {
		case "(":
		case "+", "-", "*", "/", "sqrt":
			calContainer.pushOps(char)
		case ")":
			v := calContainer.popVal()
			op := calContainer.popOps()
			switch op {
			case "+":
				v = calContainer.popVal() + v
			case "-":
				v = calContainer.popVal() - v
			case "*":
				v = calContainer.popVal() * v
			case "/":
				v = calContainer.popVal() / v
			case "sqrt":
				v = math.Sqrt(v)
			}
			calContainer.pushVal(fmt.Sprint(v))
		default:
			calContainer.pushVal(char)
		}
	}
}

type calContainer struct {
	ops  []string
	vals []float64
}

func (o *calContainer) pushVal(valStr string) {
	val, _ := strconv.ParseFloat(valStr, 2)
	o.vals = append(o.vals, val)
}

func (o *calContainer) popVal() float64 {
	l := len(o.vals)
	popVal := o.vals[l-1]
	o.vals = o.vals[:l-1]
	return popVal
}

func (o *calContainer) pushOps(op string) {
	o.ops = append(o.ops, op)
}

func (o *calContainer) popOps() string {
	l := len(o.ops)
	popOp := o.ops[l-1]
	o.vals = o.vals[:l-1]
	return popOp
}
