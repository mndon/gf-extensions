// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
)

type IAdminDictType interface {
	List(ctx context.Context, req *api.DictTypeSearchReq) (res *api.DictTypeSearchRes, err error)
	Add(ctx context.Context, req *api.DictTypeAddReq, userId uint64) (err error)
	Edit(ctx context.Context, req *api.DictTypeEditReq, userId uint64) (err error)
	Get(ctx context.Context, req *api.DictTypeGetReq) (dictType *entity.AdminDictType, err error)
	ExistsDictType(ctx context.Context, dictType string, dictId ...int64) (err error)
	Delete(ctx context.Context, dictIds []int) (err error)
	GetAllDictType(ctx context.Context) (list []*entity.AdminDictType, err error)
}

var localAdminDictType IAdminDictType

func AdminDictType() IAdminDictType {
	if localAdminDictType == nil {
		panic("implement not found for interface IAdminDictType, forgot register?")
	}
	return localAdminDictType
}

func RegisterAdminDictType(i IAdminDictType) {
	localAdminDictType = i
}
