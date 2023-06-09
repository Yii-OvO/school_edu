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

type TeamController[TIRes co_model.ITeamRes] struct {
	i_controller.ITeam[TIRes]
}

// GetTeamById 根据id获取班级或小组信息
func (c *TeamController[ITTeamRes]) GetTeamById(ctx context.Context, req *school_v1.GetTeamByIdReq) (ITTeamRes, error) {
	return c.ITeam.GetTeamById(ctx, &co_company_api.GetTeamByIdReq{Id: req.Id})
}

// HasTeamByName 判断名称是否存在
func (c *TeamController[ITTeamRes]) HasTeamByName(ctx context.Context, req *school_v1.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return c.ITeam.HasTeamByName(ctx, &co_company_api.HasTeamByNameReq{Name: req.Name})
}

// QueryTeamList 查询班级或小组列表
func (c *TeamController[ITTeamRes]) QueryTeamList(ctx context.Context, req *school_v1.QueryTeamListReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.QueryTeamList(ctx, &req.QueryTeamListReq)
	return kconv.Struct(ret, &api_v1.MapRes{}), err
}

// CreateTeam 创建班级或小组
func (c *TeamController[ITTeamRes]) CreateTeam(ctx context.Context, req *school_v1.CreateTeamReq) (ITTeamRes, error) {
	return c.ITeam.CreateTeam(ctx, kconv.Struct(req, &co_company_api.CreateTeamReq{}))
}

// UpdateTeam 更新班级或小组
func (c *TeamController[ITTeamRes]) UpdateTeam(ctx context.Context, req *school_v1.UpdateTeamReq) (ITTeamRes, error) {
	return c.ITeam.UpdateTeam(ctx, kconv.Struct(req, &co_company_api.UpdateTeamReq{}))
}

// DeleteTeam 删除班级或小组
func (c *TeamController[ITTeamRes]) DeleteTeam(ctx context.Context, req *school_v1.DeleteTeamReq) (api_v1.BoolRes, error) {
	return c.ITeam.DeleteTeam(ctx, kconv.Struct(req, &co_company_api.DeleteTeamReq{Id: req.Id}))
}

// QueryTeamListByTeacher 根据教师查询班级或小组｜列表
func (c *TeamController[ITTeamRes]) QueryTeamListByTeacher(ctx context.Context, req *school_v1.QueryTeamListByTeacherReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.QueryTeamListByEmployee(ctx, kconv.Struct(req, &co_company_api.QueryTeamListByEmployeeReq{}))
	return kconv.Struct(ret, &api_v1.MapRes{}), err
}

// QueryTeamListByStudent 根据学生查询班级或小组｜列表
func (c *TeamController[ITTeamRes]) QueryTeamListByStudent(ctx context.Context, req *school_v1.QueryTeamListByStudentReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.QueryTeamListByEmployee(ctx, kconv.Struct(req, &co_company_api.QueryTeamListByEmployeeReq{}))
	//ret, err := share_consts.Student().IModules.Team().QueryTeamListByEmployee(ctx, req.StudentId, req.UnionMainId)
	return kconv.Struct(ret, &api_v1.MapRes{}), err
}

// SetTeamMember 设置班级或小组成员
func (c *TeamController[ITTeamRes]) SetTeamMember(ctx context.Context, req *school_v1.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamMember(ctx, kconv.Struct(req, &co_company_api.SetTeamMemberReq{}))
}

// SetTeamOwner 设置班级或小组管理员
func (c *TeamController[ITTeamRes]) SetTeamOwner(ctx context.Context, req *school_v1.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamOwner(ctx, kconv.Struct(req, &co_company_api.SetTeamOwnerReq{}))
}

//// SetClassLeader 设置班级班主任
//func (c *TeamController[ITTeamRes]) SetClassLeader(ctx context.Context, req *school_v1.SetClassLeaderReq) (api_v1.BoolRes, error) {
//	return c.ITeam.SetTeamOwner(ctx, kconv.Struct(req, &co_company_api.SetTeamOwnerReq{}))
//}
//
//// SetClassMonitor 设置小组管理员
//func (c *TeamController[ITTeamRes]) SetClassMonitor(ctx context.Context, req *school_v1.SetClassMonitorReq) (api_v1.BoolRes, error) {
//	return c.ITeam.SetTeamOwner(ctx, kconv.Struct(req, &co_company_api.SetTeamOwnerReq{}))
//}

// SetTeamCaptain 设置班长或小组组长
func (c *TeamController[ITTeamRes]) SetTeamCaptain(ctx context.Context, req *school_v1.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamCaptain(ctx, kconv.Struct(req, &co_company_api.SetTeamCaptainReq{}))
}
