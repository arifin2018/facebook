package models

type RolePermission struct {
	RoleId       int `gorm:"column:Role_id" validate:"required"`
	PermissionId int `gorm:"column:Permissions_id" validate:"required"`
}

func (RolePermission) TableName() string {
	return "role_permisssions"
}
