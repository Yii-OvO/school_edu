package school_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
	model "github.com/kysion/edu-share/share_model"
)

type GetStudentByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取学生|信息" tags:"学校学生"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"学生ID"`
}

type HasStudentByNameReq struct {
	g.Meta      `method:"post" summary:"判断名称是否存在" tags:"学校学生"`
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" v:"required#名称不能为空" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的学生ID"`
}

type HasStudentByNoReq struct {
	g.Meta    `method:"post" summary:"判断学号是否存在" tags:"学校学生"`
	No        string `json:"no" dc:"学号"`
	ExcludeId int64  `json:"excludeId" dc:"要排除的学生ID"`
}

type QueryStudentListReq struct {
	g.Meta `method:"post" summary:"查询学生|列表" tags:"学校学生"`
	co_company_api.QueryEmployeeListReq
}

type CreateStudentReq struct {
	g.Meta `method:"post" summary:"创建学生|信息" tags:"学校学生"`
	model.Student
}

type UpdateStudentReq struct {
	g.Meta `method:"post" summary:"更新学生|信息" tags:"学校学生"`
	model.Student
}

type DeleteStudentReq struct {
	g.Meta `method:"post" summary:"删除学生|信息" tags:"学校学生"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"学生ID"`
}

type GetStudentDetailByIdReq struct {
	g.Meta `method:"post" summary:"获取学生详情|信息" tags:"学校学生"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"学生ID"`
}

type GetStudentListByRoleIdReq struct {
	g.Meta `method:"post" summary:"根据角色ID获取所有所属学生|列表" tags:"学校学生"`
	co_company_api.GetEmployeeListByRoleIdReq
}

type SetStudentRolesReq struct {
	g.Meta `method:"post" summary:"设置学生角色" dc:"设置学生所属角色" tags:"学校学生"`
	co_company_api.SetEmployeeRolesReq
}

type SetStudentStateReq struct {
	g.Meta `method:"post" summary:"设置学生状态" dc:"设置学生状态" tags:"学校学生"`
	model.SetStudentStateReq
}
