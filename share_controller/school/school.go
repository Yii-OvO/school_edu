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

type SchoolController[TIRes co_model.ICompanyRes] struct {
	i_controller.ICompany[TIRes]
}

// GetSchoolById 通过ID获取学校信息
func (c *SchoolController[TIRes]) GetSchoolById(ctx context.Context, req *school_v1.GetSchoolByIdReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.GetCompanyById(ctx, &co_company_api.GetCompanyByIdReq{Id: req.Id})
	return ret.Data(), err
}

// HasSchoolByName 学校名称是否存在
func (c *SchoolController[TIRes]) HasSchoolByName(ctx context.Context, req *school_v1.HasSchoolByNameReq) (api_v1.BoolRes, error) {
	return c.ICompany.HasCompanyByName(ctx, &co_company_api.HasCompanyByNameReq{Name: req.Name})
}

// QuerySchoolList 查询学校列表
func (c *SchoolController[TIRes]) QuerySchoolList(ctx context.Context, req *school_v1.QuerySchoolListReq) (*co_model.CompanyListRes, error) {
	ret, err := c.ICompany.QueryCompanyList(ctx, &co_company_api.QueryCompanyListReq{
		SearchParams: req.SearchParams,
	})
	return kconv.Struct(ret, &co_model.CompanyListRes{}), err
}

// CreateSchool 创建学校信息
func (c *SchoolController[TIRes]) CreateSchool(ctx context.Context, req *school_v1.CreateSchoolReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.CreateCompany(ctx, &co_company_api.CreateCompanyReq{
		Company: *kconv.Struct(req.School, &co_model.Company{}),
	})
	return ret.Data(), err
}

// UpdateSchool 更新学校信息
func (c *SchoolController[TIRes]) UpdateSchool(ctx context.Context, req *school_v1.UpdateSchoolReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.UpdateCompany(ctx, &co_company_api.UpdateCompanyReq{
		Company: *kconv.Struct(req.School, &co_model.Company{}),
	})
	return ret.Data(), err
}

// GetSchoolDetail 查看更多信息含完整手机号
func (c *SchoolController[TIRes]) GetSchoolDetail(ctx context.Context, req *school_v1.GetSchoolDetailReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.GetCompanyDetail(ctx, &co_company_api.GetCompanyDetailReq{Id: req.Id})
	return ret.Data(), err
}
