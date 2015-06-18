package gephi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const (
	defaultHost      = "localhost:8080"
	defaultWorkspace = "workspace0"
	version          = 1

	addNode    = "an"
	changeNode = "cn"
	deleteNode = "dn"

	addEdge    = "ae"
	changeEdge = "ce"
	deleteEdge = "de"
)

var (
	opsOrder = []string{"an", "ae", "ce", "cn", "de", "dn"}
)

type Action struct {
	Type, ID string
}

type Gephi struct {
	client   *http.Client
	URL      url.URL
	Entities map[Action]Attributes
}

func NewGephi(host string, workspace string) (g *Gephi) {
	g = new(Gephi)
	if host == "" {
		host = defaultHost
	}
	if workspace == "" {
		workspace = defaultWorkspace
	}
	g.Entities = make(map[Action]Attributes)
	g.URL.Scheme = "http"
	g.URL.Host = host
	g.URL.Path = fmt.Sprintf("/%s/", workspace)
	values := url.Values{}
	values.Add("operation", "updateGraph")
	g.URL.RawQuery = values.Encode()
	g.client = &http.Client{}
	return g
}

func (g *Gephi) Commit() (err error) {
	payload := make(map[string]map[string]Attributes)
	for _, op := range opsOrder {
		payload[op] = make(map[string]Attributes)
	}
	for _, op := range opsOrder {
		for action, e := range g.Entities {
			if action.Type == op {
				payload[action.Type][action.ID] = e
			}
		}
	}
	b, err := json.Marshal(payload)
	err = g.post(b)
	if err != nil {
		log.Error(string(b))
	}
	g.Entities = make(map[Action]Attributes)
	return
}

func (g *Gephi) AddNode(n Node) {
	g.Entities[Action{Type: "an", ID: n.ID}] = n.GetAttributes()
}

func (g *Gephi) ChangeNode(n Node) {
	g.Entities[Action{Type: "cn", ID: n.ID}] = n.GetAttributes()
}

func (g *Gephi) DeleteNode(n Node) {
	g.Entities[Action{Type: "dn", ID: n.ID}] = n.GetAttributes()
}

func (g *Gephi) AddEdge(e Edge) {
	g.Entities[Action{Type: "ae", ID: e.ID}] = e.GetAttributes()
}

func (g *Gephi) ChangeEdge(e Edge) {
	g.Entities[Action{Type: "ce", ID: e.ID}] = e.GetAttributes()
}

func (g *Gephi) DeleteEdge(e Edge) {
	g.Entities[Action{Type: "de", ID: e.ID}] = e.GetAttributes()
}

func (g *Gephi) post(payload []byte) (err error) {
	req, err := http.NewRequest("POST", g.URL.String(), bytes.NewBuffer(payload))
	log.Debugf("POST - %s", g.URL.String())
	return g.do(req, g.URL)
}

func (g *Gephi) do(req *http.Request, u url.URL) (err error) {
	req.Header.Set("User-Agent", fmt.Sprintf("Golang Gephi Streaming Clientv%s", version))
	req.Header.Set("Content-Type", "application/json")
	resp, err := g.client.Do(req)
	if err != nil {
		log.Errorf("error when querying Gephi Streaming API - %v", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Errorf("bad response from Gephi Streaming API - %s - %s - %s - %s", resp.Status, u.String(), string(body))
	}
	return
}
