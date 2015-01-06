package main

import(
  "log"
  "net/http"
  "time"
  "fmt"
)

type DelayHandler struct {
  delayMs int32
  proxy http.Handler
}

func (h *DelayHandler) ServeHTTP(writer http.ResponseWriter,
                                request *http.Request) {
  fmt.Printf("Delaying %d ms...\n", h.delayMs)
  time.Sleep(time.Duration(h.delayMs) * time.Millisecond)
  fmt.Println("Serving...");
  h.proxy.ServeHTTP(writer, request)
}

func delayFileHandler(folder string, delay int32) http.Handler {
  return &DelayHandler{delay, http.FileServer(http.Dir(folder))}
}

func main() {
  http.Handle("/", delayFileHandler("/Users/ycsteven/workspace/go/files", 5000))
  log.Fatal(http.ListenAndServe(":8899", nil))
}
