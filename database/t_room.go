package database

import (
	"github.com/wonderivan/logger"
)

type t_room struct {
	Id				int
	User_id			int
	Room_type		int
	Im_group		string
	Room_name		string
	Like_num		int
	Is_open			int
	Open_time		[]uint8
	Close_time		[]uint8
}

type TRoom t_room

func Room_Get(roomid int) (*TRoom, error) {
	t := &TRoom{}
	err := Get(t, "SELECT * FROM room WHERE id = ?", roomid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Room_Find(imGroup string) (*TRoom, error) {
	t := &TRoom{}
	err := Get(t, "SELECT * FROM room WHERE im_group = ?", imGroup)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Room_AllList(index, maxcount, roomtype int) ([]*TRoom, error) {
	t := []*TRoom{}
	err := Select(&t, "SELECT * FROM room WHERE room_type = ? AND is_open = 1 ORDER BY open_time DESC LIMIT ?,?",
		roomtype, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Room_HotList(index, maxcount, roomtype int) ([]*TRoom, error) {
	t := []*TRoom{}
	err := Select(&t, "SELECT * FROM room WHERE room_type = ? AND is_open = 1 ORDER BY like_num DESC LIMIT ?,?",
		roomtype, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Room_FocusList(userid, index, maxcount, roomtype int) ([]*TRoom, error) {
	t := []*TRoom{}
	err := Select(&t, `SELECT * FROM room WHERE room_type = ? AND is_open = 1 
		AND user_id IN (SELECT to_user_id FROM focuslist WHERE user_id = ?)
		ORDER BY open_time DESC LIMIT ?,?`,
		roomtype, userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Room_Insert(userid int) error {
	_, err := Exec("INSERT INTO room(user_id) VALUES(?)",
		userid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func Room_GetLikeNum(userid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT like_num FROM room WHERE user_id = ?", userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0
		}

		logger.Error(err)
		return 0
	}

	return cnt
}

func Room_Like(roomid int) error {
	_, err := Exec("UPDATE room SET like_num = like_num + 1 WHERE id = ?",
		roomid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}