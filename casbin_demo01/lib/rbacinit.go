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
