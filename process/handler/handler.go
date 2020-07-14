package handler

import (
	"errors"
	"fmt"
	"usermanager/process/schema"
	"usermanager/util/db"
)

func CreateHandler(user schema.User) error {
	flag := 0
	errmsg := "missing field or field with inappropriate value:"
	if user.Name == "" { // valiating fields
		errmsg = errmsg + " " + "name"
		flag = 1
	}
	if user.Age == 0 {
		flag = 1
		errmsg = errmsg + " " + "age"
	}
	if len(user.Phone) < 10 {
		fmt.Println(user.Phone)
		flag = 1
		errmsg = errmsg + " " + "phone"
	}
	if user.City == "" {
		flag = 1
		errmsg = errmsg + " " + "city"
	}
	if flag == 1 {
		fmt.Println(errmsg)
		return errors.New(errmsg) //if validation is failed
	}
	_, err := db.Exec(fmt.Sprintf("insert into users(id,name,age,phone,city) values('"+user.Id+"','"+user.Name+"',%d,'%s','"+user.City+"');", user.Age, user.Phone))
	if err != nil {
		return err
	}
	return nil
}

//List users
func ListHandler() ([]schema.User, error) {
	rows := db.List("select * from users;")
	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone, &user.City)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

//Read user by id
func ReadHandler(id string) (schema.User, error) {
	rows := db.Read("select * from users where id='" + id + "';")
	var user schema.User
	err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone, &user.City)
	if err != nil {
		return user, err
	}
	return user, nil
}

//Update by id
func UpdateHandler(user schema.User) (schema.User, error) {
	middle := ""
	flag := 0
	if len(user.Name) > 0 { //Generating query to change only necessary field
		if flag == 0 {
			flag = 1
			middle = middle + "name = '" + user.Name + "' "
		} else {
			middle = middle + ",name = '" + user.Name + "' "
		}
	}
	if user.Age > 0 {
		if flag == 0 {
			middle = middle + "age = " + fmt.Sprintf("%d", user.Age)
			flag = 1
		} else {
			middle = middle + ",age = " + fmt.Sprintf("%d", user.Age)
		}
	}
	if len(user.Phone) >= 10 {
		if flag == 0 {
			flag = 1
			middle = middle + "phone = '" + user.Phone + "' "
		} else {
			middle = middle + ",phone = '" + user.Phone + "' "
		}
	} else if len(user.Phone) > 0 {
		return user, errors.New("invalid phone number")
	}
	if len(user.City) > 0 {
		if flag == 0 {
			flag = 1
			middle = middle + "city = '" + user.City + "' "
		} else {
			middle = middle + ",city = '" + user.City + "' "
		}
	}
	sql := "update users SET " + middle + " where id='" + user.Id + "';"
	result, err := db.Exec(sql)
	if err != nil {
		return user, err
	}
	affectedRow, _ := result.RowsAffected()
	if affectedRow == 0 {
		return user, errors.New("id not found")
	}
	return user, nil
}

//Delete by id
func DeleteHandler(id string) error {
	result, err := db.Exec("delete from users where id='" + id + "';")
	if err != nil {
		return err
	}
	affectedRow, _ := result.RowsAffected()
	if affectedRow == 0 {
		return errors.New("id not found")
	}
	return nil
}
