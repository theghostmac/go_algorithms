package main

import "fmt"

func TowersOfHanoi(n int) {
	ToHAlgorithm(n, "Source", "Destination", "Auxiliary")
}

func ToHAlgorithm(n int, fromTower, toTower, auxTower string) {
	// if n = 1 disk, move from source to destination and return
	if n == 1 {
		fmt.Println("Move disk", n, "from ", fromTower, " tower to ", toTower, " tower.")
		return
	}
	// move the top n-1 disks from source tower to auxiliary tower using the destination as decoy
	ToHAlgorithm(n-1, fromTower, auxTower, toTower)

	// move the remaining disks from source to destination
	fmt.Println("Move disk", n, " from ", fromTower, " to ", toTower, ".")

	//move n-1 disks from auxiliary tower to destination tower using source tower as decoy.
	ToHAlgorithm(n-1, auxTower, toTower, fromTower)
}

func main() {
	TowersOfHanoi(3)
}