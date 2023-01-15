package DataService

type Persona struct {
	Id     int
	Nombre string
	Correo string
}

// Insert de valores en tabla
func InsertDataPersonas(nombre string, apellido string) {
	var query string = "Insert into personas(nombre,correo) values ('" + nombre + "' ,'" + apellido + "')"
	EjecucionConsulta(query)
}

func BorrarPersona(id string) {
	var query string = "DELETE FROM personas WHERE personas.id = " + id
	EjecucionConsulta(query)
}

func EditarPersona(id string, nombre string, correo string) {
	var query string = "UPDATE personas SET nombre = '" + nombre + "', correo = '" + correo + "' WHERE personas.id = " + id
	// fmt.Println(query)
	EjecucionConsulta(query);
}

/*Esta funcion trae solo 1 registro producto de un sola seleccion, ideal para editar registros*/
func VerPersona(idPersona string) Persona {
	var query string = "select id,nombre,correo from personas where personas.id = " + idPersona
	persona := Persona{}
	registro := ExtraerRegistro(query)

	for registro.Next() {
		var id int
		var nombre, correo string
		err := registro.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		persona.Id = id
		persona.Nombre = nombre
		persona.Correo = correo
	}

	return persona
}

/*
Esta funcion sirve para traer una lista de registros de la base de datos
ideal para hacer un read.
*/
func ListaPersonas() (listaPersonas []Persona) {
	persona := Persona{}
	arregloPersonas := []Persona{}
	var query string = "select * from personas"

	registros := ExtraerRegistro(query)

	for registros.Next() {
		var id int
		var nombre, correo string
		err := registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		persona.Id = id
		persona.Nombre = nombre
		persona.Correo = correo

		arregloPersonas = append(arregloPersonas, persona)

	}
	return arregloPersonas
}
