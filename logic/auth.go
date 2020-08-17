package logic

import (
	"VPartyServer/config"
	"VPartyServer/database"
	"VPartyServer/im"
	"VPartyServer/public"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
)

//获取验证码
func GetCaptcha(phone string) interface{} {
	//检查手机号格式
	if !public.ValidPhoneNumber(phone) {
		return ErrorResult("请输入正确的手机号")
	}

	//获取验证码数据对象
	tcaptcha, err := database.Captcha_Get(phone)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//插入/更新验证码
	if tcaptcha == nil {
		err = database.Captcha_Insert(phone, "123456")
	} else {
		err = tcaptcha.UpdateCaptcha("123456")
	}
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

//登录/注册
func Login(phone, captcha, ip string) interface{} {
	//检查手机号格式
	if !public.ValidPhoneNumber(phone) {
		return ErrorResult("请输入正确的手机号")
	}

	//检查验证码
	tcaptcha, err := database.Captcha_Get(phone)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	timenow := public.GetNowTimestamp()
	if tcaptcha == nil || tcaptcha.Is_used > 0 || tcaptcha.Expire_time < timenow {
		return ErrorResult("请先获取验证码")
	}
	if tcaptcha.Captcha != captcha {
		return ErrorResult("验证码错误")
	}

	//读取用户数据
	tuser, err := database.User_GetByPhone(phone)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	userkey := public.GetRandString(32)
	bRegister := tuser == nil
	//不存在就注册
	if bRegister {
		tuser, err = database.User_Insert(phone, ip, userkey)
		if tuser == nil || err != nil {
			return ErrorResult("数据库异常")
		}
	} else {	//更新最后登录信息
		tuser.UpdateLogin(ip, userkey)
	}

	//使用验证码
	tcaptcha.SetUsed()

	//IM相关
	imconfig := config.GetIMConfig()
	sig, err := im.GenSig(imconfig.AppId, imconfig.Key, fmt.Sprintf("%s%d", imconfig.Pre, tuser.Id), 60*60*24*180)
	if err != nil {
		logger.Error("IM生成sig失败：", err)
	}

	if bRegister {
		im.AccountImport(tuser.Id, OnAccountImport)
	}

	//返回成功消息
	result := struct{
		Result bool
		UserId int
		UserKey string
		Nickname string
		AvatarFile string
		AvatarAudit int
		Signature string
		Sex int
		Age int
		Coins int
		IMAppid int
		IMSig string
		IMPre string
	}{
		true,
		tuser.Id,
		userkey,
		tuser.Nickname,
		MakeImageUrl(tuser.AvatarFile),
		tuser.AvatarAudit,
		tuser.Signature,
		tuser.Sex,
		tuser.Age,
		tuser.Coins,
		imconfig.AppId,
		sig,
		imconfig.Pre,
	}

	return result
}

//导入IM账号结果
func OnAccountImport(resultData string, err error, user_data interface{}) {
	userdata := user_data.(int)

	if err != nil {
		logger.Error("导入IM账号失败,userid=%d,err=%v", userdata, err)
		return
	}

	logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("导入IM账号失败,userid=%d,err=%v,resultData=%v", userdata, err, resultData)
		return
	}

	//logger.Debug(revData)

	ActionStatus := revData["ActionStatus"].(string)
	if ActionStatus != "OK" {
		ErrorCode := int(revData["ErrorCode"].(float64))
		ErrorInfo := revData["ErrorInfo"].(string)
		logger.Error("导入IM账号失败,userid=%d,errcode=%d,errinfo=%s", userdata, ErrorCode, ErrorInfo)
		return
	}

	//发送欢迎消息
	msg := "欢迎来到海星！"
	im.SendMsg(1, userdata, 0, msg)
	msg = "这里是海星的反馈与帮助，如果你遇到了任何问题，都可以告知我们。良好的交友环境需要大家共同营造，祝你在海星玩得开心，找到陪伴~"
	im.SendMsg(2, userdata, 0, msg)
}