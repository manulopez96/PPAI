package pantalla

import (
	"net/http"
	"ppai/internal/modelo"

	"github.com/gin-gonic/gin"
)

type PantallaRegistrarResultado struct {
	listEventosSismicos  []*modelo.ESString
	listSeriesTemporales []*modelo.SerieTemporal
}

func NewPantallaRegistrarResultado() *PantallaRegistrarResultado {
	return &PantallaRegistrarResultado{
		listEventosSismicos:  make([]*modelo.ESString, 0),
		listSeriesTemporales: make([]*modelo.SerieTemporal, 0),
	}
}

func (p *PantallaRegistrarResultado) PresentarEventos(listaEventosSismicos []modelo.ESString, sesion string, c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":              "Listado de Eventos Sismicos",
		"cardsEventoSismico": listaEventosSismicos,
		"empleado":           sesion,
		"templ":              "lista-es",
	})
}
