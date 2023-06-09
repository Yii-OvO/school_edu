package share_enum

import (
	"github.com/kysion/edu-share/share_model/share_enum/internal/user"
)

type (
	UserType user.TypeEnum

	/*
		DeviceState  device.StateTypeEnum
		DeviceAction device.ActionEnum

		VideoAuditType video.AuditTypeEnum

		ConsumerState consumer.StateTypeEnum

		FeatureType feature.FeatureTypeEnum

		SubAccountSchemeState sub_account_scheme.StateTypeEnum

		SubAccountAction     sub_account_record.ActionEnum
		SubAccountState      sub_account_record.StateTypeEnum
		SubAccountAuditState sub_account_record.AuditStateEnum

		// DeviceTransferAction 设备流转行为类型
		DeviceTransferAction device_transfer_record.ActionTypeEnum

		// AuthState 授权状态
		AuthState auth.StateTypeEnum

		// InviteState  邀约记录
		InviteState invite_record.StateEnum

		// InviteType  邀约类型
		InviteType invite_record.StateEnum

		// ServiceState 业务关系状态
		ServiceState service_record.StateEnum
	*/
)

var (
	// User 用户
	User = user.User
)
