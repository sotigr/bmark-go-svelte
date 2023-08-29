package emitter

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Lobby struct {
	mu          sync.Mutex
	Name        string
	Key         *string
	Connections []ConnectionSession
}

type ConnectionSession struct {
	id  string
	ctx *gin.Context
}

type EventEmitter interface {
	Init(path string)
	Send(lobby string, event_name string, data string) bool
	CreateLobby(name string, key *string)
	Close()
}

type HTTPEmitter struct {
	path    string
	Lobbies *[]*Lobby
	stop    bool
}

func (this HTTPEmitter) Init(server *gin.Engine, path string) *HTTPEmitter {
	this.Lobbies = &[]*Lobby{}
	this.path = path
	this.stop = false
	server.GET(path, func(c *gin.Context) {

		id := uuid.NewString()
		lobbiesStr := c.DefaultQuery("lobbies", "")
		if lobbiesStr == "" {
			c.JSON(int(http.StatusBadRequest), gin.H{
				"message": "no_lobbies",
			})
			return
		}
		if this.stop {
			c.JSON(int(http.StatusServiceUnavailable), gin.H{
				"message": "closed",
			})
			return
		}

		lobbies := strings.Split(lobbiesStr, ",")
		if len(lobbies) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no_lobbies",
			})
			return
		}
		subbed_lobbies := []*Lobby{}

		for _, lobbyStr := range lobbies {
			values := strings.Split(lobbyStr, ":")
			lobbyName := values[0]
			var lobbyKey string
			if len(values) > 1 {
				lobbyKey = values[1]
			}
			fmt.Println(this.Lobbies)
			cur_lobby := this.get_lobby_by_name(lobbyName)
			subbed_lobbies = append(subbed_lobbies, cur_lobby)
			if cur_lobby == nil {
				c.JSON(int(http.StatusNoContent), gin.H{
					"message": "lobby_not_found",
					"lobby":   lobbyName,
				})
				return
			}

			if cur_lobby.Key != nil {
				if *cur_lobby.Key != lobbyKey {
					c.JSON(int(http.StatusForbidden), gin.H{
						"message": "forbidden",
						"lobby":   lobbyName,
					})
					return
				}
			}
			cur_lobby.mu.Lock()
			c.Writer.Header().Set("Content-Type", "text/event-stream")
			c.Writer.Header().Set("Cache-Control", "no-cache")
			c.Writer.Flush()
			cur_lobby.Connections = append(cur_lobby.Connections, ConnectionSession{ctx: c, id: id})
			cur_lobby.mu.Unlock()
		}

		fmt.Println("new connection")

		for c.Request.Context().Err() == nil && this.stop == false {
		}

		this._send_to_connection(c, "close", "ok", false)

		for _, l := range subbed_lobbies {

			l.mu.Lock()
			clean_connections := []ConnectionSession{}
			for _, c := range l.Connections {
				if c.id != id {
					clean_connections = append(clean_connections, c)
				}
			}
			l.Connections = clean_connections
			l.mu.Unlock()
		}

		fmt.Println("connection closed")
	})
	return &this
}

func (this HTTPEmitter) get_lobby_by_name(name string) *Lobby {
	for _, lobby := range *this.Lobbies {
		if lobby.Name == name {
			return lobby
		}
	}
	return nil
}

func (this HTTPEmitter) _send_to_connection(c *gin.Context, event_name string, data string, flush bool) {
	c.Writer.Write([]byte(fmt.Sprintf("id: %s\n", uuid.NewString())))
	c.Writer.Write([]byte("event: " + event_name + "\n"))
	c.Writer.Write([]byte(fmt.Sprintf("data: %s\n", data)))
	c.Writer.Write([]byte("\n"))
	if flush {
		c.Writer.Flush()
	}
}

func (this HTTPEmitter) CreateLobby(name string, key *string) {

	*this.Lobbies = append(*this.Lobbies, &Lobby{Name: name, Key: key, Connections: []ConnectionSession{}})

}

func (this HTTPEmitter) Send(lobbyName string, event_name string, data string) {
	lobby := this.get_lobby_by_name(lobbyName)
	if lobby != nil {
		for _, c := range lobby.Connections {
			this._send_to_connection(c.ctx, lobbyName+":"+event_name, data, true)
		}
	}
}

func CloseEmitter(e *HTTPEmitter) {
	e.stop = true
}
