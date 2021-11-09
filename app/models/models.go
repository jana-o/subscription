package models

import (
	"github.com/google/uuid"
	"time"
)

// User represents data about a User.
type User struct {
	ID        uuid.UUID  `json:"id"`
	Name  	  string   	`json:"name"`
	Email 	  string 	 `json:"email"`
	Active 	  bool 	 	`json:"active"`
}


// Product represents data about a Product.
type Product struct {
	ID       	 string  `json:"id"`
	Name  	 	 string  `json:"name"`
	Duration 	 int32   `json:"duration"` // months or days?
	Price  		 float32   `json:"price"`
	Description  string   `json:"description"`
}

type ProductList struct {
	Products []Product `json:"productss"`
}

// Products slice to seed Product data
var Products = []Product{
	{ID: "1", Name: "Annual payment", Duration: 12, Price: 6.99, Description: "83.99 annual payment"},
	{ID: "2", Name: "Semi-annual payment", Duration: 6, Price: 9.99, Description: "59.99 semi-annual payment"},
	{ID: "3", Name: "Quarterly payment", Duration: 3, Price: 12.99, Description: "38.99 quarterly payment"},
}

// Contract represents data about a Subscription Contract.
type Contract struct {
	ID       	  uuid.UUID  `json:"id"`
	UserID  	  uuid.UUID  `json:"user_id"`
	ProductID  	  string 	 `json:"product_id"`
	TrialStart	  time.Time  `json:"trial_start"`
	TrialEnd	  time.Time  `json:"trial_end"`
	StartDate	  time.Time  `json:"start_date"`
	EndDate	  	  time.Time  `json:"end_date"`
	Discount 	  float64  	 `json:"discount"`
	Tax 	  	  float64    `json:"tax"`
	Active 	  	  bool  	 `json:"active"`
}

// Contracts slice to seed Product data
var Contracts = []Contract{
	{ID: uuid.MustParse("3f9c34c7-7cfe-460c-a145-39e79a09087d"),
		UserID: uuid.MustParse("3f9c34c7-7cfe-460c-a145-39e79a09087d"),
		ProductID: "3",
		TrialStart: time.Now(),
		TrialEnd: time.Now().AddDate(0,0,7),
		StartDate: time.Now().AddDate(0,0,8),
		EndDate: time.Now().AddDate(0,3,0),
		Tax: 0.19,
		Active: true,
	},
}
