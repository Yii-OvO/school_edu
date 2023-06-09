package share_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/edu-share/share_model/share_entity"
)

type TeacherUser struct {
	g.Meta   `orm:"table:sys_user"`
	Id       int64  `json:"id"           dc:"ID，保持与USERID一致"`
	Username string `json:"username"  dc:"账号"`
	State    int    `json:"state"     dc:"状态： -1已离职，0待确认，1已入职"`
	Type     int    `json:"type"      dc:"用户类型: -1超级管理员，0匿名，1学生，2教师，4家长、8广告主、16运营商、32公益机构，64平台"`
}

type TeacherRes struct {
	share_entity.EduSchoolTeacher
	User     TeacherUser              `orm:"with:id" json:"user"`
	Detail   sys_entity.SysUserDetail `orm:"with:id" json:"detail"`
	TeamList []Team                   `json:"teamList"`
}

type TeacherListRes base_model.CollectRes[*TeacherRes]

func (m *TeacherRes) Data() *TeacherRes {
	return m
}

type IETeacherRes interface {
	Data() *TeacherRes
}
