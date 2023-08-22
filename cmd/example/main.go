package main

func main() {
	mt := newMTClient()
	mt.GetUsers(ReadFile("who.txt", "\n")).SendMessage([]string{"asd"})
}
