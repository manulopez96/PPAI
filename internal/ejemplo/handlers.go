package example

import (
	"fmt"
	"net/http"
	"ppai/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("api/properties", auth.MiddlewareJWT())
	{
		r.GET("/all", getPropertiesOfAllAccountsHandler)
		r.POST("/", createPropertyHandler)
		r.PUT("/", updatePropertyByIDHandler)
		r.DELETE("/:id", deletePropertyByIDHandler)

		r.GET("/types", GetAllPropertyTypesHandler)
		r.POST("/types", CreatePropertyTypeHandler)
		r.PUT("/types", UpdatePropertyTypeHandler)
		r.DELETE("/types/:id", DeletePropertyTypeByIDHandler)
	}
	s := router.Group("api/properties")
	{
		s.GET("/:id", getPropertyByIDHandler)
		s.GET("/", getAllPropertyHandler)
	}
}

func createPropertyHandler(c *gin.Context) {
	var property Property
	id, permission := getDatosJWT(c)

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": "Property recuperado correctamente", "property": property})

	// Verificación de permisos
	if property.AccountID != uint(id) && permission < 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}

	// Insertar en la base de datos
	if err := CreatePropertyService(&property); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Propiedad creada correctamente", "property": property})
}

func getAllPropertyHandler(c *gin.Context) {
	idQuery := c.Query("AccountID")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Falta c.Query(AccountID)"})
		return
	}
	properties, err := GetAllPropertiesByAccountIDService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fallo el servicio"})
		return
	}
	c.JSON(http.StatusOK, &properties)
}

func getPropertiesOfAllAccountsHandler(c *gin.Context) {
	_, permission := getDatosJWT(c)
	if permission < 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}

	properties, err := GetAllPropertiesService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fallo el servicio"})
		return
	}
	c.JSON(http.StatusOK, &properties)
}

func getPropertyByIDHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	property, err := GetPropertyByIDService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property no encontrada"})
		return
	}
	c.JSON(http.StatusOK, &property)
}

func updatePropertyByIDHandler(c *gin.Context) {
	var property Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountID := property.AccountID
	id, permission := getDatosJWT(c)
	if (accountID != uint(id) && permission >= 5) || property.ID != 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}

	if err := UpdatePropertyByIDService(&property); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Property actualizado correctamente", "property": property})
}

func deletePropertyByIDHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	id, permission := getDatosJWT(c)
	if idParam != id && permission < 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
	err = DeletePropertyByIDService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property no encontrada"})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Property %d eliminado", idParam))

}

/*----------------------------------------------------------------------------------------------------------------*/

// CRUD TIPOS DE PROPIEDADES
//
// Crear tipo de propiedad
func CreatePropertyTypeHandler(c *gin.Context) {
	var propertyType PropertyType
	_, permission := getDatosJWT(c)
	if err := c.ShouldBindJSON(&propertyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if permission < 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
	if err := CreatePropertyTypeService(&propertyType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PropertyType creado correctamente", "propertyType": propertyType})
}

// Obtener tipo de propiedad por ID
func GetPropertyTypeByIDHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	propertyType, err := GetPropertyTypeByIDService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PropertyType no encontrada"})
		return
	}
	c.JSON(http.StatusOK, &propertyType)
}

// Obtener todos los tipos de propiedades
func GetAllPropertyTypesHandler(c *gin.Context) {
	propertyTypes, err := GetAllPropertyTypesService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fallo el servicio"})
		return
	}
	c.JSON(http.StatusOK, &propertyTypes)
}

// Borrar tipo de propiedad por ID
func DeletePropertyTypeByIDHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	_, permission := getDatosJWT(c)
	if permission >= 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
	err = DeletePropertyTypeByIDService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property no encontrada"})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Property %d eliminado", idParam))
}

// Actualizar tipo de propiedad
func UpdatePropertyTypeHandler(c *gin.Context) {
	var propertyType PropertyType
	if err := c.ShouldBindJSON(&propertyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, permission := getDatosJWT(c)
	if permission >= 5 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
	if err := UpdatePropertyTypeService(&propertyType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PropertyType actualizado correctamente", "propertyType": propertyType})
}

/*----------------------------------------------------------------------------------------------------------------*/

/*----------------------------------------------------------------------------------------------------------------*/

func getDatosJWT(c *gin.Context) (int, int) {
	permissionAny, exists := c.Get("userpermission")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso no encontrado en context"})
		return -1, -1
	}
	strPermition := fmt.Sprintf("%v", permissionAny)
	permission, err := strconv.Atoi(strPermition)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso sin formato"})
		return -1, -1
	}
	idAny, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Id no encontrado en context"})
		return -1, -1
	}
	strId := fmt.Sprintf("%v", idAny)
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Id sin formato"})
		return -1, -1
	}
	return id, permission
}
