package main

import(
  "fmt"
  "io/ioutil"
  "net/http"
)

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
  req.Header.Add("User-Agent", via[0].UserAgent())
  fmt.Printf("Redirected to %s, user agent: %s\n", req.URL, req.UserAgent())
  return nil 
}

func main() {
  url := "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.26 Safari/537.11"
  client := &http.Client{CheckRedirect: redirectPolicyFunc,}
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("User-Agent", url)
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("error")
    return
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}
