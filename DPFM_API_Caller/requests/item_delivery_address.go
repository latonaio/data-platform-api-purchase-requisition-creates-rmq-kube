package requests

type ItemDeliveryAddress struct {
	BusinessPartner         *int   `json:"BusinessPartner"`
	PurchaseRequisition     *int   `json:"PurchaseRequisition"`
	PurchaseRequisitionItem *int   `json:"PurchaseRequisitionItem"`
	AddressID               *int   `json:"AddressID"`
	PostalCode              string `json:"PostalCode"`
	LocalRegion             string `json:"LocalRegion"`
	Country                 string `json:"Country"`
	StreetName              string `json:"StreetName"`
	CityName                string `json:"CityName"`
}
