package user

import "github.com/kysion/base-library/utility/enum"

type TypeEnum enum.IEnumCode[int]

type userType struct {
	Student      TypeEnum
	Operator     TypeEnum
	Anonymous    TypeEnum
	Parents      TypeEnum
	Teacher      TypeEnum
	Organization TypeEnum
	Advertiser   TypeEnum
	Admin        TypeEnum
	SuperAdmin   TypeEnum
}

var Type = userType{
	//用户类型，-1超级管理员，0匿名，1学生，2教师，4家长，8广告主，16运营商，32公益机构，64平台
	Anonymous:    enum.New[TypeEnum](0, "匿名"),
	Student:      enum.New[TypeEnum](1, "学生"),
	Teacher:      enum.New[TypeEnum](2, "教师"),
	Parents:      enum.New[TypeEnum](4, "家长"),
	Advertiser:   enum.New[TypeEnum](8, "广告主"),
	Organization: enum.New[TypeEnum](16, "公益机构"),
	Operator:     enum.New[TypeEnum](32, "运营商"),
	Admin:        enum.New[TypeEnum](64, "平台"),
	SuperAdmin:   enum.New[TypeEnum](-1, "超级管理员"),
}

func (e userType) New(code int) TypeEnum {
	if code == Type.Anonymous.Code() {
		return Type.Anonymous
	}
	if (code & Type.Student.Code()) == Type.Student.Code() {
		return Type.Student
	}
	if (code & Type.Teacher.Code()) == Type.Teacher.Code() {
		return Type.Teacher
	}
	if (code & Type.Parents.Code()) == Type.Parents.Code() {
		return Type.Parents
	}
	if (code & Type.Advertiser.Code()) == Type.Advertiser.Code() {
		return Type.Advertiser
	}
	if (code & Type.Organization.Code()) == Type.Organization.Code() {
		return Type.Organization
	}
	if (code & Type.Operator.Code()) == Type.Operator.Code() {
		return Type.Operator
	}
	if (code & Type.Admin.Code()) == Type.Admin.Code() {
		return Type.Admin
	}
	if (code & Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
		return Type.SuperAdmin
	}

	panic("User.Type.New: error")
}
