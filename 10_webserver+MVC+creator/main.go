package main

import "creator/webservergo"

//type ControllersSet interface {
//	Controllers() []Controller
//}
//
//type Controller interface {
//	Path() string
//	Name() string
//	DoAction(string)
//}

func main() {
	webservergo.StartServer()
}
