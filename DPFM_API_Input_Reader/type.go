package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string          `json:"connection_key"`
	Result           bool            `json:"result"`
	RedisKey         string          `json:"redis_key"`
	Filepath         string          `json:"filepath"`
	APIStatusCode    int             `json:"api_status_code"`
	RuntimeSessionID string          `json:"runtime_session_id"`
	BusinessPartner  *int            `json:"business_partner"`
	ServiceLabel     string          `json:"service_label"`
	APIType          string          `json:"api_type"`
	Header           Header          `json:"PurchaseRequisition"`
	APISchema        string          `json:"api_schema"`
	Accepter         []string        `json:"accepter"`
	Deleted          bool            `json:"deleted"`
}

type Header struct {
	PurchaseRequisition                int     `json:"PurchaseRequisition"`
	PurchaseRequisitionDate            *string `json:"PurchaseRequisitionDate"`
	Buyer                              *int    `json:"Buyer"`
	PurchaseRequisitionType            *string `json:"PurchaseRequisitionType"`
	PlannedOrder                       *int    `json:"PlannedOrder"`
	PlannedOrderItem                   *int    `json:"PlannedOrderItem"`
	ProductionOrder                    *int    `json:"ProductionOrder"`
	ProductionOrderItem                *int    `json:"ProductionOrderItem"`
	PrecedingOrderID                   *int    `json:"PrecedingOrderID"`
	PrecedingOrderItem                 *int    `json:"PrecedingOrderItem"`
	Project                            *int    `json:"Project"`
	WBSElement                         *int    `json:"WBSElement"`
	HeaderOrderStatus				   *string `json:"HeaderOrderStatus"`
	HeaderCompleteOrderIsDefined	   *bool   `json:"HeaderCompleteOrderIsDefined"`
	CreationDate                       *string `json:"CreationDate"`
	CreationTime                       *string `json:"CreationTime"`
	LastChangeDate                     *string `json:"LastChangeDate"`
	LastChangeTime                     *string `json:"LastChangeTime"`
	IsReleased                         *bool  `json:"IsReleased"`
	IsCancelled                        *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                *bool   `json:"IsMarkedForDeletion"`
	Item                               []Item  		`json:"Item"`
	Partner                            []Partner	`json:"Partner"`
	Address                            []Address	`json:"Address"`
}

type Item struct {
	PurchaseRequisition                           int      `json:"PurchaseRequisition"`
	PurchaseRequisitionItem                       int      `json:"PurchaseRequisitionItem"`
	PurchaseRequisitionItemCategory               *string  `json:"PurchaseRequisitionItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                         int      `json:"Buyer"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	Product                                       *string  `json:"Product"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	RequestedQuantityInBaseUnit                   *float32 `json:"RequestedQuantityInBaseUnit"`
	RequestedQuantityInDeliveryUnit               *float32 `json:"RequestedQuantityInDeliveryUnit"`
	BaseUnit                                      *string  `json:"BaseUnit"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	PlannedOrder                                  *int     `json:"PlannedOrder"`
	PlannedOrderItem                              *int     `json:"PlannedOrderItem"`
	ProductionOrder                    			  *int     `json:"ProductionOrder"`
	ProductionOrderItem                			  *int     `json:"ProductionOrderItem"`
	PrecedingOrderID                              *int     `json:"PrecedingOrderID"`
	PrecedingOrderItem                            *int     `json:"PrecedingOrderItem"`
	FollowingOrderID                              *int     `json:"FollowingOrderID"`
	FollowingOrderItem                            *int     `json:"FollowingOrderItem"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	PurchaseRequisitionItemText                   *string  `json:"PurchaseRequisitionItemText"`
	PurchaseRequisitionItemTextByBuyer            *string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller           *string  `json:"PurchaseRequisitionItemTextBySeller"`
	PurchaseRequisitionItemPrice                  *float32 `json:"PurchaseRequisitionItemPrice"`
	PurchaseRequisitionItemPriceQuantity          *int     `json:"PurchaseRequisitionItemPriceQuantity"`
	ProductPlannedDeliveryDuration                *float32 `json:"ProductPlannedDeliveryDuration"`
	ProductPlannedDeliveryDurationUnit            *int     `json:"ProductPlannedDeliveryDurationUnit"`
	OrderedQuantityInBaseUnit                     *float32 `json:"OrderedQuantityInBaseUnit"`
	OrderedQuantityInDeliveryUnit                 *float32 `json:"OrderedQuantityInDeliveryUnit"`
	DeliveryDate                                  *string  `json:"DeliveryDate"`
	ItemCompleteOrderIsDefined             		  *bool    `json:"ItemCompleteOrderIsDefined"`
	TransactionCurrency               			  *string  `json:"TransactionCurrency"`
	MRPArea                                       *string  `json:"MRPArea"`
	MRPController                                 *string  `json:"MRPController"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	GLAccount                                     *string  `json:"GLAccount"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	CreationDate                       			  *string  `json:"CreationDate"`
	CreationTime                       			  *string  `json:"CreationTime"`
	LastChangeDate                     			  *string  `json:"LastChangeDate"`
	LastChangeTime                     			  *string  `json:"LastChangeTime"`
	IsReleased                         			  *bool    `json:"IsReleased"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}

type Partner struct {
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
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
	PurchaseRequisition     int     `json:"PurchaseRequisition"`
	AddressID   			int     `json:"AddressID"`
	PostalCode  			*string `json:"PostalCode"`
	LocalRegion 			*string `json:"LocalRegion"`
	Country     			*string `json:"Country"`
	District    			*string `json:"District"`
	StreetName 				*string `json:"StreetName"`
	CityName    			*string `json:"CityName"`
	Building    			*string `json:"Building"`
	Floor       			*int    `json:"Floor"`
	Room        			*int    `json:"Room"`
}
