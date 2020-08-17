package logic

import "VPartyServer/database"

type queItem struct {
	action string
	param map[string]interface{}
	retchan chan interface{}
}

var QueChan chan queItem

func PushQue(action string, param map[string]interface{}, retchan chan interface{}) {
	QueChan <- queItem {
		action: action,
		param:   param,
		retchan: retchan,
	}
}

func Queue() {
	QueChan = make(chan queItem, 500)

	for {
		data := <- QueChan
		if data.retchan == nil {
			continue
		}

		switch data.action {
		case "CheckUser":	//检查用户
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := CheckUser(userid, userkey)
			data.retchan <- result		
		case "GetCaptcha":	//获取验证码
			phone := data.param["phone"].(string)
			result := GetCaptcha(phone)
			data.retchan <- result
		case "Login":	//登录/注册
			phone := data.param["phone"].(string)
			captcha := data.param["captcha"].(string)
			ip := data.param["ip"].(string)
			result := Login(phone, captcha, ip)
			data.retchan <- result
		case "SetInfo":		//设置信息
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			nickname := data.param["nickname"].(string)
			sex := data.param["sex"].(int)
			year := data.param["year"].(int)
			month := data.param["month"].(int)
			day := data.param["day"].(int)
			result := SetInfo(userid, userkey, nickname, sex, year, month, day)
			data.retchan <- result
		case "EditInfo":	//编辑信息
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			nickname := data.param["nickname"].(string)
			signature := data.param["signature"].(string)
			status := data.param["status"].(string)
			purpose := data.param["purpose"].(string)
			hobbies := data.param["hobbies"].(string)
			photolist := data.param["photolist"].([]int)
			result := EditInfo(userid, userkey, nickname, signature, status, purpose, hobbies, photolist)
			data.retchan <- result
		case "GetMyMenu":	//我的菜单
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := GetMyMenu(userid, userkey)
			data.retchan <- result
		case "GetMyInfo":	//我的资料
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := GetMyInfo(userid, userkey)
			data.retchan <- result
		case "GetUserDetail":		//用户详情
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			result := GetUserDetail(userid, userkey, touserid)
			data.retchan <- result
		case "GetUserCard":		//用户名片
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			result := GetUserCard(userid, userkey, touserid)
			data.retchan <- result
		case "GetUserInfoList":		//批量用户详情
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touseridlist := data.param["touseridlist"].([]int)
			result := GetUserInfoList(userid, userkey, touseridlist)
			data.retchan <- result
		case "IsBlacklist":		//是否拉黑
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			result := IsBlacklist(userid, userkey, touserid)
			data.retchan <- result
		case "GetFocusList":	//获取关注列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetFocusList(userid, userkey, page)
			data.retchan <- result
		case "GetFansList":		//获取粉丝列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetFansList(userid, userkey, page)
			data.retchan <- result
		case "GetBlacklist":	//获取黑名单
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetBlacklist(userid, userkey, page)
			data.retchan <- result
		case "GetFriendList":	//获取好友列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetFriendList(userid, userkey, page)
			data.retchan <- result
		case "GetApplyList":	//获取好友申请列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetApplyList(userid, userkey, page)
			data.retchan <- result
		case "GetReceiveGiftList":	//获取收礼列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			page := data.param["page"].(int)
			result := GetReceiveGiftList(userid, userkey, page)
			data.retchan <- result
		case "GetMyWallet":	//我的钱包
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := GetMyWallet(userid, userkey)
			data.retchan <- result
		case "PayOrder":	//提交充值订单
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			money := data.param["money"].(int)
			result := PayOrder(userid, userkey, money)
			data.retchan <- result
		case "PayFinish":	//完成充值
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			orderid := data.param["orderid"].(string)
			status := data.param["status"].(int)
			result := PayFinish(userid, userkey, orderid, status)
			data.retchan <- result
		case "SendGift":		//送礼物
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			scene := data.param["scene"].(string)
			sceneid := data.param["sceneid"].(int)
			giftid := data.param["giftid"].(int)
			result := SendGift(userid, userkey, scene, sceneid, giftid)
			data.retchan <- result
		case "UploadImage":	//上传图片
			user := data.param["user"].(*database.TUser)
			file := data.param["file"].(string)
			filetype := data.param["filetype"].(string)
			usetype := data.param["usetype"].(string)
			index := data.param["index"].(int)
			result := UploadImage(user, file, filetype, usetype, index)
			data.retchan <- result
		case "UploadVideo":		//上传视频
			user := data.param["user"].(*database.TUser)
			file := data.param["file"].(string)
			filetype := data.param["filetype"].(string)
			cover := data.param["cover"].(string)
			covertype := data.param["covertype"].(string)
			usetype := data.param["usetype"].(string)
			rotation := data.param["rotation"].(int)
			index := data.param["index"].(int)
			result := UploadVideo(user, file, filetype, cover, covertype, usetype, rotation, index)
			data.retchan <- result
		case "GetMatchUser":	//1V1
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := GetMatchUser(userid, userkey)
			data.retchan <- result
		case "CallUp":	//打电话
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			result := CallUp(userid, userkey, touserid)
			data.retchan <- result
		case "HangUp":	//挂电话
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			result := HangUp(userid, userkey)
			data.retchan <- result
		case "GetRankList":	//排行榜
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			tag := data.param["tag"].(string)
			page := data.param["page"].(int)
			result := GetRankList(userid, userkey, tag, page)
			data.retchan <- result
		case "Focus":	//关注/取消关注
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			action := data.param["action"].(int)
			result := Focus(userid, userkey, touserid, action)
			data.retchan <- result
		case "Blacklist":	//拉黑/取消拉黑
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			action := data.param["action"].(int)
			result := Blacklist(userid, userkey, touserid, action)
			data.retchan <- result
		case "Denounce":	//举报
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			Type := data.param["type"].(string)
			content := data.param["content"].(string)
			result := Denounce(userid, userkey, touserid, Type, content)
			data.retchan <- result
		case "DynamicList":		//获取动态列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			filetype := data.param["filetype"].(string)
			tag := data.param["tag"].(string)
			page := data.param["page"].(int)
			result := DynamicList(userid, userkey, filetype, tag, page)
			data.retchan <- result
		case "DynamicLike":		//点赞/取消点赞动态
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			dynamicid := data.param["dynamicid"].(int)
			action := data.param["action"].(int)
			result := DynamicLike(userid, userkey, dynamicid, action)
			data.retchan <- result
		case "DynamicCommentList":		//获取评论列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			dynamicid := data.param["dynamicid"].(int)
			page := data.param["page"].(int)
			result := DynamicCommentList(userid, userkey, dynamicid, page)
			data.retchan <- result
		case "DynamicComment":		//评论动态
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			dynamicid := data.param["dynamicid"].(int)
			content := data.param["content"].(string)
			result := DynamicComment(userid, userkey, dynamicid, content)
			data.retchan <- result
		case "DynamicLikeComment":		//点赞/取消点赞评论
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			commentid := data.param["commentid"].(int)
			action := data.param["action"].(int)
			result := DynamicLikeComment(userid, userkey, commentid, action)
			data.retchan <- result
		case "DynamicPost":		//发布动态
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			description := data.param["description"].(string)
			filetype := data.param["filetype"].(string)
			filelist := data.param["filelist"].([]int)
			result := DynamicPost(userid, userkey, description, filetype, filelist)
			data.retchan <- result
		case "DynamicUserList":		//用户动态列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			touserid := data.param["touserid"].(int)
			filetype := data.param["filetype"].(string)
			page := data.param["page"].(int)
			result := DynamicUserList(userid, userkey, touserid, page, filetype)
			data.retchan <- result
		case "DynamicDelete":		//删除动态
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			dynamicid := data.param["dynamicid"].(int)
			result := DynamicDelete(userid, userkey, dynamicid)
			data.retchan <- result
		case "RoomList":	//获取房间列表
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomtype := data.param["roomtype"].(int)
			tag := data.param["tag"].(string)
			page := data.param["page"].(int)
			result := RoomList(userid, userkey, roomtype, page, tag)
			data.retchan <- result			
		case "RoomEnter":		//进入房间
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomid := data.param["roomid"].(int)
			result := RoomEnter(userid, userkey, roomid)
			data.retchan <- result
		case "RoomLeave":		//退出房间
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomid := data.param["roomid"].(int)
			result := RoomLeave(userid, userkey, roomid)
			data.retchan <- result
		case "RoomCreate":		//申请创建房间
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomtype := data.param["roomtype"].(int)
			result := RoomCreate(userid, userkey, roomtype)
			data.retchan <- result
		case "RoomLike":		//点赞房间
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomid := data.param["roomid"].(int)
			result := RoomLike(userid, userkey, roomid)
			data.retchan <- result
		case "RoomSeat":		//申请上座
			userid := data.param["userid"].(int)
			userkey := data.param["userkey"].(string)
			roomid := data.param["roomid"].(int)
			result := RoomSeat(userid, userkey, roomid)
			data.retchan <- result
		default:
			result := ErrorResult("异常错误")
			data.retchan <- result
		}
	}
}
