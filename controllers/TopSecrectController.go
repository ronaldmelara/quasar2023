package controllers

import (
	"errors"
	"fmt"
	"strings"
	"meliQuasar/dto"
	"meliQuasar/services"
	"meliQuasar/util"
)


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

	
	x, y, err := services.GetLocation(distances...)
	if err != nil{
		return dto.ResponseTopSecret{} , err
	}
	resp.Position.X = x
	resp.Position.Y = y

	er := services.SaveMessages(messages...)
	if(er != nil){
		return dto.ResponseTopSecret{}  ,  er
	}

	m, er := services.GetMessage()
	if er != nil{
		return dto.ResponseTopSecret{}  ,  er
	}

	resp.Message =  strings.Join(m," ")


	return resp, nil

}

func GetTopSecretSplit(rq dto.Entry)(dto.ResponseTopSecret, error){
	b, s := services.CheckExistsSatellite(rq.Name)
	if !b{
		msgErr :=  fmt.Sprintf("There is no information for the %s satellite", rq.Name)
		return dto.ResponseTopSecret{} , &util.Exception{
			StatusCode: 404,
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

	message, err := services.GetMessage()

	if err != nil{
		return dto.ResponseTopSecret{}  ,  err
	}

	resp.Message =  strings.Join(message," ")


	return resp, nil
}

func GetTopSecretSplitByName(name string)(dto.ResponseTopSecret, error){
	b, s := services.CheckExistsSatellite(name)
	if !b{
		msgErr :=  fmt.Sprintf("Could not find satellite %s", name)
		return dto.ResponseTopSecret{} , &util.Exception{
			StatusCode: 502,
			Err : errors.New(msgErr),
		}
	}

	
	var resp dto.ResponseTopSecret
	resp.Position.X = s.X
	resp.Position.Y = s.Y

	message, err := services.GetMessageBySatellite(s.Id)

	if err != nil{
		return dto.ResponseTopSecret{}  ,  err
	}

	resp.Message =  strings.Join(message,"***")


	return resp, nil
}