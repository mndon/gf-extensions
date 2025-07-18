// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mndon/gf-extensions/adminx/internal/dao/internal"
)

// internalAdminUserOnlineDao is internal type for wrapping internal DAO implements.
type internalAdminUserOnlineDao = *internal.AdminUserOnlineDao

// adminUserOnlineDao is the data access object for table admin_user_online.
// You can define custom methods on it to extend its functionality as you wish.
type adminUserOnlineDao struct {
	internalAdminUserOnlineDao
}

var (
	// AdminUserOnline is globally public accessible object for table admin_user_online operations.
	AdminUserOnline = adminUserOnlineDao{
		internal.NewAdminUserOnlineDao(),
	}
)

// Fill with you ideas below.
