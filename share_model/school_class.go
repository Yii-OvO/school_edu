package share_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/edu-share/share_model/share_entity"
)

type Team struct {
	Id                int64  `json:"id"                dc:"ID"`
	Name              string `json:"name"              v:"required|max-length:128#名称不能为空|名称长度超128字符出限定范围" dc:"班级名称，学校维度下唯一"`
	OwnerEmployeeId   int64  `json:"ownerEmployeeId"   dc:"班主任/班长"`
	CaptainEmployeeId int64  `json:"captainEmployeeId" dc:"班长/小组组长"`
	ParentId          int64  `json:"parentId" dc:"班级或小组父级ID"`
	Remark            string `json:"remark"            dc:"备注"`
}

type TeamRes struct {
	share_entity.EduSchoolClass
	OwnerT    *TeacherRes `json:"ownerT" dc:"班主任"`
	OwnerS    *StudentRes `json:"ownerS" dc:"班长"`
	Captain   *StudentRes `json:"captain" dc:"班长/小组组长"`
	UnionMain *SchoolRes  `json:"unionMain" dc:"关联主体"`
	Parent    *TeamRes    `json:"parent" dc:"班级或小组父级ID"`
}

func (m *TeamRes) Data() *TeamRes {
	return m
}

type ITeamRes interface {
	Data() *TeamRes
}

type TeamMemberRes struct {
	co_entity.CompanyTeamMember
	Employee   *StudentRes `json:"employee"   description:"成员"`
	InviteUser *TeacherRes `json:"inviteUser" description:"邀约人"`
	UnionMain  *SchoolRes  `json:"unionMain"  description:"关联主体"`
}
