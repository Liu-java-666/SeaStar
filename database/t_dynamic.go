package database

import (
	"fmt"
	"github.com/wonderivan/logger"
)

type t_dynamic struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	Description		string
	Topic			string
	Filetype		string
	Filelist		string
	Is_audit		int
	Audit_time		[]uint8
}

type TDynamic t_dynamic

func Dynamic_Get(id int) (*t_dynamic, error) {
	t := &t_dynamic{}
	err := Get(t, "SELECT * FROM dynamic WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_Insert(userid int, description, topic, filetype, filelist string) error {
	_, err := Exec("INSERT INTO dynamic(user_id,description,topic,filetype,filelist) VALUES(?,?,?,?,?)",
		userid, description, topic, filetype, filelist)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func Dynamic_GetCount(userid int, filetype string) int {
	cnt := 0

	var sqlstr string
	if filetype == "" {
		sqlstr = fmt.Sprintf("SELECT COUNT(*) FROM dynamic WHERE user_id = %d AND is_audit = 1",
			userid)
	} else {
		sqlstr = fmt.Sprintf("SELECT COUNT(*) FROM dynamic WHERE user_id = %d AND filetype = '%s' AND is_audit = 1",
			userid, filetype)
	}

	err := Get(&cnt, sqlstr)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func GetFileTypeString(filetype string) string {
	str := ""
	if filetype == "" {
		str = "('image','video')"
	} else {
		str = fmt.Sprintf("('%s')", filetype)
	}
	return str
}

func Dynamic_AllList(index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_TopicList(topic string, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND topic = %s AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), topic, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_FocusList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf(`SELECT * FROM dynamic WHERE filetype IN %s AND is_audit = 1 
		AND user_id IN (SELECT to_user_id FROM focuslist WHERE user_id = %d)
		ORDER BY audit_time DESC, id DESC LIMIT %d,%d`,
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_HotList(index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf(`SELECT a.* FROM dynamic AS a LEFT JOIN (
			SELECT COUNT(*) AS likenum, postuser_id AS likeuser FROM dynamic_like GROUP BY likeuser
		) AS b ON a.user_id = b.likeuser 
		WHERE filetype IN %s AND is_audit = 1 ORDER BY likenum DESC, audit_time DESC LIMIT %d,%d`,
		GetFileTypeString(filetype), index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_MyList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND user_id = %d ORDER BY id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_UserList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND user_id = %d AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_AuditList(index, maxcount int) ([]*t_dynamic, error) {
	t := []*t_dynamic{}
	err := Select(&t, "SELECT * FROM dynamic WHERE is_audit = 0 ORDER BY id DESC LIMIT ?,?",
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func (t *t_dynamic) SetAudit(audit int) error {
	_, err := Exec("UPDATE dynamic SET is_audit = ?, audit_time = CURRENT_TIMESTAMP() WHERE `id` = ?",
		audit, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_dynamic) Delete() error {
	_, err := Exec("DELETE FROM dynamic WHERE `id` = ?",
		t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}