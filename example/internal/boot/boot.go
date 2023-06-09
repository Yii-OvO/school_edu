package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/kysion/edu-share/share_consts"
	"github.com/kysion/edu-share/share_controller"
	"github.com/kysion/oss-library/oss_consts"
	"github.com/kysion/oss-library/oss_global"
	"github.com/kysion/oss-library/oss_router"
	"github.com/kysion/sms-library/sms_consts"
	"github.com/kysion/sms-library/sms_global"
	"github.com/kysion/sms-library/sms_router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			{
				// OpenApi自定义信息
				oai.Info.Title = `API Reference`
				oai.Config.CommonResponse = api_v1.JsonRes{}
				oai.Config.CommonResponseDataField = `Data`
			}

			{
				// 静态目录设置
				uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
				if uploadPath == "" {
					g.Log().Fatal(ctx, "文件上传配置路径不能为空!")
				}
				if !gfile.Exists(uploadPath) {
					_ = gfile.Mkdir(uploadPath)
				}
				// 上传目录添加至静态资源
				s.AddStaticPath("/upload", uploadPath)
			}

			{
				// HOOK, 开发阶段禁止浏览器缓存,方便调试
				if gmode.IsDevelop() {
					s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
						r.Response.Header().Set("Cache-Control", "no-store")
					})
				}
			}

			{
				// sysPermission 初始化  (基础骨架权限树)
				sys_service.SysPermission().ImportPermissionTree(ctx, sys_consts.Global.PermissionTree, nil)

				// ImportPermissionTree 导入权限结构 (Company骨架权限树)
				sys_service.SysPermission().ImportPermissionTree(ctx, co_consts.PermissionTree, nil)

				// 导入财务服务权限结构 (可选) (Company骨架财务权限树)
				sys_service.SysPermission().ImportPermissionTree(ctx, co_consts.FinancialPermissionTree, nil)

				// 导入oss + sms权限树
				sys_service.SysPermission().ImportPermissionTree(ctx, oss_consts.PermissionTree, nil)
				sys_service.SysPermission().ImportPermissionTree(ctx, sms_consts.PermissionTree, nil)

				// CASBIN 初始化
				sys_service.Casbin().Enforcer()
			}

			apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
			s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
				// 注册中间件
				group.Middleware(
					// sys_service.Middleware().Casbin,
					sys_service.Middleware().CTX,
					sys_service.Middleware().ResponseHandler,
				)

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 鉴权：登录，注册，找回密码等
					group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
					// 图型验证码、短信验证码、地区
					group.Group("/common", func(group *ghttp.RouterGroup) {
						group.Bind(
							// 图型验证码
							sys_controller.Captcha,
							// 短信验证码
							sys_controller.SysSms,
							// 地区
							sys_controller.SysArea,
						)
					})
				})

				// 权限路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						sys_service.Middleware().Auth,
					)

					// sys公共骨架相关
					{
						// 角色
						group.Group("/role", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysRole) })
						// 权限
						group.Group("/permission", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysPermission) })
						// 系统配置
						group.Group("/system/config", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysConfig) })
						// 我的
						group.Group("/my", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysMy) }) // 需替换成自己的
						// 资质管理
						group.Group("/license", func(group *ghttp.RouterGroup) {
							group.Bind(
								// 主体资质
								sys_controller.SysAudit,
								// 资质审核
								sys_controller.SysLicense,
							)
						})

						// 文件上传
						group.Group("/common/file", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysFile) })

					}
					{
						// oss 服务路由绑定
						oss_router.ModulesGroup(oss_global.Global.Modules, group)

						// sms_v1 短信路由绑定
						sms_router.ModulesGroup(sms_global.Global.Modules, group)

						// Sms
						//router.ModulesGroup(sms_global.Global.Modules, group) // key：场景类型 + 验证码
					}

					// 本项目拓展，学校接口路由注册
					{

						//share_router.NewMainRouteGroup(share_controller.NewSchoolController(share_consts.School().IModules), group)
						//share_router.NewSchoolRouteGroup(share_controller.NewSchoolController(share_consts.School().IModules), group)
						//share_router.NewSchoolRouteGroup(share_controller.NewSchoolController(share_consts.Teacher().IModules), group)

						{
							schoolSchoolController := share_controller.NewSchoolController(share_consts.School().IModules)
							group.POST("/school/createSchool", schoolSchoolController.School.CreateSchool)
							group.POST("/school/updateSchool", schoolSchoolController.School.UpdateSchool)
							group.POST("/school/hasSchoolByName", schoolSchoolController.School.HasSchoolByName)
							group.POST("/school/getSchoolById", schoolSchoolController.School.GetSchoolById)
							group.POST("/school/querySchoolList", schoolSchoolController.School.QuerySchoolList)
							group.POST("/school/getSchoolDetail", schoolSchoolController.School.GetSchoolDetail)
						}

						{
							schoolTeacherlController := share_controller.NewSchoolController(share_consts.Teacher().IModules)
							group.POST("/teacher/getTeacherById", schoolTeacherlController.Teacher.GetTeacherById)
							group.POST("/teacher/getTeacherDetailById", schoolTeacherlController.Teacher.GetTeacherDetailById)
							group.POST("/teacher/hasTeacherByName", schoolTeacherlController.Teacher.HasTeacherByName)
							group.POST("/teacher/hasTeacherByNo", schoolTeacherlController.Teacher.HasTeacherByNo)
							group.POST("/teacher/queryTeacherList", schoolTeacherlController.Teacher.QueryTeacherList)
							group.POST("/teacher/createTeacher", schoolTeacherlController.Teacher.CreateTeacher)
							group.POST("/teacher/updateTeacher", schoolTeacherlController.Teacher.UpdateTeacher)
							group.POST("/teacher/deleteTeacher", schoolTeacherlController.Teacher.DeleteTeacher)
							group.POST("/teacher/getTeacherListByRoleId", schoolTeacherlController.Teacher.GetTeacherListByRoleId)
							group.POST("/teacher/setTeacherRoles", schoolTeacherlController.Teacher.SetTeacherRoles)
							group.POST("/teacher/setTeacherState", schoolTeacherlController.Teacher.SetTeacherState)

							{
								group.POST("/class/getTeamById", schoolTeacherlController.Class.GetTeamById)
								group.POST("/class/hasTeamByName", schoolTeacherlController.Class.HasTeamByName)
								group.POST("/class/queryTeamList", schoolTeacherlController.Class.QueryTeamList)
								group.POST("/class/createTeam", schoolTeacherlController.Class.CreateTeam)
								group.POST("/class/updateTeam", schoolTeacherlController.Class.UpdateTeam)
								group.POST("/class/queryTeamListByTeacher", schoolTeacherlController.Class.QueryTeamListByTeacher)
								group.POST("/class/deleteTeam", schoolTeacherlController.Class.DeleteTeam)
								//group.POST("/class/setClassLeader", schoolTeacherlController.Class.SetClassLeader)
								group.POST("/class/setTeamOwner", schoolTeacherlController.Class.SetTeamOwner)

							}
						}

						{
							schoolStudentController := share_controller.NewSchoolController(share_consts.Student().IModules)
							group.POST("/student/getStudentById", schoolStudentController.Student.GetStudentById)
							group.POST("/student/getStudentDetailById", schoolStudentController.Student.GetStudentDetailById)
							group.POST("/student/hasStudentByName", schoolStudentController.Student.HasStudentByName)
							group.POST("/student/hasStudentByNo", schoolStudentController.Student.HasStudentByNo)
							group.POST("/student/queryStudentList", schoolStudentController.Student.QueryStudentList)
							group.POST("/student/createStudent", schoolStudentController.Student.CreateStudent)
							group.POST("/student/updateStudent", schoolStudentController.Student.UpdateStudent)
							group.POST("/student/deleteStudent", schoolStudentController.Student.DeleteStudent)
							group.POST("/student/getStudentListByRoleId", schoolStudentController.Student.GetStudentListByRoleId)
							group.POST("/student/setStudentRoles", schoolStudentController.Student.SetStudentRoles)
							group.POST("/student/setStudentState", schoolStudentController.Student.SetStudentState)

							{

								group.POST("/class/queryTeamListByStudent", schoolStudentController.Class.QueryTeamListByStudent)
								group.POST("/class/setTeamMember", schoolStudentController.Class.SetTeamMember)
								//group.POST("/class/setClassMonitor", schoolStudentController.Class.SetClassMonitor)
								//group.POST("/class/setTeamMonitor", schoolStudentController.Class.SetTeamOwner)
								group.POST("/class/setTeamCaptain", schoolStudentController.Class.SetTeamCaptain)

							}
						}
					}

					{
						// 教师
						//group.Group("/teacher", func(group *ghttp.RouterGroup) {
						//	group.Bind(school_controller.Teacher)
						//	group.POST("/getTeacherById", school_controller.Teacher.GetTeacherById)
						//group.POST("/getTeacherDetailById", v.GetTeacherDetailById)
						//group.POST("/hasTeacherByName", v.HasTeacherByName)
						//group.POST("/hasTeacherByNo", v.HasTeacherByNo)
						//group.POST("/queryTeacherList", v.QueryTeacherList)
						//group.POST("/createTeacher", v.CreateTeacher)
						//group.POST("/updateTeacher", v.UpdateTeacher)
						//group.POST("/deleteTeacher", v.DeleteTeacher)
						//group.POST("/getTeacherListByRoleId", v.GetTeacherListByRoleId)
						//group.POST("/setTeacherRoles", v.SetTeacherRoles)
						//group.POST("/setTeacherState", v.SetTeacherState)
						//})
					}

					/*
						// 消费者
						{
							// 消费者无需团队员工，只需要财务
							//share_router.NewMainRouteGroup(share_controller.NewConsumerController(share_consts.Global.Consumer), group, share_enum.Feature.Type.Financial)
							// 消费者
							group.Group("/consumer", func(group *ghttp.RouterGroup) { // 共享库可删除
								group.Bind(consumer.Consumer)

								// 后续把消费者的财务写成 NewFinancialRouterGroup() 这样的形式，不要使用下面直接Bind绑定、不然出来的接口就是根据方法名字生成接口path
								group.Group("/financial", func(group *ghttp.RouterGroup) {
									group.POST("/getAccountBalance", share_controller.Financial(share_consts.Consumer().IModules).GetAccountBalance)
									group.POST("/invoiceRegister", share_controller.Financial(share_consts.Consumer().IModules).InvoiceRegister)
									group.POST("/queryInvoice", share_controller.Financial(share_consts.Consumer().IModules).QueryInvoice)
									group.POST("/deletesFdInvoiceById", share_controller.Financial(share_consts.Consumer().IModules).DeletesFdInvoiceById)
									group.POST("/invoiceDetailRegister", share_controller.Financial(share_consts.Consumer().IModules).InvoiceDetailRegister)
									group.POST("/queryInvoiceDetailList", share_controller.Financial(share_consts.Consumer().IModules).QueryInvoiceDetailList)
									group.POST("/makeInvoiceDetailReq", share_controller.Financial(share_consts.Consumer().IModules).MakeInvoiceDetailReq)
									group.POST("/auditInvoiceDetail", share_controller.Financial(share_consts.Consumer().IModules).AuditInvoiceDetail)
									group.POST("/bankCardRegister", share_controller.Financial(share_consts.Consumer().IModules).BankCardRegister)
									group.POST("/deleteBankCard", share_controller.Financial(share_consts.Consumer().IModules).DeleteBankCard)
									group.POST("/queryBankCardList", share_controller.Financial(share_consts.Consumer().IModules).QueryBankCardList)

									group.POST("/getAccountDetail", share_controller.Financial(share_consts.Consumer().IModules).GetAccountDetail)
									group.POST("/updateAccountIsEnabled", share_controller.Financial(share_consts.Consumer().IModules).UpdateAccountIsEnabled)
									group.POST("/updateAccountLimitState", share_controller.Financial(share_consts.Consumer().IModules).UpdateAccountLimitState)

									group.POST("/getAccountDetailByAccountId", share_controller.Financial(share_consts.Consumer().IModules).GetAccountDetailById)
									group.POST("/increment", share_controller.Financial(share_consts.Consumer().IModules).Increment)
									group.POST("/decrement", share_controller.Financial(share_consts.Consumer().IModules).Decrement)
								})

								// 消费者支付
								group.Bind(consumer.ConsumerAlipay)

							})
						}*/

					// share_consts.Global.Consumer.Account().CreateAccount()  具体的财务方法
					// 拓展 ，，， 下面可以自定义share_controller.XXX XXXapi_v1 拓展业务功能

					// 用法2, 选择性注册路由
					//share_router.BindCompanyGroup("facilitator", share_controller.Company(consts.Global.Facilitator), group)
					//share_router.BindCompanyGroup("company", share_controller.Company(consts.Global.Facilitator), group)

					//share_router.BindCompanyGroup("company", share_controller.Company(consts.Global.Facilitator), group)

				})

				//

			})

			s.Run()
			return nil
		},
	}
)
