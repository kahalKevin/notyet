package db_handler

// import "time"
// import "github.com/jinzhu/gorm"

type Post struct {
	// gorm.Model 
	// idposts     int       `gorm:"primary_key"`
 //    title      	string    
 //    content     string      
 //    created_by  string
	// PageId string                 `bson:"pageId" json:"pageId"`
 //    Meta   map[string]interface{} `bson:"meta" json:"pageId"`
	Idposts     int       
    Title      	string    `json:"title"`
    Content     string    `json:"content"`
    Created_by  string	  `json:"created_by"`
}
type Posts []Post

type Comment struct {
	Id     		int    `json:"comment_id"`
    Post_id     int    `json:"post_id"`
    Comment     string `json:"comment"`
    Like_count  int	   
}
type Comments []Comment

// func (Post) TableName() string {
//   return "posts"
// }