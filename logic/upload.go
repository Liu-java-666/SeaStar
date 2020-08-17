package logic

import (
	"VPartyServer/database"
)

//上传图片
func UploadImage(user *database.TUser, file, filetype, usetype string, index int) interface{} {
	//插入图片表
	id, err := database.Image_Insert(user.Id, file, filetype, usetype)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		UseType string
		Index int
		Id int
	}{
		true,
		usetype,
		index,
		id,
	}

	return result
}

//上传视频
func UploadVideo(user *database.TUser, file, filetype, cover, covertype, usetype string, rotation, index int) interface{} {
	//插入视频表
	id, err := database.Video_Insert(user.Id, file, filetype, cover, covertype, usetype, rotation)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		UseType string
		Index int
		Id int
	}{
		true,
		usetype,
		index,
		id,
	}

	return result
}
