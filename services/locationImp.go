package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"meliQuasar/model"
	"meliQuasar/repository"
	"meliQuasar/util"
)

func GetLocation(distances ...float32)(x, y float32){
	fmt.Println("Calculate Trilateration")
	var lstSat []model.Satellite
	lstSat = repository.GetSatellites()


	_, err := checkLenDistanceList(lstSat, distances...)
	if(err != nil){
		fmt.Println(err)
		return
	}

	//distance Kenobi and Skywalter
	kenobi :=lstSat[0]
	skywalker := lstSat[1]
	var a float32 = -2*kenobi.X + 2*skywalker.X
	var b float32 = -2*kenobi.Y + 2*skywalker.Y
	var c float32 = float32(
					math.Pow(float64(distances[0]),2)-
					math.Pow(float64(distances[1]),2)-
					math.Pow(float64(kenobi.X),2)+
					math.Pow(float64(skywalker.X),2)-
					math.Pow(float64(kenobi.Y),2)+
					math.Pow(float64(skywalker.Y),2))
	
	//distance Skywalker and Sato
	sato := lstSat[2]
	var e float32 = -2*skywalker.X + 2*sato.X
	var f float32 = -2*skywalker.Y + 2*sato.Y
	var g float32 = float32(
		math.Pow(float64(distances[1]),2)-
		math.Pow(float64(distances[2]),2)-
		math.Pow(float64(skywalker.X),2)+
		math.Pow(float64(sato.X),2)-
		math.Pow(float64(skywalker.Y),2)+
		math.Pow(float64(sato.Y),2))
	
	//Resolviendo por Determinante o regla de Cramer
	var d float32 = a*f-b*e

	var dx float32 = c*f-b*g
	var dy float32 = a*g-c*e

	var pointX = dx/d
	var pointY = dy/d

	return pointX,pointY
}

func checkLenDistanceList(s []model.Satellite, arr ...float32) (string, error) {
	log.Printf("Distances %v \n", arr)
	if(len(arr) == len(s)){
		return "", nil
	}else{
		return  "", &util.Exception{
			StatusCode: 502,
			Err: errors.New("Three distances are required"),
		}
	}
}