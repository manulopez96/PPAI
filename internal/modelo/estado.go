package modelo

type Estado struct {
	ambito       string
	NombreEstado string
}

func NewEstado(ambito, nombre string) Estado {
	return Estado{
		ambito:       ambito,
		NombreEstado: nombre,
	}
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
func (e *Estado) EsEstado(estado string) bool {
	return e.NombreEstado == estado
}

func AutoConfirmado() Estado {
	return NewEstado("Evento sismico", "Auto Confirmado")
}
func AutoDetectado() Estado {
	return NewEstado("Evento sismico", "Auto Detectado")
}
func PendienteDeRevision() Estado {
	return NewEstado("Evento sismico", "Pendiente de revision")
}
func BloquearEvento() Estado {
	return NewEstado("Evento sismico", "Bloqueado")
}
func RechazarEvento() Estado {
	return NewEstado("Evento sismico", "Rechazado")
}
func DerivarEvento() Estado {
	return NewEstado("Evento sismico", "Derivado")
}
func AceptarEvento() Estado {
	return NewEstado("Evento sismico", "Aceptado")
}
func PendienteDeCierre() Estado {
	return NewEstado("Evento sismico", "Pendiente de cierre")
}
func CerrarEvento() Estado {
	return NewEstado("Evento sismico", "Cerrado")
}
func SinRevision() Estado {
	return NewEstado("Evento sismico", "Sin revision")
}
