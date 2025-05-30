package modelo

import (
	"strconv"
	"time"
)

type EventoSismico struct {
	id                  int
	FechaHoraFin        time.Time
	fechaHoraOcurrencia time.Time
	latitudEpicentro    float64
	longitudEpicentro   float64
	hipocentro          float64
	valorMagnitud       float64
	analistaSupervisor  Empleado
	clasificacion       ClasificacionSismo
	origenDeGeneracion  OrigenDeGeneracion
	alcanceSismo        AlcanceSismo
	estadoActual        Estado
	estado              []CambioEstado
	SerieTemporal       []SerieTemporal
}

func NewEventoSismico(
	id int,
	fechaHoraOcurrencia time.Time,
	latitudEpicentro float64,
	longitudEpicentro float64,
	hipocentro float64,
	valorMagnitud float64,
	analistaSupervisor Empleado,
	clasificacion ClasificacionSismo,
	origenDeGeneracion OrigenDeGeneracion,
	alcanceSismo AlcanceSismo,
) *EventoSismico {
	estadoInicial := AutoDetectado()
	if valorMagnitud >= 4.0 {
		estadoInicial = AutoConfirmado()
	}
	var estado []CambioEstado
	estado = append(estado, NewCambioEstado(estadoInicial, analistaSupervisor, fechaHoraOcurrencia))
	return &EventoSismico{
		id:                  id,
		fechaHoraOcurrencia: fechaHoraOcurrencia,
		latitudEpicentro:    latitudEpicentro,
		longitudEpicentro:   longitudEpicentro,
		hipocentro:          hipocentro,
		valorMagnitud:       valorMagnitud,
		analistaSupervisor:  analistaSupervisor,
		clasificacion:       clasificacion,
		origenDeGeneracion:  origenDeGeneracion,
		alcanceSismo:        alcanceSismo,
		estadoActual:        estadoInicial,
		estado:              estado,
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

func (e *EventoSismico) SetAnalistaSupervisor(a Empleado) {
	e.analistaSupervisor = a
}
func (e *EventoSismico) GetAnalistaSupervisor() Empleado {
	return e.analistaSupervisor
}
func (e *EventoSismico) SetClasificacion(c ClasificacionSismo) {
	e.clasificacion = c
}
func (e *EventoSismico) GetClasificacion() ClasificacionSismo {
	return e.clasificacion
}
func (e *EventoSismico) SetOrigenDeGeneracion(o OrigenDeGeneracion) {
	e.origenDeGeneracion = o
}
func (e *EventoSismico) GetOrigenDeGeneracion() OrigenDeGeneracion {
	return e.origenDeGeneracion
}
func (e *EventoSismico) SetAlcanceSismo(a AlcanceSismo) {
	e.alcanceSismo = a
}
func (e *EventoSismico) GetAlcanseSismo() AlcanceSismo {
	return e.alcanceSismo
}
func (e *EventoSismico) GetId() int {
	return e.id
}

func (e *EventoSismico) SetEstadoActual(estado Estado, responsable Empleado, fin time.Time) {
	i := len(e.estado) - 1
	e.estado[i].SetFechaHoraFin(&fin)
	e.estado = append(e.estado, NewCambioEstado(estado, responsable, fin))
	e.estadoActual = estado
}
func (e *EventoSismico) GetEstadoActual() Estado {
	return e.estadoActual
}
func (e *EventoSismico) GetCambioDeEstado() []CambioEstado {
	return e.estado
}

func (e *EventoSismico) SosAutoDetectado() bool {
	return e.estadoActual.EsEstado("Auto Detectado")
}

func (e *EventoSismico) String() string {
	return "\nEvento Sismico: " + e.fechaHoraOcurrencia.String() + "\nLatitud epicentro: " + strconv.FormatFloat(e.latitudEpicentro, 'f', 2, 64) + "\nLongitud epicentro:  " + strconv.FormatFloat(e.longitudEpicentro, 'f', 2, 64) + "\nHipocentro:  " + strconv.FormatFloat(e.hipocentro, 'f', 2, 64) + "\nAnalista supervisor: " + e.analistaSupervisor.Nombre + " " + e.analistaSupervisor.Apellido + "\nValor magnitud: " + strconv.FormatFloat(e.valorMagnitud, 'f', 2, 64) + "\n"
}

func (e *EventoSismico) GetCardEventoSismico() ESCard {

	var cardEstados []CECard
	for _, cambioEstado := range e.estado {
		cardEstados = append(cardEstados, cambioEstado.GetCardCambioEstado())
	}

	cardEventoSismico := ESCard{
		// Id:                         strconv.Itoa(e.id),
		Id:                         (e.id),
		FechaHoraOcurrencia:        e.fechaHoraOcurrencia.Format("2006-01-02 15:04:05"),
		LatitudEpicentro:           strconv.FormatFloat(e.latitudEpicentro, 'f', 2, 64),
		LongitudEpicentro:          strconv.FormatFloat(e.longitudEpicentro, 'f', 2, 64),
		Hipocentro:                 strconv.FormatFloat(e.hipocentro, 'f', 2, 64),
		ValorMagnitud:              strconv.FormatFloat(e.valorMagnitud, 'f', 2, 64),
		AnalistaSupervisorNombre:   e.analistaSupervisor.Nombre,
		AnalistaSupervisorApellido: e.analistaSupervisor.Apellido,
		Clasificacion:              e.clasificacion.GetNombre(),
		OrigenDeGeneracion:         e.origenDeGeneracion.GetNombre(),
		AlcanceSismo:               e.alcanceSismo.GetNombre(),
		EstadoActual:               e.estadoActual.GetNombre(),
		Estado:                     cardEstados,
	}
	return cardEventoSismico
}

type ESCard struct {
	Id                         int
	FechaHoraFin               string
	FechaHoraOcurrencia        string
	LatitudEpicentro           string
	LongitudEpicentro          string
	Hipocentro                 string
	ValorMagnitud              string
	AnalistaSupervisorNombre   string
	AnalistaSupervisorApellido string
	Clasificacion              string
	OrigenDeGeneracion         string
	AlcanceSismo               string
	EstadoActual               string
	Estado                     []CECard
}

func (e *EventoSismico) GetCalificacion(hipocentro float64, c ClasificacionSismo) bool {
	return c.EsClasificacion(hipocentro)
}

// TODO: Implementar relacion magnitud
// TODO: Implementar relacion serieTemporal

// Fecha y hora de ocurrencia.
// Hipocentro: el punto en la profundidad de la tierra desde donde se origina el sismo.
// Clasificación: basada en la profundidad a la que se origina (superficial, intermedio, profundo).
// Epicentro: la latitud y longitud geográfica en la superficie terrestre directamente encima del hipocentro. El sistema estima la localización (epicentro) a partir del procesamiento de datos sísmicos con Machine Learning.
// Origen de generación: por ejemplo, sismo interplaca, volcánico, provocado por explosiones de minas, etc..
// Alcance: la distancia epicentral entre el epicentro y el punto de observación (Estación Sismológica). Se clasifica como local, regional o telesismo.
// Magnitud: un parámetro que caracteriza el tamaño y la energía liberada por un sismo. El sistema estima la magnitud mediante el procesamiento de datos sísmicos con Machine Learning. Si la magnitud estimada es mayor o igual a 4.0 en la escala Richter, el evento se registra como auto confirmado; si es menor, se registra como auto detectado. La Magnitud local, conocida como Magnitud Richter, es un tipo de magnitud utilizada.
// El sistema utiliza un proceso automático basado en Machine Learning para fusionar datos de sismógrafos cercanos y estimar la localización (epicentro) y la magnitud del evento sísmico
