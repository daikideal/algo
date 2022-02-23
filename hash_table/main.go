package main

import (
	"fmt"
)

type HashTable struct {
	Size  int
	Table [][][]string // [key, value]を格納するスライス
}

func NewHashTable(size int) *HashTable {
	t := make([][][]string, size)
	h := &HashTable{Size: size, Table: t}

	return h
}

// キーのハッシュ値を求める
//
// 参考:
// 	- https://selfnote.work/20210923/programming/golang-algorithm-hash/
// 	- https://stackoverflow.com/questions/28128285/best-way-to-convert-an-md5-to-decimal-in-golang
func (h HashTable) hash(key string) int {
	var checksum int
	for _, c := range key {
		checksum = checksum + int(c)
	}
	index := checksum % h.Size

	return index
}

func (h *HashTable) Add(key string, value string) {
	index := h.hash(key)
	if h.Table[index] != nil {
		for _, m := range h.Table[index] {
			if m[0] == key {
				m[1] = value
				return
			}
		}
	}

	newData := []string{key, value}
	h.Table[index] = append(h.Table[index], newData)

	return
}

func (h HashTable) Get(key string) string {
	index := h.hash(key)
	for _, m := range h.Table[index] {
		if m[0] == key {
			return m[1]
		}
	}

	return ""
}

func (h HashTable) Print() {
	fmt.Print("{")
	for i := range h.Table {
		for _, kv := range h.Table[i] {
			fmt.Printf("%v->%v,", kv[0], kv[1])
		}
	}
	fmt.Print("}")

	return
}

func main() {
	h := NewHashTable(3)
	h.Add("やま", "かわ")
	h.Add("空", "青")
	h.Add("Apple", "iPhone")
	h.Add("やま", "うみ")

	fmt.Println(h.Get("Apple"))
	fmt.Println(h.Get("やま"))

	h.Print()
}
