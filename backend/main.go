package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/purchases", GetPurchases)
	http.ListenAndServe(":8080", r)
}

func GetPurchases(w http.ResponseWriter, r *http.Request) {
	render.RenderList(w, r, NewPurchaseListResponse(purchases))
}

func NewPurchaseListResponse(purchases []*Purchase) []render.Renderer {
	list := []render.Renderer{}
	for _, purchase := range purchases {
		list = append(list, NewPurchaseResponse(purchase))
	}
	return list
}

func NewPurchaseResponse(purchase *Purchase) *PurchaseResponse {
	resp := &PurchaseResponse{Purchase: purchase}
	return resp
}

func (rd *PurchaseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PurchaseResponse struct {
	*Purchase
}

type UserRole int

const (
	Employee UserRole = iota
	Manager
	Admin
)

type User struct {
	ID   string
	Name string
	Role UserRole
	Area string
}

type PaymentType int

const (
	Cash PaymentType = iota
	Card
)

type ClientType int

const (
	Personal ClientType = iota
	Managed
)

type Purchase struct {
	ID                  string
	PaymentType         PaymentType
	CreatedAt           time.Time
	Creator             User
	Supplier            string
	VolumeInMilliLiters int
	PricePerMilliLiter  int
	Area                string
	TotalPrice          int
	InvoiceINN          *string
	InvoiceAccount      *string
	ClientType          ClientType
	Approved            bool
	ApprovedAt          time.Time
	Approver            User
}

var purchases = []*Purchase{
	{
		ID:                  "1",
		PaymentType:         Card,
		CreatedAt:           time.Now(),
		Creator:             User{ID: "u1", Name: "John", Role: Manager, Area: "Sales"},
		Supplier:            "Supplier1",
		VolumeInMilliLiters: 1000,
		PricePerMilliLiter:  5,
		Area:                "Sales",
		TotalPrice:          5000,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Managed,
		Approved:            true,
		ApprovedAt:          time.Now(),
		Approver:            User{ID: "u2", Name: "Manager1", Role: Manager, Area: "Sales"},
	},
	{
		ID:                  "2",
		PaymentType:         Cash,
		CreatedAt:           time.Now().Add(-24 * time.Hour),
		Creator:             User{ID: "u3", Name: "Alice", Role: Employee, Area: "Marketing"},
		Supplier:            "Supplier2",
		VolumeInMilliLiters: 800,
		PricePerMilliLiter:  6,
		Area:                "Marketing",
		TotalPrice:          4800,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Personal,
		Approved:            false,
		ApprovedAt:          time.Time{},
		Approver:            User{},
	},
	{
		ID:                  "3",
		PaymentType:         Card,
		CreatedAt:           time.Now().Add(-48 * time.Hour),
		Creator:             User{ID: "u4", Name: "Bob", Role: Admin, Area: "Finance"},
		Supplier:            "Supplier3",
		VolumeInMilliLiters: 1200,
		PricePerMilliLiter:  4,
		Area:                "Finance",
		TotalPrice:          4800,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Managed,
		Approved:            true,
		ApprovedAt:          time.Now(),
		Approver:            User{ID: "u5", Name: "Admin1", Role: Admin, Area: "Finance"},
	},
	{
		ID:                  "4",
		PaymentType:         Cash,
		CreatedAt:           time.Now().Add(-72 * time.Hour),
		Creator:             User{ID: "u6", Name: "Eve", Role: Employee, Area: "HR"},
		Supplier:            "Supplier4",
		VolumeInMilliLiters: 1500,
		PricePerMilliLiter:  3,
		Area:                "HR",
		TotalPrice:          4500,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Personal,
		Approved:            false,
		ApprovedAt:          time.Time{},
		Approver:            User{},
	},
	{
		ID:                  "5",
		PaymentType:         Card,
		CreatedAt:           time.Now().Add(-96 * time.Hour),
		Creator:             User{ID: "u7", Name: "Charlie", Role: Manager, Area: "Sales"},
		Supplier:            "Supplier5",
		VolumeInMilliLiters: 900,
		PricePerMilliLiter:  7,
		Area:                "Sales",
		TotalPrice:          6300,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Managed,
		Approved:            true,
		ApprovedAt:          time.Now(),
		Approver:            User{ID: "u8", Name: "Manager2", Role: Manager, Area: "Sales"},
	},
	{
		ID:                  "6",
		PaymentType:         Cash,
		CreatedAt:           time.Now().Add(-120 * time.Hour),
		Creator:             User{ID: "u9", Name: "David", Role: Employee, Area: "Marketing"},
		Supplier:            "Supplier6",
		VolumeInMilliLiters: 1100,
		PricePerMilliLiter:  5,
		Area:                "Marketing",
		TotalPrice:          5500,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Personal,
		Approved:            false,
		ApprovedAt:          time.Time{},
		Approver:            User{},
	},
	{
		ID:                  "7",
		PaymentType:         Card,
		CreatedAt:           time.Now().Add(-144 * time.Hour),
		Creator:             User{ID: "u10", Name: "Frank", Role: Admin, Area: "Finance"},
		Supplier:            "Supplier7",
		VolumeInMilliLiters: 1300,
		PricePerMilliLiter:  6,
		Area:                "Finance",
		TotalPrice:          7800,
		InvoiceINN:          nil,
		InvoiceAccount:      nil,
		ClientType:          Managed,
		Approved:            true,
		ApprovedAt:          time.Now(),
		Approver:            User{ID: "u11", Name: "Admin2", Role: Admin, Area: "Finance"},
	},
}
