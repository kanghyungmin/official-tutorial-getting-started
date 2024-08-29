// package main

// import (
// 	"example/fedex"
// 	koreapost "example/koreaPost"
// 	"fmt"
// )

// type Stringer interface {
// 	String() string
// }

// type Student struct {
// 	Name string
// 	Age  int
// }

// func (s Student) String() string {
// 	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
// }

// type Sender interface {
// 	Send(parcel string)
// }

// func SendBook(name string, sender Sender) {
// 	sender.Send(name)
// }

// func main() {
// 	koreapost := koreapost.PostSender{}
// 	SendBook("어린왕상", koreapost)

// 	fedexSender := fedex.FedexSender{}
// 	SendBook("어린왕상", fedexSender)

// }

package main

import "fmt"

type Stringer interface {
	// String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "student"
}

type Actor struct {
}

func (s *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {

	if student, ok := stringer.(*Student); ok {
		fmt.Println(student)
	}
	fmt.Printf("%T", stringer)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
}
