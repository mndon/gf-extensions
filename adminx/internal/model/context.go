/*
* @desc:context-model
* @company:
* @Author:
* @Date:   2022/3/16 14:45
 */

package model

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*LoginUserRes
}
