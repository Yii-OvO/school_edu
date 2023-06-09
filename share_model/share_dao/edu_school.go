// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package share_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/edu-share/share_model/share_dao/internal"
)

type EduSchoolDao = dao_interface.TIDao[internal.EduSchoolColumns]

func NewEduSchool(dao ...dao_interface.IDao) EduSchoolDao {
	return (EduSchoolDao)(internal.NewEduSchoolDao(dao...))
}

var (
	EduSchool = NewEduSchool()
)
