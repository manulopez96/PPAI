package object

type AlcanceSismo struct {
	Nombre      string
	Descripcion string
}

func NewAlcanceSismo(nombre, descripcion string) AlcanceSismo {
	return AlcanceSismo{
		Nombre:      nombre,
		Descripcion: descripcion,
	}
}

func (a *AlcanceSismo) SetNombre(nombre string) {
	a.Nombre = nombre
}
func (a *AlcanceSismo) GetNombre() string {
	return a.Nombre
}
func (a *AlcanceSismo) SetDescripcion(descripcion string) {
	a.Descripcion = descripcion
}
func (a *AlcanceSismo) GetDescripcion() string {
	return a.Descripcion
}
