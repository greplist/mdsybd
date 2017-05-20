package models

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
