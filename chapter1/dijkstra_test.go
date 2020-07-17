package chapter1

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test_dijkstra(t *testing.T) {
	cal := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"
	cal2 := "( ( 1 + sqrt ( 5.0 ) ) / 2.0 )"
	log.Println(cal, "=", calulate(cal))
	log.Println(cal2, "=", calulate(cal2))

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
	return calContainer.popVal()
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
	o.ops = o.ops[:l-1]
	return popOp
}
