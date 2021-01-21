package main

import "fmt"

func calculateBill(price int, num int) int {
	var totalPrice = price * num
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter

}
func main() {
	price, num := 90, 6
	totalPrice := calculateBill(price, num)
	fmt.Println("Total Price is ", totalPrice)
	area, perimeter := rectProps(10.8, 10.1)
	fmt.Println("Area is ", area)
	fmt.Println("Perimeter is ", perimeter)
}
