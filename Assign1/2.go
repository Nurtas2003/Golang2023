package main

import (
	"fmt"
	"os"
	"sort"
)

type User struct {
	Username string
	Password string
}

type Item struct {
	Name       string
	Rating     float64
	numOfRates int
}

var Users = []User{{"Serikkanov", "Nurtas"}, {"Madiyarov", "Madi"}}
var Items = []Item{{"History book", 5, 4}, {"Snikers", 3.2, 2}}

func RegisterUser(username, password string) {
	exists := false
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username {
			exists = true
			break
		}
	}
	if exists {
		fmt.Println("user already exists")
	} else {
		Users = append(Users, User{Username: username, Password: password})
	}
}

func AuthorizeUser(username, password string) *User {
	//var u *User = nil
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username && Users[i].Password == password {
			return &Users[i]
		}
	}
	return nil
}

func AddItem(name string, rating float64) {
	Items = append(Items, Item{Name: name, Rating: rating})
}

func filteringByRating() {
	sort.Slice(Items, func(i, j int) bool {
		return Items[i].Rating < Items[j].Rating
	})
}

func filteringByName() {
	sort.Slice(Items, func(i, j int) bool {
		return Items[i].Name < Items[j].Name
	})
}

func GiveItemRating(name string, rating float64) {
	for i := 0; i < len(Items); i++ {
		if Items[i].Name == name {
			Items[i].Rating += rating
			Items[i].numOfRates++
		}
	}
}

func SearchItem(name string) *Item {
	for i := 0; i < len(Items); i++ {
		if Items[i].Name == name {
			return &Items[i]
		}
	}
	return nil
}

func bye() {
	fmt.Print("---------------------------------------------------------\n\n")
	fmt.Println("                    GOODBYE")
}

func main() {
	for true {
		fmt.Println("1.Registration\n2.Authorization")
		var choise int
		fmt.Scan(&choise)

		fmt.Println("Enter login: ")
		var login string
		fmt.Scan(&login)

		fmt.Println("Enter password: ")
		var password string
		fmt.Scan(&password)
		var u *User = nil
		if choise == 1 {
			RegisterUser(login, password)
		} else {
			u = AuthorizeUser(login, password)
			if u == nil {
				fmt.Println("Wrong login or password")
			} else {
				for true {
					fmt.Println("1.Searching\n2.Give rating\n3.Filtering by Name\n4.Filtering by rating\n5.AddItem \n6.Bye")
					fmt.Scan(&choise)
					if choise == 1 {
						fmt.Println("Enter name: ")
						var name string
						fmt.Scan(&name)
						SearchItem(name)
						var it *Item = SearchItem(name)
						if it == nil {
							fmt.Println("Not found")
						} else {
							fmt.Println(it)
						}

					} else if choise == 2 {
						fmt.Println("Enter name of item")
						var name string
						fmt.Scan(&name)
						fmt.Println("Enter rating you want to add")
						var rating float64
						fmt.Scan(&rating)
						GiveItemRating(name, rating)

					} else if choise == 3 {
						filteringByName()
						for i := 0; i < len(Items); i++ {
							fmt.Println(Items[i])
						}
					} else if choise == 4 {
						filteringByRating()
						for i := 0; i < len(Items); i++ {
							fmt.Println(Items[i])
						}
					} else if choise == 6 {
						bye()
						os.Exit(1)

					} else if choise == 5 {
						fmt.Println("Enter name of item")
						var name string
						fmt.Scan(&name)
						fmt.Println("Enter rating of item")
						var rating float64
						fmt.Scan(&rating)
						AddItem(name, rating)
					}
				}
			}
		}
	}
}
