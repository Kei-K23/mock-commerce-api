package dto

type UserRequest struct {
	ID            int     `json:"id"`
	Username      string  `json:"username"`
	FirstName     string  `json:"firstname"`
	LastName      string  `json:"lastname"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Phone         string  `json:"phone"`
	City          *string `json:"city"`
	Street        *string `json:"street"`
	AddressNumber *string `json:"address_number"`
	ZipCode       *string `json:"zip_code"`
	Lat           *string `json:"lat"`
	Long          *string `json:"long"`
}
