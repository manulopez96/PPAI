package modelo

import "time"

type MuestraSismica struct {
	FechaHoraMuestra time.Time
	DetalleMuestraSismica []DetalleMuestraSismica
}

type DetalleMuestraSismica struct {
	Valor float64
	TipoDeDato TipoDeDato
}

type TipoDeDato struct {
	Denominacion        string
	NombreUnidadMedidad string
	ValorUmbral         float64
}
