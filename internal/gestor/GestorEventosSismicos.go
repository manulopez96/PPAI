package gestor

import (
	"fmt"
	"math/rand"
	"ppai/internal/modelo"
	"time"
)

type GestorEventosSismicos struct {
	SesionActual *modelo.Empleado
	Eventos      []*modelo.EventoSismico
	Sismografos  []*modelo.Sismografo
	Estados      []*modelo.Estado
}

func NewGestorEventos() *GestorEventosSismicos {
	return &GestorEventosSismicos{
		Eventos: make([]*modelo.EventoSismico, 0),
	}
}
func (g *GestorEventosSismicos) SetSesionActual(sesion *modelo.Empleado) {
	g.SesionActual = sesion
}
func (g *GestorEventosSismicos) CerrarSesionActual() {
	g.SesionActual = &modelo.Empleado{}
}

func (g *GestorEventosSismicos) AddEvento(evento *modelo.EventoSismico) {
	g.Eventos = append(g.Eventos, evento)
}
func (g *GestorEventosSismicos) CrearEvento(id int, fecha time.Time, lat, lon, hipo, magn float64, emp modelo.Empleado, clas modelo.ClasificacionSismo, origen modelo.OrigenDeGeneracion, alcance modelo.AlcanceSismo) {
	evento := modelo.NewEventoSismico(id, fecha, lat, lon, hipo, magn, emp, clas, origen, alcance)
	g.Eventos = append(g.Eventos, evento)
}

func (g *GestorEventosSismicos) GetEventos() []*modelo.EventoSismico {
	return g.Eventos
}

func (g *GestorEventosSismicos) GetCantidadEventos() int {
	return len(g.Eventos)
}

func (g *GestorEventosSismicos) GetUltimoEvento() *modelo.EventoSismico {
	if len(g.Eventos) == 0 {
		return nil
	}
	return g.Eventos[len(g.Eventos)-1]
}
func (g *GestorEventosSismicos) ExisteSesionActiva() bool {
	return g.SesionActual != nil && g.SesionActual.Nombre != ""
}
func (g *GestorEventosSismicos) ExistenEventos() bool {
	return g.GetCantidadEventos() > 0
}

func (g *GestorEventosSismicos) GetEventoPorID(id int) *modelo.EventoSismico {
	for _, evento := range g.Eventos {
		if evento.GetId() == id {
			return evento
		}
	}
	return nil
}
func (g *GestorEventosSismicos) GetSeriesTemporales(id int) []*modelo.SerieTemporal {
	evento := g.GetEventoPorID(id)
	if evento != nil {
		return evento.GetSerieTemporal()
	}
	return nil

}
func (g *GestorEventosSismicos) BuscarDatosSismicosRegistrados(id int) modelo.ESCard {
	datos := g.GetEventoPorID(id).GetDatos()
	return datos
}
func (g *GestorEventosSismicos) BloquearEventoPorID(id int, responsable modelo.Empleado, fin time.Time) {
	estado := g.GetEstado("Evento sismico", "Bloqueado")
	g.Eventos[id].SetEstadoActual(*estado, responsable, fin)

}
func (g *GestorEventosSismicos) RechazarEventoPorID(id int, responsable modelo.Empleado, fin time.Time) {
	estado := g.GetEstado("Evento sismico", "Rechazado")
	g.Eventos[id].SetEstadoActual(*estado, responsable, fin)

}
func (g *GestorEventosSismicos) GetEstacionSismologica(serie *modelo.SerieTemporal) *modelo.EstacionSismologica {
	for _, sismografo := range g.Sismografos {
		if sismografo.ContieneSerieTemporal(serie) {
			return sismografo.EstacionSismologica
		}
	}
	return nil
}
func (g *GestorEventosSismicos) LlamarCUGenerarSismograma() {
	fmt.Println("Llamado al caso de uso: Generar Sismograma")

}
func (g *GestorEventosSismicos) GetCardEventosSismicos(estado string) []modelo.ESCard {
	cardEventosSismicos := []modelo.ESCard{}

	for _, evento := range g.Eventos {

		// t1 := evento.GetFechaHoraOcurrencia()
		// diff := time.Since(t1)
		// if diff.Minutes() >= 5 && (evento.GetEstadoActual() == estados[1] || evento.GetEstadoActual() == estados[2]) {
		// 	evento.SetEstadoActual(estados[2], sesionActual) // pendiente revision
		// }

		if estado == "all" {
			cardEventosSismicos = append(cardEventosSismicos, evento.GetDatos())
		} else if estado == "Registrar resultado de revisión manual" {
			if evento.SosAutoDetectado() {
				cardEventosSismicos = append(cardEventosSismicos, evento.GetDatos())
			}
		} else if estado == evento.GetEstadoActual().NombreEstado {
			cardEventosSismicos = append(cardEventosSismicos, evento.GetDatos())
		}
	}
	return cardEventosSismicos
}
func (g *GestorEventosSismicos) AddSismografo(sismografo *modelo.Sismografo) {
	g.Sismografos = append(g.Sismografos, sismografo)
}
func (g *GestorEventosSismicos) GetEstado(ambito, nombre string) *modelo.Estado {
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

func (g *GestorEventosSismicos) GenerarEventoSismicoAleatorio(tipo string) modelo.EventoSismico {

	clasificaciones := []modelo.ClasificacionSismo{}
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(0, 70, "Superficial"))
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(70, 300, "Intermedio"))
	clasificaciones = append(clasificaciones, modelo.NewClasificacionSismo(300, 700, "Profundo"))

	origenDeGeneracion := []modelo.OrigenDeGeneracion{
		modelo.NewOrigenDeGeneracion("Tectonico", "Movimiento de placas tectonicas"),
		modelo.NewOrigenDeGeneracion("Volcanico", "Actividad volcanica"),
		modelo.NewOrigenDeGeneracion("Colapso", "Colapso de cavernas o minas"),
		modelo.NewOrigenDeGeneracion("Artificial", "Actividad humana"),
		modelo.NewOrigenDeGeneracion("Desconocido", "Origen desconocido"),
	}

	alcanceSismo := []modelo.AlcanceSismo{
		modelo.NewAlcanceSismo("Sismo local", "Hasta 100 km"),
		modelo.NewAlcanceSismo("Sismo regional", "Hasta 1000 km"),
		modelo.NewAlcanceSismo("Tele sismo", "Mas de 1000 km"),
	}

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
	analistaSupervisor := *g.SesionActual
	clasificacion := clasificaciones[randomInt(len(clasificaciones))]
	origen := origenDeGeneracion[randomInt(len(origenDeGeneracion))]
	alcance := alcanceSismo[randomInt(len(alcanceSismo))]
	for _, clasificacionItem := range clasificaciones {
		if clasificacionItem.EsClasificacion(hipocentro) {
			clasificacion = clasificacionItem
		}
	}

	g.CrearEvento(g.GetCantidadEventos(), fechaHoraOcurrencia, latitudEpicentro, longitudEpicentro, hipocentro, valorMagnitud, analistaSupervisor, clasificacion, origen, alcance)
	g.GetUltimoEvento().AddSerieTemporal(modelo.SerieTemporal1)
	return *g.GetUltimoEvento()
}
func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
func floatAleatorio(limite int) float64 {
	rand.Seed(time.Now().UnixNano())

	num := float64(rand.Intn(limite)) + rand.Float64()
	return num
}
