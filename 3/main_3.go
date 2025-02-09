package main

import "fmt"

type StringIntMap struct {
	db map[string]int
}

func main() {
	myMap := NewStringIntMap()
	myMap.Add("first", 1)
	myMap.Add("second", 2)
	fmt.Println(myMap)
	myMap.Remove("first")
	fmt.Println(myMap)
	fmt.Println(myMap.Copy())
	fmt.Println(myMap.Exists("second"))
	fmt.Println(myMap.Get("second"))
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		db: make(map[string]int),
	}
}

func (str *StringIntMap) Add(key string, value int) {
	str.db[key] = value
}

func (str *StringIntMap) Remove(key string) {
	delete(str.db, key)
}

func (str *StringIntMap) Copy() map[string]int {
	copiedMap := make(map[string]int)
	for key, v := range str.db {
		copiedMap[key] = v
	}
	return copiedMap
}

func (str *StringIntMap) Exists(key string) bool {
	_, ok := str.db[key]
	return ok
}

func (str *StringIntMap) Get(key string) (int, bool) {
	v, ok := str.db[key]
	return v, ok
}
