package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"meliQuasar/controllers"
	"meliQuasar/dto"
	"net/http"
)


func main(){
	http.HandleFunc("/", topsecret)
	http.HandleFunc("/topsecret_split/", getSatellite)
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

func getSatellite(w http.ResponseWriter, r *http.Request){

	// Verificamos que sea una petición POST o GET
	if r.Method == http.MethodPost || r.Method == http.MethodGet {

		id := strings.TrimPrefix(r.URL.Path, "/topsecret_split/")
		//path := strings.Split(r.URL.Path, "/")
		fmt.Println(len(id))
		if len(id) <= 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		var rq dto.Entry
		err := json.NewDecoder(r.Body).Decode(&rq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rq.Name = id
		
		result, err := controllers.GetTopSecretSplit(rq)

		if err != nil{
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
		}else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

		//satelliteName := id
		//fmt.Fprintf(w, "Top secret information about satellite %s", satelliteName)
	}
	
}