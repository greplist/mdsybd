package models

import "time"

const (
	contentFormatFields = "id, created_at, updated_at, text_align, title_color, title_background_color, " +
		"author_color, category_color, background_color, background_image"
	contentfields = "id, created_at, updated_at, name, visible"
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

// ContentFormatCreate insert row to category table
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
