package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "golang"
	password = "t$hZw!Kz"
)

var (
	db  *sql.DB
	err error
)

type User struct {
	ID       uint
	Name     string
	Surname  string
	Username string
	Password string
}

type Item struct {
	ID     uint
	Name   string
	Price  float64
	Rating float64
}

func home_page(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("home_page.html")
	temp.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("index.html")
	temp.Execute(w, nil)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()
	if r.Method == "POST" {

		name := r.FormValue("name")
		surname := r.FormValue("surname")
		Username := r.FormValue("username")
		Password := r.FormValue("password")

		_, err = db.Exec(`INSERT INTO users (name, surname, username, password) VALUES ($1, $2, $3, $4)`, name, surname, Username, Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User %s created successfully!", Username, Password)
	} else {
		http.ServeFile(w, r, "register.html")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, nil)
		return
	}

	Username := r.FormValue("username")
	Password := r.FormValue("password")

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name=$1 and password=$2", Username, Password).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, "Invalid username or password")
		return
	}
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT name, price, rating FROM items `)
	CheckError(err)

	defer rows.Close()
	var data []Item

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.Name, &i.Price, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, " <h2> All items in store </h2><tr><th>Name</th><th>Price</th><th>Raiting</th></tr>")
	for _, i := range data {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%v</td><td>%v</td></tr>", i.Name, i.Price, i.Rating)
	}
	fmt.Fprintf(w, "</table>")
}

func filterItems(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT name, price, rating FROM items order by price`)
	CheckError(err)

	defer rows.Close()
	var data []Item

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.Name, &i.Price, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, " <h2> Filtering by price </h2> <tr><th>Name</th><th>Price</th><th>Raiting</th></tr>")
	for _, i := range data {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%v</td><td>%v</td></tr>", i.Name, i.Price, i.Rating)
	}
	fmt.Fprintf(w, "</table>")
}

var items []Item

func searchByName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GETgo run" {
		tmpl, err := template.New("search").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<meta charset="UTF-8">
				<title>Search for Item</title>
			</head>
			<body>
				<h1>Search for Item</h1>
				<form method="POST">
					<label>Item Name:</label>
					<input type="text" name="name" required>
					<button type="submit">Search</button>
				</form>
			</body>
			</html>
		`)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if r.Method == "POST" {
		itemName := r.FormValue("name")

		db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = sql.Open("postgres", db_connection)
		CheckError(err)
		defer db.Close()

		stmt, err := db.Prepare("SELECT * FROM items WHERE Name LIKE $1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query("%" + itemName + "%")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var item Item
			err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Rating)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}
	}
	tmpl, err := template.ParseFiles("search.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, items)
	CheckError(err)
}

func rateItemHandler(w http.ResponseWriter, r *http.Request) {

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	if r.Method == "POST" {
		id := r.FormValue("id")
		newValue := r.FormValue("new_value")
		_, err = db.Exec(`UPDATE items SET rating = (rating + $1)/2 WHERE id = $2`, newValue, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.ServeFile(w, r, "rate.html")
	}
}

func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/register", registrationHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/show", showHandler)
	http.HandleFunc("/filter", filterItems)
	http.HandleFunc("/search", searchByName)
	http.HandleFunc("/rate", rateItemHandler)
	http.ListenAndServe(":8080", nil)
}
