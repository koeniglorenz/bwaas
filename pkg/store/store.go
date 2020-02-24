package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Store interface {
	GetHTML() string
	GetJSON() ([]byte, error)
}

type buzzwords struct {
	Adj []string `json:"adj"`
	Sub []string `json:"sub"`
	App []string `json:"app"`
}

type buzzword struct {
	Buzzwords string `json:"buzzwords"`
}

type jsonstore struct {
}

var b buzzwords

var adjCount int
var subCount int
var appCount int

func New(p string) (*jsonstore, error) {
	c, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(c, &b)
	if err != nil {
		return nil, err
	}

	adjCount = len(b.Adj)
	subCount = len(b.Sub)
	appCount = len(b.App)

	s := jsonstore{}

	rand.Seed(time.Now().Unix())
	return &s, nil
}

func getRandomWords() string {
	adj := b.Adj[rand.Intn(adjCount)]
	sub := b.Sub[rand.Intn(subCount)]
	app := b.App[rand.Intn(appCount)]

	return fmt.Sprintf("%s %s %s", adj, sub, app)
}

func (store *jsonstore) GetHTML() string {
	w := getRandomWords()
	html := fmt.Sprintf(`
		<html>
			<head>
				<title>Buzzwords As a Service</title>
			</head>
			<body style="height: 100%%; display: grid">
				<h1 style="font-size:4vw; font-style=Sans-Serif; margin: auto;">%s</h1>
			</body>
		</html>`, w)
	return html
}

func (store *jsonstore) GetJSON() ([]byte, error) {
	w := getRandomWords()
	o := buzzword{w}
	j, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return j, nil
}
