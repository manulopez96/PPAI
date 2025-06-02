package modelo

import "time"

type Sismografo struct{
	FechaAdquisicion time.Time
	Identificador int
	NroSerie string
	SerieTemporal []*SerieTemporal
	EstacionSismologica *EstacionSismologica
}

func NewSismografo(fechaAdquisicion time.Time, id int, nroSerie string, serieTemporal *SerieTemporal, estacion *EstacionSismologica) *Sismografo{
	serieTemporalArray := []*SerieTemporal{serieTemporal}
	return &Sismografo{
		FechaAdquisicion: fechaAdquisicion,
		Identificador: id,
		NroSerie: nroSerie,
		SerieTemporal: serieTemporalArray,
		EstacionSismologica: estacion,
	}
}

func (s *Sismografo) ContieneSerieTemporal(serie *SerieTemporal) bool {
	for _, sSerie := range s.SerieTemporal {
		if sSerie == serie {
			return true
		}
	}
	return false
}