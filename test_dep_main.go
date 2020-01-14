package test_dep_main

import "fmt"
import "github.com/460s/test_dep"

func TestDepMain() {
	fmt.Println("Hello i'm testDepMain")
	test_dep.TestDep()
}
