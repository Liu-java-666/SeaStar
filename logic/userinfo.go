package logic

import (
	"VPartyServer/database"
	"VPartyServer/public"
	"fmt"
)

//检查生日参数
func CheckBirthday(year, month, day int) bool {
	if year <= 1900 || year >= 2020 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}

	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}

	if month == 2 && day > 29 {
		return false
	}

	if month == 2 && (year % 4 != 0 || year % 100 == 0) && day > 28 {
		return false
	}

	return true
}

//设置信息
func SetInfo(userid int, userkey, nickname string, sex, year, month, day int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//昵称已存在就不能修改
	if ck.User.Nickname != "" {
		return ErrorResult("不能修改性别和生日")
	}

	//检查参数
	if nickname == "" {
		return ErrorResult("昵称不能为空")
	}
	if sex != 0 && sex != 1 {
		return ErrorResult("性别参数错误")
	}
	if CheckBirthday(year, month, day) == false {
		return ErrorResult("生日参数错误")
	}

	//保存资料
	err := ck.User.SetInfo(nickname, fmt.Sprintf("%04d-%02d-%02d", year, month, day), sex)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}

//编辑信息
func EditInfo(userid int, userkey, nickname, signature, status, purpose, hobbies string, photolist []int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//检查参数
	if nickname == "" {
		return ErrorResult("昵称不能为空")
	}

	//现有照片
	plist := ""
	t, err := database.PhotoList_MyList(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if t != nil {
		plist = t.Photolist
	}
	plistnew := public.MakeFileIdList(photolist)

	//没有修改
	if ck.User.Nickname == nickname && ck.User.Signature == signature && ck.User.Relationship_status == status &&
		ck.User.Friends_purpose == purpose && ck.User.Hobbies == hobbies && plist == plistnew {
		return ErrorResult("没有修改")
	}

	//保存资料
	err = ck.User.UpdateInfo(nickname, signature, status, purpose, hobbies)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	//保存照片列表
	if plist != plistnew {
		_, err = database.PhotoList_Insert(userid, plistnew)
		if err != nil {
			return ErrorResult("数据库异常")
		}
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}

//我的菜单
func GetMyMenu(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Coins int
		FocusNum int
		FansNum int
		LikeNum int
		GiftList []int
		IsGiftMore bool
	}{
		true,
		ck.User.Nickname,
		MakeImageUrl(ck.User.AvatarFile),
		ck.User.AvatarAudit,
		ck.User.Sex,
		ck.User.Age,
		ck.User.Coins,
		database.FocusList_GetFocus(userid),
		database.FocusList_GetFans(userid),
		database.User_GetLikeCount(userid),
		[]int{},
		false,
	}

	tlist, _ := database.GiftLog_LastList(userid, 6)
	for _, v := range tlist {
		result.GiftList = append(result.GiftList, v.Gift_id)
		if len(result.GiftList) >= 5 {
			break
		}
	}
	result.IsGiftMore = len(tlist) > 5

	return result
}

//我的资料
func GetMyInfo(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	type tagPhotoItem struct {
		Id int
		Url string
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Birthday string
		Signature string
		Status string
		Purpose	string
		Hobbies	string
		PhotoList []tagPhotoItem
	}{
		true,
		ck.User.Nickname,
		MakeImageUrl(ck.User.AvatarFile),
		ck.User.AvatarAudit,
		ck.User.Sex,
		ck.User.Age,
		string(ck.User.Birthday),
		ck.User.Signature,
		ck.User.Relationship_status,
		ck.User.Friends_purpose,
		ck.User.Hobbies,
		[]tagPhotoItem{},
	}

	to, _ := database.PhotoList_MyList(userid)
	if to != nil {
		pidlist := public.GetFileIdList(to.Photolist)
		for _, v := range pidlist {
			tp, _ := database.Image_Get(v)
			if tp != nil {
				Item := tagPhotoItem{
					Id:  v,
					Url: MakeImageUrl(tp.File_name),
				}
				result.PhotoList = append(result.PhotoList, Item)
			}
		}
	}

	return result
}

//用户详情
func GetUserDetail(userid int, userkey string, touserid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	var tu *database.TUser
	var err error
	if userid == touserid {
		tu = ck.User
	} else {
		tu, err = database.User_GetById(touserid, false)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		if tu == nil {
			return ErrorResult("用户不存在")
		}
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Birthday string
		Signature string
		Status string
		Purpose	string
		Hobbies	string
		IsFocus bool
		FocusNum int
		FansNum int
		LikeNum int
		CoinsUsed int
		PhotoList []string
		GiftList []int
		IsGiftMore bool
		DynamicVideoNum int
		DynamicVideoList []string
		DynamicImageNum int
		DynamicImageList []string
	}{
		true,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		tu.AvatarAudit,
		tu.Sex,
		tu.Age,
		string(tu.Birthday),
		tu.Signature,
		tu.Relationship_status,
		tu.Friends_purpose,
		tu.Hobbies,
		database.FocusList_Get(userid, touserid),
		database.FocusList_GetFocus(touserid),
		database.FocusList_GetFans(touserid),
		database.User_GetLikeCount(touserid),
		tu.Coins_used,
		[]string{},
		[]int{},
		false,
		database.Dynamic_GetCount(touserid, "video"),
		[]string{},
		database.Dynamic_GetCount(touserid, "image"),
		[]string{},
	}
	if tu.Photolist_id > 0 {
		to, _ := database.PhotoList_Get(tu.Photolist_id)
		if to != nil {
			pidlist := public.GetFileIdList(to.Photolist)
			for _, v := range pidlist {
				tp, _ := database.Image_Get(v)
				if tp != nil {
					result.PhotoList = append(result.PhotoList, MakeImageUrl(tp.File_name))
				}
			}
		}
	}

	tlist, _ := database.GiftLog_LastList(touserid, 6)
	for _, v := range tlist {
		result.GiftList = append(result.GiftList, v.Gift_id)
		if len(result.GiftList) >= 5 {
			break
		}
	}
	result.IsGiftMore = len(tlist) > 5

	if result.DynamicVideoNum > 0 {
		dylist, _ := database.Dynamic_UserList(touserid, 0, 3, "video")
		for _, v := range dylist {
			pidlist := public.GetFileIdList(v.Filelist)
			for _, v2 := range pidlist {
				tv, _ := database.Video_Get(v2)
				if tv != nil {
					result.DynamicVideoList = append(result.DynamicVideoList, MakeVideoUrl(tv.Cover_name))
				}
			}
		}
	}

	if result.DynamicImageNum > 0 {
		dylist, _ := database.Dynamic_UserList(touserid, 0, 1, "image")
		if len(dylist) > 0 {
			pidlist := public.GetFileIdList(dylist[0].Filelist)
			for _, v := range pidlist {
				ti, _ := database.Image_Get(v)
				if ti != nil {
					result.DynamicImageList = append(result.DynamicImageList, MakeImageUrl(ti.File_name))
				}
			}
		}
	}

	return result
}

//用户名片
func GetUserCard(userid int, userkey string, touserid int) interface{} {
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

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Signature string
		IsFocus bool
		FocusNum int
		FansNum int
		LikeNum int
		CoinsUsed int
	}{
		true,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		tu.AvatarAudit,
		tu.Sex,
		tu.Age,
		tu.Signature,
		database.FocusList_Get(userid, touserid),
		database.FocusList_GetFocus(touserid),
		database.FocusList_GetFans(touserid),
		database.User_GetLikeCount(touserid),
		tu.Coins_used,
	}

	return result
}

//批量用户详情
func GetUserInfoList(userid int, userkey string, touseridlist []int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	type tagUserData struct {
		UserId int
		Nickname string
		AvatarFile string
		IsFocus bool
	}

	UserInfoList := []tagUserData{}
	for _, v := range touseridlist {
		Item := tagUserData{
			UserId:     v,
		}
		tu, err := database.User_GetById(v, false)
		if err == nil && tu != nil {
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
			Item.IsFocus = database.FocusList_Get(userid, v)
		}
		UserInfoList = append(UserInfoList, Item)
	}

	//返回成功消息
	result := struct{
		Result bool
		UserInfoList []tagUserData
	}{
		true,
		UserInfoList,
	}

	return result
}

//是否拉黑
func IsBlacklist(userid int, userkey string, touserid int) interface{} {
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

	//返回成功消息
	result := struct{
		Result bool
		IsBlacklist bool
		IsBeBlacklist bool
	}{
		true,
		database.Blacklist_Get(userid, touserid),
		database.Blacklist_Get(touserid, userid),
	}

	return result
}

//获取关注列表
func GetFocusList(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.FocusList_GetFocusList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取粉丝列表
func GetFansList(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.FocusList_GetFansList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		IsFocus		bool
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
				database.FocusList_Get(userid, tu.Id),
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取黑名单
func GetBlacklist(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.Blacklist_GetList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取好友列表
func GetFriendList(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.FocusList_GetFriendList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Signature	string
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Signature,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取好友申请列表
func GetApplyList(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.FocusNotice_GetList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Cdate		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				public.StrToTimestamp(string(v.Cdate)),
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取收礼列表
func GetReceiveGiftList(userid int, userkey string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}
	tlist, err := database.GiftLog_List(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		GiftId		int
		Cdate		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		Item := tagUserData{
			UserId: v.User_id,
			GiftId: v.Gift_id,
			Cdate:  public.StrToTimestamp(string(v.Cdate)),
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
			Item.Sex = tu.Sex
			Item.Age = tu.Age
		}

		userlist = append(userlist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取排行榜
func GetRankList(userid int, userkey, tag string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 10
	index := page * perpage
	if index < 0 {
		index = 0
	}

	var tlist []*database.RankData
	var err error

	switch tag {
	case "star":
		tlist, err = database.User_StarList(index, perpage)
	case "charm":
		tlist, err = database.User_CharmList(index, perpage)
	default:
		tlist, err = database.User_RichList(index, perpage)
	}

	if err != nil {
		return ErrorResult("数据库异常")
	}


	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		Amount		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		Item := tagUserData{
			v.Id,
			v.Nickname,
			MakeImageUrl(v.AvatarFile),
			v.Sex,
			v.Age,
			v.Num,
		}
		userlist = append(userlist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}