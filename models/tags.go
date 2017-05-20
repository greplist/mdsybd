package models

import "time"

const (
	tagfields = "tags.id, tags.name, tags.taggings_count"
)

// Tag -main tag struct
type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"taggings_count"`
}

// TagCreate insert row to tag table
func (c *Client) TagCreate(tag *Tag) error {
	_, err := c.oracle.Exec("INSERT INTO tags (id, name) VALUES (ids.nextval, :1)", tag.Name)
	return err
}

// Tags - list all tags
func (c *Client) Tags() (tags []Tag, err error) {
	rows, err := c.oracle.Query("select " + tagfields + " from tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags = make([]Tag, 0, 128)

	var tag Tag
	for rows.Next() {
		rows.Scan(&tag.ID, &tag.Name, &tag.Count)
		tags = append(tags, tag)
	}
	return tags, rows.Err()
}

// TagDelete - delete tag by id
func (c *Client) TagDelete(id int64) error {
	_, err := c.oracle.Exec("delete from tags where id = :1", id)
	return err
}

// TagDeleteByName - delete tag by name
func (c *Client) TagDeleteByName(name string) error {
	_, err := c.oracle.Exec("delete from tags where name = :1", name)
	return err
}

const (
	taggingfields = "taggings.id, taggings.created_at, taggings.tag_id, taggings.content_id"
)

// Tagging - main struct for taggings model
type Tagging struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	TagID     int64     `json:"tag_id"`
	ContentID int64     `json:"content_id"`
}

// TaggingCreate - create m2m relation between tag and content (add tag to content)
func (c *Client) TaggingCreate(t *Tagging) error {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now().UTC()
	}
	_, err := c.oracle.Exec("INSERT INTO taggings ("+taggingfields+") VALUES (ids.nextval, :1, :2, :3)",
		t.CreatedAt, t.TagID, t.ContentID,
	)
	return err
}

// TaggingDelete - delete m2m relation between tag and content (delete tag from content)
func (c *Client) TaggingDelete(content int64, tag string) error {
	_, err := c.oracle.Exec("delete from taggings where content_id = :1 and tag = :2", content, tag)
	return err
}
