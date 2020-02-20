package main

import (
	"fmt"
)

type USB interface {
	Name() string //Return USB's name
	Connect() // the connecting method
}

/*Or use on-board interface:
type USB interface {
	Name() string //Return USB's name
	Connecter // Connecter is a on-board struct, USB has connect() methode
}

type Connecter interface {
	connect()
}
*/

type PhoneConnecter struct{ //realize the interface, so that we can connect to an phone
	name string
}

func (pc PhoneConnecter) Name() string{ //add method for struct 
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB //a is USB type, so it's an interface
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnect.")//seccesfully call , so that we've realized the interface
}