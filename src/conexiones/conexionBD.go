package DataService

import (
	"database/sql"
	"fmt"
)
/*
Este parametro representa a la conexion abierta a la base de datos.
*/
var ConexionAbierta *sql.DB
/*
	Se establece una conexión a MySQL con los siguientes campos:
		driver := "mysql"
		user := "root"
		password := ""
		db_name := "curso_golang"
*/
func ConexionMysql() {
	driver := "mysql"
	user := "root"
	password := ""
	db_name := "curso_golang"

	conexion, err := sql.Open(driver, user+":"+password+"@tcp(127.0.0.1)/"+db_name)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexion establecida con exito!")
		ConexionAbierta = conexion
	}
}



