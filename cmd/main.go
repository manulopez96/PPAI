package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"ppai/config"
	"ppai/internal/empleado"
	"ppai/object"
	"ppai/pkg/login"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func imprimirSesion(s object.Empleado) {
	fmt.Println("Nombre:", s.Nombre)
	fmt.Println("Apellido:", s.Apellido)
	fmt.Println("Emanil:", s.Email)
	fmt.Println("Telefono:", s.Telefono)
}

var sesionActual object.Empleado
var clasificaciones []object.ClasificacionSismo

func main() {

	// Codigo Hardcodeado para pruebas
	eventosSismicos := []object.EventoSismico{}

	clasificaciones = []object.ClasificacionSismo{}
	clasificaciones = append(clasificaciones, object.NewClasificacionSismo(0, 70, "Superficial"))
	clasificaciones = append(clasificaciones, object.NewClasificacionSismo(70, 300, "Intermedio"))
	clasificaciones = append(clasificaciones, object.NewClasificacionSismo(300, 700, "Profundo"))

	r := gin.Default() // Crea una instancia del router con middlewares por defecto
	// gin.SetMode(gin.ReleaseMode)

	// Configuración de la base de datos
	config.ConnectDB()
	config.DB.AutoMigrate(&empleado.User{})

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

	//Cargando rutas para modificacion datos en DB
	empleado.RegisterRoutes(r)
	login.RegisterRoutes(r)

	// Index principal, donde se cargan todas las plantillas
	inicio := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Proyecto PPAI",
			"templ":    "principal",
			"sesion":   sesionActual,
			"empleado": sesionActual.Nombre,
		})
	}

	// Ruta inicio
	r.GET("/", inicio)
	r.GET("/inicio", inicio)

	// Login template
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"empleado": sesionActual.Nombre,
			"templ":    "login",
		})
	})
	// Procesar login
	r.POST("/login", func(c *gin.Context) {
		sesionActual.Nombre = c.PostForm("nombre")
		sesionActual.Apellido = c.PostForm("apellido")
		sesionActual.Email = c.PostForm("email")
		sesionActual.Telefono = c.PostForm("telefono")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"mensaje":  "Formulario enviado correctamente",
			"nombre":   sesionActual.Nombre,
			"apellido": sesionActual.Apellido,
			"email":    sesionActual.Email,
			"telefono": sesionActual.Telefono,
			"templ":    "login",
		})
		imprimirSesion(sesionActual)

	})

	// Cerrar sesion
	r.GET("/cerrarsesion", func(c *gin.Context) {
		imprimirSesion(sesionActual)
		sesionActual = object.Empleado{}
		fmt.Println("Cerrando sesión")
		imprimirSesion(sesionActual)
	}, inicio)

	// Crear E.S.
	r.POST("/sim-es-a", func(c *gin.Context) {
		eventosSismicos = append(eventosSismicos, generarEventoSismicoAleatorio(c.PostForm("sim-tipo")))
		fmt.Println("Evento generado:", eventosSismicos[len(eventosSismicos)-1].String())
		ultimoEvento := eventosSismicos[len(eventosSismicos)-1]
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":                      "Simulacion evento sismico aleatorio",
			"cardTitle":                  len(eventosSismicos),
			"eventoSismico":              ultimoEvento.CardDatos(),
			"eventoSismicoValorMagnitud": ultimoEvento.GetValorMagnitud(),
			"eventoSismicoHipocentro":    ultimoEvento.GetHipocentro(),
			"empleado":                   sesionActual.Nombre,
			"templ":                      "sim-es-a",
		})
	})

	// Listar E.S.
	r.GET("/list-es", func(c *gin.Context) {

		if sesionActual.Nombre == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"empleado": sesionActual.Nombre,
				"templ":    "login",
			})
		} else if len(eventosSismicos) == 0 {
			inicio(c)
		} else {
			cardEventosSismicos := make([][]object.ESCard, len(eventosSismicos))
			fmt.Println("lista eventos sismicos")
			for i, evento := range eventosSismicos {
				cardEventosSismicos[i] = evento.CardDatos()
			}
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title":              "Listado de Eventos Sismicos",
				"cardsEventoSismico": cardEventosSismicos,
				"empleado":           sesionActual.Nombre,
				"templ":              "list-es",
			})
		}
	})

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}

func generarEventoSismicoAleatorio(tipo string) object.EventoSismico {

	var magnitudMaxima int
	var magnitudMinima int
	if tipo == "aleatorio" {
		magnitudMaxima = 7
		magnitudMinima = 0
	} else if tipo == "mayor4.0" {
		magnitudMaxima = 3
		magnitudMinima = 4
	} else if tipo == "menor4.0" {
		magnitudMaxima = 3
		magnitudMinima = 0
	}
	// Generar un evento sísmico aleatorio
	fechaHoraOcurrencia := time.Now()
	latitudEpicentro := floatAleatorio(2000)
	longitudEpicentro := floatAleatorio(2000)
	hipocentro := floatAleatorio(700)
	valorMagnitud := floatAleatorio(magnitudMaxima) + float64(magnitudMinima)
	analistaSupervisor := sesionActual
	clasificacion := clasificaciones[0]
	for _, clasificacionItem := range clasificaciones {
		if clasificacionItem.EsClasificacion(hipocentro) {
			clasificacion = clasificacionItem
		}
	}
	// if tipo == "aleatorio" {	}
	return *object.NewEventoSismico(fechaHoraOcurrencia, latitudEpicentro, longitudEpicentro, hipocentro, valorMagnitud, analistaSupervisor, clasificacion)
}

func floatAleatorio(limite int) float64 {
	rand.Seed(time.Now().UnixNano())

	num := float64(rand.Intn(limite)) + rand.Float64()
	return num
}

// eq (equal) → {{if eq .Estado "activo"}}
// ne (not equal)
// lt (less than)
// gt (greater than)
// le (less or equal)
// ge (greater or equal)
// not
