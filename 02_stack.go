package main

import "fmt"

//todo: доделать (домашка)

func main() {
	st := stack{values:[]int{1,2,3}}

	for i := 0; i < 10; i++ {
		st.push(i)
	}

	fmt.Println(st)

	//delete element with value
	st.pop(3)

	fmt.Println(st)
	//retrive top element
	st.peek()

	fmt.Println(st)
}

type stack struct {
	values []int
}

// Это МЕТОД! Не функция
func (amt *stack) push(value int) {
	amt.values = append(amt.values, value)
}

func (st *stack) pop(value int)  {
	//st.values = st.values[0:len(st.values)- 1]

	st_new := stack{}

	for i := 0; i < len(st.values); i++{


		if(st.values[i] != value){

			st_new.values = append(st_new.values, st.values[i])
		}

	}
	st.values = st_new.values

}

func (st *stack) peek() {

	st.values= st.values[len(st.values)-1:len(st.values)]
}