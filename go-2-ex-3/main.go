package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "m104",
		117: "m117",
		346: "m346",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[114] = "m114"
	// TODO: replace one
	modules[117] = "not m117"
	fmt.Println(modules)
}
