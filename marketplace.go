package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type User struct {
	Email    string
	Password string
}

type Item struct {
	ID    int
	Name  string
	Price float64
}

type Message struct {
	Sender   string
	Receiver string
	Content  string
}

type App struct {
	Users    []User
	Items    []Item
	Messages []Message
}

var app App
var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	// Initialize the Stripe API key
	stripe.Key = "sk_test_1234567890"

	// Create some users and items
	user1 := User{Email: "user1@example.com", Password: "password1"}
	user2 := User{Email: "user2@example.com", Password: "password2"}
	app.Users = append(app.Users, user1, user2)

	item1 := Item{ID: 1, Name: "Item 1", Price: 10.0}
	item2 := Item{ID: 2, Name: "Item 2", Price: 20.0}
	app.Items = append(app.Items, item1, item2)

	// Create the router
	r := mux.NewRouter()

	// Set up the routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/signup", signupHandler).Methods("GET")
	r.HandleFunc("/signup", signupPostHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/items", itemsHandler).Methods("GET")
	r.HandleFunc("/item/{id}", itemHandler).Methods("GET")
	r.HandleFunc("/item/{id}/buy", buyHandler).Methods("GET")
	r.HandleFunc("/messages", messagesHandler).Methods("GET")
	r.HandleFunc("/messages/new", newMessageHandler).Methods("GET")
	r.HandleFunc("/messages/new", newMessagePostHandler).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the marketplace app!")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sign up page")
}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := User{Email: email, Password: password}
	app.Users = append(app.Users, user)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Log in page")
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	for _, user := range app.Users {
		if user.Email == email && user.Password == password {
			session

      		// Set a session cookie
		session, _ := store.Get(r, "session")
		session.Values["email"] = email
		session.Save(r, w)

		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}
}
http.Error(w, "Invalid login credentials", http.StatusUnauthorized)

  }

func logoutHandler(w http.ResponseWriter, r *http.Request) {
// Delete the session cookie
session, _ := store.Get(r, "session")
session.Options.MaxAge = -1
session.Save(r, w)
  
  http.Redirect(w, r, "/", http.StatusSeeOther)

  
  }

func itemsHandler(w http.ResponseWriter, r *http.Request) {
// Get the user's email from the session cookie
session, _ := store.Get(r, "session")
email, ok := session.Values["email"].(string)
if !ok {
http.Redirect(w, r, "/login", http.StatusSeeOther)
return
}
  
  // Get the user's items
var userItems []Item
for _, item := range app.Items {
	if item.Seller == email {
		userItems = append(userItems, item)
	}
}

// Render the template with the user's items
data := struct {
	Items []Item
}{
	Items: userItems,
}
renderTemplate(w, "items.html", data)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
// Get the item ID from the URL
vars := mux.Vars(r)
id, _ := strconv.Atoi(vars["id"])
  
  // Get the item with the specified ID
var item Item
for _, i := range app.Items {
	if i.ID == id {
		item = i
		break
	}
}

// Render the template with the item
data := struct {
	Item Item
}{
	Item: item,
}
renderTemplate(w, "item.html", data)

  }

func buyHandler(w http.ResponseWriter, r *http.Request) {
// Get the item ID from the URL
vars := mux.Vars(r)
id, _ := strconv.Atoi(vars["id"])
  
  // Get the item with the specified ID
var item Item
for _, i := range app.Items {
	if i.ID == id {
		item = i
		break
	}
}

// Get the user's email from the session cookie
session, _ := store.Get(r, "session")
email, ok := session.Values["email"].(string)
if !ok {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}

// Create a Stripe charge for the item's price
params := &stripe.ChargeParams{
	Amount:   stripe.Int64(int64(item.Price * 100)),
	Currency: stripe.String(string(stripe.CurrencyUSD)),
	Desc:     stripe.String(fmt.Sprintf("Payment for %s", item.Name)),
}
params.SetSource(r.FormValue("stripeToken"))
ch, err := charge.New(params)
if err != nil {
	http.Error(w, "Payment failed", http.StatusInternalServerError)
	return
}

// Update the item's seller and buyer
for i, it := range app.Items {
	if it.ID == item.ID {
		app.Items[i].Seller = email
		app.Items[i].Buyer = ch.BillingDetails.Email
		break
	}
}

http.Redirect(w, r, "/item/"+vars["id"], http.StatusSeeOther)

  }

func messagesHandler(w http.ResponseWriter, r *http.Request) {
// Get the user's email from the session cookie
session, _ := store.Get
  
  
  email, ok := session.Values["email"].(string)
if !ok {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}

// Get the user's messages
var userMessages []Message
for _, msg := range app.Messages {
	if msg.To == email {
		userMessages = append(userMessages, msg)
	}
}

// Render the template with the user's messages
data := struct {
	Messages []Message
}{
	Messages: userMessages,
}
renderTemplate(w, "messages.html", data)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
// Get the message ID from the URL
vars := mux.Vars(r)
id, _ := strconv.Atoi(vars["id"])
  
  
  // Get the message with the specified ID
var msg Message
for _, m := range app.Messages {
	if m.ID == id {
		msg = m
		break
	}
}

// Render the template with the message
data := struct {
	Message Message
}{
	Message: msg,
}
renderTemplate(w, "message.html", data)

  }

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
// Get the message ID from the URL
vars := mux.Vars(r)
id, _ := strconv.Atoi(vars["id"])
  
  
  // Get the item with the specified ID
var item Item
for _, i := range app.Items {
	if i.ID == id {
		item = i
		break
	}
}

// Get the user's email from the session cookie
session, _ := store.Get(r, "session")
email, ok := session.Values["email"].(string)
if !ok {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}

// Get the recipient's email from the form data
to := r.FormValue("to")

// Create the message
msg := Message{
	ID:      len(app.Messages) + 1,
	From:    email,
	To:      to,
	Subject: r.FormValue("subject"),
	Body:    r.FormValue("body"),
}

// Add the message to the app's messages
app.Messages = append(app.Messages, msg)

// Redirect to the item page
http.Redirect(w, r, "/item/"+vars["id"], http.StatusSeeOther)

  }

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
tmpl = filepath.Join("templates", tmpl)
t, err := template.ParseFiles(tmpl)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
  
  
  err = t.Execute(w, data)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
}

func sellHandler(w http.ResponseWriter, r *http.Request) {
// Get the user's email from the session cookie
session, _ := store.Get(r, "session")
email, ok := session.Values["email"].(string)
if !ok {
http.Redirect(w, r, "/login", http.StatusSeeOther)
return
}
  
  // Parse the form data
err := r.ParseForm()
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

// Get the form values
name := r.FormValue("name")
description := r.FormValue("description")
price, err := strconv.Atoi(r.FormValue("price"))
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

// Create a new item
item := Item{
	ID:          len(app.Items) + 1,
	Name:        name,
	Description: description,
	Price:       price,
	Seller:      email,
}

// Add the item to the app's items
app.Items = append(app.Items, item)

// Redirect to the user's items page
http.Redirect(w, r, "/items", http.StatusSeeOther)

  }

func main() {
// Initialize the app
app = App{
Items: make([]Item, 0),
Messages: make([]Message, 0),
}
  
  // Register the routes
r := mux.NewRouter()
r.HandleFunc("/", indexHandler).Methods("GET")
r.HandleFunc("/login", loginHandler).Methods("GET", "POST")
r.HandleFunc("/logout", logoutHandler).Methods("GET")
r.HandleFunc("/register", registerHandler).Methods("GET", "POST")
r.HandleFunc("/items", itemsHandler).Methods("GET")
r.HandleFunc("/item/{id}", itemHandler).Methods("GET")
r.HandleFunc("/item/{id}/message", messageHandler).Methods("GET")
r.HandleFunc("/item/{id}/message/send", sendMessageHandler).Methods("POST")
r.HandleFunc("/sell", sellHandler).Methods("GET", "POST")

// Serve the app
http.ListenAndServe(":8080", r)

}
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
      
      
      
      
      
      
      
      
