// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package share_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EduSchoolStudent is the golang structure of table edu_school_student for DAO operations like Where/Data.
type EduSchoolStudent struct {
	g.Meta       `orm:"table:edu_school_student, do:true"`
	Id           interface{} // ID，保持与USERID一致
	No           interface{} // 学号
	Avatar       interface{} // 头像
	Name         interface{} // 姓名
	Mobile       interface{} // 手机号
	UnionMainId  interface{} // 所属主体
	State        interface{} // 状态： -1毕业，0待确认，1在读，2休学，3退学
	LastActiveIp interface{} // 最后活跃IP
	HiredAt      *gtime.Time // 入学时间
	CreatedBy    interface{} //
	CreatedAt    *gtime.Time //
	UpdatedBy    interface{} //
	UpdatedAt    *gtime.Time //
	DeletedBy    interface{} //
	DeletedAt    *gtime.Time //
	Sex          interface{} // 性别：0女 1男
}
