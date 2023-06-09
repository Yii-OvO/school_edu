// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package share_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EduSchoolStudent is the golang structure for table edu_school_student.
type EduSchoolStudent struct {
	Id           int64       `json:"id"           description:"ID，保持与USERID一致"`
	No           string      `json:"no"           description:"学号"`
	Avatar       string      `json:"avatar"       description:"头像"`
	Name         string      `json:"name"         description:"姓名"`
	Mobile       string      `json:"mobile"       description:"手机号"`
	UnionMainId  int64       `json:"unionMainId"  description:"所属主体"`
	State        int         `json:"state"        description:"状态： -1毕业，0待确认，1在读，2休学，3退学"`
	LastActiveIp string      `json:"lastActiveIp" description:"最后活跃IP"`
	HiredAt      *gtime.Time `json:"hiredAt"      description:"入学时间"`
	CreatedBy    int64       `json:"createdBy"    description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedBy    int64       `json:"updatedBy"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	DeletedBy    int64       `json:"deletedBy"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`
	Sex          int         `json:"sex"          description:"性别：0女 1男"`
}