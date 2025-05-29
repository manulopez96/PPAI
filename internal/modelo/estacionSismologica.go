package modelo

type EstacionSismologica struct {
	CodigoEstacion string
	Latitud        float64
	Longitud       float64
	Nombre         string
}

func NewEstacionSismologica(codigoEstacion string, latitud float64, longitud float64, nombre string) *EstacionSismologica {
	return &EstacionSismologica{
		CodigoEstacion: codigoEstacion,
		Latitud:        latitud,
		Longitud:       longitud,
		Nombre:         nombre,
	}
}
