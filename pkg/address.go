package pub

import "fmt"
//Address is the address
type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
}
//FullAddress returns the full address
func (a Address) FullAddress() string {
	return fmt.Sprintf("%s %s %s %s",a.AddressLine1,a.AddressLine2,a.City,a.State)
}
