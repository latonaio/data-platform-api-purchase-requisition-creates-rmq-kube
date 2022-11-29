package requests

type Item struct {
	BusinessPartner                *int     `json:"BusinessPartner"`
	PurchaseRequisition            *int     `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        *int     `json:"PurchaseRequisitionItem"`
	PurchasingDocument             *int     `json:"PurchasingDocument"`
	PurchasingDocumentItem         *int     `json:"PurchasingDocumentItem"`
	PurReqnReleaseStatus           string   `json:"PurReqnReleaseStatus"`
	PurchasingDocumentItemCategory string   `json:"PurchasingDocumentItemCategory"`
	PurchaseRequisitionItemText    string   `json:"PurchaseRequisitionItemText"`
	Product                        string   `json:"Product"`
	ProductGroup                   string   `json:"ProductGroup"`
	RequestedQuantity              string   `json:"RequestedQuantity"`
	BaseUnit                       string   `json:"BaseUnit"`
	PurchaseRequisitionPrice       *float32 `json:"PurchaseRequisitionPrice"`
	PurReqnPriceQuantity           *float32 `json:"PurReqnPriceQuantity"`
	ProductGoodsReceiptDuration    *int     `json:"ProductGoodsReceiptDuration"`
	PurchaseRequisitionReleaseDate *string  `json:"PurchaseRequisitionReleaseDate"`
	Organization                   string   `json:"Organization"`
	ShippingPlant                  string   `json:"ShippingPlant"`
	ShippingPlantStorageLocation   string   `json:"ShippingPlantStorageLocation"`
	ReceivingPlant                 string   `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  string   `json:"ReceivingPlantStorageLocation"`
	SourceOfSupplyIsAssigned       *bool    `json:"SourceOfSupplyIsAssigned"`
	OrderedQuantity                *float32 `json:"OrderedQuantity"`
	DeliveryDate                   *string  `json:"DeliveryDate"`
	CreationDate                   *string  `json:"CreationDate"`
	ProcessingStatus               string   `json:"ProcessingStatus"`
	PurchasingInfoRecord           *int     `json:"PurchasingInfoRecord"`
	Supplier                       *int     `json:"Supplier"`
	RequisitionerName              string   `json:"RequisitionerName"`
	PurReqnItemCurrency            string   `json:"PurReqnItemCurrency"`
	ProductPlannedDeliveryDurn     *int     `json:"ProductPlannedDeliveryDurn"`
	IsPurReqnBlocked               *bool    `json:"IsPurReqnBlocked"`
	GLAccount                      string   `json:"GLAccount"`
	BusinessArea                   string   `json:"BusinessArea"`
	Batch                          string   `json:"Batch"`
	MRPController                  string   `json:"MRPController"`
	Reservation                    *int     `json:"Reservation"`
	LastChangeDateTime             *string  `json:"LastChangeDateTime"`
	IsDeleted                      *bool    `json:"IsDeleted"`
}
