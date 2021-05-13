package robotname

import (
	"math/rand"
	"strconv"
)

type Robot struct {
	name string
}

// var existingRobotNames map[string]int

func (r *Robot) Name() (string, error) {
	//Initialize if name is empty
	if r.name == "" {
		//Below commented block of code checks if the name is already used.
		// var n string
		// for {
		// 	n = RandomChar() + RandomChar() + strconv.Itoa(rand.Int())[:3]
		// 	_, chk := existingRobotNames[n]
		// 	if !chk {
		// 		break
		// 	}
		// }
		// r.name = n
		r.name = RandomChar() + RandomChar() + strconv.Itoa(rand.Int())[:3]
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func RandomChar() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, 1)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
