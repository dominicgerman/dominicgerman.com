package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID          int
	Title       string
	Description string
	Tags        []string
	Content     string
	Created     time.Time
	Updated     time.Time
}

// Define a PostModel type which wraps a sql.DB connection pool.
type PostModel struct {
	DB *sql.DB
}

// This will insert a new post and its tags into the database.
func (m *PostModel) Insert(title string, description string, tags []string, content string) (int, error) {
	postsStmt := `INSERT INTO posts (title, description, content, created, updated)
    VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	tagQuery := "INSERT OR IGNORE INTO tags (name) VALUES (?)"

	tagsStmt, err := m.DB.Prepare(tagQuery)
	if err != nil {
		return 0, err
	}
	defer tagsStmt.Close()

	postResult, err := m.DB.Exec(postsStmt, title, description, content)
	if err != nil {
		return 0, err
	}

	postID, err := postResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Prepare the statement for associating posts with tags.
	postTagQuery := "INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)"
	postTagStmt, err := m.DB.Prepare(postTagQuery)
	if err != nil {
		return 0, err
	}
	defer postTagStmt.Close()

	// Loop over tags array, insert tags, and associate them with the post.
	for _, tag := range tags {
		// Insert the tag or do nothing if it already exists.
		_, err = tagsStmt.Exec(tag)
		if err != nil {
			return 0, err
		}

		var tagID int
		err = m.DB.QueryRow("SELECT id FROM tags WHERE name = ?", tag).Scan(&tagID)
		if err != nil {
			return 0, err
		}

		// Insert the association into the post_tags table.
		_, err = postTagStmt.Exec(postID, tagID)
		if err != nil {
			return 0, err
		}
	}

	// The ID returned has the type int64, so we convert it to an int type
	return int(postID), nil

}

// This will return a specific post based on its id.
func (m *PostModel) Get(id int) (Post, error) {
	stmt := `SELECT id, title, description, content, created, updated FROM posts 
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	p := Post{}

	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Content, &p.Created, &p.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		} else {
			return Post{}, err
		}
	}

	// Query to select the tags associated with the post.
	tagsStmt := `SELECT t.name FROM tags t
		INNER JOIN post_tags pt ON t.id = pt.tag_id
		WHERE pt.post_id = ?`

	// Execute the query.
	rows, err := m.DB.Query(tagsStmt, id)
	if err != nil {
		return Post{}, err
	}
	defer rows.Close()

	// Loop over the rows and append each tag to the Post's Tags slice.
	var tags []string
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return Post{}, err
		}
		tags = append(tags, tag)
	}

	// Check for errors during the row iteration.
	if err = rows.Err(); err != nil {
		return Post{}, err
	}

	// Assign the tags to the Post struct.
	p.Tags = tags
	return p, nil
}

// This will return the 10 most recently created posts.
func (m *PostModel) Latest() ([]Post, error) {
	stmt := `SELECT id, title, description, content, created, updated FROM posts 
    ORDER BY updated DESC LIMIT 10`

	// Execute the SQL statement and get the result set.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	// Iterate over each row in the result set.
	for rows.Next() {
		var p Post
		err = rows.Scan(&p.ID, &p.Title, &p.Description, &p.Content, &p.Created, &p.Updated)
		if err != nil {
			return nil, err
		}

		// Retrieve the tags for the current post.
		tagsStmt := `SELECT t.name FROM tags t
			INNER JOIN post_tags pt ON t.id = pt.tag_id
			WHERE pt.post_id = ?`

		// Query the tags for the current post.
		tagRows, err := m.DB.Query(tagsStmt, p.ID)
		if err != nil {
			return nil, err
		}
		defer tagRows.Close()

		// Append each tag to the Post's Tags slice.
		var tags []string
		for tagRows.Next() {
			var tag string
			err = tagRows.Scan(&tag)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}

		// Check for any error encountered during iteration.
		if err = tagRows.Err(); err != nil {
			return nil, err
		}

		// Assign the tags to the Post struct.
		p.Tags = tags

		// Append the post to the posts slice.
		posts = append(posts, p)
	}

	// Check for any error encountered during iteration of the rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) Update(id int, title string, description string, tags []string, content string) (int, error) {
	postsStmt := `UPDATE posts
	SET title = ?, description = ?, content = ?, updated = CURRENT_TIMESTAMP
    WHERE id = ?;`

	_, err := m.DB.Exec(postsStmt, title, description, content, id)
	if err != nil {
		return 0, err
	}

	tagQuery := "INSERT OR IGNORE INTO tags (name) VALUES (?)"

	tagsStmt, err := m.DB.Prepare(tagQuery)
	if err != nil {
		return 0, err
	}
	defer tagsStmt.Close()

	deleteTagsStmt := `DELETE FROM post_tags WHERE post_id = ?`
	_, err = m.DB.Exec(deleteTagsStmt, id)
	if err != nil {
		return 0, err
	}

	// Now, associate new tags with the post.
	postTagStmt := `INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)`
	for _, tag := range tags {
		// Insert the tag if it doesn't already exist.
		_, err = tagsStmt.Exec(tag)
		if err != nil {
			return 0, err
		}

		// Get the tag ID.
		var tagID int
		err = m.DB.QueryRow("SELECT id FROM tags WHERE name = ?", tag).Scan(&tagID)
		if err != nil {
			return 0, err
		}

		// Insert the association between the post and the tag.
		_, err = m.DB.Exec(postTagStmt, id, tagID)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
