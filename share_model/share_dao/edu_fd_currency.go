// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package share_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/edu-share/share_model/share_dao/internal"
)

type EduFdCurrencyDao = dao_interface.TIDao[internal.EduFdCurrencyColumns]

func NewEduFdCurrency(dao ...dao_interface.IDao) EduFdCurrencyDao {
	return (EduFdCurrencyDao)(internal.NewEduFdCurrencyDao(dao...))
}

var (
	EduFdCurrency = NewEduFdCurrency()
)