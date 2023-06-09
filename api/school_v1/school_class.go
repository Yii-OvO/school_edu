package school_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
	model "github.com/kysion/edu-share/share_model"
)

type GetTeamByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取班级或小组|信息" tags:"学校班级|小组"`
	Id     int64 `json:"id" v:"required#班级或小组ID校验失败" dc:"班级或小组ID"`
}

type HasTeamByNameReq struct {
	g.Meta      `method:"post" summary:"判断名称是否存在" tags:"学校班级|小组"`
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的班级或小组ID"`
}

type QueryTeamListReq struct {
	g.Meta `method:"post" summary:"查询班级或小组|列表" tags:"学校班级|小组"`
	co_company_api.QueryTeamListReq
}

type CreateTeamReq struct {
	g.Meta `method:"post" summary:"创建班级或小组|信息" tags:"学校班级|小组"`
	model.Team
}

type UpdateTeamReq struct {
	g.Meta `method:"post" summary:"更新班级或小组|信息" tags:"学校班级|小组"`
	Id     int64  `json:"id" v:"required#班级或小组ID校验失败" dc:"班级或小组ID"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Remark string `json:"remark" dc:"备注"`
}

type DeleteTeamReq struct {
	g.Meta `method:"post" summary:"删除班级或小组|信息" tags:"学校班级|小组"`
	Id     int64 `json:"id" v:"required#班级或小组ID校验失败" dc:"班级或小组ID"`
}

type QueryTeamListByTeacherReq struct {
	g.Meta      `method:"post" summary:"根据教师查询班级|列表" tags:"学校班级|小组"`
	EmployeeId  int64 `json:"employeeId" v:"required#教师ID校验失败" dc:"教师ID"`
	UnionMainId int64 `json:"unionMainId" dc:"关联主体，默认当前主体"`
}

type QueryTeamListByStudentReq struct {
	g.Meta      `method:"post" summary:"根据学生查询班级|列表" tags:"学校班级|小组"`
	EmployeeId  int64 `json:"employeeId" v:"required#学生ID校验失败" dc:"学生ID"`
	UnionMainId int64 `json:"unionMainId" dc:"关联主体，默认当前主体"`
}

type SetTeamMemberReq struct {
	g.Meta      `method:"post" summary:"设置班级或小组成员" tags:"学校班级|小组"`
	Id          int64   `json:"id" v:"required#班级ID校验失败" dc:"班级或小组ID"`
	EmployeeIds []int64 `json:"employeeIds" dc:"班级成员"`
}

type SetTeamOwnerReq struct {
	g.Meta     `method:"post" summary:"设置班主任或班长" tags:"学校班级|小组"`
	Id         int64 `json:"id" v:"required#班级或小组ID校验失败" dc:"班级或小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#班主任或班长ID校验失败" dc:"班主任或班长ID"`
}

//type SetClassLeaderReq struct {
//	g.Meta     `method:"post" summary:"设置班主任" tags:"学校班级|小组"`
//	Id         int64 `json:"id" v:"required#班级ID校验失败" dc:"班级ID"`
//	EmployeeId int64 `json:"employeeId" v:"required#班主任ID校验失败" dc:"教师ID"`
//}
//
//type SetClassMonitorReq struct {
//	g.Meta     `method:"post" summary:"设置小组管理员" tags:"学校班级|小组"`
//	Id         int64 `json:"id" v:"required#小组ID校验失败" dc:"小组ID"`
//	EmployeeId int64 `json:"employeeId" v:"required#学生ID校验失败" dc:"学生ID,一般为班长"`
//}

type SetTeamCaptainReq struct {
	g.Meta     `method:"post" summary:"设置班长或小组组长" tags:"学校班级|小组"`
	Id         int64 `json:"id" v:"required#班级或小组ID校验失败" dc:"班级或小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#班长或小组组长ID校验失败" dc:"班长或小组组长ID"`
}
