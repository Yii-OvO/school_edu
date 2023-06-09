package school

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
	"github.com/kysion/edu-share/share_model/share_dao"
	"github.com/kysion/edu-share/share_model/share_enum"
	"math"
)

type Global struct {
	// Modules co_interface.IModules

	co_interface.IModules[
		*co_model.CompanyRes,
		*co_model.EmployeeRes,
		*co_model.TeamRes,
		*co_model.FdAccountRes,
		*co_model.FdAccountBillRes,
		*co_model.FdBankCardRes,
		*co_model.FdCurrencyRes,
		*co_model.FdInvoiceRes,
		*co_model.FdInvoiceDetailRes,
	]
	// PermissionTree 可自定义也可直接使用company骨架项目定义的权限树
	PermissionTree []*sys_model.SysPermissionTree

	// FinancialPermissionTree 财务服务权限树 (可选)
	FinancialPermissionTree []*sys_model.SysPermissionTree
}

var global *Global

func School() *Global {
	if global != nil {
		return global
	}

	global = &Global{
		// 学校
		IModules: co_module.NewModules[
			*co_model.CompanyRes,
			*co_model.EmployeeRes,
			*co_model.TeamRes,
			*co_model.FdAccountRes,
			*co_model.FdAccountBillRes,
			*co_model.FdBankCardRes,
			*co_model.FdCurrencyRes,
			*co_model.FdInvoiceRes,
			*co_model.FdInvoiceDetailRes,
		](
			&co_model.Config{
				AllowEmptyNo:                   true,
				IsCreateDefaultEmployeeAndRole: true,
				HardDeleteWaitAt:               math.MaxInt32,
				KeyIndex:                       "School",
				RoutePrefix:                    "/school",
				StoragePath:                    "./resources/school",
				UserType:                       share_enum.User.Type.Admin,
				Identifier: co_model.Identifier{
					Company:  "edu_school",
					Employee: "edu_school_teacher",
					Team:     "edu_school_class",

					FdAccount:       "edu_fd_account",
					FdAccountBill:   "edu_fd_account_bill",
					FdBankCard:      "edu_fd_bank_card",
					FdInvoice:       "edu_fd_invoice",
					FdInvoiceDetail: "edu_fd_invoice_detail",
				},
			},
			&co_dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去，相同属性前缀需要配合使用不能拆开应用
				Company:  co_dao.NewCompany(share_dao.NewEduSchool()),
				Employee: co_dao.NewCompanyEmployee(share_dao.NewEduSchoolTeacher()),

				Team:       co_dao.NewCompanyTeam(share_dao.NewEduSchoolClass()),
				TeamMember: co_dao.NewCompanyTeamMember(share_dao.NewEduSchoolClassMember()),
				//
				FdAccount:       co_dao.NewFdAccount(share_dao.NewEduFdAccount()),
				FdAccountBill:   co_dao.NewFdAccountBill(share_dao.NewEduFdAccountBill()),
				FdInvoice:       co_dao.NewFdInvoice(share_dao.NewEduFdInvoice()),
				FdInvoiceDetail: co_dao.NewFdInvoiceDetail(share_dao.NewEduFdInvoiceDetail()),
				FdCurrency:      co_dao.NewFdCurrency(share_dao.NewEduFdCurrency()),
				FdBankCard:      co_dao.NewFdBankCard(share_dao.NewEduFdBankCard()),
				FdAccountDetail: co_dao.NewFdAccountDetail(share_dao.NewEduFdAccountDetail()),
			},
		),
	}
	return global
}
