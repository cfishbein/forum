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
	id, err := strconv.Atoi(c.Param("userId"))
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

// AddThread adds a new thread
func AddThread(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		invalidRequest(c, "Invalid Category ID")
		return
	}
	category, err := db.GetCategory(categoryID)
	if err != nil {
		invalidRequest(c, "Category ID not found")
		return
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
	thread, err := model.NewThread(title, *author)
	if err != nil {
		invalidRequest(c, err.Error())
		return
	}

	// Add the thread
	err = db.AddThread(*category, thread)
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

	err = db.AddPost(thread.ID, *post)
	if err != nil {
		serverError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// ListThreads lists all Threads for a CategoryID
func ListThreads(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		invalidRequest(c, "Invalid Category ID")
	} else {
		threads, err := db.ListThreads(id)
		if err != nil {
			serverError(c, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"threads": threads,
		})
	}
}

// ListPosts lists all posts for a thread ID in the path param
func ListPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		invalidRequest(c, "Invalid Thread ID")
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

type addPostReq struct {
	threadID int
	author   model.User
	content  string
}

// AddPost adds a post to the database
func AddPost(c *gin.Context) {
	req, err := newAddPostReq(c)
	if err != nil {
		invalidRequest(c, err.Error())
		return
	}

	// TODO FK's not being enforce in sqlite3, so thread ID isn't validated
	post, err := model.NewPost(req.content, req.author)
	if err != nil {
		invalidRequest(c, err.Error())
	}
	err = db.AddPost(req.threadID, *post)
	if err != nil {
		serverError(c, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func newAddPostReq(c *gin.Context) (*addPostReq, error) {
	tID, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return nil, err
	}

	uID, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		return nil, err
	}
	_content := c.PostForm("content")
	_author, err := db.GetUser(uID)
	if err != nil {
		return nil, err
	}
	return &addPostReq{threadID: tID, author: *_author, content: _content}, nil
}

func invalidRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": msg})
}

func serverError(c *gin.Context, msg string) {
	log.Println(msg)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
}
