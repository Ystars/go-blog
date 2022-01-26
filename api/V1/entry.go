package V1

type EntryGroup struct {
	Admin      AdminApi
	Permission PermissionApi
}

var Entry = new(EntryGroup)
