// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package share_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/edu-share/share_model/share_dao/internal"
)

type EduFdAccountDetailDao = dao_interface.TIDao[internal.EduFdAccountDetailColumns]

func NewEduFdAccountDetail(dao ...dao_interface.IDao) EduFdAccountDetailDao {
	return (EduFdAccountDetailDao)(internal.NewEduFdAccountDetailDao(dao...))
}

var (
	EduFdAccountDetail = NewEduFdAccountDetail()
)
