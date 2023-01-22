package main

type book struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Currency string `json:"currency"`
	Author   string `json:"author"`
}

var books = []book{
	{
		ID:       1,
		Name:     "The Alchemist",
		Price:    2000,
		Currency: "$",
		Author:   "Paulo Coelho",
	},
	{
		ID:       2,
		Name:     "1984",
		Price:    1500,
		Currency: "$",
		Author:   "George Orwell",
	},
	{
		ID:       3,
		Name:     "The Catcher in the Rye",
		Price:    1000,
		Currency: "$",
		Author:   "J.D. Salinger",
	},
	{
		ID:       4,
		Name:     "To Kill a Mockingbird",
		Price:    900,
		Currency: "$",
		Author:   "Harper Lee",
	},
	{
		ID:       5,
		Name:     "The Lord of the Rings",
		Price:    3000,
		Currency: "$",
		Author:   "J.R.R. Tolkien",
	},
}
