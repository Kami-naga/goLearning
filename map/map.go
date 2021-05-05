package main

import "fmt"

func main() {
	m := map[string]string{
		"name":     "aaa",
		"language": "GO",
		"time":     "morning",
		"location": "JP",
	}

	m2 := make(map[string]int) // m2=empty map

	var m3 map[string]int // m3==nil

	//unlike null in other languages, nil in GO can participate in computing
	fmt.Println(m, m2, m3)

	fmt.Println("Traverse map")
	//map is a hash map, so the map is unordered
	//each time you traverse it, the order may change
	//if you want it get in order, you need to put keys into a slice and sort the slice,
	//then traverse the map with the sorted key
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	lang := m["language"]
	fmt.Println(lang)
	//if you are getting a key which  does not exist?
	fmt.Println(m["notExist"]) //you will get empty string, not error

	//how to know a key exists or not?
	lang2, ok := m["language"]
	fmt.Println(lang2, ok)
	if notExist, ok := m["notExist"]; ok {
		fmt.Println(notExist)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("delete element")
	name, ok := m["name"] //here though ok has benn defined before, but name not, so here can use :=
	fmt.Println(name, ok, len(m))
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok, len(m))

	//what type can a map's key be?
	//map uses hash, so it should be compared
	//in GO, except for slice, map, function, other types can all be keys
	//for custom type, as long as no slice,map or function inside, it's also OK
	//no need to define hashCode, equals functions like java, GO do this for us
}
