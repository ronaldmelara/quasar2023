package repository

import (
	_"encoding/json"
	"meliQuasar/model"
	_ "meliQuasar/model"
	_"log"
	"fmt"
	"database/sql"
	// Driver para SQLite3
    _ "github.com/mattn/go-sqlite3"
)

func conexionBD()(conexion *sql.DB){
	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		fmt.Println(err)
		return
	}

	return db
}


func GetSatellites() ([]model.Satellite,error){

	err := createSatellites()
	if(err != nil){
		return nil, err
	}
	satelites, err:= dbLoadSatellites()
	if (err != nil){
		return nil,err
	}

	return satelites, err

}

func SaveMessages(menssages ...[]string)error{
	db:= conexionBD()

	for idx:=0; idx < len(menssages); idx++{
		stmt, err := db.Prepare("UPDATE sateliteMensage SET mensaje = ? WHERE id = ?")
		if err != nil {
			fmt.Println(err)
			return err
		}

		blob, err := encodeBlob(menssages[idx])
		if err != nil {
			panic(err)
		}


		_ , err = stmt.Exec(blob, idx+1)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func GetMessages()([]model.SatelliteMessage, [][]string){

	db:= conexionBD()
	var messages []model.SatelliteMessage
	var arrMsg [][]string
	rows, err := db.Query("SELECT id, mensaje FROM sateliteMensage ORDER BY id")
	if err != nil {
		fmt.Println(err)
		
	}
	for rows.Next() {
		var blob []byte
		var id int
		var sm  model.SatelliteMessage

		if err := rows.Scan(&id, &blob); err != nil {
			fmt.Println(err)
		}
		
		sm.Id = id
		// Decodifica el blob en un slice de strings
		strs, err := decodeBlob(blob)
		if err != nil {
			panic(err)
		}

		sm.Message = strs

		messages = append(messages,sm)
		arrMsg = append(arrMsg, strs)
	}
	return messages, arrMsg
}



func createSatellites()(error){
	db:= conexionBD()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS satelites (id INTEGER, nombre TEXT, x REAL, y REAL)")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO satelites (id,nombre, x, y) SELECT ?, ?, ? ,? WHERE NOT EXISTS (SELECT 1 FROM satelites WHERE nombre = ?)")
	if err != nil {
		panic(err)
	}

	_ , err = stmt.Exec(1, "kenobi", -500.0, -200.0, "kenobi")
	if err != nil {
		panic(err)
	}

	_ , err = stmt.Exec(2, "skywalker", 100.0, -100.0, "skywalker")
	if err != nil {
		panic(err)
	}

	_ , err = stmt.Exec(3, "sato", 500.0, 100.0, "sato")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sateliteMensage (id INTEGER, mensaje BLOB)")
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("INSERT INTO sateliteMensage (id,mensaje) SELECT ?, ? WHERE NOT EXISTS (SELECT 1 FROM sateliteMensage WHERE id=?)")
	if err != nil {
		panic(err)
	}

	blob, err := encodeBlob([]string{""})
    if err != nil {
        panic(err)
    }

	_ , err = stmt.Exec(1, blob, 1)
	if err != nil {
		panic(err)
	}

	_ , err = stmt.Exec(2, blob, 2)
	if err != nil {
		panic(err)
	}

	_ , err = stmt.Exec(3, blob, 3)
	if err != nil {
		panic(err)
	}

	return nil
}

func dbLoadSatellites()([]model.Satellite, error){

	var satellites []model.Satellite

	db:= conexionBD()
	var sat model.Satellite
	rows, err := db.Query("SELECT id, nombre, x, y FROM satelites")
	if err != nil {
		panic(err)
		
	}
	for rows.Next() {
		var nombre string
		var id int
		var x float32
		var y float32
		if err := rows.Scan(&id, &nombre, &x,&y); err != nil {
			panic(err)
		}

		sat.Id = id
		sat.Name = nombre
		sat.X = x
		sat.Y = y

		satellites = append(satellites,sat)
	}
	return satellites,nil
}

// Función para codificar un slice de strings en un blob
func encodeBlob(strs []string) ([]byte, error) {
    var blob []byte
    for _, str := range strs {
        blob = append(blob, []byte(str)...)
        blob = append(blob, 0) // Agrega un byte nulo después de cada string
    }
    return blob, nil
}

// Función para decodificar un blob en un slice de strings
func decodeBlob(blob []byte) ([]string, error) {
    var strs []string
    str := ""
    for _, b := range blob {
        if b == 0 {
            strs = append(strs, str)
            str = ""
        } else {
            str += string(b)
        }
    }
    return strs, nil
}

func GetMessageBySatellite(id int)([]string,error){

	db:= conexionBD()
	var blob []byte
	err := db.QueryRow("SELECT mensaje FROM sateliteMensage WHERE id=?", id).Scan(&blob)

	if err != nil {
		panic(err)
	}

	strs, err := decodeBlob(blob)
	if err != nil {
		panic(err)
	}

	return strs,nil

}