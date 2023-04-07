package main

//import (
//	"fmt"
//	"sort"
//)
//
//type User struct {
//	Login    string
//	Password string
//}
//
//type Item struct {
//	Name       string
//	Price      float64
//	Rating     float64
//	NumOfRates int
//}
//
//var users = []User{{"n_serikkanov@kbtu.kz", "t$hZw!Kz"},
//	{"moldaspan@mail.ru", "ers123"},
//	{"madimadiyarov@gmail.com", "madi123"}}
//var items = []Item{{"Snikers", 36.990, 4.7, 1},
//	{"Kitkat", 45.990, 4.9, 2},
//	{"Bounty", 30.000, 4.5, 2},
//	{"Albini", 20.990, 4.0, 1}}
//
//func RegisterUser(login, password string) {
//	for i := 0; i < len(users); i++ {
//		if users[i].Login == login {
//			fmt.Println("User already exists")
//			return
//		}
//	}
//	users = append(users, User{Login: login, Password: password})
//}
//
//func AuthorizeUser(login, password string) *User {
//	for i := 0; i < len(users); i++ {
//		if users[i].Login == login && users[i].Password == password {
//			return &users[i]
//		}
//	}
//	return nil
//}
//
//func Search(name string) *Item {
//	for i := 0; i < len(items); i++ {
//		if items[i].Name == name {
//			return &items[i]
//		}
//	}
//	return nil
//}
//
//func FilterByPrice() {
//	sort.Slice(items, func(i, j int) bool {
//		return items[i].Price < items[j].Price
//	})
//}
//func FilterByRating() {
//	sort.Slice(items, func(i, j int) bool {
//		return items[i].Rating < items[j].Rating
//	})
//}
//
//func (u *User) GiveRate(name string, rate float64) {
//	for i := 0; i < len(items); i++ {
//		if items[i].Name == name {
//			items[i].Rating += rate
//			items[i].NumOfRates++
//		}
//	}
//}
//
//func main() {
//	for true {
//		fmt.Println("1.Registration\n2.Authorization")
//		var choise int
//		fmt.Scan(&choise)
//
//		fmt.Println("Enter login: ")
//		var login string
//		fmt.Scan(&login)
//
//		fmt.Println("Enter password: ")
//		var password string
//		fmt.Scan(&password)
//		var u *User = nil
//		if choise == 1 {
//			RegisterUser(login, password)
//		} else {
//			u = AuthorizeUser(login, password)
//			if u == nil {
//				fmt.Println("Wrong login p=or password")
//			} else {
//				fmt.Println("Authorised.")
//				for true {
//					fmt.Println("3. Searching")
//					var name string
//					fmt.Scan(&name)
//					var it *Item = Search(name)
//					if it == nil {
//						fmt.Println("Not found")
//					} else {
//						fmt.Println(it)
//					}
//					fmt.Println("4.Filtering by price")
//					FilterByPrice()
//					for i := 0; i < len(items); i++ {
//						fmt.Println(items[i])
//
