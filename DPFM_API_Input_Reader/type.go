package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string              `json:"connection_key"`
	Result              bool                `json:"result"`
	RedisKey            string              `json:"redis_key"`
	Filepath            string              `json:"filepath"`
	APIStatusCode       int                 `json:"api_status_code"`
	RuntimeSessionID    string              `json:"runtime_session_id"`
	BusinessPartnerID   *int                `json:"business_partner"`
	ServiceLabel        string              `json:"service_label"`
	PurchaseRequisition PurchaseRequisition `json:"PurchaseRequisition"`
	APISchema           string              `json:"api_schema"`
	Accepter            []string            `json:"accepter"`
	OrderID             *int                `json:"order_id"`
	Deleted             bool                `json:"deleted"`
	SQLUpdateResult     *bool               `json:"sql_update_result"`
	SQLUpdateError      string              `json:"sql_update_error"`
	SubfuncResult       *bool               `json:"subfunc_result"`
	SubfuncError        string              `json:"subfunc_error"`
	ExconfResult        *bool               `json:"exconf_result"`
	ExconfError         string              `json:"exconf_error"`
	APIProcessingResult *bool               `json:"api_processing_result"`
	APIProcessingError  string              `json:"api_processing_error"`
}

type PurchaseRequisition struct {
	BusinessPartner         *int   `json:"BusinessPartner"`
	PurchaseRequisition     *int   `json:"PurchaseRequisition"`
	PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
	SourceDetermination     *bool  `json:"SourceDetermination"`
	Language                string `json:"Language"`
	Item                    Item   `json:"Item"`
}

type Item struct {
	BusinessPartner                *int                `json:"BusinessPartner"`
	PurchaseRequisition            *int                `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        *int                `json:"PurchaseRequisitionItem"`
	PurchasingDocument             *int                `json:"PurchasingDocument"`
	PurchasingDocumentItem         *int                `json:"PurchasingDocumentItem"`
	PurReqnReleaseStatus           string              `json:"PurReqnReleaseStatus"`
	PurchasingDocumentItemCategory string              `json:"PurchasingDocumentItemCategory"`
	PurchaseRequisitionItemText    string              `json:"PurchaseRequisitionItemText"`
	Product                        string              `json:"Product"`
	ProductGroup                   string              `json:"ProductGroup"`
	RequestedQuantity              string              `json:"RequestedQuantity"`
	BaseUnit                       string              `json:"BaseUnit"`
	PurchaseRequisitionPrice       *float32            `json:"PurchaseRequisitionPrice"`
	PurReqnPriceQuantity           *float32            `json:"PurReqnPriceQuantity"`
	ProductGoodsReceiptDuration    *int                `json:"ProductGoodsReceiptDuration"`
	PurchaseRequisitionReleaseDate *string             `json:"PurchaseRequisitionReleaseDate"`
	Organization                   string              `json:"Organization"`
	ShippingPlant                  string              `json:"ShippingPlant"`
	ShippingPlantStorageLocation   string              `json:"ShippingPlantStorageLocation"`
	ReceivingPlant                 string              `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  string              `json:"ReceivingPlantStorageLocation"`
	SourceOfSupplyIsAssigned       *bool               `json:"SourceOfSupplyIsAssigned"`
	OrderedQuantity                *float32            `json:"OrderedQuantity"`
	DeliveryDate                   *string             `json:"DeliveryDate"`
	CreationDate                   *string             `json:"CreationDate"`
	ProcessingStatus               string              `json:"ProcessingStatus"`
	PurchasingInfoRecord           *int                `json:"PurchasingInfoRecord"`
	Supplier                       *int                `json:"Supplier"`
	RequisitionerName              string              `json:"RequisitionerName"`
	PurReqnItemCurrency            string              `json:"PurReqnItemCurrency"`
	ProductPlannedDeliveryDurn     *int                `json:"ProductPlannedDeliveryDurn"`
	IsPurReqnBlocked               *bool               `json:"IsPurReqnBlocked"`
	GLAccount                      string              `json:"GLAccount"`
	BusinessArea                   string              `json:"BusinessArea"`
	Batch                          string              `json:"Batch"`
	MRPController                  string              `json:"MRPController"`
	Reservation                    *int                `json:"Reservation"`
	LastChangeDateTime             *string             `json:"LastChangeDateTime"`
	IsDeleted                      *bool               `json:"IsDeleted"`
	ItemDeliveryAddress            ItemDeliveryAddress `json:"ItemDeliveryAddress"`
}

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
