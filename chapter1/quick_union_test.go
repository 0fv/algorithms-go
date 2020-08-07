package chapter1

import (
	"log"
	"testing"
)

type QU struct {
	id    []int
	count int
}

func NewQU(i int) QU {
	u := QU{}
	u.id = make([]int, 0, i)
	for x := 0; x < i; x++ {
		u.id = append(u.id, x)
	}
	u.count = i
	return u
}

func (o *QU) Count() int {
	return o.count
}

func (o *QU) Connected(p, q int) bool {
	return o.Find(p) == o.Find(q)
}

func (o *QU) Find(p int) int {
	for {
		if p != o.id[p] {
			p = o.id[p]
		} else {
			return p
		}
	}
}

func (o *QU) Union(p, q int) {
	if o.Connected(p, q) {
		return
	}
	o.id[o.Find(p)] = o.Find(q)
	o.count--

}

func Test_QU(t *testing.T) {
	u := NewQU(6)
	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(2, 3)
	// u.Union(0, 4)
	u.Union(4, 5)
	log.Println(u.Count())
	log.Println(u.Connected(1, 5))
	log.Println(u.Connected(0, 3))
	log.Println(u.id)
}
