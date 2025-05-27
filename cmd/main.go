package main

import (
	"os/exec"
	"runtime"

	"ppai/internal/gestor"
	"ppai/internal/modelo"
	"ppai/internal/pantalla"

	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

var sesionActual modelo.Empleado
var clasificaciones []modelo.ClasificacionSismo
var origenDeGeneracion []modelo.OrigenDeGeneracion
var alcanceSismo []modelo.AlcanceSismo
var estados []modelo.Estado

func main() {
	r := gin.Default() // Crea una instancia del router con middlewares por defecto

	// funciones simples para las plantillas (para las funciones mas delicadas usare js)
	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"concatenate": func(a, b string) string {
			return a + b
		},
	})

	// Archivos estáticos: "/static" servirá los archivos en "./static"
	r.Static("/static", "./static")
	// Cargar plantillas HTML
	r.LoadHTMLGlob("templates/*")

	gestorEventos := gestor.NewGestorEventos()
	pantallaEventos := pantalla.NewPantalla(gestorEventos)

	//--------------------------------------------------------------------------------------------------------------------
	// Codigo Hardcodeado para pruebas

	clasificaciones = []modelo.ClasificacionSismo{}
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(0, 70, "Superficial"))
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(70, 300, "Intermedio"))
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(300, 700, "Profundo"))
	origenDeGeneracion = []modelo.OrigenDeGeneracion{
		modelo.NewOrigenDeGeneracion("Tectonico", "Movimiento de placas tectonicas"),
		modelo.NewOrigenDeGeneracion("Volcanico", "Actividad volcanica"),
		modelo.NewOrigenDeGeneracion("Colapso", "Colapso de cavernas o minas"),
		modelo.NewOrigenDeGeneracion("Artificial", "Actividad humana"),
		modelo.NewOrigenDeGeneracion("Desconocido", "Origen desconocido"),
	}
	alcanceSismo = []modelo.AlcanceSismo{
		modelo.NewAlcanceSismo("Sismo local", "Hasta 100 km"),
		modelo.NewAlcanceSismo("Sismo regional", "Hasta 1000 km"),
		modelo.NewAlcanceSismo("Tele sismo", "Mas de 1000 km"),
	}
	estados = modelo.GetEstadosMuestra()

	sesionActual = modelo.Empleado{
		Nombre:   "Juan",
		Apellido: "Test",
		Email:    "juan@Test.com",
		Telefono: "123456789",
	}

	gestorEventos.SetSesionActual(&sesionActual)
	gestorEventos.CrearEvento(0, time.Now().Add(-time.Hour*4), 900.0, 20.0, 50.0, 3.0, sesionActual, clasificaciones[0], origenDeGeneracion[0], alcanceSismo[0])
	gestorEventos.CrearEvento(1, time.Now().Add(-time.Hour*2), 500.0, 350.0, 100.0, 2.5, sesionActual, clasificaciones[1], origenDeGeneracion[1], alcanceSismo[1])
	gestorEventos.CrearEvento(2, time.Now().Add(-time.Hour), 150.0, 125.0, 150.0, 2.5, sesionActual, clasificaciones[1], origenDeGeneracion[1], alcanceSismo[1])

	//--------------------------------------------------------------------------------------------------------------------
	
	// Ruta inicio
	r.GET("/", *pantallaEventos.Principal(gestorEventos))
	r.GET("/inicio", *pantallaEventos.Principal(gestorEventos))

	// Login template
	r.GET("/login", *pantallaEventos.Login(gestorEventos))

	// Procesar login
	r.POST("/login", *pantallaEventos.PostLogin(gestorEventos))

	// Cerrar sesion
	r.GET("/cerrarsesion", *pantallaEventos.CerrarSesion(gestorEventos), *pantallaEventos.Principal(gestorEventos))

	// Crear E.S.
	r.POST("/sim-es-a", *pantallaEventos.CrearEventosAleatorios(gestorEventos))

	// Listar E.S.
	r.POST("/list-es", *pantallaEventos.ListarEventos(gestorEventos))

	// Revision manual.
	r.POST("/review-es", *pantallaEventos.RevisionManual(gestorEventos))

	openBrowser("http://localhost:8080/inicio")
	r.Run(":8080") // Inicia el servidor en el puerto 8080
}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	}

	cmd.Start()
}

// eq (equal) → {{if eq .Estado "activo"}}
// ne (not equal)
// lt (less than)
// gt (greater than)
// le (less or equal)
// ge (greater or equal)
// not
