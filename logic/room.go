package logic

import (
	"VPartyServer/database"
	"VPartyServer/im"
	"VPartyServer/public"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"sort"
)

//获取房间列表
func RoomList(userid int, userkey string, roomtype, page int, tag string) interface{} {
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

	var tlist []*database.TRoom
	var err error

	switch tag {
	case "hot":
		tlist, err = database.Room_HotList(index, perpage, roomtype)
	case "focus":
		tlist, err = database.Room_FocusList(userid, index, perpage, roomtype)
	default:
		tlist, err = database.Room_AllList(index, perpage, roomtype)
	}

	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagRoomItem struct {
		UserId			int
		RoomId			int
		AvatarFile		string
		RoomName		string
		LikeNum			int
	}

	rlist := []tagRoomItem{}
	for _, v := range tlist {
		Item := tagRoomItem{
			UserId: v.User_id,
			RoomId: v.Id,
			RoomName: v.Room_name,
			LikeNum: v.Like_num,
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
		}

		rlist = append(rlist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagRoomItem
	}{
		true,
		len(rlist) < perpage,
		rlist,
	}

	return result
}

//进入房间
func RoomEnter(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	tu, err := database.User_GetById(r.User_id, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("房主不存在")
	}

	tru, err := database.RoomUser_Get(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	if tru == nil {
		//保存房间号
		database.RoomUser_Insert(userid, roomid, r.Im_group)
	} else if tru.Room_id != roomid {
		//退出上个房间群
		if tru.Room_id > 0 {
			im.DeleteGroupMember(tru.Im_group, userid)
		}
		//保存现在的房间号
		tru.SetRoom(roomid, r.Im_group)
	}

	//加IM群
	go im.AddGroupMember(r.Im_group, ck.User.Nickname, userid, OnAddGroupMember)

	//返回成功消息
	result := struct{
		Result 			bool
		ImGroup 		string
		UserId			int
		Nickname		string
		AvatarFile		string
		GiftValue 		int
		LikeNum 		int
	}{
		true,
		r.Im_group,
		r.User_id,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		database.GiftLog_GetValue("room", roomid),
		r.Like_num,
	}

	return result
}

//增加群组成员结果
func OnAddGroupMember(resultData string, err error, user_data interface{}) {
	userdata := user_data.([]string)

	if err != nil {
		logger.Error("增加群组成员失败,group=%s,account=%s,err=%v", userdata[0], userdata[1], err)
		return
	}

	//logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("增加群组成员失败,group=%s,account=%s,err=%v,resultData=%v", userdata[0], userdata[1], err, resultData)
		return
	}

	//logger.Debug(revData)

	ActionStatus := revData["ActionStatus"].(string)
	if ActionStatus != "OK" {
		ErrorCode := int(revData["ErrorCode"].(float64))
		ErrorInfo := revData["ErrorInfo"].(string)
		logger.Error("增加群组成员失败,group=%s,account=%s,errcode=%d,errinfo=%s", userdata[0], userdata[1], ErrorCode, ErrorInfo)
		return
	}

	//发送欢迎消息
	msg := fmt.Sprintf("%s来了！", userdata[2])
	im.SendGroupSysNotice(userdata[0], 0, msg)
}

//退出房间
func RoomLeave(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, _ := database.Room_Get(roomid)
	if r != nil {
		//退IM群
		im.DeleteGroupMember(r.Im_group, userid)
	}

	tru, _ := database.RoomUser_Get(userid)
	if tru != nil {
		//清除房间号
		tru.SetRoom(0, "")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//申请创建房间
func RoomCreate(userid int, userkey string, roomtype int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if roomtype < 0 || roomtype > 1 {
		return ErrorResult("房间类型错误")
	}

	err := database.Room_Insert(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//点赞房间
func RoomLike(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	err = database.Room_Like(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	RoomUpdateData(r)

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//申请上座
func RoomSeat(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	result := struct{
		Result bool
		IsRepeat bool
	}{
		true,
		false,
	}

	if database.RoomSeat_Get(userid, roomid) {
		result.IsRepeat = true
		return result
	}

	err = database.RoomSeat_Insert(userid, roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	return result
}

//房间在线人数更新
func RoomUpdateOnline(imGroup string) {
	r, _ := database.Room_Find(imGroup)
	if r == nil {
		logger.Error("查询房间失败，imGroup=", imGroup)
		return
	}

	im.GroupInfo(imGroup, OnGroupInfo)
}

type tagMember struct {
	Account string
	JoinTime int
}
type memberListSlice []tagMember
func (s memberListSlice) Len() int {return len(s)}
func (s memberListSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s memberListSlice) Less(i, j int) bool { return s[i].JoinTime < s[j].JoinTime }

//获取群详细资料结果
func OnGroupInfo(resultData string, err error, user_data interface{}) {
	imGroup := user_data.(string)

	if err != nil {
		logger.Error("获取群详细资料失败,group=%s,err=%v", imGroup, err)
		return
	}

	//logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("获取群详细资料失败,group=%s,err=%v,resultData=%v", imGroup, err, resultData)
		return
	}

	//logger.Debug(revData)

	ActionStatus := revData["ActionStatus"].(string)
	if ActionStatus != "OK" {
		ErrorCode := int(revData["ErrorCode"].(float64))
		ErrorInfo := revData["ErrorInfo"].(string)
		logger.Error("获取群详细资料失败,group=%s,errcode=%d,errinfo=%s", imGroup, ErrorCode, ErrorInfo)
		return
	}

	r, _ := database.Room_Find(imGroup)
	if r == nil {
		logger.Error("获取群详细资料结果错误：房间不存在，group=%s", imGroup)
		return
	}

	GroupInfoList, ok := revData["GroupInfo"].([]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回GroupInfoList数据有误,group=%s", imGroup)
		return
	}
	if len(GroupInfoList) <= 0 {
		logger.Error("获取群详细资料结果错误：返回GroupInfo数据为空,group=%s", imGroup)
		return
	}
	GroupInfo, ok := GroupInfoList[0].(map[string]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回GroupInfo数据有误,group=%s", imGroup)
		return
	}

	MemberNum := int(GroupInfo["MemberNum"].(float64))
	logger.Debug("MemberNum=", MemberNum)

	MemberList, ok := GroupInfo["MemberList"].([]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回MemberList数据有误,group=%s", imGroup)
		return
	}

	memberList := memberListSlice{}
	for _, v := range MemberList {
		item, ok := v.(map[string]interface{})
		if !ok {
			logger.Error("获取群详细资料结果错误：返回Member数据有误,group=%s", imGroup)
			continue
		}
		memberList = append(memberList, tagMember{
			Account:  item["Member_Account"].(string),
			JoinTime: int(item["JoinTime"].(float64)),
		})
	}
	sort.Sort(sort.Reverse(memberList))

	type tagOnlineUser struct {
		UserId int
		AvatarFile string
	}
	onlineList := []tagOnlineUser{}
	for _, v := range memberList {
		userid := public.GetIMUserID(v.Account)
		if userid <= 0 {
			continue
		}
		tu, _ := database.User_GetById(userid, false)
		if tu == nil {
			continue
		}
		onlineList = append(onlineList, tagOnlineUser{
			UserId:     userid,
			AvatarFile: MakeImageUrl(tu.AvatarFile),
		})
		if len(onlineList) >= 5 {
			break
		}
	}

	//发送在线列表消息
	msgdata := struct{
		RoomID int
		OnlineUserCnt int
		OnlineUserList []tagOnlineUser
	}{
		r.Id,
		len(onlineList),
		onlineList,
	}
	go im.SendGroupSysNotice(imGroup, 2, msgdata)
}

//更新房间数据
func RoomUpdateData(r *database.TRoom) {
	//发送房间更新消息
	msgdata := struct{
		RoomID int
		GiftValue int
		LikeNum int
	}{
		r.Id,
		database.GiftLog_GetValue("room", r.Id),
		r.Like_num,
	}
	go im.SendGroupSysNotice(r.Im_group, 3, msgdata)
}