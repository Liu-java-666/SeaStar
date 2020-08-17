package server

import (
	"VPartyServer/config"
	"VPartyServer/logic"
	"VPartyServer/manage"
	"VPartyServer/network"
	"VPartyServer/public"
	"encoding/base64"
	"github.com/wonderivan/logger"
	"net/http"
	"strings"
)

//使用统一规则解密字符串
func DecodeString(routerStr string) string {
	bodys, err := base64.StdEncoding.DecodeString(routerStr)
	if err != nil {
		logger.Error("解密失败")
		return routerStr
	}
	param := public.AesDecryptECB(bodys, config.GetEncryptKey())
	paramstr := string(param)
	return paramstr
}
//解密
func Decode(r *http.Request) {
	data := r.Form.Get(config.GetEncryptPre())

	logger.Info(data)

	if data != "" {
		bodys, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			logger.Error(err)
			return
		}

		param := public.AesDecryptECB(bodys, config.GetEncryptKey())

		paramstr := string(param)
		logger.Info("Decrypted data:", paramstr)
		paramlist := strings.Split(paramstr, "&")
		for _, v := range paramlist {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				r.Form.Add(kv[0], kv[1])
			}
		}
	}
}

type MyHandler struct{}
var routerMap map[string]func(w http.ResponseWriter, r *http.Request)
func (mh MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	if r.Method != "POST" {
		network.SendResult(w, logic.ErrorResult("不支持的方法"))
		return
	}

	r.ParseMultipartForm(100*1024*1024)

	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))

	phoneBrand := r.Header.Get("phoneBrand") // 手机品牌
	phoneSystem := r.Header.Get("phoneSystem")//操作系统
	phoneModels := r.Header.Get("phoneModels")//手机型号

	routerUrl := r.URL.Path[strings.Index(r.URL.Path,"/")+1:] //将开头/去掉

	if config.IsTest() == false {
		Decode(r)
	}

	//打印路由日志
	logger.Info(routerUrl, r.RemoteAddr,phoneBrand,phoneSystem,phoneModels)

	fn, ok := routerMap[routerUrl]
	if ok {
		fn(w, r)
	}else{
		network.SendResult(w, logic.ErrorResult("未知路径"))
	}
}

func SetRouter() {

	http.Handle("/", MyHandler{})
	//初始化路由
	routerMap = make(map[string]func(w http.ResponseWriter, r *http.Request), 100)
	routerMap["imcallback"] = network.OnIMCallback
	routerMap["vparty/config"] = network.OnCheckVersion
	routerMap["vparty/phone-code"] = network.OnGetCaptcha
	routerMap["vparty/phone-login"] = network.OnLogin
	routerMap["vparty/set-info"] = network.OnSetInfo
	routerMap["vparty/edit-info"] = network.OnEditInfo
	routerMap["vparty/my-menu"] = network.OnGetMyMenu
	routerMap["vparty/my-info"] = network.OnGetMyInfo
	routerMap["vparty/user-detail"] = network.OnGetUserDetail
	routerMap["vparty/user-card"] = network.OnGetUserCard
	routerMap["vparty/userinfo-list"] = network.OnGetUserInfoList
	routerMap["vparty/is-hate"] = network.OnIsBlacklist
	routerMap["vparty/love-list"] = network.OnGetFocusList
	routerMap["vparty/fans-list"] = network.OnGetFansList
	routerMap["vparty/hate-list"] = network.OnGetBlacklist
	routerMap["vparty/friend-list"] = network.OnGetFriendList
	routerMap["vparty/apply-list"] = network.OnGetApplyList
	routerMap["vparty/receive-giftlist"] = network.OnGetReceiveGiftList
	routerMap["vparty/wallet"] = network.OnGetMyWallet
	routerMap["vparty/pay-order"] = network.OnPayOrder
	routerMap["vparty/pay-finish"] = network.OnPayFinish
	routerMap["vparty/send-gift"] = network.OnSendGift
	routerMap["vparty/upload-image"] = network.OnUploadImage
	routerMap["vparty/upload-video"] = network.OnUploadVideo
	routerMap["vparty/1v1"] = network.OnGetMatchUser
	routerMap["vparty/call-up"] = network.OnCallUp
	routerMap["vparty/hang-up"] = network.OnHangUp
	routerMap["vparty/ranklist"] = network.OnGetRankList
	routerMap["vparty/love"] = network.OnFocus
	routerMap["vparty/hate"] = network.OnBlacklist
	routerMap["vparty/report"] = network.OnDenounce
	routerMap["vparty/dynamic-list"] = network.OnDynamicList
	routerMap["vparty/dynamic-like"] = network.OnDynamicLike
	routerMap["vparty/dynamic-commentlist"] = network.OnDynamicCommentList
	routerMap["vparty/dynamic-comment"] = network.OnDynamicComment
	routerMap["vparty/dynamic-likecomment"] = network.OnDynamicLikeComment
	routerMap["vparty/dynamic-post"] = network.OnDynamicPost
	routerMap["vparty/dynamic-userlist"] = network.OnDynamicUserList
	routerMap["vparty/dynamic-delete"] = network.OnDynamicDelete
	routerMap["vparty/room-list"] = network.OnRoomList
	routerMap["vparty/room-enter"] = network.OnRoomEnter
	routerMap["vparty/room-leave"] = network.OnRoomLeave
	routerMap["vparty/room-create"] = network.OnRoomCreate
	routerMap["vparty/room-like"] = network.OnRoomLike
	routerMap["vparty/room-seat"] = network.OnRoomSeat
}

func SetRouter_Manager() {
	http.HandleFunc("/vpartymanager/im/import", manage.OnIMImport)
	http.HandleFunc("/vpartymanager/avatar/list", manage.OnAvatarList)
	http.HandleFunc("/vpartymanager/avatar/audit", manage.OnAvatarAudit)
	http.HandleFunc("/vpartymanager/dynamic/list", manage.OnDynamicList)
	http.HandleFunc("/vpartymanager/dynamic/audit", manage.OnDynamicAudit)
	http.HandleFunc("/vpartymanager/photo/list", manage.OnPhotoList)
	http.HandleFunc("/vpartymanager/photo/audit", manage.OnPhotoAudit)
}