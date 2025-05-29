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

	// Tipos de datos
	tipoVelocidad := modelo.TipoDeDato{Denominacion: "Velocidad de Onda", NombreUnidadMedidad: "m/s", ValorUmbral: 500}
	tipoFrecuencia := modelo.TipoDeDato{Denominacion: "Frecuencia de Onda", NombreUnidadMedidad: "Hz", ValorUmbral: 50}
	tipoLongitud := modelo.TipoDeDato{Denominacion: "Longitud de Onda", NombreUnidadMedidad: "m", ValorUmbral: 200}

	// Estaciones
	estacionNorte := &modelo.EstacionSismologica{Nombre: "Estación Norte"}
	estacionSur := &modelo.EstacionSismologica{Nombre: "Estación Sur"}

	// Series temporales
	serieTemporal1 := modelo.GenerarSerieTemporal(tipoVelocidad, tipoFrecuencia, tipoLongitud, time.Now())
	serieTemporal2 := modelo.GenerarSerieTemporal(tipoVelocidad, tipoFrecuencia, tipoLongitud, time.Now())

	// Sismógrafos
	sismografo1 := modelo.NewSismografo(time.Now().AddDate(-2, 0, 0), 1, "SN123", serieTemporal1, estacionNorte)
	sismografo2 := modelo.NewSismografo(time.Now().AddDate(-1, 0, 0), 2, "SS456", serieTemporal2, estacionSur)

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
	gestorEventos.AddSismografo(sismografo1)
	gestorEventos.AddSismografo(sismografo2)

	//--------------------------------------------------------------------------------------------------------------------

	// Ruta inicio
	r.GET("/", *pantallaEventos.MostrarPaginaPrincipal(gestorEventos))
	r.GET("/inicio", *pantallaEventos.MostrarPaginaPrincipal(gestorEventos))

	// Login template
	r.GET("/login", *pantallaEventos.MostrarLogin(gestorEventos))

	// Procesar login
	r.POST("/login", *pantallaEventos.HabilitarLogin(gestorEventos))

	// Cerrar sesion
	r.GET("/cerrarsesion", *pantallaEventos.HabilitarCerrarSesion(gestorEventos), *pantallaEventos.MostrarPaginaPrincipal(gestorEventos))

	// Crear E.S.
	r.POST("/sim-es-a", *pantallaEventos.OpcionCrearEventosAleatorios(gestorEventos))

	// Listar E.S.
	r.POST("/list-es", *pantallaEventos.MostrarListaEventos(gestorEventos))

	// Revision manual.
	r.POST("/review-es", *pantallaEventos.MostrarRevisionManual(gestorEventos))

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
