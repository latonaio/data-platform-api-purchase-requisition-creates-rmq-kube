package dpfm_api_processing_formatter

type HeaderUpdates struct {
	PurchaseRequisition                int    `json:"PurchaseRequisition"`
	PurchaseRequisitionDate            string `json:"PurchaseRequisitionDate"`
	Project                            *int   `json:"Project"`
	WBSElement                         *int   `json:"WBSElement"`
}

type ItemUpdates struct {
	PurchaseRequisition                           int      `json:"PurchaseRequisition"`
	PurchaseRequisitionItem                       int      `json:"PurchaseRequisitionItem"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                int      `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                string   `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	Product                                       string   `json:"Product"`
	RequestedQuantityInBaseUnit                   float32  `json:"RequestedQuantityInBaseUnit"`
	RequestedDeliveryDate                         string   `json:"RequestedDeliveryDate"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	PurchaseRequisitionItemText                   *string  `json:"PurchaseRequisitionItemText"`
	PurchaseRequisitionItemTextByBuyer            *string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller           *string  `json:"PurchaseRequisitionItemTextBySeller"`
	PurchaseRequisitionItemPrice                  *float32 `json:"PurchaseRequisitionItemPrice"`
	PurchaseRequisitionItemPriceQuantity          *int     `json:"PurchaseRequisitionItemPriceQuantity"`
	ProductPlannedDeliveryDuration                *float32 `json:"ProductPlannedDeliveryDuration"`
	ProductPlannedDeliveryDurationUnit            *int     `json:"ProductPlannedDeliveryDurationUnit"`
	TransactionCurrency               			  *string  `json:"TransactionCurrency"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
}

type PartnerUpdates struct {
	PurchaseRequisition     int     `json:"PurchaseRequisition"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
	PurchaseRequisition     int     `json:"OrderID"`
	AddressID   			int     `json:"AddressID"`
	PostalCode  			*string `json:"PostalCode"`
	LocalRegion 			*string `json:"LocalRegion"`
	Country     			*string `json:"Country"`
	District    			*string `json:"District"`
	StreetName  			*string `json:"StreetName"`
	CityName    			*string `json:"CityName"`
	Building    			*string `json:"Building"`
	Floor       			*int    `json:"Floor"`
	Room        			*int    `json:"Room"`
}
