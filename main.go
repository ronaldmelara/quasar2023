package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"meliQuasar/controllers"
	"meliQuasar/dto"
	"net/http"
	_ "meliQuasar/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           API Operación Fuego de Quasar
// @version         1.0
// @description     Esta API corresponde al desarrollo del desafío tecnico a desarrollar en GO
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main(){
	http.HandleFunc("/api/v1/topsecret/", topsecret)
	http.HandleFunc("/api/v1/topsecret_split/", getSatelliteGet)
	http.HandleFunc("/api/v1/p/topsecret_split/", getSatellitePost)


	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Servidor corriendo.....")
	log.Fatal(http.ListenAndServe(":8080", nil))

	
	
}

// topsecret godoc
// @Summary      Get position and message 
// @Description  Obtains the position of the enemy ship and deciphers the message received on the satellites
// @Tags         TopSecret
// @Accept       json
// @Produce      json
// @Param        Body   body   dto.TopSecret true "Se debe enviar por el body la siguiente estructura JSON"
// @Success      200  {object}  dto.ResponseTopSecret
// @Router       /topsecret/ [post]
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


// topsecret_split godoc
// @Summary      Get position and message for a specific satellite
// @Description  Obtains the position of the enemy ship and deciphers the message received on the satellites
// @Tags         TopSecretSplit
// @Accept       json
// @Produce      json
// @Param        satellite_name path string   true   "Nombre del Satélite"
// @Success      200  {object}  dto.ResponseTopSecret
// @Router       /topsecret_split/{satellite_name} [get]
func getSatelliteGet(w http.ResponseWriter, r *http.Request){

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/topsecret_split/")

	
	fmt.Println(id)
		
	//path := strings.Split(r.URL.Path, "/")
	fmt.Println(len(id))
	if len(id) <= 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	
	if r.Method == http.MethodGet{
		

		result, err := controllers.GetTopSecretSplitByName(id)

		if err != nil{
			//fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
		}else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}
	}	
}

// topsecret_split godoc
// @Summary      Get position and message for a specific satellite
// @Description  Obtains the position of the enemy ship and deciphers the message received on the satellites
// @Tags         TopSecretSplit
// @Accept       json
// @Produce      json
// @Param        satellite_name path string   true   "Nombre del Satélite"
// @Param        Data body dto.TopSecret   true   "Distancia de la nave enemiga"
// @Success      200  {object}  dto.ResponseTopSecret
// @Router       /p/topsecret_split/{satellite_name} [post]
func getSatellitePost(w http.ResponseWriter, r *http.Request){

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/p/topsecret_split/")


	if len(id) <= 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost{
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
			json.NewEncoder(w).Encode(err.Error())
		}else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

	}
	
}

