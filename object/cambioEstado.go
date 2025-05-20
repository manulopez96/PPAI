package object

import "time"

type CambioEstado struct {
	fechaHoraInicio       time.Time
	fechaHoraFin          *time.Time
	estado                Estado
	responsableInspeccion Empleado
}

func NewCambioEstado(e Estado, ri Empleado) CambioEstado {
	return CambioEstado{
		fechaHoraInicio:       time.Now(),
		fechaHoraFin:          nil,
		estado:                e,
		responsableInspeccion: ri,
	}
}

func (c *CambioEstado) GetFechaHoraInicio() time.Time {
	return c.fechaHoraInicio
}
func (c *CambioEstado) SetFechaHoraInicio(t time.Time) {
	c.fechaHoraInicio = t
}
func (c *CambioEstado) GetFechaHoraFin() *time.Time {
	return c.fechaHoraFin
}
func (c *CambioEstado) SetFechaHoraFin(t *time.Time) {
	c.fechaHoraFin = t
}
func (c *CambioEstado) GetEstado() Estado {
	return c.estado
}
func (c *CambioEstado) SetEstado(e Estado) {
	c.estado = e
}
func (c *CambioEstado) GetResponsableInspeccion() Empleado {
	return c.responsableInspeccion
}
func (c *CambioEstado) SetResponsableInspeccion(e Empleado) {
	c.responsableInspeccion = e
}

func (c *CambioEstado) GetCardCambioEstado() CECard {
	fin := ""
	if c.GetFechaHoraFin() == nil {
		fin = "Vigente"
	} else {
		fin = c.GetFechaHoraFin().Format("2006-01-02 15:04:05")
	}
	return CECard{
		FechaHoraInicio:       c.fechaHoraInicio.Format("2006-01-02 15:04:05"),
		FechaHoraFin:          fin,
		Estado:                c.estado.GetNombre(),
		ResponsableInspeccion: c.responsableInspeccion.Nombre,
	}
}

type CECard struct {
	FechaHoraInicio       string
	FechaHoraFin          string
	Estado                string
	ResponsableInspeccion string
}
