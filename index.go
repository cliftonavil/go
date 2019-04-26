package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	Name  string
	Price int
}

type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}

func main() {
	var template = template.Must(template.ParseFiles("products.html"))
	userDB := map[string]int{
		"java":   20,
		"php":    12,
		"python": 50,
		"golang": 5,
	}
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/login_data", logindata)
	router.HandleFunc("/home", home)
	router.HandleFunc("/upload", upload)
	router.HandleFunc("/createcookie", createcookie)
	router.HandleFunc("/upload_handler", uploadHandler)
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		userDB := map[string]int{
			"java":   20,
			"php":    12,
			"python": 50,
			"golang": 5,
		}
		name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Print(name)
		fmt.Print(userDB[name])
		fmt.Fprintf(w, "%s age is %d ", name, age)
	})
	router.HandleFunc("/lanaguages/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "%s age is %d ", name, age)
	}).Methods("GET")
	router.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "name.txt")
	})
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		myProducts := Product{"Laptop", 27}
		template.ExecuteTemplate(w, "products.html", myProducts)
	})
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Post Already used !!!")
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}
func logindata(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Method:", r.Method)
	//r.ParseForm()
	//r.Form["username"]
	//r.Form["password"]
	username := r.FormValue("username")
	password := r.FormValue("password")
	//fmt.Println("Username :", username)
	//fmt.Println("Password :", password)
	usernameDB := "clifton"
	passwordDB := "password"
	if username == usernameDB && password == passwordDB {
		fmt.Fprintln(w, "Login Sucess!!")
	} else {
		fmt.Fprintln(w, "Username and password Invalied!!!")
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "upload.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Fprintf(w, "file Header is %v", handle.Header)
	// fmt.Fprintf(w, "file Size is %v", handle.Size)
	// fmt.Fprintf(w, "file Filename is %v", handle.Filename)
	f, err := os.OpenFile("./uploadedfiles/"+handle.Filename, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, "Upload Complete")
}

func createcookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.January)
	fmt.Println(time.Now())
	expire := time.Now().Add(time.Hour * 24 * 365)
	cookie := http.Cookie{Name: "Clifton", Value: "cliftonavil", Expires: expire}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Create cookie")
}

//hii
