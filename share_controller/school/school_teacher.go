package school_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/edu-share/api/school_v1"
)

type TeacherController[TIRes co_model.IEmployeeRes] struct {
	i_controller.IEmployee[TIRes]
}

// GetTeacherById 根据ID获取教师信息
func (c *TeacherController[TIRes]) GetTeacherById(ctx context.Context, req *school_v1.GetTeacherByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeById(ctx, &co_company_api.GetEmployeeByIdReq{Id: req.Id})
}

// GetTeacherDetailById 获取教师详情信息
func (c *TeacherController[TIRes]) GetTeacherDetailById(ctx context.Context, req *school_v1.GetTeacherDetailByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeDetailById(ctx, &co_company_api.GetEmployeeDetailByIdReq{Id: req.Id})
}

// HasTeacherByName 教师名称是否存在
func (c *TeacherController[TIRes]) HasTeacherByName(ctx context.Context, req *school_v1.HasTeacherByNameReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByName(ctx, kconv.Struct(req, &co_company_api.HasEmployeeByNameReq{}))
}

// HasTeacherByNo 教师工号是否存在
func (c *TeacherController[TIRes]) HasTeacherByNo(ctx context.Context, req *school_v1.HasTeacherByNoReq) (api_v1.BoolRes, error) {

	return c.IEmployee.HasEmployeeByNo(ctx, &co_company_api.HasEmployeeByNoReq{
		No:        req.No,
		ExcludeId: req.ExcludeId,
	})
}

// QueryTeacherList 查询教师列表
func (c *TeacherController[TIRes]) QueryTeacherList(ctx context.Context, req *school_v1.QueryTeacherListReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.QueryEmployeeList(ctx, &req.QueryEmployeeListReq)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// CreateTeacher 创建教师信息
func (c *TeacherController[TIRes]) CreateTeacher(ctx context.Context, req *school_v1.CreateTeacherReq) (TIRes, error) {
	return c.IEmployee.CreateEmployee(ctx, &req.CreateEmployeeReq)
}

// UpdateTeacher 更新教师信息
func (c *TeacherController[TIRes]) UpdateTeacher(ctx context.Context, req *school_v1.UpdateTeacherReq) (TIRes, error) {
	return c.IEmployee.UpdateEmployee(ctx, &req.UpdateEmployeeReq)
}

// DeleteTeacher 删除教师信息
func (c *TeacherController[TIRes]) DeleteTeacher(ctx context.Context, req *school_v1.DeleteTeacherReq) (api_v1.BoolRes, error) {
	return c.IEmployee.DeleteEmployee(ctx, &co_company_api.DeleteEmployeeReq{Id: req.Id})
}

// GetTeacherListByRoleId 根据角色ID获取教师列表
func (c *TeacherController[TIRes]) GetTeacherListByRoleId(ctx context.Context, req *school_v1.GetTeacherListByRoleIdReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.GetEmployeeListByRoleId(ctx, &req.GetEmployeeListByRoleIdReq)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// SetTeacherRoles 设置教师角色
func (c *TeacherController[TIRes]) SetTeacherRoles(ctx context.Context, req *school_v1.SetTeacherRolesReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeRoles(ctx, kconv.Struct(req, &co_company_api.SetEmployeeRolesReq{}))
}

// SetTeacherState 设置教师状态
func (c *TeacherController[TIRes]) SetTeacherState(ctx context.Context, req *school_v1.SetTeacherStateReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeState(ctx, kconv.Struct(req, &co_company_api.SetEmployeeStateReq{}))
}
