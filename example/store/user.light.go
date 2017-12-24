package store

import (
	"bytes"

	"github.com/arstd/light/example/model"
	"github.com/arstd/light/light"
	"github.com/arstd/log"
)

type UserStore struct{}

func (*UserStore) Insert(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`INSEST INTO users(username, phone, address, status, birthday, created, updated) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`)
	args = append(args, u.Username, light.String(&u.Phone), u.Address, light.Uint8(&u.Status), u.Birthday)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
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
func (*UserStore) Update(u *model.User) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`UPDATE UPDATE users SET `)
	if u.Username != "" {
		buf.WriteString(`username=?, `)
		args = append(args, u.Username)
	}
	if u.Phone != "" {
		buf.WriteString(`phone=?, `)
		args = append(args, light.String(&u.Phone))
	}
	if u.Address != nil {
		buf.WriteString(`address=?, `)
		args = append(args, u.Address)
	}
	if u.Status != 0 {
		buf.WriteString(`status=?, `)
		args = append(args, light.Uint8(&u.Status))
	}
	if u.Birthday != nil {
		buf.WriteString(`birthday=?, `)
		args = append(args, u.Birthday)
	}
	buf.WriteString(`updated=CURRENT_TIMESTAMP WHERE id=?`)
	args = append(args, u.Id)
	query := buf.String()
	log.Debug(query)
	log.Debug(args...)
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
func (*UserStore) Delete(id uint64) (int64, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`DELETE FROM users WHERE id=?`)
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
	return res.LastInsertId()
}
func (*UserStore) Get(id uint64) (*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`SELECT id,username,phone,address,status,birthday,created,updated FROM users WHERE id=?`)
	args = append(args, id)
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
		xdst := []interface{}{&xu.Id, &xu.Username, light.String(&xu.Phone), &xu.Address, light.Uint8(&xu.Status), &xu.Birthday, &xu.Created, &xu.Updated}
		err = rows.Scan(xdst...)
		if err != nil {
			log.Error(query)
			log.Error(args...)
			log.Error(err)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	return data, nil
}
func (*UserStore) List(u *model.User, offset int, size int) ([]*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`SELECT id,username,phone,address,status,birthday,created,updated FROM users WHERE username LIKE ? `)
	args = append(args, u.Username)
	if u.Phone != "" {
		buf.WriteString(`AND address = ?`)
		args = append(args, u.Address)
		if u.Phone != "" {
			buf.WriteString(`AND phone LIKE ? `)
			args = append(args, light.String(&u.Phone))
		}
		buf.WriteString(`AND created > ? `)
		args = append(args, u.Created)
	}
	buf.WriteString(`AND status != ? `)
	args = append(args, light.Uint8(&u.Status))
	if !u.Updated.IsZero() {
		buf.WriteString(`AND updated > ? `)
		args = append(args, u.Updated)
	}
	buf.WriteString(`AND birthday IS NOT NULL `)
	buf.WriteString(`ORDER BY updated DESC LIMIT ?, ?`)
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
		xdst := []interface{}{&xu.Id, &xu.Username, light.String(&xu.Phone), &xu.Address, light.Uint8(&xu.Status), &xu.Birthday, &xu.Created, &xu.Updated}
		err = rows.Scan(xdst...)
		if err != nil {
			log.Error(query)
			log.Error(args...)
			log.Error(err)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	return data, nil
}
func (*UserStore) Page(u *model.User, page int, size int) (int64, []*model.User, error) {
	var buf bytes.Buffer
	var args []interface{}
	buf.WriteString(`SELECT id,username,phone,address,status,birthday,created,updated FROM users WHERE username LIKE ? `)
	args = append(args, u.Username)
	if u.Phone != "" {
		buf.WriteString(`AND phone LIKE ? `)
		args = append(args, light.String(&u.Phone))
	}
	buf.WriteString(`AND status != ? `)
	args = append(args, light.Uint8(&u.Status))
	if !u.Updated.IsZero() {
		buf.WriteString(`AND updated > ? `)
		args = append(args, u.Updated)
	}
	buf.WriteString(`ORDER BY updated DESC LIMIT ?, ?`)
	args = append(args, page, size)
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
		xdst := []interface{}{&xu.Id, &xu.Username, light.String(&xu.Phone), &xu.Address, light.Uint8(&xu.Status), &xu.Birthday, &xu.Created, &xu.Updated}
		err = rows.Scan(xdst...)
		if err != nil {
			log.Error(query)
			log.Error(args...)
			log.Error(err)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
		return nil, err
	}
	return data, nil
}
