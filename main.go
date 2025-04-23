package main

import (
	"fmt"
	"net/http"
)

// Función para manejar la solicitud GET
func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Manejo de GET
		fmt.Fprintf(w, "<h1>Soy alumno de la UOC</h1>")
		fmt.Fprintf(w, `<img src="/images/logo.jpg" alt="Imagen UOC" />`)
	case http.MethodPost:
		// Manejo de POST 
		r.ParseForm()
		fmt.Fprintf(w, "<h1>No tenemos datos a enviar</h1>")
		fmt.Fprintf(w, "<h1>Soy alumno de la UOC</h1>")
		fmt.Fprintf(w, `<img src="/images/logo.jpg" alt="Imagen UOC" />`)

	case http.MethodPut:
		// Manejo de PUT 
		r.ParseForm()
		fmt.Fprintf(w, "<h1>No hay datos a actualizar</h1>")
		fmt.Fprintf(w, "<h1>Soy alumno de la UOC</h1>")
		fmt.Fprintf(w, `<img src="/images/logo.jpg" alt="Imagen UOC" />`)

	case http.MethodDelete:
		// Manejo de DELETE 
		fmt.Fprintf(w, "<h1>Nada a eliminar</h1>")
		fmt.Fprintf(w, "<h1>Soy alumno de la UOC</h1>")
		fmt.Fprintf(w, `<img src="/images/logo.jpg" alt="Imagen UOC" />`)

	default:
		// Si el método no está contemplado, devolver error 405
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Función principal
func main() {
	// Rutas y sus métodos
	http.HandleFunc("/", handleRoot)
	// Servir archivos estáticos
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}