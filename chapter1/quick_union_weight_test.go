package chapter1

import (
	"log"
	"testing"
)

type QUW struct {
	id    []int
	idw   []int
	count int
}

func NewQUW(i int) QUW {
	u := QUW{}
	u.id = make([]int, 0, i)
	u.idw = make([]int, i)
	for x := 0; x < i; x++ {
		u.id = append(u.id, x)
		u.idw[x] = 1
		// u.idw = append(u.idw, 0)
	}
	u.count = i
	return u
}
func (o *QUW) Count() int {
	return o.count
}

func (o *QUW) Connected(p, q int) bool {
	return o.Find(p) == o.Find(q)
}

func (o *QUW) Find(p int) int {
	for {
		if p != o.id[p] {
			p = o.id[p]
		} else {
			return p
		}
	}
}

func (o *QUW) Union(p, q int) {
	i := o.Find(p)
	j := o.Find(q)
	if i == j {
		return
	}
	if o.idw[i] > o.idw[j] {
		o.id[i] = j
		o.idw[j] += o.idw[i]
	} else {
		o.id[j] = i
		o.idw[i]+=o.idw[j]
	}
	o.count--

}

func Test_QUW(t *testing.T) {
	u := NewQUW(6)
	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(2, 3)
	// u.Union(0, 4)
	u.Union(4, 5)
	log.Println(u.Count())
	log.Println(u.Connected(1, 5))
	log.Println(u.Connected(0, 3))
	log.Println(u.id)
	log.Println(u.idw)
}
