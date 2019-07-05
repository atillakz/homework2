package main

import "fmt"

type deque struct {
	values []interface{}
}

func main() {

	dq := deque{}

	//fmt.Println(dq)

	dq.push_back(2)

	//fmt.Println(dq)

	dq.push_back("rusya")

	//fmt.Println(dq)

	//top element
	//dq.pop_front()

	//last element
	/*dq.pop_back()
	fmt.Println(dq)
	//last element
	/*dq.pop_back()
	fmt.Println(dq)

	*/

	dq.push_front("gogoog")

	dq.push_back("jack")

	dq.push_front([]int{1, 23, 4, 5})
}

// Это МЕТОД! Не функция
func (amt *deque) push_back(value interface{}) {
	amt.values = append(amt.values, value)
}

func (dq *deque) pop_front() {

	dq.values = dq.values[len(dq.values)-1 : len(dq.values)]

}

func (dq *deque) pop_back() {

	dq.values = dq.values[0:1]

}

func (dq *deque) push_front(value interface{}) {

	dq.push_back(value)

	for i := len(dq.values) - 1; i > 0; i-- {

		dq.values[i] = dq.values[i-1]
	}

	dq.values[0] = value

	fmt.Println(dq.values)
}

func (dq *deque) peek(option string) {

	if option == "front" {

		fmt.Println(dq.values[len(dq.values)-1 : len(dq.values)])
	}
	if option == "back" {

		fmt.Println(dq.values[0])
	} else {

		panic("No option")
	}
}
