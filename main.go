package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Anik.I am a learner.I am trying to software enginner")
}

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productlist []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "please give me a get request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productlist)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "please give me a post request", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("error")
		http.Error(w, "give me correct json ", 400)
		return
	}
	newProduct.Id = len(productlist) + 1
	w.WriteHeader(201)
	productlist = append(productlist, newProduct)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)
	fmt.Println("server running on : 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting on : ", err)
	}

}
func init() {
	prd1 := Product{
		Id:          1,
		Title:       "Orange",
		Description: "I love orange fruit",
		Price:       624,
		ImgUrl:      "https://cdn.vectorstock.com/i/500p/89/96/ripe-orange-fruit-vector-58996.jpg",
	}
	prd2 := Product{
		Id:          2,
		Title:       "Banana",
		Description: "I love Banana very much",
		Price:       124,
		ImgUrl:      "https://media.cnn.com/api/v1/images/stellar/prod/120604032828-fresh-ripe-bananas.jpg?q=w_3590,h_2774,x_0,y_0,c_fill",
	}
	prd3 := Product{
		Id:          3,
		Title:       "Grapes",
		Description: "I love Grapes",
		Price:       550,
		ImgUrl:      "https://images.everydayhealth.com/images/diet-nutrition/what-are-grapes-nutrition-health-benefits-risks-alt-1440x810.jpg?sfvrsn=f9f2dae1_3",
	}

	productlist = append(productlist, prd1)
	productlist = append(productlist, prd2)
	productlist = append(productlist, prd3)

}
