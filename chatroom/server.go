package chatroom

type Server struct {
	server_id   int
	name        string
	client_list []Client
}

var crIdCounter = 0

func NewChatRoom(name string) Server {
	return Server{
		server_id: generateChatRoomId(),
		name:      name,
	}
}

func generateChatRoomId() int {
	crIdCounter++
	return crIdCounter
}
