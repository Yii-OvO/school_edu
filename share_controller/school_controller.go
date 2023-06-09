package share_controller

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	schoolController "github.com/kysion/edu-share/share_controller/school"
)

func NewSchoolController[
	TICompanyRes co_model.ICompanyRes,
	TIEmployeeRes co_model.IEmployeeRes,
	TITeamRes co_model.ITeamRes,

	TIFdAccountRes co_model.IFdAccountRes,
	TIFdAccountBillRes co_model.IFdAccountBillRes,
	TIFdBankCardRes co_model.IFdBankCardRes,
	TIFdCurrencyRes co_model.IFdCurrencyRes,
	TIFdInvoiceRes co_model.IFdInvoiceRes,
	TIFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	TICompanyRes,
	TIEmployeeRes,
	TITeamRes,
	TIFdAccountRes,
	TIFdAccountBillRes,
	TIFdBankCardRes,
	TIFdCurrencyRes,
	TIFdInvoiceRes,
	TIFdInvoiceDetailRes,
]) *ModuleController[
	TICompanyRes,
	TIEmployeeRes,
	TITeamRes,
	TIFdAccountRes,
	TIFdAccountBillRes,
	TIFdBankCardRes,
	TIFdCurrencyRes,
	TIFdInvoiceRes,
	TIFdInvoiceDetailRes,
] {
	return &ModuleController[
		TICompanyRes,
		TIEmployeeRes,
		TITeamRes,
		TIFdAccountRes,
		TIFdAccountBillRes,
		TIFdBankCardRes,
		TIFdCurrencyRes,
		TIFdInvoiceRes,
		TIFdInvoiceDetailRes,
	]{
		School: &schoolController.SchoolController[TICompanyRes]{
			ICompany: co_controller.Company(modules),
		},
		Teacher: &schoolController.TeacherController[TIEmployeeRes]{
			IEmployee: co_controller.Employee(modules),
		},

		Student: &schoolController.StudentController[TIEmployeeRes]{
			IEmployee: co_controller.Employee(modules),
		},
		Class: &schoolController.TeamController[TITeamRes]{
			ITeam: co_controller.Team(modules),
		},
		Modules: modules,
	}
}
