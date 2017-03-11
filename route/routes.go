package route

import(
    "net/http"

    "service_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        service_handler.Index,
    },
    Route{
        "PostIndex",
        "GET",
        "/posts",
        service_handler.PostGetAll,
    },
    Route{
        "CommentFilter",
        "GET",
        "/comment/{postId:[0-9]+}",
        // "/products/{user:\d+}"  FOR DIGIT ONLY
        service_handler.CommentGetByPost,
    },
    Route{
        "CommentCreate",
        "POST",
        "/makecomment",
        service_handler.PostComment,
    },
    Route{
        "PostCreate",
        "POST",
        "/makepost",
        service_handler.PostCreate,
    },
}
