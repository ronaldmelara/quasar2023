package controllers

import (
	"fmt"
	"meliQuasar/services"
)


func Test(){
	fmt.Println("Este es un tesst")
	
	//x, y, err := services.GetLocation(-538.57, 141.42, 0)

	//x, y, err := services.GetLocation(538.57, 141.42, 509.90)
	x, y, err := services.GetLocation(100.0, 115.5, 142.7)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("x: %v - y: %v \n", x, y)
	}
	


	o, er:= services.GetMessage(
		[]string{"", "este", "es", "un", "mensaje"},
		[]string{"este", "", "un", "mensaje"},
		[]string{"", "", "es", "", "mensaje"})
	if er == nil{
		fmt.Println(o)
	}else{
		fmt.Println(er)
	}

	/*o, er:= services.GetMessage(
		[]string{})
	if er == nil{
		fmt.Println(o)
	}else{
		fmt.Println(er)
	}*/
}