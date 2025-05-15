package example

//	CRUD PROPIEDADES
//
// Crear una nueva propiedad
func CreatePropertyService(property *Property) error {
	return CreateProperty(property)
}

// Obtener una propiedad por ID
func GetPropertyByIDService(id uint) (*Property, error) {
	return GetPropertyByID(id)
}

// Obtener todos las propiedades
func GetAllPropertiesService() ([]Property, error) {
	return GetAllProperties()
}
func GetAllPropertiesByAccountIDService(accountId uint) ([]Property, error) {
	return GetAllPropertiesByAccountID(accountId)
}

// Borrar una propiedad por ID
func DeletePropertyByIDService(id uint) error {
	return DeletePropertyByID(id)
}

// Actualizar una propiedad
func UpdatePropertyByIDService(property *Property) error {
	return UpdateProperty(property)
}

/*------------------------------------------------------------------------*/
//	CRUD PRECIOS PROPIEDADES
//
// Crear precio
func CreatePriceService(price *Price) error {
	return CreatePrice(price)
}

// Obtener precio por ID
func GetPriceByIDService(id uint) (*Price, error) {
	return GetPriceByID(id)
}

// Obtener todos los precios
func GetAllPricesService() ([]Price, error) {
	return GetAllPrices()
}

// Borrar precio por ID
func DeletePriceByIDService(id uint) error {
	return DeletePriceByID(id)
}

// Actualizar precio
func UpdatePriceService(price *Price) error {
	return UpdatePrice(price)
}

/*------------------------------------------------------------------------*/
//	CRUD GRUPOS PROPIEDADES
//
// Crear grupo
func CreateGroupService(group *Group) error {
	return CreateGroup(group)
}

// Obtener grupo por ID
func GetGroupByIDService(id uint) (*Group, error) {
	return GetGroupByID(id)
}

// Obtener todos los grupos
func GetAllGroupsService() ([]Group, error) {
	return GetAllGroups()
}

// Borrar grupo por ID
func DeleteGroupByIDService(id uint) error {
	return DeleteGroupByID(id)
}

// Actualizar grupo
func UpdateGroupService(group *Group) error {
	return UpdateGroup(group)
}

/*------------------------------------------------------------------------*/
//	CRUD IMÁGENES PROPIEDADES
//
// Crear imagen
func CreateImageService(image *Image) error {
	return CreateImage(image)
}

// Obtener imagen por ID
func GetImageByIDService(id uint) (*Image, error) {
	return GetImageByID(id)
}

// Obtener todas las imágenes
func GetAllImagesService() ([]Image, error) {
	return GetAllImages()
}

// Borrar imagen por ID
func DeleteImageByIDService(id uint) error {
	return DeleteImageByID(id)
}

// Actualizar imagen
func UpdatePropertyImageService(image *Image) error {
	return UpdatePropertyImage(image)
}

/*------------------------------------------------------------------------*/
// CRUD TIPOS DE PROPIEDADES
//
// Crear tipo de propiedad
func CreatePropertyTypeService(propertyType *PropertyType) error {
	return CreatePropertyType(propertyType)
}

// Obtener tipo de propiedad por ID
func GetPropertyTypeByIDService(id uint) (*PropertyType, error) {
	return GetPropertyTypeByID(id)
}

// Obtener todos los tipos de propiedades
func GetAllPropertyTypesService() ([]PropertyType, error) {
	return GetAllPropertyTypes()
}

// Borrar tipo de propiedad por ID
func DeletePropertyTypeByIDService(id uint) error {
	return DeletePropertyTypeByID(id)
}

// Actualizar tipo de propiedad
func UpdatePropertyTypeService(propertyType *PropertyType) error {
	return UpdatePropertyType(propertyType)
}

/*------------------------------------------------------------------------*/
// CRUD TIPOS DE TIEMPO
//
// Crear tipo de tiempo
func CreateTimeTypeService(timeType *TimeType) error {
	return CreateTimeType(timeType)
}

// Obtener tipo de tiempo pot ID
func GettimeTypeByIDService(id uint) (*TimeType, error) {
	return GettimeTypeByID(id)
}

// Obtener todos los tipos de tiempo
func GetAllTimeTypesService() ([]TimeType, error) {
	return GetAllTimeTypes()
}

// Borrar tipo de tiempo por ID
func DeleteTimeTypeByIDService(id uint) error {
	return DeleteTimeTypeByID(id)
}

// Actualizar tipo de tiempo
func UpdateTimeTypeService(timeType *TimeType) error {
	return UpdateTimeType(timeType)
}

/*------------------------------------------------------------------------*/
