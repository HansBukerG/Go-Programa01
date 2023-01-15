package DataService

import (
	"database/sql"
	"fmt"
)

/*
Esta función ejecuta un query entrante
*/
func EjecucionConsulta(query string) {
	// ConexionAbierta.Ping();
	insertarRegistro, err := ConexionAbierta.Prepare(query)

	if err != nil {
		panic(err.Error())
	} else {
		response, responseError := insertarRegistro.Exec()
		if responseError != nil {
			panic("Ha ocurrido un error: " + responseError.Error())
		} else {
			fmt.Println("El comando: ")
			fmt.Println("	" + query)
			fmt.Println("Ha sido ejecutado con exito!")
			fmt.Println("Respuesta:")
			fmt.Println(response)
		}
	}
}

/*
Esta función retorna un array de registros en base a la query que se este solicitando
*/
func ExtraerRegistro(query string) *sql.Rows {

	registrosEncontrados, err := ConexionAbierta.Query(query)
	if err != nil {
		panic("Fallo en la ejecución de comando: " + query + " Error: " + err.Error())
	}
	return registrosEncontrados
}
