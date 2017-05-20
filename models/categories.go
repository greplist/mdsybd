package models

import "time"

const (
	categoryfields = "id, created_at, updated_at, name, visible"
)

// Category - main struct for category model
type Category struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Visible   int       `json:"visible"`
}

// IsVisible check is category visible
func (c *Category) IsVisible() bool {
	return c.Visible == 1
}

// CategoryCreate insert row to category table
func (c *Client) CategoryCreate(category *Category) error {
	now := time.Now().UTC()
	if category.CreatedAt.IsZero() {
		category.CreatedAt = now
	}
	if category.UpdatedAt.IsZero() {
		category.UpdatedAt = now
	}
	_, err := c.oracle.Exec("INSERT INTO categories ("+categoryfields+") VALUES (ids.nextval, :1, :2, :3, :4)",
		category.CreatedAt, category.UpdatedAt, category.Name, category.Visible)
	return err
}

// Category - get category by id
func (c *Client) Category(id int64) (category *Category, err error) {
	category = &Category{}
	err = c.oracle.QueryRow("select "+categoryfields+" from categories where id = :1", id).Scan(
		&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.Name, &category.Visible,
	)
	return
}

// CategoryByName - get category by name or email
func (c *Client) CategoryByName(name string) (category *Category, err error) {
	category = &Category{}
	err = c.oracle.QueryRow("select "+categoryfields+" from categories where name = :1", name).Scan(
		&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.Name, &category.Visible,
	)
	return
}

// CategoryDelete delete category by id
func (c *Client) CategoryDelete(id int64) error {
	_, err := c.oracle.Exec("DELETE FROM categories where id = :1", id)
	return err
}
