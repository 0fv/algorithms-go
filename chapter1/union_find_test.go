package chapter1

import (
	"log"
	"testing"
)

type UF struct {
	id    []int
	count int
}

func NewUF(i int) UF {
	u := UF{}
	u.id = make([]int, 0, i)
	for x := 0; x < i; x++ {
		u.id = append(u.id, x)
	}
	u.count = i
	return u
}

func (o *UF) Count() int {
	return o.count
}

func (o *UF) Connected(p, q int) bool {
	return o.Find(p) == o.Find(q)
}

func (o *UF) Find(p int) int {
	return o.id[p]
}

func (o *UF) Union(p, q int) {
	if o.Connected(p, q) {
		return
	}
	for i := range o.id {
		if o.id[i] == o.Find(p) {
			o.id[i] = o.Find(q)
		}
	}
	o.count--

}
func Test_01(t *testing.T) {
	u := NewUF(6)
	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(2, 3)
	// u.Union(0, 4)
	u.Union(4, 5)
	log.Println(u.Count())
	log.Println(u.Connected(1, 5))
	log.Println(u.id)
}
