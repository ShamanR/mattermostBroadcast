package mattermostBroadcast

type Mt struct{}

func (m *Mt) GetPostById(post string) *Post {
	return nil
}

func (m *Mt) GetUsers(names []string) *Users {
	return nil
}

type User struct{}

func (u *User) SendMessage(m string) {
}

type Users []User

func (u *Users) SendMessage(messages []string) {
	for i, v := range *u {
		v.SendMessage(messages[i])
	}
}

type Post struct{}

func (p *Post) GetAuthor() *Users     { return nil }
func (p *Post) GetTreadUsers() *Users { return nil }

func ReadFile(path, separator string) []string {
	return nil
}
func ReadCsv(path string, column int) []string {
	return nil
}

func ReadGoogleTable(path string, column int) []string {
	return nil
}
func main() {
	mt := &Mt{}
	mt.GetPostById("asd").GetTreadUsers().SendMessage([]string{"sda"})
	mt.GetUsers(ReadFile("who.txt", "\n")).SendMessage([]string{"asd"})
	mt.GetUsers(ReadCsv("who.txt", 0)).SendMessage(ReadCsv("who.txt", 1))
	mt.GetUsers(ReadGoogleTable("who.txt", 0)).SendMessage(ReadGoogleTable("who.txt", 1))
}
