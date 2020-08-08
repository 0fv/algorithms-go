package chapter1

import (
	"log"
	"testing"
)

func ShellSort(n1 []int) {
	N := len(n1)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && n1[j] < n1[j-h]; j -= h {
				n1[j], n1[j-h] = n1[j-h], n1[j]
			}
		}
		h = h / 3
	}
}

func Test_ShellSort(t *testing.T) {
	n1 := []int{5, 4, 3, 2, 1,10,9,8,7,6}
	ShellSort(n1)
	log.Println(n1)
}
