package data

import (
	"goproject/post_service/models"
)

func FetchPostsFromDB() ([]models.Post, error) {
	rows, err := DB.Query("SELECT id, userId, title, body FROM Posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func SavePostToDB(post models.Post) (models.Post, error) {
	result, err := DB.Exec("INSERT INTO Posts (userId, title, body) VALUES (?, ?, ?)", post.UserID, post.Title, post.Body)
	if err != nil {
		return models.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Post{}, err
	}

	post.ID = int(id)
	return post, nil
}
