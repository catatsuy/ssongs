package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/k0kubun/pp/v3"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tmp := []string{"a", "a", "a", "a", "b", "b", "c", "c", "d"}
	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})

	r := shuffle(tmp)
	pp.Println(r)
	si := make([]string, 0, len(tmp))
	for i := r.next; i != nil; i = i.next {
		si = append(si, i.value)
	}
	fmt.Println(si)
	fmt.Println(tmp)
}

type item struct {
	value string
	order float64
	next  *item
}

type rootItem struct {
	next *item
}

func (i *item) insert(newI item) {
	if i.next == nil {
		i.next = &newI
		return
	}

	if i.next.order < newI.order {
		i.next.insert(newI)
		return
	}

	tmp := i.next
	i.next = &newI
	i.next.next = tmp
}

func (r *rootItem) insert(newI item) {
	if r.next == nil {
		r.next = &newI
		return
	}

	if r.next.order < newI.order {
		r.next.insert(newI)
		return
	}

	tmp := r.next
	r.next = &newI
	r.next.next = tmp
}

func shuffle(tmp []string) (r rootItem) {
	m := make(map[string]int)
	for _, v := range tmp {
		c := m[v]
		m[v] = c + 1
	}

	for k, c := range m {
		advL := 1.0 / float64(c)
		o := rand.Float64() * advL
		r.insert(item{k, o, nil})
		for i := 1; i < c; i++ {
			o = o + advL*0.8 + rand.Float64()*advL*0.4
			r.insert(item{k, o, nil})
		}
	}

	return r
}
