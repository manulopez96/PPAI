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
func Bloqueado() Estado {
	return NewEstado("Evento sismico", "Bloqueado")
}
func Rechazado() Estado {
	return NewEstado("Evento sismico", "Rechazado")
}
func Derivado() Estado {
	return NewEstado("Evento sismico", "Derivado")
}
func Aceptado() Estado {
	return NewEstado("Evento sismico", "Aceptado")
}
func PendienteDeCierre() Estado {
	return NewEstado("Evento sismico", "Pendiente de cierre")
}
func Cerrado() Estado {
	return NewEstado("Evento sismico", "Cerrado")
}
func SinRevision() Estado {
	return NewEstado("Evento sismico", "Sin revision")
}
func GetEstadosMuestra() []Estado {
	return []Estado{
		NewEstado("Evento sismico", "Auto Confirmado"),       // [0]
		NewEstado("Evento sismico", "Auto Detectado"),        // [1]
		NewEstado("Evento sismico", "Pendiente de revision"), // [2]
		NewEstado("Evento sismico", "Bloqueado"),             // [3]
		NewEstado("Evento sismico", "Rechazado"),             // [4]
		NewEstado("Evento sismico", "Derivado"),              // [5]
		NewEstado("Evento sismico", "Aceptado"),              // [6]
		NewEstado("Evento sismico", "Pendiente de cierre"),   // [7]
		NewEstado("Evento sismico", "Cerrado"),               // [8]
		NewEstado("Evento sismico", "Sin revision"),          // [9]
	}
}
