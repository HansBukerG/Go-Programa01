package main

import (
	DataService "front-ejemplo/src/conexiones"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var plantillas = template.Must(template.ParseGlob("./plantillas/*"))

/*Este comando junto a la importacion de template, sirve para modularizar nuestra aplicacion en diferentes carpetas*/

func main() {
	DataService.ConexionMysql()
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/formularioCrear", formularioCrear)
	http.HandleFunc("/insertar", insertarData)
	http.HandleFunc("/borrar", borrarData)
	http.HandleFunc("/detalle", detalleData)
	http.HandleFunc("/actualizar", actualizarData)

	http.ListenAndServe(":8081", nil)
}

type Persona struct {
	Id     int
	Nombre string
	Correo string
}

func Inicio(writer http.ResponseWriter, request *http.Request) {
	// plantillas.ExecuteTemplate(w, "cabecera", nil)
	plantillas.ExecuteTemplate(writer, "inicio", DataService.ListaPersonas())
	// plantillas.ExecuteTemplate(w, "footer", nil)
}

func formularioCrear(writer http.ResponseWriter, request *http.Request) {
	plantillas.ExecuteTemplate(writer, "formularioCrear", nil)
}

func insertarData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		nombre := request.FormValue("nombre")
		correo := request.FormValue("correo")
		DataService.InsertDataPersonas(nombre, correo)
		http.Redirect(writer, request, "/", http.StatusMovedPermanently)
	}
}

func borrarData(writer http.ResponseWriter, request *http.Request) {
	// ConexionEstablecida := DataService.ConexionAbierta

	idPersona := request.URL.Query().Get("id")
	DataService.BorrarPersona(idPersona)
	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}

func detalleData(writer http.ResponseWriter, request *http.Request) {
	idPersona := request.URL.Query().Get("id")
	plantillas.ExecuteTemplate(writer, "formularioEditar", DataService.VerPersona(idPersona))
}

func actualizarData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id := request.FormValue("id")
		nombre := request.FormValue("nombre")
		correo := request.FormValue("correo")
		DataService.EditarPersona(id, nombre, correo)
		http.Redirect(writer, request, "/", http.StatusMovedPermanently)
	}
}
