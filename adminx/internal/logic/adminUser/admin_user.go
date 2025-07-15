/*
* @desc:用户处理
* @company:
* @Author:
* @Date:   2022/9/23 15:08
 */

package adminUser

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/libUtils"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/mssola/user_agent"
)

func init() {
	service.RegisterAdminUser(New())
}

type sAdminUser struct {
	casBinUserPrefix string //CasBin 用户id前缀
}

func New() *sAdminUser {
	return &sAdminUser{
		casBinUserPrefix: "u_",
	}
}

func (s *sAdminUser) GetCasBinUserPrefix() string {
	return s.casBinUserPrefix
}

func (s *sAdminUser) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "api.notCheckAuthAdminIds")
	if !g.IsNil(ids) {
		return gset.NewFrom(ids)
	}
	return gset.New()
}

func (s *sAdminUser) GetAdminUserByUsernamePassword(ctx context.Context, req *api.UserLoginReq) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, req.Username)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(user, "账号密码错误")
		//验证密码
		if libUtils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			liberr.ErrIsNil(ctx, gerror.New("账号密码错误"))
		}
		//账号状态
		if user.UserStatus == 0 {
			liberr.ErrIsNil(ctx, gerror.New("账号已被冻结"))
		}
	})
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sAdminUser) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.LoginUserRes{}
		err = dao.AdminUser.Ctx(ctx).Fields(user).Where(dao.AdminUser.Columns().UserName, userName).Scan(user)
		liberr.ErrIsNil(ctx, err, "账号密码错误")
	})
	return
}

// GetUserById 通过用户名获取用户信息
func (s *sAdminUser) GetUserById(ctx context.Context, id uint64) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.LoginUserRes{}
		err = dao.AdminUser.Ctx(ctx).Fields(user).WherePri(id).Scan(user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
	})
	return
}

// LoginLog 记录登录日志
func (s *sAdminUser) LoginLog(ctx context.Context, params *model.LoginLogParams) {
	ua := user_agent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.AdminLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: libUtils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.AdminLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sAdminUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminUser.Ctx(ctx).WherePri(id).Unscoped().Update(g.Map{
			dao.AdminUser.Columns().LastLoginIp:   ip,
			dao.AdminUser.Columns().LastLoginTime: gtime.Now(),
		})
		liberr.ErrIsNil(ctx, err, "更新用户登录信息失败")
	})
	return
}

// GetAdminRules 获取用户菜单数据
func (s *sAdminUser) GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//是否超管
		isSuperAdmin := false
		//获取无需验证权限的用户id
		s.NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
			if gconv.Uint64(v) == userId {
				isSuperAdmin = true
				return false
			}
			return true
		})
		//获取用户菜单数据
		allRoles, err := service.AdminRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err)
		roles, err := s.GetAdminRole(ctx, userId, allRoles)
		liberr.ErrIsNil(ctx, err)
		name := make([]string, len(roles))
		roleIds := make([]uint, len(roles))
		for k, v := range roles {
			name[k] = v.Name
			roleIds[k] = v.Id
		}
		//获取菜单信息
		if isSuperAdmin {
			//超管获取所有菜单
			permissions = []string{"*/*/*"}
			menuList, err = s.GetAllMenus(ctx)
			liberr.ErrIsNil(ctx, err)
		} else {
			menuList, err = s.GetAdminMenusByRoleIds(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
			permissions, err = s.GetPermissions(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// GetAdminRole 获取用户角色
func (s *sAdminUser) GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.AdminRole) (roles []*entity.AdminRole, err error) {
	var roleIds []uint
	roleIds, err = s.GetAdminRoleIds(ctx, userId)
	if err != nil {
		return
	}
	roles = make([]*entity.AdminRole, 0, len(allRoleList))
	for _, v := range allRoleList {
		for _, id := range roleIds {
			if id == v.Id {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

// GetAdminRoleIds 获取用户角色ids
func (s *sAdminUser) GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error) {
	enforcer, e := service.CasbinEnforcer(ctx)
	if e != nil {
		err = e
		return
	}
	//查询关联角色规则
	groupPolicy := enforcer.GetFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, userId))
	if len(groupPolicy) > 0 {
		roleIds = make([]uint, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			roleIds[k] = gconv.Uint(v[1])
		}
	}
	return
}

func (s *sAdminUser) GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error) {
	//获取所有开启的菜单
	var allMenus []*model.AdminAuthRuleInfoRes
	allMenus, err = service.AdminAuthRule().GetIsMenuList(ctx)
	if err != nil {
		return
	}
	menus = make([]*model.UserMenus, len(allMenus))
	for k, v := range allMenus {
		var menu *model.UserMenu
		menu = s.setMenuData(menu, v)
		menus[k] = &model.UserMenus{UserMenu: menu}
	}
	menus = s.GetMenusTree(menus, 0)
	return
}

func (s *sAdminUser) GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error) {
	//获取角色对应的菜单id
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			//查询当前权限
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				mid := gconv.Int64(p[1])
				menuIds[mid] = mid
			}
		}
		//获取所有开启的菜单
		allMenus, err := service.AdminAuthRule().GetIsMenuList(ctx)
		liberr.ErrIsNil(ctx, err)
		menus = make([]*model.UserMenus, 0, len(allMenus))
		for _, v := range allMenus {
			if _, ok := menuIds[gconv.Int64(v.Id)]; gstr.Equal(v.Condition, "nocheck") || ok {
				var roleMenu *model.UserMenu
				roleMenu = s.setMenuData(roleMenu, v)
				menus = append(menus, &model.UserMenus{UserMenu: roleMenu})
			}
		}
		menus = s.GetMenusTree(menus, 0)
	})
	return
}

func (s *sAdminUser) GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus {
	returnList := make([]*model.UserMenus, 0, len(menus))
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = s.GetMenusTree(menus, menu.Id)
			returnList = append(returnList, menu)
		}
	}
	return returnList
}

func (s *sAdminUser) setMenuData(menu *model.UserMenu, entity *model.AdminAuthRuleInfoRes) *model.UserMenu {
	menu = &model.UserMenu{
		Id:        entity.Id,
		Pid:       entity.Pid,
		Name:      gstr.CaseCamelLower(gstr.Replace(entity.Name, "/", "_")),
		Component: entity.Component,
		Path:      entity.Path,
		MenuMeta: &model.MenuMeta{
			Icon:        entity.Icon,
			Title:       entity.Title,
			IsLink:      "",
			IsHide:      entity.IsHide == 1,
			IsKeepAlive: entity.IsCached == 1,
			IsAffix:     entity.IsAffix == 1,
			IsIframe:    entity.IsIframe == 1,
		},
	}
	if menu.MenuMeta.IsIframe || entity.IsLink == 1 {
		menu.MenuMeta.IsLink = entity.LinkUrl
	}
	return menu
}

func (s *sAdminUser) GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//获取角色对应的菜单id
		enforcer, err := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			//查询当前权限
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				mid := gconv.Int64(p[1])
				menuIds[mid] = mid
			}
		}
		//获取所有开启的按钮
		allButtons, err := service.AdminAuthRule().GetIsButtonList(ctx)
		liberr.ErrIsNil(ctx, err)
		userButtons = make([]string, 0, len(allButtons))
		for _, button := range allButtons {
			if _, ok := menuIds[gconv.Int64(button.Id)]; gstr.Equal(button.Condition, "nocheck") || ok {
				userButtons = append(userButtons, button.Name)
			}
		}
	})
	return
}

// List 用户列表
func (s *sAdminUser) List(ctx context.Context, req *api.UserSearchReq) (total interface{}, userList []*entity.AdminUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminUser.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
		}
		if req.DeptId != "" {
			deptIds, e := s.getSearchDeptIds(ctx, gconv.Uint64(req.DeptId))
			liberr.ErrIsNil(ctx, e)
			m = m.Where("dept_id in (?)", deptIds)
		}
		if req.Status != "" {
			m = m.Where("user_status", gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where("mobile like ?", "%"+req.Mobile+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		err = m.FieldsEx(dao.AdminUser.Columns().UserPassword, dao.AdminUser.Columns().UserSalt).
			Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		liberr.ErrIsNil(ctx, err, "获取用户列表失败")
	})
	return
}

// GetUsersRoleDept 获取多个用户角色 部门信息
func (s *sAdminUser) GetUsersRoleDept(ctx context.Context, userList []*entity.AdminUser) (users []*model.AdminUserRoleDeptRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		allRoles, e := service.AdminRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, e)
		users = make([]*model.AdminUserRoleDeptRes, len(userList))
		for k, u := range userList {
			users[k] = &model.AdminUserRoleDeptRes{
				AdminUser: u,
			}
			roles, e := s.GetAdminRole(ctx, u.Id, allRoles)
			liberr.ErrIsNil(ctx, e)
			for _, r := range roles {
				users[k].RoleInfo = append(users[k].RoleInfo, &model.AdminUserRoleInfoRes{RoleId: r.Id, Name: r.Name})
			}
		}
	})
	return
}

func (s *sAdminUser) getSearchDeptIds(ctx context.Context, deptId uint64) (deptIds []uint64, err error) {
	return
}

func (s *sAdminUser) Add(ctx context.Context, req *api.UserAddReq) (err error) {
	err = s.UserNameOrMobileExists(ctx, req.UserName, req.Mobile)
	if err != nil {
		return
	}
	req.UserSalt = grand.S(10)
	req.Password = libUtils.EncryptPassword(req.Password, req.UserSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			userId, e := dao.AdminUser.Ctx(ctx).TX(tx).InsertAndGetId(do.AdminUser{
				UserName:     req.UserName,
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserPassword: req.Password,
				UserSalt:     req.UserSalt,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, e, "添加用户失败")
			e = s.addUserRole(ctx, req.RoleIds, userId)
			liberr.ErrIsNil(ctx, e, "设置用户权限失败")
			e = s.AddUserPost(ctx, tx, req.PostIds, userId)
			liberr.ErrIsNil(ctx, e)
		})
		return err
	})
	return
}

func (s *sAdminUser) Edit(ctx context.Context, req *api.UserEditReq) (err error) {
	err = s.UserNameOrMobileExists(ctx, "", req.Mobile, req.UserId)
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AdminUser.Ctx(ctx).TX(tx).WherePri(req.UserId).Update(do.AdminUser{
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, err, "修改用户信息失败")
			//设置用户所属角色信息
			err = s.EditUserRole(ctx, req.RoleIds, req.UserId)
			liberr.ErrIsNil(ctx, err, "设置用户权限失败")
			err = s.AddUserPost(ctx, tx, req.PostIds, req.UserId)
			liberr.ErrIsNil(ctx, err)
		})
		return err
	})
	return
}

// AddUserPost 添加用户岗位信息
func (s *sAdminUser) AddUserPost(ctx context.Context, tx gdb.TX, postIds []int64, userId int64) (err error) {
	return
}

// AddUserRole 添加用户角色信息
func (s *sAdminUser) addUserRole(ctx context.Context, roleIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		for _, v := range roleIds {
			_, e = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.casBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, e)
		}
	})
	return
}

// EditUserRole 修改用户角色信息
func (s *sAdminUser) EditUserRole(ctx context.Context, roleIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)

		//删除用户旧角色信息
		enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, userId))
		for _, v := range roleIds {
			_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.casBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sAdminUser) UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error {
	user := (*entity.AdminUser)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminUser.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.AdminUser.Columns().Id+" != ", id)
		}
		m = m.Where(fmt.Sprintf("%s='%s' OR %s='%s'",
			dao.AdminUser.Columns().UserName,
			userName,
			dao.AdminUser.Columns().Mobile,
			mobile))
		err := m.Limit(1).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if user == nil {
			return
		}
		if user.UserName == userName {
			liberr.ErrIsNil(ctx, gerror.New("用户名已存在"))
		}
		if user.Mobile == mobile {
			liberr.ErrIsNil(ctx, gerror.New("手机号已存在"))
		}
	})
	return err
}

// GetEditUser 获取编辑用户信息
func (s *sAdminUser) GetEditUser(ctx context.Context, id uint64) (res *api.UserGetEditRes, err error) {
	res = new(api.UserGetEditRes)
	err = g.Try(ctx, func(ctx context.Context) {
		//获取用户信息
		res.User, err = s.GetUserInfoById(ctx, id)
		liberr.ErrIsNil(ctx, err)
		//获取已选择的角色信息
		res.CheckedRoleIds, err = s.GetAdminRoleIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
		res.CheckedPosts, err = s.GetUserPostIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

// GetUserInfoById 通过Id获取用户信息
func (s *sAdminUser) GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.AdminUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if len(withPwd) > 0 && withPwd[0] {
			//用户用户信息
			err = dao.AdminUser.Ctx(ctx).Where(dao.AdminUser.Columns().Id, id).Scan(&user)
		} else {
			//用户用户信息
			err = dao.AdminUser.Ctx(ctx).Where(dao.AdminUser.Columns().Id, id).
				FieldsEx(dao.AdminUser.Columns().UserPassword, dao.AdminUser.Columns().UserSalt).Scan(&user)
		}
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")
	})
	return
}

// GetUserPostIds 获取用户岗位
func (s *sAdminUser) GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error) {
	return
}

// ResetUserPwd 重置用户密码
func (s *sAdminUser) ResetUserPwd(ctx context.Context, req *api.UserResetPwdReq) (err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminUser.Ctx(ctx).WherePri(req.Id).Update(g.Map{
			dao.AdminUser.Columns().UserSalt:     salt,
			dao.AdminUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "重置用户密码失败")
	})
	return
}

func (s *sAdminUser) ChangeUserStatus(ctx context.Context, req *api.UserStatusReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminUser.Ctx(ctx).WherePri(req.Id).Update(do.AdminUser{UserStatus: req.UserStatus})
		liberr.ErrIsNil(ctx, err, "设置用户状态失败")
	})
	return
}

// Delete 删除用户
func (s *sAdminUser) Delete(ctx context.Context, ids []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AdminUser.Ctx(ctx).TX(tx).Where(dao.AdminUser.Columns().Id+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除用户失败")
			//删除对应权限
			enforcer, e := service.CasbinEnforcer(ctx)
			liberr.ErrIsNil(ctx, e)
			for _, v := range ids {
				enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, v))
			}
		})
		return err
	})
	return
}

// GetUsers 通过用户ids查询多个用户信息
func (s *sAdminUser) GetUsers(ctx context.Context, ids []int) (users []*model.AdminUserSimpleRes, err error) {
	if len(ids) == 0 {
		return
	}
	idsSet := gset.NewIntSetFrom(ids).Slice()
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminUser.Ctx(ctx).Where(dao.AdminUser.Columns().Id+" in(?)", idsSet).
			Order(dao.AdminUser.Columns().Id + " ASC").Scan(&users)
	})
	return
}
