package system

import (
	"github.com/google/uuid"
	"oms-fiber/global"
)

type SysUser struct {
	global.GModel
	UserID      uint      `gorm:"primarykey"`                                                                                                                         // 主键ID
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                                                                   // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                                                                // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                                                                           // 用户登录密码
	NickName    string    `json:"nick_name" gorm:"default:系统用户;comment:用户昵称"`                                                                                         // 用户昵称
	SideMode    string    `json:"side_mode" gorm:"default:dark;comment:用户侧边主题"`                                                                                       // 用户侧边主题
	HeaderImg   string    `json:"header_img" gorm:"default:https://img0.baidu.com/it/u=3427037332,481490977&fm=253&fmt=auto&app=138&f=JPEG?w=190&h=190;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"base_color" gorm:"default:#fff;comment:基础颜色"`                                                                                        // 基础颜色
	ActiveColor string    `json:"active_color" gorm:"default:#1890ff;comment:活跃颜色"`                                                                                   // 活跃颜色
	AuthorityId uint      `json:"authority_id" gorm:"default:888;comment:用户角色ID"`                                                                                     // 用户角色ID
	//Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	//Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone  string `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email  string `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
