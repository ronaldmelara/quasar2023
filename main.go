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

// @title           API Operación Fuego de Quasar - MELI
// @version         1.0
// @description     Esta API corresponde al desarrollo planteado para el desafío técnico de Mercado Libre. La API REST ha sido desarrollada en GO. Es posible ver visualizar el código alojado en el github https://github.com/ronaldmelara/quasar2023

// @contact.name   Ronald Melara
// @contact.email  ronald.melara@gmail.com

// @host      quasar2023-jnswrwco3q-uc.a.run.app
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
// @Summary      Cálculo de Trilateración y Decifrar mensaje 
// @Description  Este servicio permite calcular la trilateración de la nave enemiga en base a las 3 posiciones de los satélites, la distancia del objecto (nave enemiga) a cada satélite. Adicional, realiza un merge de cada fragmento de los mensajes recepcionado en cada satélite para entregarlo en un unico mensaje.
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
		fmt.Println(err)
		return
	}

	// procesa los datos de la solicitud aquí
	//log.Printf("Received data: %v", request.Satellites)

	result, err := controllers.GetTopSecret(rq)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}


// topsecret_split godoc
// @Summary      Obtener la información de un satélite en particular
// @Description  Enviando el nombre de un satélite se puede obtener la posición donde se encuentra en el plano.
// @Tags         TopSecretSplit
// @Accept       json
// @Produce      json
// @Param        satellite_name path string   true   "Nombre del Satélite"
// @Success      200  {object}  dto.ResponseTopSecret
// @Router       /topsecret_split/{satellite_name} [get]
func getSatelliteGet(w http.ResponseWriter, r *http.Request){

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/topsecret_split/")

	if len(id) <= 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	
	if r.Method == http.MethodGet{
		result, err := controllers.GetTopSecretSplitByName(id)

		if err != nil{
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
		}else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}	
}

// topsecret_split godoc
// @Summary      Obtener la información de un satélite en particular y su mensaje
// @Description  Enviando el nombre de un satélite, mas la distancia se puede obtener la posición donde se encuentra en el plano y saber si la distancia ingresada se encuentra en el radio de alcance del satélite.
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
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
}

