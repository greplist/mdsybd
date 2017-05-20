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

// Personality - main struct for personality model
type Personality struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Admin     int       `json:"is_admin"`
	ContendID int64     `json:"content_id"`
	UserID    int64     `json:"user_id"`
}

// CreatePersonality - create m2m relation beetwen user & content
func (c *Client) CreatePersonality(p *Personality) error {
	now := time.Now().UTC()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}
	_, err := c.oracle.Exec(
		"INSERT INTO personalities (id, created_at, updated_at, is_admin, content_id, user_id) "+
			"VALUES (ids.nextval, :1, :2, :3, :4, :5)",
		p.CreatedAt, p.UpdatedAt, p.Admin, p.ContendID, p.UserID,
	)
	return err
}

// Authors - returns authors by content id
func (c *Client) Authors(content int64) (authors []string, err error) {
	rows, err := c.oracle.Query(
		" select users.name from users "+
			" INNER JOIN personalities ON personalities.user_id = users.id "+
			" INNER JOIN contents ON personalities.content_id = contents.id "+
			" WHERE contents.id = :1 ",
		content,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors = make([]string, 0, 32)

	var author string
	for rows.Next() {
		rows.Scan(&author)
		authors = append(authors, author)
	}
	return authors, rows.Err()
}

// DeletePersonality - delete m2m rekation beetwen user & content (drop author)
func (c *Client) DeletePersonality(user, content int64) error {
	_, err := c.oracle.Exec("DELETE FROM personalities where user_id = :1 and content_id = :2", user, content)
	return err
}
