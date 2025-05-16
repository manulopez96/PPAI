package object

import (
	"strconv"
	"time"
)

type EventoSismico struct {
	FechaHoraFin        time.Time
	fechaHoraOcurrencia time.Time
	latitudEpicentro    float64
	longitudEpicentro   float64
	hipocentro          float64
	valorMagnitud       float64
	analistaSupervisor  Empleado
	clasificacion       ClasificacionSismo
}

func NewEventoSismico(
	fechaHoraOcurrencia time.Time,
	latitudEpicentro float64,
	longitudEpicentro float64,
	hipocentro float64,
	valorMagnitud float64,
	analistaSupervisor Empleado,
	clasificacion ClasificacionSismo) *EventoSismico {
	return &EventoSismico{
		fechaHoraOcurrencia: fechaHoraOcurrencia,
		latitudEpicentro:    latitudEpicentro,
		longitudEpicentro:   longitudEpicentro,
		hipocentro:          hipocentro,
		valorMagnitud:       valorMagnitud,
		analistaSupervisor:  analistaSupervisor,
		clasificacion:       clasificacion,
	}
}

func (e *EventoSismico) GetValorMagnitud() float64 {
	alcance := e.valorMagnitud
	return alcance
}
func (e *EventoSismico) GetFechaHoraOcurrencia() time.Time {
	return e.fechaHoraOcurrencia
}
func (e *EventoSismico) GetFecha() string {
	return e.fechaHoraOcurrencia.Format("2006-01-02")
}
func (e *EventoSismico) GetHora() string {
	return e.fechaHoraOcurrencia.Format("15:04:05")
}
func (e *EventoSismico) GetLatitudEpicentro() float64 {
	return e.latitudEpicentro
}

func (e *EventoSismico) GetLongitudEpicentro() float64 {
	return e.longitudEpicentro
}

func (e *EventoSismico) GetHipocentro() float64 {
	return e.hipocentro
}
func (e *EventoSismico) String() string {
	return "\nEvento Sismico: " + e.fechaHoraOcurrencia.String() + "\nLatitud epicentro: " + strconv.FormatFloat(e.latitudEpicentro, 'f', 2, 64) + "\nLongitud epicentro:  " + strconv.FormatFloat(e.longitudEpicentro, 'f', 2, 64) + "\nHipocentro:  " + strconv.FormatFloat(e.hipocentro, 'f', 2, 64) + "\nAnalista supervisor: " + e.analistaSupervisor.Nombre + " " + e.analistaSupervisor.Apellido + "\nValor magnitud: " + strconv.FormatFloat(e.valorMagnitud, 'f', 2, 64) + "\n"
}

func (e *EventoSismico) CardDatos() []ESCard {

	titulos := []string{
		"Fecha y hora",
		"Latitud epicentro",
		"Longitud epicentro",
		"Hipocentro",
		"Valor magnitud",
		"Analista supervisor nombre",
		"Analista supervisor apellido",
		"Clasificación",
	}
	datos := []string{
		e.fechaHoraOcurrencia.Format("2006-01-02 15:04:05"),
		strconv.FormatFloat(e.latitudEpicentro, 'f', 2, 64),
		strconv.FormatFloat(e.longitudEpicentro, 'f', 2, 64),
		strconv.FormatFloat(e.hipocentro, 'f', 2, 64),
		strconv.FormatFloat(e.valorMagnitud, 'f', 2, 64),
		e.analistaSupervisor.Nombre,
		e.analistaSupervisor.Apellido,
		e.clasificacion.GetNombre(),
	}

	var pares []ESCard
	for i := range titulos {
		pares = append(pares, ESCard{
			Titulo: titulos[i],
			Dato:   datos[i],
		})
	}
	return pares
}

type ESCard struct {
	Titulo string
	Dato   string
}

func (e *EventoSismico) GetCalificacion(hipocentro float64, c ClasificacionSismo) bool {
	return c.EsClasificacion(hipocentro)
}

// TODO: Implementar GetOrigen
// TODO: Implementar GetAlcance
// TODO: Implementar relacion magnitud
// TODO: Implementar relacion clasificacion
// TODO: Implementar relacion origenGeneracion
// TODO: Implementar relacion alcanceSismo
// TODO: Implementar relacion estadoActual
// TODO: Implementar relacion cambioEstado
// TODO: Implementar relacion serieTemporal

// Fecha y hora de ocurrencia.
// Hipocentro: el punto en la profundidad de la tierra desde donde se origina el sismo.
// Clasificación: basada en la profundidad a la que se origina (superficial, intermedio, profundo).
// Epicentro: la latitud y longitud geográfica en la superficie terrestre directamente encima del hipocentro. El sistema estima la localización (epicentro) a partir del procesamiento de datos sísmicos con Machine Learning.
// Origen de generación: por ejemplo, sismo interplaca, volcánico, provocado por explosiones de minas, etc..
// Alcance: la distancia epicentral entre el epicentro y el punto de observación (Estación Sismológica). Se clasifica como local, regional o telesismo.
// Magnitud: un parámetro que caracteriza el tamaño y la energía liberada por un sismo. El sistema estima la magnitud mediante el procesamiento de datos sísmicos con Machine Learning. Si la magnitud estimada es mayor o igual a 4.0 en la escala Richter, el evento se registra como auto confirmado; si es menor, se registra como auto detectado. La Magnitud local, conocida como Magnitud Richter, es un tipo de magnitud utilizada.
// El sistema utiliza un proceso automático basado en Machine Learning para fusionar datos de sismógrafos cercanos y estimar la localización (epicentro) y la magnitud del evento sísmico
