package empleado

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model          // Agrega campos ID, CreatedAt, UpdatedAt, DeletedAt automáticamente
	Name      string   `gorm:"default:null"`
	LastName  string   `gorm:"default:null"`
	PersonalId string   `gorm:"type:varchar(255);not null;unique;index"` //DNI, legajo, algun numero de identificacion
	Email      string   `gorm:"index;unique;not null"`
	Password   string   `gorm:"not null"`
	Active     bool     `gorm:"default:true"`
	ProfileID  *uint    `gorm:"not null"`
	Profile    *Profile `gorm:"foreignKey:ProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Profile struct { //Datos completos utilizados para facturarle
	gorm.Model        // Incluye ID, CreatedAt, UpdatedAt y DeletedAt
	Name       string `gorm:"type:varchar(255);index;unique;not null"`
	Permission uint   `gorm:"default:null"`
}
