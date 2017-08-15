pakcage main

type Hub struct {
	redis      *redis.Client
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}


type Client struct {
	conn *websocket.Conn
	send chan []byte
}