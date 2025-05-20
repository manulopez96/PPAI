package object

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
