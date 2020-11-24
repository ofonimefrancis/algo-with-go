package Stack

//Stack Stack Data structure
type Stack struct {
	data []int
}

//New Returns a new
func New() *Stack {
	return new(Stack)
}

//Push Adds a new integer into the stack
func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

//Pop removes the last element in the array
func (s *Stack) Pop() int {
	if len(s.data) <= 0 {
		return -1
	}

	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element
}

// func main() {
// 	s := New()

// 	s.Push(2)
// 	s.Push(3)
// 	s.Push(4)

// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

// }
