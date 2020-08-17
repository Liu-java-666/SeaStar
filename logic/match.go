package logic

import (
	"VPartyServer/database"
)

//1V1
func GetMatchUser(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	list, err := database.User_GetMatchUser(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}


	type User struct{
		UserId int
		Nickname string
		AvatarFile string //头像地址
		Sex int //（0女，1男）
	}

	//返回成功消息
	result := struct{
		Result bool
		UserList []User

	}{
		Result: true,
	}


	for _,v := range list {
		item := User{
			v.Id,
			v.Nickname,
			MakeImageUrl(v.AvatarFile),
			v.Sex,
		}
		result.UserList = append(result.UserList,item)
	}

	//写日志
	//database.MatchLog_Insert(userid, t.Id)


	return result
}

//打电话
func CallUp(userid int, userkey string, touserid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	tu, err := database.User_GetById(touserid, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("用户不存在")
	}

	if database.Blacklist_Get(userid, touserid) {
		return ErrorResult("你已将TA拉入黑名单")
	} else if database.Blacklist_Get(touserid, userid) {
		return ErrorResult("TA已将您拉入黑名单")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}

//挂电话
func HangUp(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}