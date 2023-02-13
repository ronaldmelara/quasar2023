package controllers

import (
	"errors"
	"fmt"
	"strings"
	"meliQuasar/dto"
	"meliQuasar/services"
	"meliQuasar/util"
)

//const MESSAGE_ERROR_SAT_NOT_FOUND string = "Could not find satellite %s"

func GetTopSecret(rq dto.TopSecret)(dto.ResponseTopSecret, error){

	var resp dto.ResponseTopSecret
	var distances []float32
	var messages [][]string
	for _, v := range rq.Satellites{
		b, _ := services.CheckExistsSatellite(v.Name)
		if !b{
			msgErr :=  fmt.Sprintf("Could not find satellite %s", v.Name)
			return dto.ResponseTopSecret{} , &util.Exception{
				StatusCode: 502,
				Err : errors.New(msgErr),
			}
		}
		distances = append(distances, v.Distance)
		messages = append(messages, v.Message)
	}

	fmt.Println(distances)
	x, y, err := services.GetLocation(distances...)
	if err != nil{
		return dto.ResponseTopSecret{} , err
	}
	resp.Position.X = x
	resp.Position.Y = y

	fmt.Println(messages)
	m, er:= services.GetMessage(messages...)
	if er != nil{
		return dto.ResponseTopSecret{}  ,  er
	}

	resp.Message =  strings.Join(m," ")


	return resp, nil

}

func GetTopSecretSplit(rq dto.Entry)(dto.ResponseTopSecret, error){
	b, s := services.CheckExistsSatellite(rq.Name)
	if !b{
		msgErr :=  fmt.Sprintf("Could not find satellite %s", rq.Name)
		return dto.ResponseTopSecret{} , &util.Exception{
			StatusCode: 502,
			Err : errors.New(msgErr),
		}
	}

	isOk, err := services.CheckDistanceVsRadiusRange(rq.Distance, s)

	if err != nil && !isOk{
		return dto.ResponseTopSecret{}, err
	}
	
	var resp dto.ResponseTopSecret
	resp.Position.X = s.X
	resp.Position.Y = s.Y
	resp.Message = "test"

	return resp, nil
}