package system

import (
	"oms-fiber/global"
	"oms-fiber/model/system"
	"oms-fiber/utils"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	//if nil == global.GVA_DB {
	//	return nil, fmt.Errorf("db not init")
	//}

	var user system.SysUser
	err = global.Gorm.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}
