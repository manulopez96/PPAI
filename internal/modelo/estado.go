package modelo

type Estado struct {
	ambito       string
	NombreEstado string
}

const AutoDetectado = "Auto Detectado"
const AutoConfirmado = "Auto Confirmado"

func NewEstado(ambito, nombre string) *Estado {
	return &Estado{
		ambito:       ambito,
		NombreEstado: nombre,
	}
}

func (e *Estado) EsAmbito(ambito string) bool {
	return e.ambito == ambito
}
func (e *Estado) EsAmbitoEventoSismico() bool {
	return e.ambito == "Evento sismico"
}

func (e *Estado) EsEstado(estado string) bool {
	return e.NombreEstado == estado
}
func (e *Estado) EsAutodetectado() bool {
	return e.NombreEstado == AutoDetectado
}
func (e *Estado) EsAutoConfirmado() bool {
	return e.NombreEstado == AutoConfirmado
}

func (e *Estado) SetNombre(nombre string) {
	e.NombreEstado = nombre
}
func (e *Estado) GetNombre() string {
	return e.NombreEstado
}
func (e *Estado) SetAmbito(ambito string) {
	e.ambito = ambito
}
func (e *Estado) GetAmbito() string {
	return e.ambito
}

func GetEstadoAutoConfirmado() Estado {
	return *NewEstado("Evento sismico", "Auto Confirmado")
}
func GetEstadoAutoDetectado() Estado {
	return *NewEstado("Evento sismico", "Auto Detectado")
}
func GetEstadoPendienteDeRevision() Estado {
	return *NewEstado("Evento sismico", "Pendiente de revision")
}
func GetEstadoBloqueado() Estado {
	return *NewEstado("Evento sismico", "Bloqueado")
}
func GetEstadoRechazado() Estado {
	return *NewEstado("Evento sismico", "Rechazado")
}
func GetEstadoDerivado() Estado {
	return *NewEstado("Evento sismico", "Derivado")
}
func GetEstadoAceptado() Estado {
	return *NewEstado("Evento sismico", "Aceptado")
}
func GetEstadoPendienteDeCierre() Estado {
	return *NewEstado("Evento sismico", "Pendiente de cierre")
}
func GetEstadoCerrado() Estado {
	return *NewEstado("Evento sismico", "Cerrado")
}
func GetEstadoSinRevision() Estado {
	return *NewEstado("Evento sismico", "Sin revision")
}
