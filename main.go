package main

import (
	"encoding/json"
	
	_ "fmt"
	"log"
	"net/http"
	_"io/ioutil"
	"meliQuasar/dto"
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

	var request dto.TopSecret
	
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// procesa los datos de la solicitud aquí
	log.Printf("Received data: %v", request.Satellites)

	response := dto.ResponseTopSecret{
		Message: "Solicitud procesada exitosamente",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	
}