package main

import (
	"encoding/json"
	"fmt"
	"log"
	"meliQuasar/controllers"
	"meliQuasar/dto"
	"net/http"
)


func main(){
	http.HandleFunc("/", topsecret)
	//fmt.Println("test")
	//controllers.Test();

	log.Println("Servidor corriendo.....")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func topsecret(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var rq dto.TopSecret
	
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// procesa los datos de la solicitud aquí
	//log.Printf("Received data: %v", request.Satellites)

	result, err := controllers.GetTopSecret(rq)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

	
	
}