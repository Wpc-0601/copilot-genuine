package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

type Level = int8
type String string
type Array []int
type method func(str string) string

const (
	age            Level  = 0
	bug            string = "initialization"
	interpretation String = "String"
	crucial        String = "YCM"
)

func TestMethod(test *testing.T) {
	intValue := rand.Intn(100)
	fmt.Println("origin int value was:", intValue)
	stringValue := strconv.Itoa(intValue)
	fmt.Println("present string value is: ", stringValue)
	quizMethod(func(str string) string {
		return str + "about it"
	})
}

func quizMethod(m method) {
	result := m("How do you feel")
	fmt.Println(result)
}

func TestSlice(test *testing.T) {
	array := make(Array, 5)
	plus := append(array, 1, 2, 3)
	plus.quizSlice()
}

func (array Array) quizSlice() {
	fmt.Print("the slice is :", array[5:8])
	i := len(array)
	fmt.Println(i)
}
