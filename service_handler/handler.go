package service_handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"
    "github.com/gorilla/mux"
    "db_handler"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan Job)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to notyet!")
}

func PostGetAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    _posts := db_handler.GetPost()
    if err := json.NewEncoder(w).Encode(_posts); err != nil {
        panic(err)
    }
}

func CommentGetByPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postId := vars["postId"]
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    // _postid,_ := strconv.Atoi(postId)
    // _comments := db_handler.GetComment(_postid) 
    _comments := db_handler.GetComment(postId) 
    if err := json.NewEncoder(w).Encode(_comments); err != nil {
        panic(err)
    }
    // fmt.Fprintln(w, "Todo show:", todoId)
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
    var post db_handler.Post
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &post); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    //do concurrently
    // go db_handler.CreatePost(post)

    job := Job{JobType:RequestPost , Data:body}
    WorkQueue <- job

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func PostComment(w http.ResponseWriter, r *http.Request) {
    var comment db_handler.Comment
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &comment); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    if(db_handler.CheckPostExist(strconv.Itoa(comment.Post_id))){
        //do concurrently
        // go db_handler.CreateComment(comment)
        job := Job{JobType:RequestComment, Data:body}
        WorkQueue <- job
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
    var comment db_handler.Comment
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &comment); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    if(db_handler.CheckCommentExist(strconv.Itoa(comment.Id))){
        //do concurrently
        // go db_handler.IncrementLike(strconv.Itoa(comment.Id))
        job := Job{JobType:RequestLike, Data:body}
        WorkQueue <- job
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}