package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cfishbein/forum/internal/db"
	"github.com/cfishbein/forum/internal/model"
	"github.com/gin-gonic/gin"
)

var categories []model.Category

// RegisterCategories registers the DB stored list of categories to the router
func RegisterCategories() {
	cats, err := db.ListCategories()
	if err != nil {
		panic(err)
	}
	categories = cats
}

// AddUser attempts to add a new User
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	user := model.User{Name: name}
	err := db.AddUser(user)
	if err != nil {
		serverError(c, err.Error())
	} else {
		c.JSON(http.StatusCreated, gin.H{})
	}
}

// GetUser attempts to get an existing User
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidRequest(c, "Invalid User ID")
		return
	}
	user, err := db.GetUser(id)
	if err != nil {
		serverError(c, "User not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// ListUsers lists all users
func ListUsers(c *gin.Context) {
	users, err := db.ListUsers()
	if err != nil {
		serverError(c, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

// AddTopic adds a new topic
func AddTopic(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		invalidRequest(c, "Invalid Category ID")
		return
	}
	category, err := db.GetCategory(categoryID)
	if err != nil {
		invalidRequest(c, "Category ID not found")
	}

	userID, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		invalidRequest(c, "Invalid User ID")
		return
	}
	author, err := db.GetUser(userID)
	if err != nil {
		invalidRequest(c, "User not found")
		return
	}

	title := c.PostForm("title")
	topic, err := model.NewTopic(title, *author)
	if err != nil {
		invalidRequest(c, err.Error())
		return
	}

	// Add the Topic
	err = db.AddTopic(*category, topic)
	if err != nil {
		serverError(c, err.Error())
		return
	}

	// Add the Post
	content := c.PostForm("content")
	post, err := model.NewPost(content, *author)
	if err != nil {
		invalidRequest(c, err.Error())
		return
	}

	err = db.AddPost(topic.ID, *post)
	if err != nil {
		serverError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// ListTopics lists all Topics for a CategoryID
func ListTopics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		invalidRequest(c, "Invalid Category ID")
	} else {
		topics, err := db.ListTopics(id)
		if err != nil {
			serverError(c, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"topics": topics,
		})
	}
}

// GetPosts gets all posts for a topic ID in the path param
func GetPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidRequest(c, "Invalid Post ID")
		return
	}
	posts, err := db.GetPosts(id)
	if err != nil {
		serverError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// AddPost adds a post to the database
func AddPost(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidRequest(c, "Invalid Topic ID")
		return
	}
	userID, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		invalidRequest(c, "Invalid User ID")
		return
	}
	content := c.PostForm("content")
	author, err := db.GetUser(userID)
	if err != nil {
		invalidRequest(c, "User ID not found")
		return
	}

	// TODO FK's not being enforce in sqlite3, so topic ID isn't validated
	post, err := model.NewPost(content, *author)
	if err != nil {
		invalidRequest(c, err.Error())
	}
	err = db.AddPost(topicID, *post)
	if err != nil {
		serverError(c, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func invalidRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": msg})
}

func serverError(c *gin.Context, msg string) {
	log.Println(msg)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
}
