package lru

import (
	"fmt"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}
func Re(s string, v Value) {
	fmt.Println(11111)
}
func TestCache_RemoveOldest(t *testing.T) {
	c := New(int64(0), Re)
	c.Add("key1", String("123456"))
	c.Add("k2", String("k2"))
	c.Add("k3", String("k3"))
	c.Add("k4", String("k4"))
	fmt.Println(c.Len())
}
