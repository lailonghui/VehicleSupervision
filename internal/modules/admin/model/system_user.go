package model

import "time"

type SystemUser struct {
	// 是否绑定IP
	IsBindIP *bool `json:"Is_bind_ip"`
	// 客户端版本号
	AppVersion *string `json:"app_version"`
	// 审核等级
	AuditLevel *int `json:"audit_level"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 创建时间
	CreatedAt *time.Time `json:"created_at"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 部门ID
	DepartmentID *string `json:"department_id"`
	// 邮箱
	Email *string `json:"email"`
	// 企业ID
	EnterpriseID *string `json:"enterprise_id"`
	// 级别
	Grade *int `json:"grade"`
	// ID
	ID int64 `json:"id"`
	// ip地址
	IPAddress *string `json:"ip_address"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 是否有效
	IsValid *bool `json:"is_valid"`
	// 手机串号
	Mkey *string `json:"mkey"`
	// 手机号码
	Mobile *string `json:"mobile"`
	// 密码
	Password  string  `json:"password"`
	ProxyUser *string `json:"proxy_user"`
	// 备注
	Remarks *string `json:"remarks"`
	// 电话号码
	Telephone *string `json:"telephone"`
	// 用户名
	Username string `json:"username"`
	// 加密串码
	Ukey *string `json:"ukey"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 用户ID
	UserID string `json:"user_id"`
	// 状态
	UserState *int `json:"user_state"`
	// 用户类型
	UserType *int `json:"user_type"`
}

func (s SystemUser) TableName() string {
	return "system_user"
}
