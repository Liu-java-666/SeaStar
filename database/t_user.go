package database

import (
	"VPartyServer/public"
	"fmt"
	"github.com/wonderivan/logger"
	"time"
)

type t_user struct {
	Id					int
	Phone_number		string
	Registration_time	[]uint8
	Nickname			string
	Sex					int
	Birthday			[]uint8
	User_key			string
	Lastlogon_time		[]uint8
	Lastlogon_ip		string
	Avatar_id			int
	Photolist_id		int
	Certification		int
	Signature			string
	Relationship_status string
	Friends_purpose		string
	Hobbies				string
	Coins				int
	Coins_used			int
}

type TUser struct {
	t_user
	AvatarFile	string
	AvatarAudit int
	Age			int
}

func GetAge(birthday []uint8) int {
	tm := public.StrToDate(string(birthday))
	if tm.Year() <= 1900 {
		return 0
	}
	now := time.Now()

	age := now.Year() - tm.Year()
	if now.Month() < tm.Month() || (now.Month() == tm.Month() && now.Day() < tm.Day()) {
		age --
	}

	return age
}

func User_GetByPhone(phone string) (*TUser, error) {
	t := TUser{}
	err := Get(&t, "SELECT * FROM user WHERE phone_number = ?", phone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	t.AvatarFile, t.AvatarAudit = Image_GetMyAvatar(t.Id)
	t.Age = GetAge(t.Birthday)

	return &t, nil
}

func User_GetById(id int, bMe bool) (*TUser, error) {
	t := TUser{}
	err := Get(&t, "SELECT * FROM user WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	if bMe {
		t.AvatarFile, t.AvatarAudit = Image_GetMyAvatar(t.Id)
	} else {
		t.AvatarFile = Image_GetOtherAvatar(t.Avatar_id)
		t.AvatarAudit = 1
	}
	t.Age = GetAge(t.Birthday)

	return &t, nil
}

func User_Insert(phone, ip, userkey string) (*TUser, error) {
	result, err := Exec("INSERT INTO user(phone_number,user_key,lastlogon_ip) VALUES(?,?,?)",
		phone, userkey, ip)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return User_GetById(int(id), true)
}

func User_GetLikeCount(userid int) int {
	cnt := 0
	err := Get(&cnt, `SELECT COALESCE(SUM(num),0) FROM (
			SELECT COUNT(*) AS num FROM dynamic_like WHERE postuser_id = ?
			UNION ALL
			SELECT COUNT(*) AS num FROM dynamic_comment_like WHERE commentuser_id = ?
			UNION ALL
			SELECT like_num AS num FROM room WHERE user_id = ?
		) AS c `,
		userid, userid, userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func User_SetAvatar(id, avatar int) {
	_, err := Exec("UPDATE user SET avatar_id = ? WHERE `id` = ?",
		avatar, id)
	if err != nil {
		logger.Error(err)
	}
}

func User_SetCertification(id int) {
	_, err := Exec("UPDATE user SET certification = 1 WHERE `id` = ?",
		id)
	if err != nil {
		logger.Error(err)
	}
}

func User_SetPhotoList(id, photolistid int) {
	_, err := Exec("UPDATE user SET photolist_id = ? WHERE `id` = ?",
		photolistid, id)
	if err != nil {
		logger.Error(err)
	}
}


func User_GetMatchUser(userid int) ([]*TUser, error) {
	t := []*TUser{}
	sqlstr := fmt.Sprintf(`SELECT * FROM user WHERE nickname != '' AND id > 10  and id <= 18 AND id != %d
		AND id NOT IN (SELECT to_user_id FROM blacklist WHERE user_id = %d)
		AND id NOT IN (SELECT user_id FROM blacklist WHERE to_user_id = %d)
		AND id NOT IN (SELECT to_user_id FROM match_log WHERE user_id = %d AND cdate >= CAST(SYSDATE()AS DATE))
		GROUP BY id ORDER BY RAND() LIMIT 4`,
		userid, userid, userid, userid)
	err := Select(&t, sqlstr)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}
	return t, nil
}

func User_GetDestined(id, cnt int) ([]*TUser, error) {
	t := []t_user{}
	err := Select(&t, "SELECT * FROM user WHERE nickname != '' AND id > 10 AND id != ? ORDER BY RAND() LIMIT ?",
		id, cnt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	T := []*TUser{}
	for _, v := range t {
		item := &TUser{
			t_user:      v,
			AvatarFile:  Image_GetOtherAvatar(v.Avatar_id),
			AvatarAudit: 1,
			Age:		 GetAge(v.Birthday),
		}
		T = append(T, item)
	}

	return T, nil
}

func User_Search(id int, keyword string, index, maxcount int) ([]*TUser, error) {
	t := []t_user{}
	sqlstr := fmt.Sprintf("SELECT * FROM user WHERE nickname like '%%%s%%' AND id > 10 AND id != %d ORDER BY id DESC LIMIT %d,%d",
		keyword, id, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	T := []*TUser{}
	for _, v := range t {
		item := &TUser{
			t_user:      v,
			AvatarFile:  Image_GetOtherAvatar(v.Avatar_id),
			AvatarAudit: 1,
			Age:		 GetAge(v.Birthday),
		}
		T = append(T, item)
	}

	return T, nil
}

type RankData struct {
	TUser
	Num int
}

func User_RichList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(coins) AS num, user_id FROM gift_log GROUP BY user_id
		) AS b ON a.id = b.user_id WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func User_StarList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(coins) AS num, to_user_id FROM gift_log GROUP BY to_user_id
		) AS b ON a.id = b.to_user_id WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func User_CharmList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(num) AS num, user_id FROM (
				SELECT SUM(like_num) AS num, user_id FROM room GROUP BY user_id
				UNION ALL
				SELECT COUNT(*) AS num, postuser_id AS user_id FROM dynamic_like GROUP BY postuser_id
				UNION ALL
				SELECT COUNT(*) AS num, commentuser_id AS user_id FROM dynamic_comment_like GROUP BY commentuser_id
			) AS c GROUP BY user_id
		) AS b ON a.id = b.user_id 
		WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func (t *TUser) UpdateLogin(ip, userkey string) {
	_, err := Exec("UPDATE user SET user_key = ?, lastlogon_time = ?, lastlogon_ip= ? WHERE `id` = ?",
		userkey, public.GetNowTimestr(), ip, t.Id)
	if err != nil {
		logger.Error(err)
	}
}

func (t *TUser) SetInfo(nickname, birthday string, sex int) error {
	_, err := Exec("UPDATE user SET nickname = ?, sex = ?, birthday = ? WHERE `id` = ?",
		nickname, sex, birthday, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *TUser) UpdateInfo(nickname, signature, status, purpose, hobbies string) error {
	_, err := Exec(`UPDATE user SET nickname = ?, signature = ?, relationship_status = ?, 
		friends_purpose = ?, hobbies = ? WHERE id = ?`,
		nickname, signature, status, purpose, hobbies, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *TUser) UseCoins(coins int) (error, int) {
	_, err := Exec("UPDATE user SET coins = coins - ?, coins_used = coins_used + ? WHERE `id` = ?",
		coins, coins, t.Id)
	if err != nil {
		logger.Error(err)
		return err, 0
	}

	cnt := 0
	_ = Get(&cnt, "SELECT coins FROM user WHERE `id` = ?", t.Id)

	return nil, cnt
}

func (t *TUser) AddCoins(coins int) (error, int) {
	_, err := Exec("UPDATE user SET coins = coins + ? WHERE `id` = ?",
		coins, t.Id)
	if err != nil {
		logger.Error(err)
		return err, 0
	}

	cnt := 0
	_ = Get(&cnt, "SELECT coins FROM user WHERE `id` = ?", t.Id)

	return nil, cnt
}