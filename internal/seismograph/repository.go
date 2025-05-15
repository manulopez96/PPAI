package seismograph

import "ppai/config"

// Crear una nueva cuenta
func CreateAccount(user *User) error {
	result := config.DB.Create(user)
	return result.Error
}

// Obtener una cuenta por ID
func GetAccountByID(id uint) (*User, error) {
	var user User
	result := config.DB.Preload("Profile").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Obtener una cuenta por email
func GetAccountByEmail(email string) (*User, error) {
	var user User
	result := config.DB.Where("email = ?", email).Preload("Profile").First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Obtener todos las cuentas
func GetAllAccounts() ([]User, error) {
	var users []User
	result := config.DB.Preload("Profile").Find(&users)
	return users, result.Error
}

// Borrar un usuario por ID
func DeleteAccountByID(id uint) error {
	return config.DB.Delete(&User{}, id).Error
}

// Actualizar cuenta
func UpdateAccount(user *User) error {
	return config.DB.Save(user).Error
}

//---------------------------------------------------------------------------------------

// crear datos completos de un usuario
func CreateDataUser(dataUser *Profile) error {
	result := config.DB.Create(dataUser)
	return result.Error
}

// Obtener datos completos de un usuario por ID
func GetDataUserByID(id uint) (*Profile, error) {
	var dataUser Profile
	result := config.DB.First(&dataUser, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataUser, nil
}

// Obtener datos completos de un usuario por su ID personal
func GetDataUserByPersonalID(personalId uint) (*Profile, error) {
	var dataUser Profile
	result := config.DB.Where("personal_id = ?", personalId).First(&dataUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataUser, nil
}

// Obtener todos los datos completos de los usuarios
func GetAllDataUsers() ([]Profile, error) {
	var dataUsers []Profile
	result := config.DB.Find(&dataUsers)
	return dataUsers, result.Error
}

// Borrar datos completos de un usuario por ID
func DeleteDataUserByID(id uint) error {
	return config.DB.Delete(&Profile{}, id).Error
}

// Actualizar datos completos de un usuario
func UpdateDataUser(dataUser *Profile) error {
	return config.DB.Save(dataUser).Error
}
