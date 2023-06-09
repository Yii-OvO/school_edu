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

type StudentController[TIRes co_model.IEmployeeRes] struct {
	i_controller.IEmployee[TIRes]
}

// GetStudentById 根据ID获取学生信息
func (c *StudentController[TIRes]) GetStudentById(ctx context.Context, req *school_v1.GetStudentByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeById(ctx, &co_company_api.GetEmployeeByIdReq{Id: req.Id})
}

// GetStudentDetailById 获取学生详情信息
func (c *StudentController[TIRes]) GetStudentDetailById(ctx context.Context, req *school_v1.GetStudentDetailByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeDetailById(ctx, &co_company_api.GetEmployeeDetailByIdReq{Id: req.Id})
}

// HasStudentByName 学生名称是否存在
func (c *StudentController[TIRes]) HasStudentByName(ctx context.Context, req *school_v1.HasStudentByNameReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByName(ctx, kconv.Struct(req, &co_company_api.HasEmployeeByNameReq{}))
}

// HasStudentByNo 学生学号是否存在
func (c *StudentController[TIRes]) HasStudentByNo(ctx context.Context, req *school_v1.HasStudentByNoReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByNo(ctx, &co_company_api.HasEmployeeByNoReq{
		No:        req.No,
		ExcludeId: req.ExcludeId,
	})
}

// QueryStudentList 查询学生列表
func (c *StudentController[TIRes]) QueryStudentList(ctx context.Context, req *school_v1.QueryStudentListReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.QueryEmployeeList(ctx, &req.QueryEmployeeListReq)

	//ret, err := share_consts.Student().IModules.Employee().QueryEmployeeList(ctx, &req.SearchParams)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// CreateStudent 创建学生信息
func (c *StudentController[TIRes]) CreateStudent(ctx context.Context, req *school_v1.CreateStudentReq) (TIRes, error) {
	return c.IEmployee.CreateEmployee(ctx, kconv.Struct(req, &co_company_api.CreateEmployeeReq{}))
}

// UpdateStudent 更新学生信息
func (c *StudentController[TIRes]) UpdateStudent(ctx context.Context, req *school_v1.UpdateStudentReq) (TIRes, error) {
	return c.IEmployee.UpdateEmployee(ctx, kconv.Struct(req, &co_company_api.UpdateEmployeeReq{}))
}

// DeleteStudent 删除学生信息
func (c *StudentController[TIRes]) DeleteStudent(ctx context.Context, req *school_v1.DeleteStudentReq) (api_v1.BoolRes, error) {
	return c.IEmployee.DeleteEmployee(ctx, &co_company_api.DeleteEmployeeReq{Id: req.Id})
}

// GetStudentListByRoleId 根据角色ID获取学生列表
func (c *StudentController[TIRes]) GetStudentListByRoleId(ctx context.Context, req *school_v1.GetStudentListByRoleIdReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.GetEmployeeListByRoleId(ctx, &req.GetEmployeeListByRoleIdReq)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// SetStudentRoles 设置学生角色
func (c *StudentController[TIRes]) SetStudentRoles(ctx context.Context, req *school_v1.SetStudentRolesReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeRoles(ctx, kconv.Struct(req, &co_company_api.SetEmployeeRolesReq{}))
}

// SetStudentState 设置学生状态
func (c *StudentController[TIRes]) SetStudentState(ctx context.Context, req *school_v1.SetStudentStateReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeState(ctx, kconv.Struct(req, &co_company_api.SetEmployeeStateReq{}))
}
