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
		if !services.CheckExistsSatellite(v.Name){
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