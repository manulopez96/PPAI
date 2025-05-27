package modelo

type ClasificacionSismo struct {
	kmProfundidadDesde float64
	kmProfundidadHasta float64
	nombre             string
}

func (c *ClasificacionSismo) GetNombre() string {
	return c.nombre
}
func NewClasificacionSismo(desde, hasta float64, nombre string) ClasificacionSismo {
	return ClasificacionSismo{
		kmProfundidadDesde: desde,
		kmProfundidadHasta: hasta,
		nombre:             nombre,
	}
}

func (c *ClasificacionSismo) EsClasificacion(hipocentro float64) bool {
	if hipocentro >= c.kmProfundidadDesde && hipocentro <= c.kmProfundidadHasta {
		return true
	}
	return false
}
