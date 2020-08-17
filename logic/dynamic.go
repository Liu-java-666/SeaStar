package logic

import (
	"VPartyServer/database"
	"VPartyServer/public"
	"fmt"
	"os"
)

//获取动态列表
func DynamicList(userid int, userkey, filetype, tag string, page int) interface{} {
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

	var tlist []*database.TDynamic
	var err error

	switch tag {
	case "hot":
		tlist, err = database.Dynamic_HotList(index, perpage, filetype)
	case "focus":
		tlist, err = database.Dynamic_FocusList(userid, index, perpage, filetype)
	default:
		tlist, err = database.Dynamic_AllList(index, perpage, filetype)
	}

	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagDynamicItem struct {
		Id				int
		UserId			int
		Nickname		string
		AvatarFile		string
		Sex				int
		Age				int
		PostTime		int
		Description		string
		Imglist			[]string
		VideoUrl		string
		VideoCover		string
		VideoRotation	int
		IsLike			bool
		LikeNum			int
		CommentNum		int
		GiftValue		int
	}

	dylist := []tagDynamicItem{}
	for _, v := range tlist {
		Item := tagDynamicItem{
			Id:			v.Id,
			UserId:      v.User_id,
			PostTime:    public.StrToTimestamp(string(v.Post_time)),
			Description: v.Description,
			IsLike:		 database.DynamicLike_Get(userid, v.Id),
			LikeNum:	 database.DynamicLike_GetCount(v.Id),
			CommentNum:	 database.DynamicComment_GetCount(v.Id),
			GiftValue:	 database.GiftLog_GetValue("dynamic", v.Id),
		}

		fileidlist := public.GetFileIdList(v.Filelist)
		if len(fileidlist) == 0 {
			continue
		}
		if v.Filetype == "video" {
			file, _ := database.Video_Get(fileidlist[0])
			if file != nil {
				Item.VideoUrl = MakeVideoUrl(file.File_name)
				Item.VideoCover = MakeVideoUrl(file.Cover_name)
				Item.VideoRotation = file.Rotation
			}
		} else {
			for _, v2 := range fileidlist {
				file, _ := database.Image_Get(v2)
				if file != nil {
					Item.Imglist = append(Item.Imglist, MakeImageUrl(file.File_name))
				}
			}
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
			Item.Sex = tu.Sex
			Item.Age = tu.Age
		}

		dylist = append(dylist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagDynamicItem
	}{
		true,
		len(dylist) < perpage,
		dylist,
	}

	return result
}

//点赞/取消点赞动态
func DynamicLike(userid int, userkey string, dynamicid, action int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	t, err := database.Dynamic_Get(dynamicid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if t == nil {
		return ErrorResult("动态不存在")
	}

	if t.Is_audit <= 0 {
		return ErrorResult("动态未审核")
	}

	bLike := database.DynamicLike_Get(userid, dynamicid)
	if bLike && action > 0 {
		return ErrorResult("重复点赞")
	} else if !bLike && action <= 0 {
		return ErrorResult("未点赞过无需取消")
	}

	if action > 0 {
		err = database.DynamicLike_Insert(userid, dynamicid, t.User_id)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, t.User_id, 1, 0, t.Id, "dynamic_like", "点赞了你的动态")
	} else {
		err = database.DynamicLike_Delete(userid, dynamicid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, t.User_id, 0, 0, t.Id, "dynamic_like", "对你的动态取消了点赞")
	}

	//返回成功消息
	result := struct{
		Result bool
		LikeNum int
	}{
		true,
		database.DynamicLike_GetCount(dynamicid),
	}

	return result
}

//获取评论列表
func DynamicCommentList(userid int, userkey string, dynamicid, page int) interface{} {
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

	tlist, err := database.DynamicComment_GetList(dynamicid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagComment struct {
		Id int
		UserId int
		Nickname string
		AvatarFile string
		Content string
		Cdate int
		IsLike bool
		LikeNum int
	}

	colist := []tagComment{}
	for _, v := range tlist {
		Item := tagComment{
			Id:			v.Id,
			Content:    v.Content,
			Cdate:      public.StrToTimestamp(string(v.Cdate)),
			IsLike:		database.DynamicCommentLike_Get(userid, v.Id),
			LikeNum:	database.DynamicCommentLike_GetCount(v.Id),
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.UserId = tu.Id
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
		}

		colist = append(colist, Item)
	}

	result := struct{
		Result	bool
		CommentNum int
		IsEnd	bool
		Data	[]tagComment
	}{
		true,
		database.DynamicComment_GetCount(dynamicid),
		len(colist) < perpage,
		colist,
	}

	return result
}

//评论动态
func DynamicComment(userid int, userkey string, dynamicid int, content string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	t, err := database.Dynamic_Get(dynamicid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if t == nil {
		return ErrorResult("动态不存在")
	}

	if userid == t.User_id {
		return ErrorResult("不能给自己评论哦")
	}

	if t.Is_audit <= 0 {
		return ErrorResult("动态未审核")
	}

	if content == "" {
		return ErrorResult("评论不能为空")
	}

	err = database.DynamicComment_Insert(dynamicid,t.User_id, userid, content)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//操作日志
	database.ActionLog_Insert(userid, t.User_id, 1, 0, t.Id, "dynamic_comment", "评论了你的动态")

	//返回成功消息
	result := struct{
		Result bool
		CommentNum int
	}{
		true,
		database.DynamicComment_GetCount(dynamicid),
	}

	return result
}

//点赞/取消点赞评论
func DynamicLikeComment(userid int, userkey string, commentid, action int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	t, err := database.DynamicComment_Get(commentid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if t == nil {
		return ErrorResult("动态不存在")
	}

	bLike := database.DynamicCommentLike_Get(userid, commentid)
	if bLike && action > 0 {
		return ErrorResult("重复点赞")
	} else if !bLike && action <= 0 {
		return ErrorResult("未点赞过无需取消")
	}

	if action > 0 {
		err = database.DynamicCommentLike_Insert(userid, commentid, t.User_id, t.Dynamic_id)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, t.User_id, 1, 0, t.Id, "dynamic_like", "点赞了你的动态")
	} else {
		err = database.DynamicCommentLike_Delete(userid, commentid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, t.User_id, 0, 0, t.Id, "dynamic_like", "对你的动态取消了点赞")
	}

	//返回成功消息
	result := struct{
		Result bool
		LikeNum int
	}{
		true,
		database.DynamicCommentLike_GetCount(commentid),
	}

	return result
}

//发布动态
func DynamicPost(userid int, userkey string, description, filetype string, filelist []int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if filetype != "image" && filetype != "video" {
		return ErrorResult("文件类型有误")
	}

	if len(filelist) == 0 {
		return ErrorResult("至少需要1张照片或1个视频")
	}

	if filetype == "video" && len(filelist) > 1 {
		return ErrorResult("最多只能上传1个视频")
	} else if filetype == "image" && len(filelist) > 9 {
		return ErrorResult("最多只能上传9张照片")
	}

	fileliststr := ""
	for k, v := range filelist {
		if v <= 0 {
			return ErrorResult("文件ID有误")
		}

		if filetype == "image" {
			timg, _ := database.Image_Get(v)
			if timg == nil {
				return ErrorResult("照片ID有误")
			}
		} else {
			tvideo, _ := database.Video_Get(v)
			if tvideo == nil {
				return ErrorResult("视频ID有误")
			}
		}

		if k == 0 {
			fileliststr = fmt.Sprintf("%d", v)
		} else {
			fileliststr = fmt.Sprintf("%s,%d", fileliststr, v)
		}
	}

	err := database.Dynamic_Insert(userid, description, "", filetype, fileliststr)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result 	bool
	}{
		true,
	}

	return result
}

//用户动态列表
func DynamicUserList(userid int, userkey string, touserid, page int, filetype string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 10000
	index := page * perpage
	if index < 0 {
		index = 0
	}

	var tlist []*database.TDynamic
	var err error
	if userid == touserid {
		tlist, err = database.Dynamic_MyList(userid, index, perpage, filetype)
	} else {
		tlist, err = database.Dynamic_UserList(touserid, index, perpage, filetype)
	}
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagDynamicItem struct {
		Id				int
		PostTime		int
		Description		string
		Imglist			[]string
		VideoUrl		string
		VideoCover		string
		VideoRotation	int
		IsLike			bool
		LikeNum			int
		CommentNum		int
		GiftValue 		int
		IsAudit 		int
	}

	dylist := []tagDynamicItem{}
	for _, v := range tlist {
		Item := tagDynamicItem{
			Id:			 v.Id,
			PostTime:    public.StrToTimestamp(string(v.Post_time)),
			Description: v.Description,
			IsLike:      database.DynamicLike_Get(userid, v.Id),
			LikeNum:	 database.DynamicLike_GetCount(v.Id),
			CommentNum:	 database.DynamicComment_GetCount(v.Id),
			GiftValue:   database.GiftLog_GetValue("dynamic", v.Id),
			IsAudit:     v.Is_audit,
		}

		fileidlist := public.GetFileIdList(v.Filelist)
		if len(fileidlist) == 0 {
			continue
		}
		if v.Filetype == "video" {
			file, _ := database.Video_Get(fileidlist[0])
			if file != nil {
				Item.VideoUrl = MakeVideoUrl(file.File_name)
				Item.VideoCover = MakeVideoUrl(file.Cover_name)
				Item.VideoRotation = file.Rotation
			}
		} else {
			for _, v2 := range fileidlist {
				file, _ := database.Image_Get(v2)
				if file != nil {
					Item.Imglist = append(Item.Imglist, MakeImageUrl(file.File_name))
				}
			}
		}

		dylist = append(dylist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagDynamicItem
	}{
		true,
		len(dylist) < perpage,
		dylist,
	}

	return result
}

//删除动态
func DynamicDelete(userid int, userkey string, dynamicid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	t, err := database.Dynamic_Get(dynamicid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if t == nil {
		return ErrorResult("动态不存在")
	}

	if t.User_id != userid {
		return ErrorResult("不能删除别人的动态")
	}

	err = t.Delete()
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//删除视频和图片
	fidlist := public.GetFileIdList(t.Filelist)
	for _, v := range fidlist {
		if t.Filetype == "video" {
			tv, _ := database.Video_Get(v)
			if tv != nil {
				tv.Delete()
				filepath := MakeVideoPath(tv.File_name)
				os.Remove(filepath)
				filepath = MakeVideoPath(tv.Cover_name)
				os.Remove(filepath)
			}
		} else {
			ti, _ := database.Image_Get(v)
			if ti != nil {
				ti.Delete()
				filepath := MakeImagePath(ti.File_name)
				os.Remove(filepath)
			}
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