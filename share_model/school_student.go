package share_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/edu-share/share_model/share_entity"
)

type Student struct {
	Id          int64       `json:"id"           dc:"ID，保持与USERID一致"`
	No          string      `json:"no"           v:"max-length:16#学号长度超出限定20字符" dc:"学号"`
	Avatar      string      `json:"avatar"       dc:"头像"`
	Name        string      `json:"name"         v:"required|max-length:16#名称不能为空|姓名长度超出限定16字符" dc:"姓名"`
	Mobile      string      `json:"mobile"       v:"phone#手机号校验失败" dc:"手机号"`
	State       int         `json:"state"        v:"in:-1,0,1,2,3#请选择学生状态" dc:"状态：-1毕业，0待确认，1在读，2休学，3退学"`
	UnionMainId int64       `json:"union_main_id" v:"required#所属主体不能为空" dc:"所属主体"`
	HiredAt     *gtime.Time `json:"hiredAt"      v:"date-format:Y-m-d#入学日期格式错误" dc:"入学日期"`
	Sex         int         `json:"sex"          dc:"性别：0女 1男"`
}

type StudentUser struct {
	g.Meta   `orm:"table:sys_user"`
	Id       int64  `json:"id"           dc:"ID，保持与USERID一致"`
	Username string `json:"username"  dc:"账号"`
	State    int    `json:"state"     dc:"状态：状态： -1毕业，0待确认，1在读，2休学，3退学"`
	Type     int    `json:"type"      dc:"用户类型，-1超级管理员，0匿名，1学生，2教师，4家长、8广告主、16运营商、32公益机构，64平台"`
}

type SetStudentStateReq struct {
	Id    int64 `json:"id"       v:"required#ID校验失败"     dc:"ID，保持与USERID一致" `
	State int   `json:"state"        v:"in:-1,0,1,2,3#请选择学生状态" dc:"状态： -1毕业，0待确认，1在读，2休学，3退学"`
}

type StudentRes struct {
	share_entity.EduSchoolStudent
	User     StudentUser              `orm:"with:id" json:"user"`
	Detail   sys_entity.SysUserDetail `orm:"with:id" json:"detail"`
	TeamList []Team                   `json:"teamList"`
}

type StudentListRes base_model.CollectRes[*StudentRes]

func (m *StudentRes) Data() *StudentRes {
	return m
}

type IStudentRes interface {
	Data() *StudentRes
}
