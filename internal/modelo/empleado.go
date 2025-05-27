package modelo

type Empleado struct {
	// DNI       string `json:"dni" gorm:"not null;unique"`
	Nombre   string `json:"nombre" gorm:"not null"`
	Apellido string `json:"apellido" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Telefono string `json:"telefono" gorm:"not null"`
}

func (e *Empleado) GetEmail() string {
	return e.Email
}
