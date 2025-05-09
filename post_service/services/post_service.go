package services

import (
	"goproject/post_service/data"
	"goproject/post_service/models"
)

func GetFilteredPosts() ([]models.Post, error) {
	posts, err := data.FetchPostsFromDB()

	if err != nil {
		return nil, err
	}

	var filtered []models.Post
	for _, post := range posts {
		if post.ID < 5 {
			filtered = append(filtered, post)
		}
	}
	return filtered, nil
}

func SavePost(post models.Post) (models.Post, error) {
	if post.Title == "" {
		post.Title = "Untitled Post"
	}
	return data.SavePostToDB(post)
}
