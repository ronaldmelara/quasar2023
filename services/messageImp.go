package services

import (
	"errors"
	//"fmt"
	//"log"
	"meliQuasar/model"
	"meliQuasar/repository"
	"meliQuasar/util"
)

func GetMessage(messages ...[]string) ([]string, error){

	var lstSat []model.Satellite
	lstSat = repository.GetSatellites()

	if len(messages) < len(lstSat){
		return []string{""} , &util.Exception{
			StatusCode: 502,
			Err: errors.New("Three messages are required"),
		}
	}
	var maxSize int

	//saco el largo maximo de los mensajes recibidos
	for m:=0; m <len(messages); m++{
		if (len(messages[m])>maxSize){
			maxSize = len(messages[m])
		}
	}
	//igual largo a aquellos arreglos que sean mas cortos
	for m:=0; m <len(messages); m++{
		if (len(messages[m])<maxSize){
			newArr := make([]string, maxSize-len(messages[m])) 

			messages[m] = append(newArr,messages[m]...)
		}
	}

	
	var t int = 0
	for l:=0; l< len(messages);l++{
		t++
		if (l < len(messages)-1){
			messages[l+1] = combinedCollection(messages[l], messages[l+1])
			l++
		}else{
			messages[t] = combinedCollection(messages[l-1], messages[t])
		}
		
	}


	return messages[0],nil
}

func combinedCollection(col1, col2 []string) []string{
	var j int = 0

	for i:=0; i < len(col1); i++{
		if (j>= len(col2) || col1[i] != col2[i]){
			//missing = append(missing, res1[i])
			col2[i] = col1[i]
		}else{
			j=j+1
		}
	}
	return col2
}