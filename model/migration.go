package model

func migration() {
	var base Base

	var admin Admin

	local := base.GetDb()

	_ = local.AutoMigrate(&admin, &CasbinRule{}, &Article{}, &Comment{})

	result := local.Take(&admin)

	if result.RowsAffected != 0 {
		return
	}

	createAdmin := Admin{Username: "admin", Password: "$2a$10$YGL5a9z7ykG6BWOo.XhJU.h8r98BD5IvAmLISBB9rFIefbDzrv58O"}

	local.Create(&createAdmin)

	createCasbin := []CasbinRule{
		{Ptype: "p", Authority: "user", Path: "/admin/users", Method: "GET"},
		{Ptype: "g", Authority: "bob", Path: "user"},
	}

	local.Create(&createCasbin)
}
