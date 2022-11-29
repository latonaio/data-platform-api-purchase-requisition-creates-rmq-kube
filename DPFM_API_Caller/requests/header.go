package requests

type Header struct {
	BusinessPartner         *int   `json:"BusinessPartner"`
	PurchaseRequisition     *int   `json:"PurchaseRequisition"`
	PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
	SourceDetermination     *bool  `json:"SourceDetermination"`
	Language                string `json:"Language"`
}
