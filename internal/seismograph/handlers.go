package seismograph

import (
	"fmt"
	"net/http"
	"ppai/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("api/users", auth.MiddlewareJWT())
	{
		r.GET("/", getAllAccountHandler)
		r.PUT("/", updateAccountHandler)
		r.GET("/:id", getAccountHandler)
		r.DELETE("/:id", deleteAccountHandler)

		r.GET("/data/", getAllDataUsersHandler)
		r.PUT("/data/", updateDataUserHandler)
		r.GET("/data/:id", getDataUserHandler)
		r.DELETE("/data/:id", deleteDataUserHandler)
	}
	s := router.Group("api/users")
	{
		s.POST("/", createAccountHandler)
		
		s.POST("/data/", createDataUserHandler)
	}
}

func createAccountHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Active = true
	if err := CreateAccountService(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func getAllAccountHandler(c *gin.Context) {
	permissionAny, exists := c.Get("userpermission")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso no encontrado en context"})
		return
	}
	strPermition := fmt.Sprintf("%v", permissionAny)
	permission, err := strconv.Atoi(strPermition)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso sin formato"})
		return
	}
	users, err := GetAllAccountsService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fallo el servicio"})
		return
	}
	if permission >= 5 {
		c.JSON(http.StatusOK, &users)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
}

func getAccountHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	id, permission := getDatosJWT(c)
	if idParam == id || permission >= 5 {
		user, err := GetAccountByIdService(uint(idParam))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User no encontrada"})
			return
		}
		c.JSON(http.StatusOK, &user)

	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}

}

func updateAccountHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, permission := getDatosJWT(c)
	if (user.ID == uint(id) || permission >= 5) && user.ID != 0 {
		if err := UpdateAccountService(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User actualizado correctamente"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}

}

func deleteAccountHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	id, permission := getDatosJWT(c)
	if idParam == id || permission >= 5 {
		err := DeleteAccountService(uint(idParam))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User no encontrada"})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("User %d eliminado", idParam))

	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permiso insuficiente", "Permiso": permission})
		return
	}
}

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

//--------------------------------------------------------------------------------------------

// crear datos completos de un usuario
func createDataUserHandler(c *gin.Context) {
	var userData Profile
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateDataUserService(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &userData)
}

// obtener datos completos de un usuario
func getDataUserHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	userData, err := GetDataUserByIdService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Datos completos no encontrados"})
		return
	}
	c.JSON(http.StatusOK, &userData)
}

// actualizar datos completos de un usuario
func updateDataUserHandler(c *gin.Context) {
	c.String(200, "updateDataUserHandler()\n")
	var userData Profile
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := UpdateDataUserService(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Datos completos actualizados correctamente"})
}

// borrar datos completos de un usuario
func deleteDataUserHandler(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = DeleteDataUserService(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Datos completos no encontrados"})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Datos completos %d eliminados", idParam))
}

// obtener todos los datos completos de los useusuarios
func getAllDataUsersHandler(c *gin.Context) {
	usersData, err := GetAllDataUsersService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fallo el servicio"})
		return
	}
	c.JSON(http.StatusOK, &usersData)
}
