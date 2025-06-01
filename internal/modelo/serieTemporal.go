package modelo

import "time"

type SerieTemporal struct {
	CondicionAlarma                 bool
	FechaHoraInicioRegistroMuestras time.Time
	FechaHoraRegistro               time.Time
	FrecuenciaMuestreo              int
	MuestraSismica                  []MuestraSismica
}

func NewSerieTemporal(condicion bool, fechaHoraInicioRegistroMuestras time.Time, fechaHoraRegistro time.Time, frecuenciaMuestreo int) *SerieTemporal {
	return &SerieTemporal{
		CondicionAlarma:                 condicion,
		FechaHoraInicioRegistroMuestras: fechaHoraInicioRegistroMuestras,
		FechaHoraRegistro:               fechaHoraRegistro,
		FrecuenciaMuestreo:              frecuenciaMuestreo,
	}
}

func GenerarSerieTemporal(tipoVel, tipoFreq, tipoLong TipoDeDato, inicio time.Time) *SerieTemporal {
	serie := NewSerieTemporal(false, inicio, inicio, 1)

	// generar
	for i := 0; i < 3; i++ {
		fecha := inicio.Add(time.Duration(i) * time.Second)
		muestra := MuestraSismica{
			FechaHoraMuestra: fecha,
			DetalleMuestraSismica: []DetalleMuestraSismica{
				{Valor: 450 + float64(i*10), TipoDeDato: tipoVel},
				{Valor: 40 + float64(i), TipoDeDato: tipoFreq},
				{Valor: 180 + float64(i*5), TipoDeDato: tipoLong},
			},
		}
		serie.MuestraSismica = append(serie.MuestraSismica, muestra)
	}

	return serie
}
