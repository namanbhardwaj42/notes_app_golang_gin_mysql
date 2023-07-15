package models

func DB_CreateUser(name string, email string, password string) error {
	entry := UserS{Name: name, Email: email, Password: password}
	DB.Create(&entry)
	return DB.Error
}

func UserFind(id uint64) *UserS {
	var user UserS
	DB.Where("id = ?", id).First(&user)
	return &user
}

func UserCheck(email string) *UserS {
	var user UserS
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}
