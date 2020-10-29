package handlers

type User struct {
	ID   int
	Room string
	Name string
}

type Users []*User

var users = []*User{}

// JoinRoom function to join a room
func JoinRoom(u *User) *User {
	u.Name = "Anonymous"
	if u.ID > 0 {
		u.Name = u.Room
	}
	user := &User{ID: u.ID, Room: u.Room, Name: u.Name}
	users = append(users, user)
	return user
}
