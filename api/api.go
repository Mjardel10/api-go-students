package api

import (
	"encoding/json"
	"net/http"
	"strings"

	/* "github.com/gorilla/mux" */
	"challenge/models"
	/* "fmt" */
)

type Data struct {
	Success bool          `json:"success"`
	Data    []models.Curso `json:"data"`
	Errors  []string      `json:"errors"`
}

type DataEstudiante struct {
	Success bool          `json:"success"`
	Data    []models.Estudiante `json:"data"`
	Errors  []string      `json:"errors"`
}

type DataEstudiante_x_Curso struct {
	Success bool          `json:"success"`
	Data    []models.Estudiante_x_Curso `json:"data"`
	Errors  []string      `json:"errors"`
}

type DataGrafica struct {
	Success bool          `json:"success"`
	Total int		`json:"total"`
	Data    []models.Grafica_Cursos `json:"data"`
	Errors  []string      `json:"errors"`
}




func DecodeCurso(req *http.Request) (models.Curso, bool) {
	var curso models.Curso
	err := json.NewDecoder(req.Body).Decode(&curso)
	if err != nil {
		
		return models.Curso{}, false
	}
	
	return curso, true
}

func DecodeEstudiante(req *http.Request) (models.Estudiante, bool) {
	var estudiante models.Estudiante
	err := json.NewDecoder(req.Body).Decode(&estudiante)
	if err != nil {
		
		return models.Estudiante{}, false
	}
	return estudiante, true
}

func IsValidCurso(nombre,horario,lugar,fecha_i string) bool {
	nom := strings.TrimSpace(nombre)
	hor := strings.TrimSpace(horario)
	lug := strings.TrimSpace(lugar)
	fech := strings.TrimSpace(fecha_i)
	if len(nom) == 0 || len(hor) == 0 || len(lug) == 0 || len(fech) == 0  {
		return false
	}

	return true
}

func IsValidEstudiante(nombres,apellidos,direccion,fecha_n, correo string) bool {
	noms := strings.TrimSpace(nombres)
	aps := strings.TrimSpace(apellidos)
	dir := strings.TrimSpace(direccion)
	fech := strings.TrimSpace(fecha_n)
	cor:= strings.TrimSpace(correo)
	if len(noms) == 0 || len(aps) == 0 || len(dir) == 0 || len(fech) == 0 || len(cor) == 0  {
		return false
	}

	return true
}

func Crear_Curso(w http.ResponseWriter, req *http.Request) {
	bodyCurso, succes := DecodeCurso(req)
	if succes != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyCurso.Nombre = strings.TrimSpace(bodyCurso.Nombre)
	if !IsValidCurso(bodyCurso.Nombre, bodyCurso.Horario,bodyCurso.Lugar,bodyCurso.Fecha_Inicio) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

		return
	}

	curso, success := models.InsertarCurso(bodyCurso.Nombre, bodyCurso.Horario, bodyCurso.Lugar, bodyCurso.Fecha_Inicio, bodyCurso.Descripcion)

	
	if success != true {
		data.Errors = append(data.Errors, "could not create tod ")
		data.Success = false
	}

	data.Success = true
	data.Data = append(data.Data, curso)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	return
}

func Crear_Estudiante(w http.ResponseWriter, req *http.Request) {
	bodyEstudiante, succes := DecodeEstudiante(req)
	/* fmt.Println(bodyEstudiante.Nombres + bodyEstudiante.Apellidos + bodyEstudiante.Fecha_Nacimiento + bodyEstudiante.Direccion +
	bodyEstudiante.Correo) */
	if succes != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
	}

	var data DataEstudiante = DataEstudiante{Errors: make([]string, 0)}
	bodyEstudiante.Nombres = strings.TrimSpace(bodyEstudiante.Nombres)
	if !IsValidEstudiante(bodyEstudiante.Nombres, bodyEstudiante.Apellidos,bodyEstudiante.Direccion,bodyEstudiante.Fecha_Nacimiento,bodyEstudiante.Correo) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

		return
	}

	estudiante, success := models.InsertarEstudiante(bodyEstudiante.Nombres, bodyEstudiante.Apellidos,bodyEstudiante.Fecha_Nacimiento,bodyEstudiante.Direccion,bodyEstudiante.Correo, bodyEstudiante.ID_Curso)

	
	if success != true {
		data.Errors = append(data.Errors, "could not create tod ")
		data.Success = false
	}else{
		data.Success = true
	}

	
	data.Data = append(data.Data,estudiante)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	return
}

func Obtener_Cursos(w http.ResponseWriter, req *http.Request) {
	/* vars := mux.Vars(req) */
	/* id := vars["id"] */
	var data Data
	var cursos [] models.Curso
	var success bool
	cursos, success = models.SelectCursos()

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, " No se han encontrado Datos")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

		return
	}

	data.Success = true

	for i := 0; i < len(cursos); i++ {
		data.Data = append(data.Data, cursos[i])
	}
	

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Obtener_Estudiantes(w http.ResponseWriter, req *http.Request) {
	/* vars := mux.Vars(req) */
	/* id := vars["id"] */
	var data DataEstudiante_x_Curso
	var estudiantes [] models.Estudiante_x_Curso
	var success bool
	estudiantes, success = models.SelectEstudiantes()

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, " No se han encontrado Datos")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

		return
	}

	data.Success = true

	for i := 0; i < len(estudiantes); i++ {
		data.Data = append(data.Data, estudiantes[i])
	}
	

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}


func Obtener_Grafica_Cursos(w http.ResponseWriter, req *http.Request)  {
	var data DataGrafica
	var cursos [] models.Grafica_Cursos
	var success bool
	var total int
	cursos, success, total = models.SelectGraficaCursos()

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, " No se han encontrado Datos")
		data.Total = 0
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

		return
	}

	data.Success = true
	data.Total = total
	for i := 0; i < len(cursos); i++ {
		data.Data = append(data.Data, cursos[i])
	}
	

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}