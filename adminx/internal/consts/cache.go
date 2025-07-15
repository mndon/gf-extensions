/*
* @desc:缓存相关
* @company:
* @Author:
* @Date:   2022/3/9 11:25
 */

package consts

const (
	// CachePrefix 应用缓存数据前缀
	CachePrefix = "APP:"

	CacheModelMem   = "memory"
	CacheModelRedis = "redis"
	CacheModelDist  = "dist"

	// CacheAdminDict 字典缓存菜单KEY
	CacheAdminDict = CachePrefix + "adminDict"

	// CacheAdminDictTag 字典缓存标签
	CacheAdminDictTag = CachePrefix + "adminDictTag"
	// CacheAdminConfigTag 系统参数配置
	CacheAdminConfigTag = CachePrefix + "adminConfigTag"
)

const (
	// CacheAdminAuthMenu 缓存菜单key
	CacheAdminAuthMenu = CachePrefix + "adminAuthMenu"
	// CacheAdminDept 缓存部门key
	CacheAdminDept = CachePrefix + "adminDept"

	// CacheAdminRole 角色缓存key
	CacheAdminRole = CachePrefix + "adminRole"
	// CacheAdminWebSet 站点配置缓存key
	CacheAdminWebSet = CachePrefix + "adminWebSet"
	// CacheAdminCmsMenu cms缓存key
	CacheAdminCmsMenu = CachePrefix + "adminCmsMenu"

	// CacheAdminAuthTag 权限缓存TAG标签
	CacheAdminAuthTag = CachePrefix + "adminAuthTag"
	// CacheAdminModelTag 模型缓存标签
	CacheAdminModelTag = CachePrefix + "adminModelTag"
	// CacheAdminCmsTag cms缓存标签
	CacheAdminCmsTag = CachePrefix + "adminCmsTag"
)
