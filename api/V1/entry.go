package V1

type EntryGroup struct {
	Admin      AdminApi
	Permission PermissionApi
	Role       RoleApi
}

var Entry = new(EntryGroup)
