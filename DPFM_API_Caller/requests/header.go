package requests

type Header struct {
	PurchaseRequisition                int    `json:"PurchaseRequisition"`
	PurchaseRequisitionDate            string `json:"PurchaseRequisitionDate"`
	Buyer                              int    `json:"Buyer"`
	PurchaseRequisitionType            string `json:"PurchaseRequisitionType"`
	PlannedOrder                       *int   `json:"PlannedOrder"`
	PlannedOrderItem                   *int   `json:"PlannedOrderItem"`
	ProductionOrder                    *int   `json:"ProductionOrder"`
	ProductionOrderItem                *int   `json:"ProductionOrderItem"`
	PrecedingOrderID                   *int   `json:"PrecedingOrderID"`
	PrecedingOrderItem                 *int   `json:"PrecedingOrderItem"`
	Project                            *int   `json:"Project"`
	WBSElement                         *int   `json:"WBSElement"`
	HeaderOrderStatus				   string `json:"HeaderOrderStatus"`
	HeaderCompleteOrderIsDefined	   *bool  `json:"HeaderCompleteOrderIsDefined"`
	CreationDate                       string `json:"CreationDate"`
	CreationTime                       string `json:"CreationTime"`
	LastChangeDate                     string `json:"LastChangeDate"`
	LastChangeTime                     string `json:"LastChangeTime"`
	IsReleased                         *bool  `json:"IsReleased"`
	IsCancelled                        *bool  `json:"IsCancelled"`
	IsMarkedForDeletion                *bool  `json:"IsMarkedForDeletion"`
}
