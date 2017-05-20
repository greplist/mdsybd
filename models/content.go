package models

import (
	"database/sql"
	"time"
)

const (
	contentFormatFields = "content_formats.id, content_formats.created_at, content_formats.updated_at, " +
		"content_formats.text_align, content_formats.title_color, content_formats.title_background_color, " +
		"content_formats.author_color, content_formats.category_color, " +
		"content_formats.background_color, content_formats.background_image"
)

// ContentFormat - main struct for content format model
type ContentFormat struct {
	ID                   int64     `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Align                int       `json:"text_align"`
	TitleColor           int       `json:"title_color"`
	TitleBackgroundColor int       `json:"title_background_color"`
	AuthorColor          int       `json:"author_color"`
	CategoryColor        int       `json:"category_color"`
	BackgroundColor      int       `json:"background_color"`
	BackgroundImage      string    `json:"background_image"`
}

// ContentFormatCreate insert row to content foramat table
func (c *Client) ContentFormatCreate(format *ContentFormat) error {
	now := time.Now().UTC()
	if format.CreatedAt.IsZero() {
		format.CreatedAt = now
	}
	if format.UpdatedAt.IsZero() {
		format.UpdatedAt = now
	}
	_, err := c.oracle.Exec("INSERT INTO content_formats ("+contentFormatFields+
		") VALUES (ids.nextval, :1, :2, :3, :4, :5, :6, :7, :8, :9)",
		format.CreatedAt, format.UpdatedAt, format.Align, format.TitleColor,
		format.TitleBackgroundColor, format.AuthorColor, format.CategoryColor,
		format.BackgroundColor, format.BackgroundImage,
	)
	return err
}

// ContentFormat - get content format by id
func (c *Client) ContentFormat(id int64) (format *ContentFormat, err error) {
	format = &ContentFormat{}
	err = c.oracle.QueryRow("select "+contentFormatFields+" from content_formats where id = :1", id).Scan(
		&format.ID, &format.CreatedAt, &format.UpdatedAt, &format.Align, &format.TitleColor,
		&format.TitleBackgroundColor, &format.AuthorColor, &format.CategoryColor,
		&format.BackgroundColor, &format.BackgroundImage,
	)
	return
}

// Formats - list all formats
func (c *Client) Formats() (formats []ContentFormat, err error) {
	rows, err := c.oracle.Query("select " + contentFormatFields + " from content_formats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	formats = make([]ContentFormat, 0, 128)

	var format ContentFormat
	for rows.Next() {
		rows.Scan(
			&format.ID, &format.CreatedAt, &format.UpdatedAt, &format.Align, &format.TitleColor,
			&format.TitleBackgroundColor, &format.AuthorColor, &format.CategoryColor,
			&format.BackgroundColor, &format.BackgroundImage,
		)
		formats = append(formats, format)
	}
	return formats, rows.Err()
}

// FormatDelete delete content format by id
func (c *Client) FormatDelete(id int64) error {
	_, err := c.oracle.Exec("DELETE FROM content_formats where id = :1", id)
	return err
}

const (
	contentfields = "contents.id, contents.created_at, contents.updated_at, contents.title, contents.content, " +
		"contents.published, contents.published_at, contents.repost_count, contents.category_id, " +
		"contents.content_format_id "
)

// Content - main struct for content model
type Content struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     []byte    `json:"content"`
	Published   int       `json:"published"`
	PublishedAt time.Time `json:"published_at"`
	RepostCount int64     `json:"repost_count"`

	CategoryID int64 `json:"category_id,ommitempty"`
	FormatID   int64 `json:"content_format_id,ommitempty"`

	Category *Category      `json:"category,ommitempty"`
	Format   *ContentFormat `json:"content_format,ommitempty"`
}

// IsPublished - check content published
func (c *Content) IsPublished() bool {
	return c.Published == 1
}

// ContentCreate insert row to content table
func (c *Client) ContentCreate(content *Content) error {
	now := time.Now().UTC()
	if content.CreatedAt.IsZero() {
		content.CreatedAt = now
	}
	if content.UpdatedAt.IsZero() {
		content.UpdatedAt = now
	}
	if content.IsPublished() && content.PublishedAt.IsZero() {
		content.PublishedAt = now
	}

	format := content.FormatID
	if content.Format != nil && content.Format.ID != 0 {
		format = content.Format.ID
	}
	category := content.CategoryID
	if content.Category != nil && content.Category.ID != 0 {
		category = content.Category.ID
	}
	_, err := c.oracle.Exec("INSERT INTO contents ("+contentfields+
		") VALUES (ids.nextval, :1, :2, :3, :4, :5, :6, :7, :8, :9)",
		content.CreatedAt, content.UpdatedAt, content.Title, content.Content,
		content.Published, content.PublishedAt, content.RepostCount,
		category, format,
	)
	return err
}

// Content - get content by id (without relations)
func (c *Client) Content(id int64) (content *Content, err error) {
	content = &Content{}
	err = c.oracle.QueryRow("select "+contentfields+" from contents where id = :1", id).Scan(
		&content.ID, &content.CreatedAt, &content.UpdatedAt, &content.Title, &content.Content,
		&content.Published, &content.PublishedAt, &content.RepostCount,
		&content.CategoryID, &content.FormatID,
	)
	return
}

// ContentRelation - select all content data with relations
func (c *Client) ContentRelation(id int64) (content *Content, err error) {
	content = &Content{}
	category, format := &Category{}, &ContentFormat{}
	content.Category, content.Format = category, format
	err = c.oracle.QueryRow("select "+contentfields+","+categoryfields+","+contentFormatFields+
		" from contents INNER JOIN categories ON contents.category_id = categories.id "+
		"INNER JOIN contents.content_format_id = content_formats_id where contents.id = :1", id).Scan(
		&content.ID, &content.CreatedAt, &content.UpdatedAt, &content.Title, &content.Content,
		&content.Published, &content.PublishedAt, &content.RepostCount,
		&content.CategoryID, &content.FormatID,
		&category.ID, &category.CreatedAt, &category.UpdatedAt, &category.Name, &category.Visible,
		&format.ID, &format.CreatedAt, &format.UpdatedAt, &format.Align, &format.TitleColor,
		&format.TitleBackgroundColor, &format.AuthorColor, &format.CategoryColor,
		&format.BackgroundColor, &format.BackgroundImage,
	)
	return
}

func execContents(rows *sql.Rows) (contents []Content, err error) {
	contents = make([]Content, 0, 128)

	var content Content
	for rows.Next() {
		rows.Scan(
			&content.ID, &content.CreatedAt, &content.UpdatedAt, &content.Title, &content.Content,
			&content.Published, &content.PublishedAt, &content.RepostCount,
			&content.CategoryID, &content.FormatID,
		)
		contents = append(contents, content)
	}
	return contents, rows.Err()
}

// Contents - select all contents
func (c *Client) Contents() (contents []Content, err error) {
	rows, err := c.oracle.Query("select " + contentfields + " from contents ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return execContents(rows)
}

// ContentsByCategory - select contents by category name
func (c *Client) ContentsByCategory(name string) (contents []Content, err error) {
	rows, err := c.oracle.Query("select "+contentfields+
		" from contents INNER JOIN categories ON contents.category_id = categories.id "+
		" where category.name = :1", name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return execContents(rows)
}

// ContentDelete delete content by id
func (c *Client) ContentDelete(id int64) error {
	_, err := c.oracle.Exec("DELETE FROM contents where id = :1", id)
	return err
}
