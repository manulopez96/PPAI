package main

import (
	"os/exec"
	"runtime"

	"ppai/internal/gestor"
	"ppai/internal/modelo"

	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

var sesionActual modelo.Empleado
var clasificaciones []modelo.ClasificacionSismo
var origenDeGeneracion []modelo.OrigenDeGeneracion
var alcanceSismo []modelo.AlcanceSismo

func main() {
	r := gin.Default() // Creo un router

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

	gestorRegistrarResultado := gestor.NewGestorEventos()
	gestorPagina := gestor.NewGestorPagina(gestorRegistrarResultado)

	//--------------------------------------------------------------------------------------------------------------------
	// Codigo Hardcodeado para pruebas

	// Estados
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Auto Confirmado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Auto Detectado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Pendiente de revision"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Bloqueado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Rechazado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Derivado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Aceptado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Pendiente de cierre"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Cerrado"))
	gestorRegistrarResultado.Estados = append(gestorRegistrarResultado.Estados, modelo.NewEstado("Evento sismico", "Sin revision"))

	// Tipos de datos
	tipoVelocidad := modelo.TipoDeDato{Denominacion: "Velocidad de Onda", NombreUnidadMedidad: "m/s", ValorUmbral: 500}
	tipoFrecuencia := modelo.TipoDeDato{Denominacion: "Frecuencia de Onda", NombreUnidadMedidad: "Hz", ValorUmbral: 50}
	tipoLongitud := modelo.TipoDeDato{Denominacion: "Longitud de Onda", NombreUnidadMedidad: "m", ValorUmbral: 200}

	// Estaciones
	estacionNorte := &modelo.EstacionSismologica{Nombre: "Estación Norte"}
	estacionSur := &modelo.EstacionSismologica{Nombre: "Estación Sur"}

	// Series temporales
	modelo.SerieTemporal1 = modelo.GenerarSerieTemporal(tipoVelocidad, tipoFrecuencia, tipoLongitud, time.Now())
	modelo.SerieTemporal2 = modelo.GenerarSerieTemporal(tipoVelocidad, tipoFrecuencia, tipoLongitud, time.Now())

	// Sismógrafos
	sismografo1 := modelo.NewSismografo(time.Now().AddDate(-2, 0, 0), 1, "SN123", modelo.SerieTemporal1, estacionNorte)
	sismografo2 := modelo.NewSismografo(time.Now().AddDate(-1, 0, 0), 2, "SS456", modelo.SerieTemporal2, estacionSur)

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
	evento1 := modelo.NewEventoSismico(0, time.Now().Add(-time.Hour*4), 900.0, 20.0, 50.0, 3.0, sesionActual, clasificaciones[0], origenDeGeneracion[0], alcanceSismo[0])
	evento2 := modelo.NewEventoSismico(1, time.Now().Add(-time.Hour*2), 500.0, 350.0, 100.0, 2.5, sesionActual, clasificaciones[1], origenDeGeneracion[1], alcanceSismo[1])
	evento3 := modelo.NewEventoSismico(2, time.Now().Add(-time.Hour), 150.0, 125.0, 150.0, 2.5, sesionActual, clasificaciones[1], origenDeGeneracion[1], alcanceSismo[1])
	evento1.AddSerieTemporal(modelo.SerieTemporal1)
	evento2.AddSerieTemporal(modelo.SerieTemporal2)
	evento3.AddSerieTemporal(modelo.SerieTemporal2)

	gestorRegistrarResultado.SetSesionActual(&sesionActual)
	gestorRegistrarResultado.AddEvento(evento1)
	gestorRegistrarResultado.AddEvento(evento2)
	gestorRegistrarResultado.AddEvento(evento3)
	gestorRegistrarResultado.AddSismografo(sismografo1)
	gestorRegistrarResultado.AddSismografo(sismografo2)

	//--------------------------------------------------------------------------------------------------------------------

	// Ruta inicio
	r.GET("/", *gestorPagina.MostrarPaginaPrincipal(gestorRegistrarResultado))
	r.GET("/inicio", *gestorPagina.MostrarPaginaPrincipal(gestorRegistrarResultado))

	// Login
	r.GET("/login", *gestorPagina.MostrarLogin(gestorRegistrarResultado))

	// Procesar login
	r.POST("/login", *gestorPagina.HabilitarLogin(gestorRegistrarResultado))

	// Cerrar sesion
	r.GET("/cerrarsesion", *gestorPagina.HabilitarCerrarSesion(gestorRegistrarResultado), *gestorPagina.MostrarPaginaPrincipal(gestorRegistrarResultado))

	// Listar E.S.
	r.POST("/lista-es", *gestorRegistrarResultado.RegistrarResultado())
	r.POST("/lista-todos-es", *gestorRegistrarResultado.MostrarTodosEventos())

	// Revision manual.
	r.POST("/revision", *gestorPagina.MostrarRevisionManual(gestorRegistrarResultado))

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
