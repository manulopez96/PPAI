package object

type Estado struct {
	ambito       string
	nombreEstado string
}

func NewEstado(ambito, nombre string) Estado {
	return Estado{
		ambito:       ambito,
		nombreEstado: nombre,
	}
}

func (e *Estado) SetNombre(nombre string) {
	e.nombreEstado = nombre
}
func (e *Estado) GetNombre() string {
	return e.nombreEstado
}
func (e *Estado) SetAmbito(ambito string) {
	e.ambito = ambito
}
func (e *Estado) GetAmbito() string {
	return e.ambito
}
func GetEstadosMuestra() []Estado {
	return []Estado{
		NewEstado("Evento sismico", "Auto Confirmado"),			//[0]
		NewEstado("Evento sismico", "Auto Detectado"),			//[1]
		NewEstado("Evento sismico", "Pendiente de revision"),	//[2]
		NewEstado("Evento sismico", "Bloqueado"),				//[3]
		NewEstado("Evento sismico", "Rechazado"),				//[4]
		NewEstado("Evento sismico", "Derivado"),				//[5]
		NewEstado("Evento sismico", "Aceptado"),				//[6]
		NewEstado("Evento sismico", "Pendiente de cierre"),		//[7]
		NewEstado("Evento sismico", "Cerrado"),					//[8]
	}
}
