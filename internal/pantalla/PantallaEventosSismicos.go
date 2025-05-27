package pantalla

import (
	"fmt"
	"net/http"
	"ppai/internal/gestor"
	"ppai/internal/modelo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Pantalla struct {
	gestor *gestor.GestorEventosSismicos
}

func NewPantalla(g *gestor.GestorEventosSismicos) *Pantalla {
	return &Pantalla{gestor: g}
}

var inicio func(c *gin.Context)

func (p *Pantalla) Principal(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
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
func (p *Pantalla) Login(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
	// Login
	login := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"empleado": gestor.SesionActual.Nombre,
			"templ":    "login",
		})
	}
	return &login
}
func (p *Pantalla) PostLogin(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
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
func (p *Pantalla) CerrarSesion(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
	// Cerrar sesión
	postLogin := func(c *gin.Context) {
		imprimirSesion(*gestor.SesionActual)
		gestor.CerrarSesionActual()
		fmt.Println("Cerrando sesión")
		imprimirSesion(*gestor.SesionActual)
	}
	return &postLogin
}
func (p *Pantalla) CrearEventosAleatorios(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
	crearEventosAleatorios := func(c *gin.Context) {
		if gestor.SesionActual.Nombre == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"empleado": gestor.SesionActual.Nombre,
				"templ":    "login",
			})
			return
		}
		gestor.GenerarEventoSismicoAleatorio(c.PostForm("sim-tipo"))
		fmt.Println("Evento generado:", gestor.GetUltimoEvento().String())
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Simulacion evento sismico aleatorio",
			"cardTitle":     gestor.GetCantidadEventos(),
			"eventoSismico": gestor.GetUltimoEvento().GetCardEventoSismico(),
			"empleado":      gestor.SesionActual.Nombre,
			"templ":         "sim-es-a",
		})
		fmt.Println("Evento sismico aleatorio:", gestor.GetUltimoEvento())
	}
	return &crearEventosAleatorios
}
func (p *Pantalla) ListarEventos(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
	listarEventos := func(c *gin.Context) {
		if gestor.SesionActual.Nombre == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"empleado": gestor.SesionActual.Nombre,
				"templ":    "login",
			})
			return
		}
		if gestor.GetCantidadEventos() == 0 {
			inicio(c)
			return
		}
		filtro := c.PostForm("filter-list")
		fmt.Println("filtro: " + filtro)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":              "Listado de Eventos Sismicos",
			"cardsEventoSismico": gestor.GetCardEventosSismicos(filtro),
			"empleado":           gestor.SesionActual.Nombre,
			"templ":              "list-es",
			"filtroEstado":       filtro,
		})

	}
	return &listarEventos
}
func (p *Pantalla) RevisionManual(gestor *gestor.GestorEventosSismicos) *func(c *gin.Context) {
	revisionManual := func(c *gin.Context) {
		accion := c.PostForm("accion")
		fmt.Println("accion: " + accion)
		idString := c.PostForm("index")
		id, _ := strconv.Atoi(idString)

		if gestor.SesionActual.Nombre == "" {
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
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.PendienteDeCierre(), *gestor.SesionActual, time.Now()) // Notificado, estado: pendiente de cierre
			fmt.Println("Evento sismico notificado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.PendienteDeCierre().NombreEstado,
			})
			return
		}
		if accion == "cerrar" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.Cerrado(), *gestor.SesionActual, time.Now()) // Estado: cerrado
			fmt.Println("Evento sismico cerrado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.Cerrado().NombreEstado,
			})
			return
		}
		if accion == "anular" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.SinRevision(), *gestor.SesionActual, time.Now()) // Estado: Sin Revision
			fmt.Println("Evento sismico anulado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.SinRevision().NombreEstado,
			})
			return
		}
		if accion == "rechazado" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.Rechazado(), *gestor.SesionActual, time.Now()) // Estado: Rechazado
			fmt.Println("Evento sismico rechazado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.Rechazado().NombreEstado,
			})
			return
		}
		if accion == "derivado" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.Derivado(), *gestor.SesionActual, time.Now()) // Estado: Derivado
			fmt.Println("Evento sismico derivado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.Derivado().NombreEstado,
			})
			return
		}
		if accion == "aceptado" {
			gestor.GetEventoPorID(id).SetEstadoActual(modelo.Aceptado(), *gestor.SesionActual, time.Now()) // Estado: Aceptado
			fmt.Println("Evento sismico aceptado: " + idString)
			c.HTML(http.StatusOK, "auto-post.html", gin.H{
				"targetURL":  "/list-es",
				"filterList": modelo.Aceptado().NombreEstado,
			})
			return
		}
		gestor.GetEventoPorID(id).SetEstadoActual(modelo.Bloqueado(), *gestor.SesionActual, time.Now()) // Estado: Bloqueado

		fmt.Println("Revision evento sismico: " + c.PostForm("index"))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":             "Revision de Evento Sismico ",
			"cardEventoSismico": gestor.GetEventoPorID(id).GetCardEventoSismico(),
			"origenGeneracion":  modelo.GetOrigenMuestra(),
			"alcanceSismo":      modelo.GetAlcanceMuestra(),
			"empleado":          gestor.SesionActual.Nombre,
			"templ":             "review-es",
			"index":             idString,
		})
	}
	return &revisionManual
}
func imprimirSesion(s modelo.Empleado) {
	fmt.Println("Nombre:", s.Nombre)
	fmt.Println("Apellido:", s.Apellido)
	fmt.Println("Email:", s.Email)
	fmt.Println("Telefono:", s.Telefono)
}
