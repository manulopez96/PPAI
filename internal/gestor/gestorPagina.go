package gestor

import (
	"fmt"
	"net/http"
	"ppai/internal/modelo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GestorPagina struct {
	gestor *GestorRegistrarResultado
}

func NewGestorPagina(g *GestorRegistrarResultado) *GestorPagina {
	return &GestorPagina{gestor: g}
}

var inicio func(c *gin.Context)

func (p *GestorPagina) MostrarPaginaPrincipal(gestor *GestorRegistrarResultado) *func(c *gin.Context) {
	// Index principal, donde se cargan todas las plantillas
	inicio = func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Proyecto PPAI",
			"templ":    "principal",
			"sesion":   gestor.SesionActual,
			"empleado": gestor.SesionActual.Nombre,
		})
	}
	return &inicio
}
func (p *GestorPagina) MostrarLogin(gestor *GestorRegistrarResultado) *func(c *gin.Context) {
	// Login
	login := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"empleado": gestor.SesionActual.Nombre,
			"templ":    "login",
		})
	}
	return &login
}
func (p *GestorPagina) HabilitarLogin(gestor *GestorRegistrarResultado) *func(c *gin.Context) {
	// Procesar login
	postLogin := func(c *gin.Context) {
		nombre := c.PostForm("nombre")
		apellido := c.PostForm("apellido")
		email := c.PostForm("email")
		telefono := c.PostForm("telefono")

		nuevaSesion := modelo.Empleado{
			Nombre:   nombre,
			Apellido: apellido,
			Email:    email,
			Telefono: telefono,
		}
		gestor.SetSesionActual(&nuevaSesion)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"mensaje":  "Formulario enviado correctamente",
			"nombre":   nombre,
			"apellido": apellido,
			"email":    email,
			"telefono": telefono,
			"templ":    "login",
		})
		imprimirSesion(nuevaSesion)

	}
	return &postLogin
}
func (p *GestorPagina) HabilitarCerrarSesion(gestor *GestorRegistrarResultado) *func(c *gin.Context) {
	// Cerrar sesión
	postLogin := func(c *gin.Context) {
		imprimirSesion(*gestor.SesionActual)
		gestor.CerrarSesionActual()
		fmt.Println("Cerrando sesión")
		imprimirSesion(*gestor.SesionActual)
	}
	return &postLogin
}
func imprimirSesion(s modelo.Empleado) {
	fmt.Println("Nombre:", s.Nombre)
	fmt.Println("Apellido:", s.Apellido)
	fmt.Println("Email:", s.Email)
	fmt.Println("Telefono:", s.Telefono)
}

func (p *GestorPagina) MostrarRevisionManual(gestor *GestorRegistrarResultado) *func(c *gin.Context) {
	revisionManual := func(c *gin.Context) {
		accion := c.PostForm("accion")
		fmt.Println("accion: " + accion)
		idString := c.PostForm("index")
		id, _ := strconv.Atoi(idString)

		if !gestor.ExisteSesionActiva() {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"empleado": gestor.SesionActual.Nombre,
				"templ":    "login",
			})
			return
		}
		if gestor.GetCantidadEventos() < id {
			inicio(c)
			return
		}
		if accion == "notificar" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.GetEstadoPendienteDeCierre(), *gestor.SesionActual, time.Now()) // Notificado, estado: pendiente de cierre
			fmt.Println("Evento sismico notificado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoPendienteDeCierre().NombreEstado,
			})
			return
		}
		if accion == "cerrar" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.GetEstadoCerrado(), *gestor.SesionActual, time.Now()) // Estado: cerrado
			fmt.Println("Evento sismico cerrado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoCerrado().NombreEstado,
			})
			return
		}
		if accion == "anular" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.GetEstadoSinRevision(), *gestor.SesionActual, time.Now()) // Estado: Sin Revision
			fmt.Println("Evento sismico anulado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoSinRevision().NombreEstado,
			})
			return
		}
		if accion == "rechazado" {
			gestor.RechazarEventoPorID(id, *gestor.SesionActual, time.Now()) // Estado: Rechazado
			fmt.Println("Evento sismico rechazado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoRechazado().NombreEstado,
			})
			return
		}
		if accion == "derivado" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.GetEstadoDerivado(), *gestor.SesionActual, time.Now()) // Estado: Derivado
			fmt.Println("Evento sismico derivado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoDerivado().NombreEstado,
			})
			return
		}
		if accion == "aceptado" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.GetEstadoAceptado(), *gestor.SesionActual, time.Now()) // Estado: Aceptado
			fmt.Println("Evento sismico aceptado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL": "/lista-es",
				"opcion":    modelo.GetEstadoAceptado().NombreEstado,
			})
			return
		}
		if accion == "revisar" {
			gestor.BloquearEventoPorID(id, *gestor.SesionActual, time.Now()) // Estado: Bloqueado

			fmt.Println("Revision evento sismico: " + c.PostForm("index"))

			datosEventoSismico := gestor.BuscarDatosSismicosRegistrados(id)
			datosSeriesTemporales := gestor.GetSeriesTemporalesEvento(id)
			estacionSismologica := gestor.GetEstacionSismologica(datosSeriesTemporales[0])
			gestor.LlamarCUGenerarSismograma()

			fmt.Println("------------------------------------------------------------------------------------------------")
			fmt.Println("Estacion Sismologica: ", estacionSismologica)

			c.HTML(http.StatusOK, "index.html", gin.H{
				"title":               "Revision de Evento Sismico ",
				"cardEventoSismico":   datosEventoSismico,
				"seriesTemporales":    datosSeriesTemporales,
				"estacionSismologica": estacionSismologica,
				"origenGeneracion":    modelo.GetOrigenMuestra(),
				"alcanceSismo":        modelo.GetAlcanceMuestra(),
				"empleado":            gestor.SesionActual.Nombre,
				"templ":               "revision",
				"index":               idString,
			})
		}
	}
	return &revisionManual
}
