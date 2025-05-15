package example

import (
	"gorm.io/gorm"
	"ppai/internal/empleado"
	"time"
)

type Property struct {
	gorm.Model
	AccountID      uint          `gorm:"not null;index"` // Relación con Account
	Account        empleado.User `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyTypeID uint          `gorm:"not null;index"` // Relación con PropertyType
	PropertyType   PropertyType  `gorm:"foreignKey:PropertyTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GroupID        *uint         `gorm:"index"` // Relación con PropertyGroup (puede ser NULL)
	Group          *Group        `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name           string        `gorm:"type:varchar(255);not null"`
	Location       string        `gorm:"type:varchar(255);"`
	Description    string        `gorm:"type:varchar(255)"`
	Prices         []Price       `gorm:"many2many:property_prices;"`
	Images         []Image       `gorm:"many2many:property_images;"`
}

type Price struct {
	gorm.Model
	AccountID   uint           `gorm:"not null"` // Clave foránea a Property
	Account     *empleado.User `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TypeTimeID  uint           `gorm:"not null"` // Clave foránea a TimeType
	TypeTime    TimeType       `gorm:"foreignKey:TypeTimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price       float64        `gorm:"not null"` // Precio de la propiedad
	Description string         `gorm:"type:varchar(255)"`
	MinRentTime uint           `gorm:"default:1"`
	MaxRentTime uint           `gorm:"default:100;check:max_rent_time >= min_rent_time"`
	StartDate   time.Time      `gorm:"not null"` // Fecha de inicio del precio
	EndDate     *time.Time     // Fecha de finalización (puede ser nula)
}
type Group struct {
	gorm.Model                // Incluye ID, CreatedAt, UpdatedAt y DeletedAt
	AccountID   uint          `gorm:"not null"` // Clave foránea a Account
	Account     empleado.User `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name        string        `gorm:"type:varchar(255);not null"`
	Description string        `gorm:"type:varchar(255)"`
}

type Image struct {
	gorm.Model
	AccountID uint          `gorm:"not null"` // Clave foránea a Account
	Account   empleado.User `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	URL       string        `gorm:"type:varchar(255);not null"` // URL de la imagen
}

type PropertyType struct {
	ID    uint   `gorm:"primaryKey"`
	Scope string `gorm:"not null"`
	Name  string `gorm:"unique;not null"` // Ejemplo: 'bedroom', 'bathroom', 'kitchen', 'living_room', 'garage'
}

type TimeType struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `gorm:"unique;not null"` // 'hourly', 'daily', 'weekly', 'monthly', 'yearly'
}

type RoomFeature struct {
	gorm.Model
	PropertyID  uint          `gorm:"not null;index"` // Relación con Property
	Description string        `gorm:"type:varchar(255)"`
	RoomTypeID  *uint         `gorm:""`
	RoomType    *PropertyType `gorm:"foreignKey:RoomTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SizeM2      int           `gorm:""`
}

// LandFeatures representa las características de un terreno
type ObjectFeatures struct {
	gorm.Model
	PropertyID  uint     `gorm:"not null;index"` // Relación con Property
	Property    Property `gorm:"foreignKey:PropertyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name        string   `gorm:"type:varchar(255);not null"`
	Description string   `gorm:"type:varchar(255)"`
	Quantity    int      `gorm:"not null;default: 1"`
}

type ParkingFeatures struct {
	gorm.Model
	PropertyID  uint     `gorm:"not null;index"` // Clave foránea referenciando Property
	Property    Property `gorm:"foreignKey:PropertyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlaceNumber int      `gorm:""` // Número de lugar de estacionamiento
	Description string
}

// Tabla Principal: VehicleFeatures (Vehículos)
type VehicleFeatures struct {
	gorm.Model
	PropertyID uint     `gorm:"not null;index"` // Relación con Property
	Property   Property `gorm:"foreignKey:PropertyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Brand      string   `gorm:"type:varchar(255);not null"`
	Modelo     string   `gorm:"type:varchar(255);not null"`
	Year       int      `gorm:"not null"`
	Color      string   `gorm:"type:varchar(100)"`
	Mileage    int      `gorm:"not null"`

	// Relación con FuelType
	FuelID uint     `gorm:"not null"`
	Type   FuelType `gorm:"foreignKey:FuelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	// Relación con TransmissionType
	TransmissionID uint             `gorm:"not null"`
	Transmission   TransmissionType `gorm:"foreignKey:TransmissionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	// Relación con ConditionType
	ConditionID uint          `gorm:"not null"`
	Condition   ConditionType `gorm:"foreignKey:ConditionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// Tabla FuelType (Combustible)
type FuelType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"` // Ejemplo: gasoline, diesel, electric, hybrid
}

// Tabla TransmissionType (Transmisión)
type TransmissionType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"` // Ejemplo: manual, automatic
}

// Tabla ConditionType (Condición del Vehículo)
type ConditionType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"` // Ejemplo: new, used, damaged
}
