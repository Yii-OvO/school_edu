package school_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetTeacherByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取教师|信息" tags:"学校教师"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"教师ID"`
}

type HasTeacherByNameReq struct {
	g.Meta      `method:"post" summary:"判断名称是否存在" tags:"学校教师"`
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的教师ID"`
}

type HasTeacherByNoReq struct {
	g.Meta    `method:"post" summary:"判断工号是否存在" tags:"学校教师"`
	No        string `json:"no" dc:"工号"`
	ExcludeId int64  `json:"excludeId" dc:"要排除的教师ID"`
}

type QueryTeacherListReq struct {
	g.Meta `method:"post" summary:"查询教师|列表" tags:"学校教师"`
	co_company_api.QueryEmployeeListReq
}

type CreateTeacherReq struct {
	g.Meta `method:"post" summary:"创建教师|信息" tags:"学校教师"`
	co_company_api.CreateEmployeeReq
}

type UpdateTeacherReq struct {
	g.Meta `method:"post" summary:"更新教师|信息" tags:"学校教师"`
	co_company_api.UpdateEmployeeReq
}

type DeleteTeacherReq struct {
	g.Meta `method:"post" summary:"删除教师|信息" tags:"学校教师"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"教师ID"`
}

type GetTeacherDetailByIdReq struct {
	g.Meta `method:"post" summary:"获取教师详情|信息" tags:"学校教师"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"教师ID"`
}

type GetTeacherListByRoleIdReq struct {
	g.Meta `method:"post" summary:"根据角色ID获取所有所属教师|列表" tags:"学校教师"`
	co_company_api.GetEmployeeListByRoleIdReq
}

type SetTeacherRolesReq struct {
	g.Meta `method:"post" summary:"设置教师角色" dc:"设置教师所属角色" tags:"学校教师"`
	co_company_api.SetEmployeeRolesReq
}

type SetTeacherStateReq struct {
	g.Meta `method:"post" summary:"设置教师状态" dc:"设置教师状态" tags:"学校教师"`
	co_company_api.SetEmployeeStateReq
}
