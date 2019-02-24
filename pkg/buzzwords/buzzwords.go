package buzzwords

import (
  "io/ioutil"
  "fmt"
  "encoding/json"
  "math/rand"
  "time"
)

type Buzzwords struct {
  Adj []string `json:adj`
  Sub []string `json:sub`
  App []string `json:app`
}

type Buzzword struct {
  Buzzwords string `json:buzzwords`
}

var b Buzzwords

var adjCount int
var subCount int
var appCount int

func Init(p string) error {
  err := readFile(p)
  if err != nil {
    return err
  }
  rand.Seed(time.Now().Unix())
  return nil
}

func readFile(p string) error {
  c, err := ioutil.ReadFile(p)
  if err != nil {
    return err
  }
  err = json.Unmarshal(c, &b)
  if err != nil {
    return err
  }

  adjCount = len(b.Adj)
  subCount = len(b.Sub)
  appCount = len(b.App)

  return nil
}

func getRandomWords() string {
  adj := b.Adj[rand.Intn(adjCount)]
  sub := b.Sub[rand.Intn(subCount)]
  app := b.App[rand.Intn(appCount)]

  return fmt.Sprintf("%s %s %s", adj, sub, app)
}

func FormatHTML() string {
  w := getRandomWords()
  html := fmt.Sprintf("<html><head><title>Buzzwords As a Service</title></head><body><center style=\"font-family: monospace\"><h1>%s</h1></center></body></html>", w)
  return html
}

func FormatJSON() ([]byte, error) {
  w := getRandomWords()
  o := Buzzword{w}
  j, err := json.Marshal(o)
  if err != nil {
    return nil, err
  }
  return j, nil
}
