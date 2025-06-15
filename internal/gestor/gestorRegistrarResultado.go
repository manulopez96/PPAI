package gestor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ppai/internal/modelo"
	"ppai/internal/pantalla"
	"time"
)

var pantallaRegistrarResultado *pantalla.PantallaRegistrarResultado

type GestorRegistrarResultado struct {
	SesionActual *modelo.Empleado
	Eventos      []*modelo.EventoSismico
	Sismografos  []*modelo.Sismografo
	Estados      []*modelo.Estado
}

func NewGestorEventos() *GestorRegistrarResultado {
	return &GestorRegistrarResultado{
		Eventos: make([]*modelo.EventoSismico, 0),
	}
}
func (g *GestorRegistrarResultado) SetSesionActual(sesion *modelo.Empleado) {
	g.SesionActual = sesion
}
func (g *GestorRegistrarResultado) CerrarSesionActual() {
	g.SesionActual = &modelo.Empleado{}
}

func (g *GestorRegistrarResultado) AddEvento(evento *modelo.EventoSismico) {
	g.Eventos = append(g.Eventos, evento)
}
func (g *GestorRegistrarResultado) CrearEvento(id int, fecha time.Time, lat, lon, hipo, magn float64, emp modelo.Empleado, clas modelo.ClasificacionSismo, origen modelo.OrigenDeGeneracion, alcance modelo.AlcanceSismo) {
	evento := modelo.NewEventoSismico(id, fecha, lat, lon, hipo, magn, emp, clas, origen, alcance)
	g.Eventos = append(g.Eventos, evento)
}

func (g *GestorRegistrarResultado) GetEventos() []*modelo.EventoSismico {
	return g.Eventos
}

func (g *GestorRegistrarResultado) GetCantidadEventos() int {
	return len(g.Eventos)
}

func (g *GestorRegistrarResultado) GetUltimoEvento() *modelo.EventoSismico {
	if len(g.Eventos) == 0 {
		return nil
	}
	return g.Eventos[len(g.Eventos)-1]
}
func (g *GestorRegistrarResultado) ExisteSesionActiva() bool {
	return g.SesionActual != nil && g.SesionActual.Nombre != ""
}
func (g *GestorRegistrarResultado) ExistenEventos() bool {
	return g.GetCantidadEventos() > 0
}

func (g *GestorRegistrarResultado) GetEventoPorID(id int) *modelo.EventoSismico {
	for _, evento := range g.Eventos {
		if evento.GetId() == id {
			return evento
		}
	}
	return nil
}
func (g *GestorRegistrarResultado) GetSeriesTemporalesEvento(id int) []*modelo.SerieTemporal {
	evento := g.GetEventoPorID(id)
	if evento != nil {
		return evento.GetSerieTemporal()
	}
	return nil

}
func (g *GestorRegistrarResultado) BuscarDatosSismicosRegistrados(id int) modelo.ESString {
	datos := g.GetEventoPorID(id).GetDatos()
	return datos
}
func (g *GestorRegistrarResultado) BloquearEventoPorID(id int, responsable modelo.Empleado, fin time.Time) {
	estado := g.GetEstado("Evento sismico", "Bloqueado")
	g.Eventos[id].SetEstadoActual(*estado, responsable, fin)

}
func (g *GestorRegistrarResultado) RechazarEventoPorID(id int, responsable modelo.Empleado, fin time.Time) {
	estado := g.GetEstado("Evento sismico", "Rechazado")
	g.Eventos[id].SetEstadoActual(*estado, responsable, fin)

}
func (g *GestorRegistrarResultado) GetEstacionSismologica(serie *modelo.SerieTemporal) *modelo.EstacionSismologica {
	for _, sismografo := range g.Sismografos {
		if sismografo.ContieneSerieTemporal(serie) {
			return sismografo.EstacionSismologica
		}
	}
	return nil
}
func (g *GestorRegistrarResultado) LlamarCUGenerarSismograma() {
	fmt.Println("Llamado al caso de uso: Generar Sismograma")
}
func (g *GestorRegistrarResultado) GetListEventosSismicos() []modelo.ESString {
	listEventosSismicos := []modelo.ESString{}
	for _, evento := range g.Eventos {
		listEventosSismicos = append(listEventosSismicos, evento.GetDatos())
	}
	return listEventosSismicos
}
func (g *GestorRegistrarResultado) BuscarEventosAutoDetectados() []modelo.ESString {
	eventosAutoDetectados := []modelo.ESString{}
	for _, evento := range g.Eventos {
		if evento.SosAutoDetectado() {
			eventosAutoDetectados = append(eventosAutoDetectados, evento.GetDatos())
		}
	}
	return eventosAutoDetectados
}
func (g *GestorRegistrarResultado) AddSismografo(sismografo *modelo.Sismografo) {
	g.Sismografos = append(g.Sismografos, sismografo)
}
func (g *GestorRegistrarResultado) GetEstado(ambito, nombre string) *modelo.Estado {
	if len(g.Estados) == 0 {
		return nil
	}
	for _, estado := range g.Estados {
		if estado.EsAmbito(ambito) && estado.EsEstado(nombre) {
			return estado
		}
	}
	return nil
}

func (g *GestorRegistrarResultado) MostrarTodosEventos() *func(c *gin.Context) {
	listarEventos := func(c *gin.Context) {
		listaEventosSismicos := g.GetListEventosSismicos()
		g.ordenarPorFechaYHora(listaEventosSismicos)
		sesion := g.SesionActual.Nombre

		pantallaRegistrarResultado.PresentarEventos(listaEventosSismicos, sesion, c)
	}
	return &listarEventos
}

func (g *GestorRegistrarResultado) ordenarPorFechaYHora(eventos []modelo.ESString) {
	for i := 0; i < len(eventos)-1; i++ {
		for j := 0; j < len(eventos)-i-1; j++ {
			if eventos[j].FechaHoraOcurrencia > eventos[j+1].FechaHoraOcurrencia {
				eventos[j], eventos[j+1] = eventos[j+1], eventos[j]
				fmt.Println("Intercambiando eventos:", eventos[j], "y", eventos[j+1])
			}
		}
	}
}

func (g *GestorRegistrarResultado) RegistrarResultado() *func(c *gin.Context) {
	listarEventos := func(c *gin.Context) {
		sesion := g.SesionActual.Nombre

		listaEventosSismicos := g.BuscarEventosAutoDetectados()
		g.ordenarPorFechaYHora(listaEventosSismicos)

		pantallaRegistrarResultado.PresentarEventos(listaEventosSismicos, sesion, c)
	}
	return &listarEventos
}
