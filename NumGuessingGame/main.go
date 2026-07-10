
package main

import ("fmt"
       "math/rand")

func main(){
	fmt.Println("Number Guessing Game!!!");
	var randomNum int = rand.Intn(100) + 1;
	var userInput int = 0;
	var count int = 0;
	for userInput != randomNum{
	  fmt.Printf("Enter the number: ");
		fmt.Scanln(&userInput);
		count++;
		if userInput > randomNum{
			fmt.Println("Too high!");
		} else if userInput < randomNum {
			fmt.Println("Too low!");
		} else{
			fmt.Println("You win!!!");
			fmt.Printf("Took you %d times\n", count);
		}

	}


}
