package database

import (
	"github.com/wonderivan/logger"
)

type t_dynamic_like struct {
	Dynamic_id	int
	User_id		int
	Postuser_id int
}

func DynamicLike_Get(userid, dynamicid int) bool {
	t := &t_dynamic_like{}
	err := Get(t, "SELECT * FROM dynamic_like WHERE dynamic_id = ? and user_id = ?",
		dynamicid, userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}

func DynamicLike_Insert(userid, dynamicid, postid int) error {
	_, err := Exec("INSERT INTO dynamic_like(dynamic_id,user_id,postuser_id) VALUES(?,?,?)",
		dynamicid, userid, postid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicLike_Delete(userid, dynamicid int) error {
	_, err := Exec("DELETE FROM dynamic_like WHERE dynamic_id = ? and user_id = ?",
		dynamicid, userid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicLike_GetCount(dynamicid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_like WHERE dynamic_id = ?",
		dynamicid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func DynamicLike_GetCountByUser(userid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_like WHERE postuser_id = ?",
		userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}