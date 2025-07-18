// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mndon/gf-extensions/adminx/internal/dao/internal"
)

// internalAdminLoginLogDao is internal type for wrapping internal DAO implements.
type internalAdminLoginLogDao = *internal.AdminLoginLogDao

// adminLoginLogDao is the data access object for table admin_login_log.
// You can define custom methods on it to extend its functionality as you wish.
type adminLoginLogDao struct {
	internalAdminLoginLogDao
}

var (
	// AdminLoginLog is globally public accessible object for table admin_login_log operations.
	AdminLoginLog = adminLoginLogDao{
		internal.NewAdminLoginLogDao(),
	}
)

// Fill with you ideas below.
