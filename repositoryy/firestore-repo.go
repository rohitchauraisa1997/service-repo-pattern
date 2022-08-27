package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	"google.golang.org/api/iterator"
)

type repo struct{}

// To create instance
// constructor function
func NewPostFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "service-repo"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatal("failed adding a new post!! ", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	fmt.Println("FindAll triggered")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}

	// defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)
	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal("failed to iterate the list of posts ", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) FindByID(id string) (*entity.Post, error) {
	var post entity.Post
	fmt.Println("Finding id ", id)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}
	defer client.Close()
	posts := client.Collection(collectionName)
	docref := posts.Doc(id)
	docsnap, err := docref.Get(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, err
		}
		log.Fatal("failed to findByID ", err)
		return nil, err
	}
	docsnap.DataTo(&post)

	return &post, nil
}
