package modelo

type OrigenDeGeneracion struct {
	Nombre      string
	Descripcion string
}

func NewOrigenDeGeneracion(nombre, descripcion string) OrigenDeGeneracion {
	return OrigenDeGeneracion{
		Descripcion: descripcion,
		Nombre:      nombre,
	}
}

func (o *OrigenDeGeneracion) SetNombre(nombre string) {
	o.Nombre = nombre
}
func (o *OrigenDeGeneracion) GetNombre() string {
	return o.Nombre
}
func (o *OrigenDeGeneracion) SetDescripcion(descripcion string) {
	o.Descripcion = descripcion
}
func (o *OrigenDeGeneracion) GetDescripcion() string {
	return o.Descripcion
}

func GetOrigenMuestra() []OrigenDeGeneracion {
	return []OrigenDeGeneracion{
		NewOrigenDeGeneracion("Tectonico", "Movimiento de placas tectonicas"),
		NewOrigenDeGeneracion("Volcanico", "Actividad volcanica"),
		NewOrigenDeGeneracion("Colapso", "Colapso de cavernas o minas"),
		NewOrigenDeGeneracion("Artificial", "Actividad humana"),
		NewOrigenDeGeneracion("Desconocido", "Origen desconocido"),
	}
}
