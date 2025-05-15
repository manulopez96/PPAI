package login

import (
	"net/http"
	"ppai/internal/empleado"
	"ppai/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("login")
	{
		r.POST("/", log)

	}
}

func log(c *gin.Context) {
	var req struct {
		Useremail string `json:"useremail"`
		Password  string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos incorrectos"})
		return
	}

	user, err := empleado.GetAccountByEmail(req.Useremail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email no existente", "req.email": req.Useremail, "err": err})
		return
	}
	if req.Useremail != user.Email || req.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}
	id := strconv.Itoa(int(user.ID))
	token, err := auth.GenerarToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
