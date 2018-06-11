package consts

import "muxin.io/chronos/models"

const (
	SessionMaxAge = 3600
)

var Success = models.Error{ErrCode: 0, ErrMsg: ""}
var NeedLogin = models.Error{ErrCode: 1, ErrMsg: "请先登录"}
var WrongInput = models.Error{ErrCode: 2, ErrMsg: "输入有误"}
