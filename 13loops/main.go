package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i:=0; i<len(days); i++{
	// 	fmt.Printf("%v : %v\n",i,days[i])
	// }

	// for i := range days{
	// 	fmt.Printf("%v : %v\n",i,days[i])
	// }
	
	// for i, day := range days{
	// 	fmt.Printf("%v : %v\n",i,day)
	// } 
	r := 1
	for r<10{
		if r==3{
			r++
			continue
		}else if r==7{
			break
		}else if r==8{
			goto redir
		}
		fmt.Println("Value is:",r)
		r++
	}
	redir:
		fmt.Println("Redirected here after",r)
}
