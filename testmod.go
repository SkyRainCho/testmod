package testmod

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello,", name)
	fmt.Println("GoodBye,", name)
	switch lang {
	case "en":
		fmt.Println("English")
	case "zh":
		fmt.Println("Chinese")
	case "fr":
		fmt.Println("Francis")
	default:
		fmt.Println("Unknows")
	}
}
