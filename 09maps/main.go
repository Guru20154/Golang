package main

import "fmt"

func main() {
	fmt.Println("Maps here")

	lang := make(map[string]string)
	lang["JS"] = "JAVASCRIPT"
	lang["PY"] = "PYTHON"
	lang["RB"] = "RUBY"
	lang["C++"] = "C++"
	lang["HTML"] = "HYPERTEXTMARKUPLANGUAGE"

	fmt.Println(lang)
	fmt.Println("JS: ",lang["JS"])

	delete(lang,"RB")
	fmt.Println("After Delete\n",lang)

	for k,v := range lang{
		fmt.Printf("Key is %v and Value is %v\n",k,v)
	}
}
