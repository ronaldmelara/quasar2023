package controllers

import (
	"fmt"
	"meliQuasar/services"
)


func Test(){
	fmt.Println("Este es un tesst")
	services.LocationImp()
	services.GetLocation(2.3,56.0,300.1)
}