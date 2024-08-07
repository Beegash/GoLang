package main

import (
	"fmt"
	"math/rand"
)

var friends = make(map[int][]int)

func addFriend(a, b int) bool {
	if _, existsA := friends[a]; !existsA {
		fmt.Println("Error: First user ID does not exist.")
		return false
	}
	if _, existsB := friends[b]; !existsB {
		fmt.Println("Error: Second user ID does not exist.")
		return false
	}
	friends[a] = append(friends[a], b)
	friends[b] = append(friends[b], a)
	return true
}

func remove(slice []int, frnd int) []int {
	for i, j := range slice {
		if j == frnd {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func delFriend(a, b int) bool {
	if _, existsA := friends[a]; !existsA {
		fmt.Println("Error: First user ID does not exist.")
		return false
	}
	if _, existsB := friends[b]; !existsB {
		fmt.Println("Error: Second user ID does not exist.")
		return false
	}
	friends[a] = remove(friends[a], b)
	friends[b] = remove(friends[b], a)
	return true
}

func createUserId(n int) []int {
	var userIds []int
	for i := 0; i < n; i++ {
		id := rand.Intn(99999999) + 1
		for _, exists := friends[id]; exists; _, exists = friends[id] {
			id = rand.Intn(99999999) + 1
		}
		friends[id] = []int{}
		userIds = append(userIds, id)
	}
	return userIds
}

func mutualFriends(a, b int) ([]int, bool) {
	if _, existsA := friends[a]; !existsA {
		fmt.Println("Error: First user ID does not exist.")
		return nil, false
	}
	if _, existsB := friends[b]; !existsB {
		fmt.Println("HError: Second user ID does not exist.")
		return nil, false
	}
	ortaklar := []int{}
	m := make(map[int]bool)
	for _, friend := range friends[a] {
		m[friend] = true
	}
	for _, friend := range friends[b] {
		if m[friend] {
			ortaklar = append(ortaklar, friend)
		}
	}
	return ortaklar, true
}

func main() {
	var n int
	fmt.Print("How many users do you want to create? ")
	fmt.Scan(&n)

	userIds := createUserId(n)
	fmt.Printf("%d users created\n", len(friends))
	fmt.Println("User IDs:", userIds)

	var choice int
	for {
		fmt.Println("\nSelect An Operation You Want To Perform:")
		fmt.Println("1. Add Friend")
		fmt.Println("2. Remove Friend")
		fmt.Println("3. Show Friends List")
		fmt.Println("4. View Mutual Friends By ID")
		fmt.Println("5. Exit")
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var a, b int
			fmt.Print("Enter the ID of the first user: ")
			fmt.Scan(&a)
			fmt.Print("Enter the ID of the second user: ")
			fmt.Scan(&b)
			if addFriend(a, b) {
				fmt.Println("Friend added.")
			}
		case 2:
			var a, b int
			fmt.Print("Enter the ID of the first user: ")
			fmt.Scan(&a)
			fmt.Print("Enter the ID of the second user: ")
			fmt.Scan(&b)
			if delFriend(a, b) {
				fmt.Println("Friend removed.")
			}
		case 3:
			fmt.Println("Friends list:")
			for id, friends := range friends {
				fmt.Printf("%d: %v\n", id, friends)
			}
		case 4:
			var a, b int
			fmt.Print("Enter the ID of the first user: ")
			fmt.Scan(&a)
			fmt.Print("Enter the ID of the second user: ")
			fmt.Scan(&b)
			if ortaklar, ok := mutualFriends(a, b); ok {
				fmt.Println("Mutual friends:", ortaklar)
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("\nInvalid choice, please try again.\n")
		}
	}
}
