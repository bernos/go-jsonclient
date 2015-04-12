package jsonclient

import (
	"testing"
	"net/http"
)

type Post struct {
	UserId int `json: "userId"`
	Id     int
	Title  string
	Body   string
}

func TestPost(t *testing.T) {
	post := Post{UserId:1,Id:2,Title:"Title",Body:"Body"}

	client, err := NewJsonClient(nil)

	if err != nil {
		t.Error("Failed to create JsonClient")
	}

	res, err := client.Post("http://jsonplaceholder.typicode.com/posts", post)

	if err != nil {
		t.Error("Failed to post request")
	}

	var resultPost Post

	res.To(&resultPost)

	if resultPost.Body != "Body" {
		t.Error("Did not receive valid post back")
	}
}

func TestGetWithDefaultHttpClient(t *testing.T) {
	if c, err := NewJsonClient(nil); err == nil {
		testGet(t, c)
	} else {
		t.Error("Failed to create JsonClient")		
	}
}

func TestGetWithCustomHttpClient(t *testing.T) {
	if c, err := NewJsonClient(&http.Client{}); err == nil {
		testGet(t, c)
	} else {
		t.Error("Failed to create JsonClient")	
	}
}

func TestDelete(t *testing.T) {
	if client, err := NewJsonClient(nil); err != nil {
		t.Error("Failes to create JsonClient")
	} else if res, err := client.Delete("http://jsonplaceholder.typicode.com/posts/1"); err != nil {
		t.Error("Delete failed")		
	} else if res.StatusCode != 204 {
		t.Errorf("Delete failed. Expected status 204, got %d", res.StatusCode)
	}
}

func testGet(t *testing.T, c *Client) {
	
	var posts []Post

	res, err := c.Get("http://jsonplaceholder.typicode.com/posts")

	if err != nil {
		t.Error(err)
	}

	res.To(&posts)

	if len(posts) != 100 {
		t.Error("Did not receive 100 posts")
	}
}