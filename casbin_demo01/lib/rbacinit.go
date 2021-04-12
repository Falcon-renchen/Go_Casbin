package lib

import "Go_Casbin/casbin_demo01/models"

type RoleRel struct {
	PRole string
	Role  string
}

func (this *RoleRel) String() string {
	return this.PRole + ":" + this.Role
}

//获取角色
func GetRoles(pid int, m *[]*RoleRel, pname string) {
	proles := make([]*models.Role, 0)
	Gorm.Where("role_pid=?", pid).Find(&proles)
	if len(proles) == 0 {
		return
	}
	for _, item := range proles {
		if pname != "" {
			*m = append(*m, &RoleRel{pname, item.RoleName})
		}
		GetRoles(item.RoleId, m, item.RoleName)
	}
}

func GetUserRoles() (users []*models.Users) {
	Gorm.Select("a.user_name,c.role_name").
		Table("users a,user_roles b,roles c").
		Where("a.user_id=b.user_id and b.role_id").
		Order("a.user_id desc").Find(&users)
	return
}
