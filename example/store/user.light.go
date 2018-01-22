// !!! DO NOT EDIT THIS FILE. It is generated by `light` tool.
// @light: https://github.com/arstd/light
// Generate from source: github.com/arstd/light/example/store/user.go
package store

import (
	"bytes"
	"fmt"

	"github.com/arstd/light/example/model"
	"github.com/arstd/light/null"
	"github.com/arstd/log"
)

var User IUser = new(StoreUser)

type StoreUser struct{}

func (*StoreUser) Create(name string) error {
	var buf bytes.Buffer
	var args []interface{}
	fmt.Fprintf(&buf, "CREATE TABLE IF NOT EXISTS %v ( id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY, username VARCHAR(32) NOT NULL UNIQUE, Phone VARCHAR(32), address VARCHAR(256), status TINYINT UNSIGNED, birth_day DATE, created TIMESTAMP default CURRENT_TIMESTAMP, updated TIMESTAMP default CURRENT_TIMESTAMP ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ", name)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
	}
	return err
}

func (*StoreUser) Insert(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("INSERT INTO users(`username`,phone,address,status,birth_day,created,updated) VALUES (?,?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP) ")
	args = append(args, u.Username, null.String(&u.Phone), u.Address, null.Uint8(&u.Status), u.BirthDay)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, err
	}
	return res.LastInsertId()
}

func (*StoreUser) Upsert(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("INSERT INTO users(username,phone,address,status,birth_day,created,updated) VALUES (?,?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP) ON DUPLICATE KEY UPDATE username=VALUES(username), phone=VALUES(phone), address=VALUES(address), status=VALUES(status), birth_day=VALUES(birth_day), updated=CURRENT_TIMESTAMP ")
	args = append(args, u.Username, null.String(&u.Phone), u.Address, null.Uint8(&u.Status), u.BirthDay)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, err
	}
	return res.LastInsertId()
}

func (*StoreUser) Replace(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("REPLACE INTO users(username,phone,address,status,birth_day,created,updated) VALUES (?,?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP) ")
	args = append(args, u.Username, null.String(&u.Phone), u.Address, null.Uint8(&u.Status), u.BirthDay)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, err
	}
	return res.LastInsertId()
}

func (*StoreUser) Update(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("UPDATE users SET ")
	if u.Username != "" {
		buf.WriteString("username=?, ")
		args = append(args, u.Username)
	}
	if u.Phone != "" {
		buf.WriteString("phone=?, ")
		args = append(args, null.String(&u.Phone))
	}
	if u.Address != nil {
		buf.WriteString("address=?, ")
		args = append(args, u.Address)
	}
	if u.Status != 0 {
		buf.WriteString("status=?, ")
		args = append(args, null.Uint8(&u.Status))
	}
	if u.BirthDay != nil {
		buf.WriteString("birth_day=?, ")
		args = append(args, u.BirthDay)
	}
	buf.WriteString("updated=CURRENT_TIMESTAMP WHERE id=? ")
	args = append(args, u.Id)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, err
	}
	return res.RowsAffected()
}

func (*StoreUser) Delete(id uint64) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("DELETE FROM users WHERE id=? ")
	args = append(args, id)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, err
	}
	return res.RowsAffected()
}

func (*StoreUser) Get(id uint64) (*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("SELECT id, username, phone, address, status, birth_day, created, updated   ")
	buf.WriteString("FROM users WHERE id=? ")
	args = append(args, id)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	row := db.QueryRow(query, args...)
	xu := new(model.User)
	xdst := []interface{}{&xu.Id, &xu.Username, null.String(&xu.Phone), &xu.Address, null.Uint8(&xu.Status), &xu.BirthDay, &xu.Created, &xu.Updated}
	err := row.Scan(xdst...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	log.Debugf("%#v", xu)
	return xu, err
}

func (*StoreUser) List(u *model.User, offset int, size int) ([]*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("SELECT (SELECT id FROM users WHERE id=a.id) AS id, `username`, phone AS phone, address, status, birth_day, created, updated   ")
	buf.WriteString("FROM users a WHERE id != -1 AND username <> 'admin' AND username LIKE ? ")
	args = append(args, u.Username)
	if (u.Phone != "") || ((u.BirthDay != nil && !u.BirthDay.IsZero()) || u.Id > 1) {
		buf.WriteString("AND address = ? ")
		args = append(args, u.Address)
		if u.Phone != "" {
			buf.WriteString("AND phone LIKE ? ")
			args = append(args, null.String(&u.Phone))
		}
		buf.WriteString("AND created > ? ")
		args = append(args, u.Created)
		if (u.BirthDay != nil && !u.BirthDay.IsZero()) || u.Id > 1 {
			if u.BirthDay != nil {
				buf.WriteString("AND birth_day > ? ")
				args = append(args, u.BirthDay)
			}
			if u.Id != 0 {
				buf.WriteString("AND id > ? ")
				args = append(args, u.Id)
			}
		}
	}
	buf.WriteString("AND status != ? ")
	args = append(args, null.Uint8(&u.Status))
	if !u.Updated.IsZero() {
		buf.WriteString("AND updated > ? ")
		args = append(args, u.Updated)
	}
	buf.WriteString("AND birth_day IS NOT NULL ")
	buf.WriteString("ORDER BY updated DESC LIMIT ?, ? ")
	args = append(args, offset, size)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	var data []*model.User
	for rows.Next() {
		xu := new(model.User)
		data = append(data, xu)
		xdst := []interface{}{&xu.Id, &xu.Username, null.String(&xu.Phone), &xu.Address, null.Uint8(&xu.Status), &xu.BirthDay, &xu.Created, &xu.Updated}
		err = rows.Scan(xdst...)
		if err != nil {
			log.Error(query)
			log.Error(args...)
			log.Error(err)
			return nil, err
		}
		log.Debugf("%#v", xu)
	}
	if err = rows.Err(); err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	return data, nil
}

func (*StoreUser) Page(u *model.User, offset int, size int) (int64, []*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString("FROM users WHERE username LIKE ? ")
	args = append(args, u.Username)
	if u.Phone != "" {
		buf.WriteString("AND address = ? ")
		args = append(args, u.Address)
		if u.Phone != "" {
			buf.WriteString("AND phone LIKE ? ")
			args = append(args, null.String(&u.Phone))
		}
		buf.WriteString("AND created > ? ")
		args = append(args, u.Created)
	}
	buf.WriteString("AND birth_day IS NOT NULL AND status != ? ")
	args = append(args, null.Uint8(&u.Status))
	if !u.Updated.IsZero() {
		buf.WriteString("AND updated > ? ")
		args = append(args, u.Updated)
	}

	var total int64
	totalQuery := "SELECT count(1) " + buf.String()
	log.Debug(totalQuery)
	log.Debug(args...)
	err := db.QueryRow(totalQuery, args...).Scan(&total)
	if err != nil {
		log.Error(totalQuery)
		log.Error(args...)
		log.Error(err)
		return 0, nil, err
	}
	buf.WriteString("ORDER BY updated DESC LIMIT ?, ? ")
	args = append(args, offset, size)
	query := `SELECT id, username, phone, address, status, birth_day, created, updated  ` + buf.String()
	log.Debug(query)
	log.Debug(args...)
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, nil, err
	}
	defer rows.Close()
	var data []*model.User
	for rows.Next() {
		xu := new(model.User)
		data = append(data, xu)
		xdst := []interface{}{&xu.Id, &xu.Username, null.String(&xu.Phone), &xu.Address, null.Uint8(&xu.Status), &xu.BirthDay, &xu.Created, &xu.Updated}
		err = rows.Scan(xdst...)
		if err != nil {
			log.Error(query)
			log.Error(args...)
			log.Error(err)
			return 0, nil, err
		}
		log.Debugf("%#v", xu)
	}
	if err = rows.Err(); err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return 0, nil, err
	}
	return total, data, nil
}
