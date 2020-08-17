package network

import (
	"VPartyServer/config"
	"VPartyServer/logic"
	"VPartyServer/public"
	"encoding/base64"
	"encoding/json"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
)

//使用统一规则加密字符串
func EncryptString(str string) string {
	// 转为byte数组
	strbyte := []byte(str)
	// aes加密
	aesStr := public.AesEncryptECB(strbyte, config.GetEncryptKey())
	// base64加密
	base64Str := base64.StdEncoding.EncodeToString(aesStr)
	return base64Str
}

func Encrypt(data []byte) []byte {
	body := public.AesEncryptECB(data, config.GetEncryptKey())
	bodys := base64.StdEncoding.EncodeToString(body)
	return []byte(bodys)
}

func SendResult(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		return
	}

	bodys := b
	if config.IsTest() == false {
		//加密
		bodys = Encrypt(b)
	}

	_, err = w.Write(bodys)
	if err != nil {
		logger.Error(err)
		return
	}
}

//IM回调
func OnIMCallback(w http.ResponseWriter, r *http.Request) {
	resultData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(string(resultData))

	revData := make(map[string]interface{})
	err = json.Unmarshal(resultData, &revData)
	if err != nil {
		logger.Error(err)
		return
	}

	CallbackCommand := revData["CallbackCommand"].(string)
	logger.Debug(CallbackCommand)
	if CallbackCommand != "Group.CallbackAfterNewMemberJoin" && CallbackCommand != "Group.CallbackAfterMemberExit" {
		return
	}

	GroupId := revData["GroupId"].(string)
	logger.Debug(GroupId)

	go logic.RoomUpdateOnline(GroupId)
}

//检查版本
func OnCheckVersion(w http.ResponseWriter, r *http.Request) {
	result := struct {
		Result bool `json:"result"`
		HasUpdate int `json:"hasUpdate"`
	}{
		true,
		0,
	}
	SendResult(w, result)
}

//获取验证码
func OnGetCaptcha(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["phone"] = r.Form.Get("phone")

	retchan := make(chan interface{})
	logic.PushQue("GetCaptcha", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//登录/注册
func OnLogin(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["phone"] = r.Form.Get("phone")
	param["captcha"] = r.Form.Get("captcha")

	param["ip"] = r.RemoteAddr
	index := strings.IndexRune(r.RemoteAddr, ':')
	if index >= 0 {
		param["ip"] = r.RemoteAddr[:index]
	}

	retchan := make(chan interface{})
	logic.PushQue("Login", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//设置信息
func OnSetInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["nickname"] = r.Form.Get("nickname")
	param["sex"], _ = strconv.Atoi(r.Form.Get("sex"))
	param["year"], _ = strconv.Atoi(r.Form.Get("year"))
	param["month"], _ = strconv.Atoi(r.Form.Get("month"))
	param["day"], _ = strconv.Atoi(r.Form.Get("day"))

	retchan := make(chan interface{})
	logic.PushQue("SetInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//编辑信息
func OnEditInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["nickname"] = r.Form.Get("nickname")
	param["signature"] = r.Form.Get("signature")
	param["status"] = r.Form.Get("status")
	param["purpose"] = r.Form.Get("purpose")
	param["hobbies"] = r.Form.Get("hobbies")
	slist := r.Form["photolist[]"]
	ilist := []int{}
	for _, v := range slist {
		id, _ := strconv.Atoi(v)
		ilist = append(ilist, id)
	}
	param["photolist"] = ilist

	retchan := make(chan interface{})
	logic.PushQue("EditInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的菜单
func OnGetMyMenu(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyMenu", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的资料
func OnGetMyInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户详情
func OnGetUserDetail(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("GetUserDetail", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户名片
func OnGetUserCard(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("GetUserCard", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//批量用户详情
func OnGetUserInfoList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	slist := r.Form["touseridlist[]"]
	ilist := []int{}
	for _, v := range slist {
		id, _ := strconv.Atoi(v)
		ilist = append(ilist, id)
	}
	param["touseridlist"] = ilist

	retchan := make(chan interface{})
	logic.PushQue("GetUserInfoList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//是否拉黑
func OnIsBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("IsBlacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取关注列表
func OnGetFocusList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFocusList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取粉丝列表
func OnGetFansList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFansList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取黑名单
func OnGetBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetBlacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取好友列表
func OnGetFriendList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFriendList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取好友申请列表
func OnGetApplyList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetApplyList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取收礼列表
func OnGetReceiveGiftList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetReceiveGiftList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的钱包
func OnGetMyWallet(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyWallet", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//提交充值订单
func OnPayOrder(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["money"], _ = strconv.Atoi(r.Form.Get("money"))

	retchan := make(chan interface{})
	logic.PushQue("PayOrder", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//完成充值
func OnPayFinish(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["orderid"] = r.Form.Get("orderid")
	param["status"], _ = strconv.Atoi(r.Form.Get("status"))

	retchan := make(chan interface{})
	logic.PushQue("PayFinish", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//送礼物
func OnSendGift(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["scene"] = r.Form.Get("scene")
	param["sceneid"], _ = strconv.Atoi(r.Form.Get("sceneid"))
	param["giftid"], _ = strconv.Atoi(r.Form.Get("giftid"))

	retchan := make(chan interface{})
	logic.PushQue("SendGift", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//上传图片
func OnUploadImage(w http.ResponseWriter, r *http.Request) {
	srcfile, srcfileheader, err := r.FormFile("file")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取文件数据错误"))
		return
	}
	defer srcfile.Close()

	if srcfileheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：无法读取文件大小"))
		return
	} else if srcfileheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：文件大小超出30M"))
		return
	}

	filetype := r.Form.Get("filetype")
	if filetype != "jpg" && filetype != "png" {
		SendResult(w, logic.ErrorResult("文件格式错误"))
		return
	}

	//检查用户
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	retchan := make(chan interface{})
	logic.PushQue("CheckUser", param, retchan)
	result := <- retchan
	ck := result.(logic.CheckUserResult)
	if !ck.Result {
		SendResult(w, ck.Error)
		return
	}

	//保存文件
	param["usetype"] = r.Form.Get("usetype")
	param["index"], _ = strconv.Atoi(r.Form.Get("index"))
	filename := public.MakeFileName(param["userid"].(int), param["index"].(int), filetype, param["usetype"].(string))
	filepath := logic.MakeImagePath(filename)
	file, err := os.Create(filepath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(file, srcfile)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	file.Close()

	//更新信息
	param["user"] = ck.User
	param["file"] = filename
	param["filetype"] = filetype
	logic.PushQue("UploadImage", param, retchan)
	result = <- retchan

	SendResult(w, result)
}

//上传视频
func OnUploadVideo(w http.ResponseWriter, r *http.Request) {
	//获取封面
	srccover, srccoverheader, err := r.FormFile("cover")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取封面文件数据错误"))
		return
	}
	defer srccover.Close()

	if srccoverheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取封面文件错误：无法读取文件大小"))
		return
	} else if srccoverheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取封面文件错误：文件大小超出30M"))
		return
	}

	covertype := r.Form.Get("covertype")
	if covertype != "jpg" && covertype != "png" {
		SendResult(w, logic.ErrorResult("封面文件格式错误"))
		return
	}

	//获取视频
	srcfile, srcfileheader, err := r.FormFile("file")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取视频文件数据错误"))
		return
	}
	defer srcfile.Close()

	if srcfileheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取视频文件错误：无法读取文件大小"))
		return
	} else if srcfileheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取视频文件错误：文件大小超出30M"))
		return
	}

	filetype := r.Form.Get("filetype")
	if filetype != "mp4" && filetype != "rmvb" {
		SendResult(w, logic.ErrorResult("视频文件格式错误"))
		return
	}

	//检查用户
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	retchan := make(chan interface{})
	logic.PushQue("CheckUser", param, retchan)
	result := <- retchan
	ck := result.(logic.CheckUserResult)
	if !ck.Result {
		SendResult(w, ck.Error)
		return
	}

	//保存封面
	param["usetype"] = r.Form.Get("usetype")
	param["index"], _ = strconv.Atoi(r.Form.Get("index"))
	covername := public.MakeFileName(param["userid"].(int), param["index"].(int), covertype, param["usetype"].(string))
	coverpath := logic.MakeVideoPath(covername)
	cover, err := os.Create(coverpath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(cover, srccover)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	cover.Close()

	//保存视频
	filename := public.ChangeFileType(covername, filetype)
	filepath := logic.MakeVideoPath(filename)
	file, err := os.Create(filepath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(file, srcfile)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	file.Close()

	//更新信息
	param["user"] = ck.User
	param["file"] = filename
	param["filetype"] = filetype
	param["cover"] = covername
	param["covertype"] = covertype
	param["rotation"], _ = strconv.Atoi(r.Form.Get("rotation"))
	logic.PushQue("UploadVideo", param, retchan)
	result = <- retchan

	SendResult(w, result)
}

//1V1
func OnGetMatchUser(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMatchUser", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//打电话
func OnCallUp(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("CallUp", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//挂电话
func OnHangUp(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("HangUp", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//排行榜
func OnGetRankList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["tag"] = r.Form.Get("tag")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetRankList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//关注/取消关注
func OnFocus(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("Focus", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//拉黑/取消拉黑
func OnBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("Blacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//举报
func OnDenounce(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["type"] = r.Form.Get("type")
	param["content"] = r.Form.Get("content")

	retchan := make(chan interface{})
	logic.PushQue("Denounce", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取动态列表
func OnDynamicList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["filetype"] = r.Form.Get("filetype")
	param["tag"] = r.Form.Get("tag")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//点赞/取消点赞动态
func OnDynamicLike(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicLike", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取评论列表
func OnDynamicCommentList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicCommentList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//评论动态
func OnDynamicComment(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["content"] = r.Form.Get("content")

	retchan := make(chan interface{})
	logic.PushQue("DynamicComment", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//点赞/取消点赞评论
func OnDynamicLikeComment(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["commentid"], _ = strconv.Atoi(r.Form.Get("commentid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicLikeComment", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//发布动态
func OnDynamicPost(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["description"] = r.Form.Get("description")
	param["filetype"] = r.Form.Get("filetype")
	slist := r.Form["filelist[]"]
	ilist := []int{}
	for _, v := range slist {
		id, _ := strconv.Atoi(v)
		ilist = append(ilist, id)
	}
	param["filelist"] = ilist

	retchan := make(chan interface{})
	logic.PushQue("DynamicPost", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户动态列表
func OnDynamicUserList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["filetype"] = r.Form.Get("filetype")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicUserList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//删除动态
func OnDynamicDelete(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicDelete", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取房间列表
func OnRoomList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomtype"], _ = strconv.Atoi(r.Form.Get("roomtype"))
	param["tag"] = r.Form.Get("tag")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("RoomList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//进入房间
func OnRoomEnter(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomEnter", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//退出房间
func OnRoomLeave(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomLeave", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//申请创建房间
func OnRoomCreate(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomtype"], _ = strconv.Atoi(r.Form.Get("roomtype"))

	retchan := make(chan interface{})
	logic.PushQue("RoomCreate", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//点赞房间
func OnRoomLike(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomLike", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//申请上座
func OnRoomSeat(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomSeat", param, retchan)
	result := <- retchan
	SendResult(w, result)
}