package share_controller

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	school_controller "github.com/kysion/edu-share/share_controller/school"
)

type ModuleController[
	TICompanyRes co_model.ICompanyRes,
	TIEmployeeRes co_model.IEmployeeRes,
	TITeamRes co_model.ITeamRes,

	TIFdAccountRes co_model.IFdAccountRes,
	TIFdAccountBillRes co_model.IFdAccountBillRes,
	TIFdBankCardRes co_model.IFdBankCardRes,
	TIFdCurrencyRes co_model.IFdCurrencyRes,
	TIFdInvoiceRes co_model.IFdInvoiceRes,
	TIFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	School  *school_controller.SchoolController[TICompanyRes]
	Teacher *school_controller.TeacherController[TIEmployeeRes]
	Student *school_controller.StudentController[TIEmployeeRes]
	Class   *school_controller.TeamController[TITeamRes]
	Modules co_interface.IModules[
		TICompanyRes,
		TIEmployeeRes,
		TITeamRes,
		TIFdAccountRes,
		TIFdAccountBillRes,
		TIFdBankCardRes,
		TIFdCurrencyRes,
		TIFdInvoiceRes,
		TIFdInvoiceDetailRes,
	]
}
