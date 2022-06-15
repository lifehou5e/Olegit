package main

import "fmt"

const ourBank float32 = 23 //it's const because we won't change this

func main() {
	var applePrice, pearPrice float32 = 5.99, 7 //i choose float32 for both cuz it's easier to operate
	//how much money do we need to buy 9 apples and 8 pears?
	buyingSomeFruits := 9*applePrice + 8*pearPrice
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблкук та 9 груш?", buyingSomeFruits, "UAH")
	//how much pears we can buy?
	pearsMax := ourBank / pearPrice
	fmt.Println("Скільки груш ми можемо купити?", int(pearsMax), "шт.")
	//how much apples we can buy?
	applesMax := ourBank / applePrice
	fmt.Println("Скільки яблук ми можемо купити?", int(applesMax), "шт.")
	//Can we buy 2 apples and 2 pears?
	conditionToBuyFruits := (2*applePrice + 2*pearPrice) > ourBank
	fmt.Println("Чи можемо купити 2 груші та 2 яблука?", conditionToBuyFruits)
}
