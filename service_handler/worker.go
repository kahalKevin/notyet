package service_handler

import (
  "log"
  "strconv"
  "db_handler"
  "encoding/json"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan Job) Worker {
  // Create, and return the worker.
  worker := Worker{
    ID:          id,
    Work:        make(chan Job),
    WorkerQueue: workerQueue,
    QuitChan:    make(chan bool)}
  
  return worker
}

type Worker struct {
  ID          int
  Work        chan Job
  WorkerQueue chan chan Job
  QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
    go func() {
      for {
        // Add ourselves into the worker queue.
        w.WorkerQueue <- w.Work
        log.Printf("adding worker to queue\n")

        select {
        case work := <-w.Work:
          // Receive a work request.
          switch work.JobType{
          case RequestPost:
            var post db_handler.Post
            json.Unmarshal(work.Data, &post)
            log.Printf("worker%d: Received work request, get Post %s \n", w.ID, post.Content)
            go db_handler.CreatePost(post)          
          case RequestComment:
            var comment db_handler.Comment
            json.Unmarshal(work.Data, &comment)
            log.Printf("worker%d: Received work request, get Post %s \n", w.ID, comment.Comment)
            go db_handler.CreateComment(comment)
          case RequestLike:
            var comment db_handler.Comment
            json.Unmarshal(work.Data, &comment)
            log.Printf("worker%d: Received work request, get Post %s \n", w.ID, comment.Id)
            go db_handler.IncrementLike(strconv.Itoa(comment.Id))
          }
          // fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())
          
          // time.Sleep(work.Delay)
          // fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)
          
        case <-w.QuitChan:
          // We have been asked to stop.
          log.Printf("worker%d stopping\n", w.ID)
          return
        }
      }
    }()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
  go func() {
    w.QuitChan <- true
  }()
}