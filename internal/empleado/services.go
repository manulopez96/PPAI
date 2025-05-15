package empleado

import (
	"fmt"
	"strings"
)

// Crear un nuevo usuario
func CreateAccountService(user *User) error {
	if user.Email == "" {
		return fmt.Errorf("email no puede estar vacio")
	}
	user.Email = strings.ToLower(user.Email)
	return CreateAccount(user)
}

// Obtener todos los usuarios
func GetAllAccountsService() ([]User, error) {
	return GetAllAccounts()
}

// Obtener user
func GetAccountByIdService(id uint) (*User, error) {
	return GetAccountByID(id)
}

// Obtener user
func GetAccountByEmailService(email string) (*User, error) {
	return GetAccountByEmail(email)
}

// Actualizar user
func UpdateAccountService(user *User) error {
	if user.Email == "" {
		return fmt.Errorf("email no puede estar vacio")
	}
	user.Email = strings.ToLower(user.Email)
	return UpdateAccount(user)
}

// Borrar user
func DeleteAccountService(id uint) error {
	return DeleteAccountByID(id)
}

//---------------------------------------------------------------------------------------

// Crear datos completos de un user
func CreateDataUserService(dataUser *Profile) error {
	return CreateDataUser(dataUser)
}

// Obtener datos completos de un user por ID
func GetDataUserByIdService(id uint) (*Profile, error) {
	return GetDataUserByID(id)
}

// Actualizar datos completos de un user
func UpdateDataUserService(dataUser *Profile) error {
	return UpdateDataUser(dataUser)
}

// Borrar datos completos de un user
func DeleteDataUserService(id uint) error {
	return DeleteDataUserByID(id)
}

// Obtener todos los datos completos de los usuarios
func GetAllDataUsersService() ([]Profile, error) {
	return GetAllDataUsers()
}
