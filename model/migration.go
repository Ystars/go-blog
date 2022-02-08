package model

import "goblog/utils/scrypt"

func migration() {
	var admin Admin

	local := DB

	_ = local.AutoMigrate(&admin, &CasbinRule{}, &Article{}, &Comment{})

	result := local.Take(&admin)

	if result.RowsAffected != 0 {
		return
	}

	createAdmin := Admin{Username: "admin", Password: scrypt.NewScrypt().GeneratePassword("123456")}

	local.Create(&createAdmin)

	createCasbin := []CasbinRule{
		{Ptype: "p", Authority: "1", Path: "/admin/users", Method: "GET"},
		{Ptype: "g", Authority: "user", Path: "1"},
	}

	local.Create(&createCasbin)
}
