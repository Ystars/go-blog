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
		{Ptype: "p", Authority: "1", Path: "/admin/users", Method: "GET"},
		{Ptype: "g", Authority: "user", Path: "1"},
	}

	local.Create(&createCasbin)
}
