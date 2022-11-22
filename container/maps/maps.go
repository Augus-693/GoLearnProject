package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    maps.go
 * @Author  Augus Lee
 * @Date    2022/10/28 16:59
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	m := map[string]string{
		"name":  "Augus Lee",
		"age":   "22",
		"sex":   "male",
		"hobby": "Coding",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map............")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting map values............")
	if age, ok := m["age"]; ok { //"age" is the key of the map
		fmt.Println(age)
	} else {
		fmt.Println("key does not exist")
	}
	if course, ok := m["course"]; ok { //"course" is not the key of the map
		fmt.Println(course)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting map values............")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)

}
