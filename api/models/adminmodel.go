package models

import (
	"database/sql"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var admin_id int64 = 0

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Admin) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Admin) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Login = html.EscapeString(strings.TrimSpace(u.Login))
	u.Password = strings.Trim(u.Password, " ")
}

func (u *Admin) GetAdmins(db *sql.DB) (*[]Admin, error) {
	defer db.Close()
	var Admins []Admin
	rows, err := db.Query("SELECT id, name, login FROM Admins ORDER BY created_at DESC")
	if err != nil {
		return &Admins, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name, &u.Login)
		if err != nil {
			return &Admins, nil
		}
		Admins = append(Admins, *u)
	}
	return &Admins, nil
}

func (u *Admin) GetAdminById(db *sql.DB, Admin_id int64) error {
	defer db.Close()
	row := db.QueryRow("SELECT id, name, login FROM Admins WHERE id=$1", Admin_id)
	err := row.Scan(&u.Id, &u.Name, &u.Login)
	return err
}

func (u *Admin) GetAdminByLogin(db *sql.DB, login string) error {
	defer db.Close()
	row := db.QueryRow("SELECT id, name, login,password FROM Admins WHERE login=$1", login)
	err := row.Scan(&u.Id, &u.Name, &u.Login, &u.Password)
	return err
}

func (u *Admin) AddAdmin(db *sql.DB) (int64, error) {
	defer db.Close()
	u.Prepare()
	u.BeforeSave()
	result, err := db.Exec("INSERT INTO Admins (name, login, password, created_at, updated_at) values ($1, $2, $3, DateTime('now'), DateTime('now') )", u.Name, u.Login, u.Password)
	if err != nil {
		return admin_id, err
	}
	admin_id, err = result.LastInsertId()
	return admin_id, err
}

func (u *Admin) EditAdmin(db *sql.DB) (int64, error) {
	var err error
	defer db.Close()
	u.Prepare()
	if u.Password != "" {
		u.BeforeSave()
		_, err = db.Exec("UPDATE Admins SET name = $1, login = $2, password = $3, updated_at = DateTime('now') WHERE id = $4", u.Name, u.Login, u.Password, u.Id)
	} else {
		_, err = db.Exec("UPDATE Admins SET name = $1, login = $2, updated_at = DateTime('now') WHERE id = $3", u.Name, u.Login, u.Id)
	}

	return u.Id, err
}

func (u *Admin) DeleteAdmin(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("DELETE FROM Admins WHERE id = $1", u.Id)
	return u.Id, err
}

func (u *Admin) DeleteAdminById(db *sql.DB, Admin_id int64) error {
	defer db.Close()
	_, err := db.Exec("DELETE FROM Admins WHERE id=$1", Admin_id)
	return err
}
