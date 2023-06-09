package school_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	model "github.com/kysion/edu-share/share_model"
)

type GetSchoolByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取学校|信息" tags:"学校"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"学校ID"`
}

type HasSchoolByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"学校"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QuerySchoolListReq struct {
	g.Meta `method:"post" summary:"查询学校|列表" tags:"学校"`
	base_model.SearchParams
}

type CreateSchoolReq struct {
	g.Meta `method:"post" summary:"创建学校|信息" tags:"学校"`
	model.School
}

type UpdateSchoolReq struct {
	g.Meta `method:"post" summary:"更新学校|信息" tags:"学校"`
	model.School
}

type GetSchoolDetailReq struct {
	g.Meta `method:"post" summary:"查看更多信息含完整手机号|信息" tags:"学校"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"学校ID"`
}
