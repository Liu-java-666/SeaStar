package logic

import (
	"VPartyServer/config"
	"VPartyServer/database"
	"fmt"
)

type tagError struct {
	Result bool
	ErrCode int
	ErrMsg string
}

type CheckUserResult struct {
	Result bool
	Error tagError
	User *database.TUser
}

//检查用户
func CheckUser(userid int, userkey string) CheckUserResult {
	result := CheckUserResult{}

	//读取用户数据
	tuser, err := database.User_GetById(userid, true)
	if err != nil {
		result.Error = ErrorResult("数据库异常")
		return result
	}

	//判断用户
	if tuser == nil {
		result.Error = ErrorResult("用户不存在")
		return result
	}

	//判断userkey
	if tuser.User_key != userkey {
		result.Error = ErrorResultSpecial(1,"用户数据异常，请重新登录")
		return result
	}

	//返回成功消息
	result.Result = true
	result.Error.Result = true
	result.User = tuser

	return result
}

func ErrorResult(msg string) tagError {
	data := tagError{
		Result: false,
		ErrMsg: msg,
	}
	return data
}

func ErrorResultSpecial(code int, msg string) tagError {
	data := tagError{
		Result: false,
		ErrCode:code,
		ErrMsg: msg,
	}
	return data
}

//生成图片地址
func MakeImageUrl(filename string) string {
	if filename == "" {
		return ""
	}
	return fmt.Sprintf("%s/image/%s", config.GetUploadPath(), filename)
}

//生成图片路径
func MakeImagePath(filename string) string {
	if filename == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.GetUploadRoot(), MakeImageUrl(filename))
}

//生成视频地址
func MakeVideoUrl(filename string) string {
	if filename == "" {
		return ""
	}
	return fmt.Sprintf("%s/video/%s", config.GetUploadPath(), filename)
}

//生成视频路径
func MakeVideoPath(filename string) string {
	if filename == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.GetUploadRoot(), MakeVideoUrl(filename))
}