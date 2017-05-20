package models

import "time"

const (
	userfields = "users.id, users.created_at, users.updated_at, users.name, users.email, users.role, users.encrypted_password"
)

// User - main struct for user model
type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"name"`
	Role      int64     `json:"-"`
	Password  string    `json:"encrypted_password"`
}

// IsAdmin check is user admin
func (u *User) IsAdmin() bool {
	return u.Role == 1
}

// UserCreate insert row to user table
func (c *Client) UserCreate(user *User) error {
	now := time.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}
	_, err := c.oracle.Exec("INSERT INTO users ("+userfields+") VALUES (ids.nextval, :1, :2, :3, :4, :5, :6)",
		user.CreatedAt, user.UpdatedAt, user.Name, user.Email)
	return err
}

// User - get user by id
func (c *Client) User(id int64) (u *User, err error) {
	u = &User{}
	err = c.oracle.QueryRow("select "+userfields+" from users where id = :1", id).Scan(
		&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name, &u.Email, &u.Role, &u.Password,
	)
	return
}

// UserByLogin - get user by name or email
func (c *Client) UserByLogin(login string) (u *User, err error) {
	u = &User{}
	err = c.oracle.QueryRow("select "+userfields+" from users where name = :1 or email = :1", login).Scan(
		&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name, &u.Email, &u.Role, &u.Password,
	)
	return
}

// UserDelete delete user by id
func (c *Client) UserDelete(id int64) error {
	_, err := c.oracle.Exec("DELETE FROM users where id = :1", id)
	return err
}
