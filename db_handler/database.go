package db_handler

import (
	"log"
	"database/sql"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func CreatePost(post Post) {	
	err := Db.Ping()
	if err != nil {
		return
	}

	query_insert := "insert into posts(`title`, `content`, `created_by`) values ('"+post.Title+"', '"+post.Content+"' ,'"+post.Created_by+"')"
	log.Printf(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		log.Fatal(err)
	}
	// lastId, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rowCnt, err := res.RowsAffected()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt) 
}

func CreateComment(comment Comment) {
	err := Db.Ping()
	if err != nil {
		return
	}

	query_insert := "insert into comments(`post_id`, `comment`) values ("+strconv.Itoa(comment.Post_id)+", '"+comment.Comment+"')"
	log.Printf(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		log.Fatal(err)
	}
}

func IncrementLike(commentid string) {
	err := Db.Ping()
	if err != nil {
		return
	}
	query_update_like_count := "update comments set like_count=like_count+1 where id ="+commentid
	log.Printf(query_update_like_count)

	_, err = Db.Exec(query_update_like_count)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPost() Posts{
	err := Db.Ping()
	if err != nil {
		return nil
	}

	rows, err := Db.Query("select idposts, title, content, created_by from posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var _Posts Posts
	var post1 Post
	for rows.Next() {
		err := rows.Scan(&post1.Idposts, &post1.Title, &post1.Content, &post1.Created_by)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(post1)
		_Posts = append(_Posts, post1)
	}
	return _Posts
}

func CheckPostExist(postid string) bool{
	err := Db.Ping()
	if err != nil {
		return false
	}

	post_exist := false
	err = Db.QueryRow("select exists(select idposts from posts where idposts = "+ postid +" limit 1) as 'exist'").Scan(&post_exist)
	if err != nil {
		log.Fatal(err)
	}
	return post_exist
}

func CheckCommentExist(commentid string) bool{
	err := Db.Ping()
	if err != nil {
		return false
	}

	comment_exist := false
	err = Db.QueryRow("select exists(select id from comments where id = "+ commentid +" limit 1) as 'exist'").Scan(&comment_exist)
	if err != nil {
		log.Fatal(err)
	}
	return comment_exist
}

func GetCurrentCommentLike(commentid string) int{
	err := Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	like_count := 0
	err = Db.QueryRow("select like_count from comments where id = "+ commentid +"").Scan(&like_count)
	if err != nil {
		log.Fatal(err)
	}
	return like_count
}

func GetComment(postid string) Comments{	
	err := Db.Ping()
	if err != nil {
		return nil
	}

	// rows, err := Db.Query("select id, post_id, comment, like_count from comments where post_id =" + strconv.Itoa(postid))
	rows, err := Db.Query("select id, post_id, comment, like_count from comments where post_id =" + postid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var _Comments Comments
	var comment1 Comment
	for rows.Next() {
		err := rows.Scan(&comment1.Id, &comment1.Post_id, &comment1.Comment, &comment1.Like_count)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(comment1)
		_Comments = append(_Comments, comment1)
	} 
	return _Comments
}
