package example

import "ppai/config"

//	CRUD PROPIEDADES
//
// Crear una nueva propiedad
func CreateProperty(property *Property) error {
	result := config.DB.Create(property)
	return result.Error
}

// Obtener una propiedad por ID
func GetPropertyByID(id uint) (*Property, error) {
	var property Property
	res := config.DB.
		Preload("Account").
		Preload("Group").
		Preload("PropertyType").
		Preload("Prices").
		Preload("Images").
		First(&property, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &property, nil
}

/*
	Preload("Group")
- Carga la relación con el grupo (Group) de la propiedad, si existe.
- Evita hacer múltiples consultas manuales a la base de datos.

	First(&property, id)
- Busca en la base de datos una propiedad cuyo ID coincida con el valor de id.
*/

// Obtener todos las propiedades
func GetAllProperties() ([]Property, error) {
	var properties []Property
	err := config.DB.
		Preload("Account").
		Preload("Group").
		Preload("PropertyType").
		Preload("Prices").
		Preload("Images").
		Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}
func GetAllPropertiesByAccountID(accountId uint) ([]Property, error) {
	var properties []Property
	err := config.DB.
		Where("account_id = ?", accountId).
		Preload("Account").
		Preload("Group").
		Preload("PropertyType").
		Preload("Prices").
		Preload("Images").
		Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

// Borrar una propiedad por ID
func DeletePropertyByID(id uint) error {
	return config.DB.Delete(&Property{}, id).Error
}

// Actualizar una propiedad
func UpdateProperty(property *Property) error {
	return config.DB.Save(property).Error
}

/*------------------------------------------------------------------------*/
//	CRUD PRECIOS PROPIEDADES
//
// Crear precio
func CreatePrice(price *Price) error {
	result := config.DB.Create(price)
	return result.Error
}

// Obtener precio por ID
func GetPriceByID(id uint) (*Price, error) {
	var Price Price
	err := config.DB.
		Preload("Account").
		First(&Price, id).Error
	if err != nil {
		return nil, err
	}
	return &Price, nil
}

// Obtener todos los precios
func GetAllPrices() ([]Price, error) {
	var Prices []Price
	err := config.DB.
		Preload("Account").
		Find(&Prices).Error
	if err != nil {
		return nil, err
	}
	return Prices, nil
}

// Borrar precio por ID
func DeletePriceByID(id uint) error {
	return config.DB.Delete(&Price{}, id).Error
}

// Actualizar precio
func UpdatePrice(price *Price) error {
	return config.DB.Save(price).Error
}

/*------------------------------------------------------------------------*/
//	CRUD GRUPOS PROPIEDADES
//
// Crear grupo
func CreateGroup(Group *Group) error {
	result := config.DB.Create(Group)
	return result.Error
}

// Obtener grupo por ID
func GetGroupByID(id uint) (*Group, error) {
	var Group Group
	err := config.DB.
		Preload("Account").
		First(&Group, id).Error
	if err != nil {
		return nil, err
	}
	return &Group, nil
}

// Obtener todos los grupos
func GetAllGroups() ([]Group, error) {
	var Group []Group
	err := config.DB.
		Preload("Account").
		Find(&Group).Error
	if err != nil {
		return nil, err
	}
	return Group, nil
}

// Borrar grupo por ID
func DeleteGroupByID(id uint) error {
	return config.DB.Delete(&Group{}, id).Error
}

// Actualizar grupo
func UpdateGroup(Group *Group) error {
	return config.DB.Save(Group).Error
}

/*------------------------------------------------------------------------*/
//	CRUD IMÁGENES PROPIEDADES
//
// Crear imagen
func CreateImage(image *Image) error {
	result := config.DB.Create(image)
	return result.Error
}

// Obtener imagen por ID
func GetImageByID(id uint) (*Image, error) {
	var Image Image
	err := config.DB.
		Preload("Account").
		First(&Image, id).Error
	if err != nil {
		return nil, err
	}
	return &Image, nil
}

// Obtener todas las imágenes
func GetAllImages() ([]Image, error) {
	var Images []Image
	err := config.DB.
		Preload("Account").
		Find(&Images).Error
	if err != nil {
		return nil, err
	}
	return Images, nil
}

// Borrar imagen por ID
func DeleteImageByID(id uint) error {
	return config.DB.Delete(&Image{}, id).Error
}

// Actualizar imagen
func UpdatePropertyImage(image *Image) error {
	return config.DB.Save(image).Error
}

/*------------------------------------------------------------------------*/
// CRUD TIPOS DE PROPIEDADES
//
// Crear tipo de propiedad
func CreatePropertyType(propertyType *PropertyType) error {
	result := config.DB.Create(propertyType)
	return result.Error
}

// Obtener tipo de propiedad por ID
func GetPropertyTypeByID(id uint) (*PropertyType, error) {
	var propertyType PropertyType
	err := config.DB.First(&propertyType, id).Error
	if err != nil {
		return nil, err
	}
	return &propertyType, nil
}

// Obtener todos los tipos de propiedades
func GetAllPropertyTypes() ([]PropertyType, error) {
	var propertyTypes []PropertyType
	err := config.DB.Find(&propertyTypes).Error
	if err != nil {
		return nil, err
	}
	return propertyTypes, nil
}

// Borrar tipo de propiedad por ID
func DeletePropertyTypeByID(id uint) error {
	return config.DB.Delete(&PropertyType{}, id).Error
}

// Actualizar tipo de propiedad
func UpdatePropertyType(propertyType *PropertyType) error {
	return config.DB.Save(propertyType).Error
}

/*------------------------------------------------------------------------*/
// CRUD TIPOS DE TIEMPO
//
// Crear tipo de tiempo
func CreateTimeType(timeType *TimeType) error {
	result := config.DB.Create(timeType)
	return result.Error
}

// Obtener tipo de tiempo pot ID
func GettimeTypeByID(id uint) (*TimeType, error) {
	var timeType TimeType
	err := config.DB.First(&timeType, id).Error
	if err != nil {
		return nil, err
	}
	return &timeType, nil
}

// Obtener todos los tipos de tiempo
func GetAllTimeTypes() ([]TimeType, error) {
	var timeTypes []TimeType
	err := config.DB.Find(&timeTypes).Error
	if err != nil {
		return nil, err
	}
	return timeTypes, nil
}

// Borrar tipo de tiempo por ID
func DeleteTimeTypeByID(id uint) error {
	return config.DB.Delete(&TimeType{}, id).Error
}

// Actualizar tipo de tiempo
func UpdateTimeType(timeType *TimeType) error {
	return config.DB.Save(timeType).Error
}

/*------------------------------------------------------------------------*/
