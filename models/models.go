package models

import (
	"challenge/database"
	/* "github.com/lib/pq" */
	"fmt"
	"log"
)

type Estudiante struct {
	ID          int    `json:"id"`
	Nombres string `json:"nombres,"`
	Apellidos string `json:"apellidos"`
	Fecha_Nacimiento string `json:"fecha_n"`
	Direccion string `json:"direccion"`
	Correo string `json:"correo"`
	ID_Curso int `json:"id_curso,omitempty"`
}

type Curso struct{
	ID          int    `json:"id"`
	Nombre string `json:"nombre"`
	Horario string `json:"horario"`
	Fecha_Inicio string `json:"fecha_i"`
	Lugar string `json:"lugar"`
	Descripcion string `json:"descripcion,omitempty"`
}

type Detalle struct{
	ID int `json:"id"`
	ID_Curso int `json:"id_curso"`
	ID_Estudiante int `json:"id_estudiante"`
}

type Estudiante_x_Curso struct {
	ID          int    `json:"id"`
	Nombres string `json:"nombres,omitempty"`
	Apellidos string `json:"apellidos,omitempty"`
	Edad int `json:"edad,omitempty"`
	Direccion string `json:"direccion,omitempty"`
	Correo string `json:"correo,omitempty"`
	ID_Curso int `json:"id_curso,omitempty"`
	Curso string `json:"curso"`
}

type Grafica_Cursos struct{
	ID int `json:"id"`
	Curso string `json:"curso"`
	Cantidad int `json:"cantidad"`
}







func InsertarCurso( nombre , horario , lugar , fecha_i , descripcion string) (Curso, bool) {
	db := database.GetConnection()

	var id_curso int
	db.QueryRow("INSERT INTO cursos(nombre,horario,fecha_inicio,lugar,descripcion) VALUES($1,$2,$3,$4,$5) RETURNING id_curso",
	nombre,horario,fecha_i,lugar,descripcion).Scan(&id_curso)


	if id_curso == 0 {
		defer db.Close()
		return Curso{}, false
	}

	defer db.Close()
	return Curso{id_curso, "","","","",""}, true
}


func InsertarEstudiante( nombres , apellidos , fecha_n , direccion , correo string, id_curso int) (Estudiante, bool) {
	db := database.GetConnection()

	var id_estudiante int
	/* db.QueryRow("INSERT INTO estudiantes(nombres,apellidos,fecha_nacimiento,direccion,correo) VALUES($1,$2,$3,$4,$5) RETURNING id_estudiante",nombres,apellidos,fecha_n,direccion,correo).Scan(&id_estudiante) */

	db.QueryRow("INSERT INTO estudiantes(nombres,apellidos,fecha_nacimiento,direccion,correo) VALUES($1,$2,$3,$4,$5) RETURNING id_estudiante",
	nombres,apellidos,fecha_n,direccion,correo).Scan(&id_estudiante)

	

	if id_estudiante == 0 {
		defer db.Close()
		return Estudiante{}, false
	}
	InsertarDetalle(id_curso, id_estudiante)

	defer db.Close()
	return Estudiante{id_estudiante, "","","","","", id_curso}, true
}

func InsertarDetalle(id_curso, id_estudiante int){
	db := database.GetConnection()

	
	/* db.QueryRow("INSERT INTO estudiantes(nombres,apellidos,fecha_nacimiento,direccion,correo) VALUES($1,$2,$3,$4,$5) RETURNING id_estudiante",nombres,apellidos,fecha_n,direccion,correo).Scan(&id_estudiante) */

	db.QueryRow("INSERT INTO detalle_estudiantes_curso(id_estudiante, id_curso) VALUES($1,$2)",id_estudiante,id_curso)
	defer db.Close()
}


func SelectCursos() ([]Curso, bool) {
	cursos_totales:=make( []Curso,0)

	db := database.GetConnection()
	rows, err := db.Query("SELECT id_curso,nombre,horario,TO_CHAR(fecha_inicio,'dd/MM/yyyy'),lugar,descripcion  FROM cursos")
	
	if err!=nil{
		fmt.Println("Error")
		return cursos_totales,false		
	}

	defer rows.Close()

	var i=0
	for rows.Next() {
		
		var curso Curso

		if err := rows.Scan(&curso.ID,&curso.Nombre,&curso.Horario,&curso.Fecha_Inicio,&curso.Lugar,
			&curso.Descripcion); err != nil {
				
				
				fmt.Println("Error_vacio")
				
		}

		cursos_totales=append(cursos_totales, curso)
		i++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return cursos_totales,false	
	}


	
	defer db.Close()
	return cursos_totales, true
}

func SelectEstudiantes() ([]Estudiante_x_Curso, bool) {
	estudiantes_totales:=make( []Estudiante_x_Curso,0)

	db := database.GetConnection()
	rows, err := db.Query("SELECT cur.id_curso, cur.nombre, e.nombres, e.apellidos,extract(year from current_date)-extract(year from e.fecha_nacimiento),e.direccion,e.id_estudiante, e.correo from estudiantes as e, cursos as cur, detalle_estudiantes_curso as d where d.id_estudiante=e.id_estudiante and cur.id_curso=d.id_curso")
	
	if err!=nil{
		fmt.Println("Error")
		return estudiantes_totales,false		
	}

	defer rows.Close()

	var i=0
	for rows.Next() {
		
		var estudiante Estudiante_x_Curso

		if err := rows.Scan(&estudiante.ID_Curso,&estudiante.Curso,&estudiante.Nombres,&estudiante.Apellidos,&estudiante.Edad,
			&estudiante.Direccion,&estudiante.ID,&estudiante.Correo); err != nil {
				fmt.Println("Error")
		}

		estudiantes_totales=append(estudiantes_totales, estudiante)
		i++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return estudiantes_totales,false	
	}


	
	defer db.Close()
	return estudiantes_totales, true
}


func SelectGraficaCursos() ([]Grafica_Cursos, bool, int) {
	cursos:=make( []Grafica_Cursos,0)
	var sum int
	db := database.GetConnection()
	rows, err := db.Query("select cursos.id_curso, cursos.nombre, count(d.id_estudiante) from cursos, detalle_estudiantes_curso as d where cursos.id_curso=d.id_curso group by cursos.id_curso,cursos.nombre")
	
	if err!=nil{
		fmt.Println("Error")
		return cursos,false,0	
	}

	defer rows.Close()

	var i=0
	for rows.Next() {
		
		var curso Grafica_Cursos
		if err := rows.Scan(&curso.ID,&curso.Curso, &curso.Cantidad); err != nil {
				fmt.Println("Campo Vacio")				
		}

		cursos=append(cursos, curso)
		sum+=curso.Cantidad
		i++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return cursos,false,0
	}


	
	defer db.Close()
	return cursos, true, sum
}