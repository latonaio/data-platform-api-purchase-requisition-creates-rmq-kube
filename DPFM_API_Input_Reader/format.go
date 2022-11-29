package dpfm_api_input_reader

import (
	"data-platform-api-purchase-requisition-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseRequisition
	return &requests.Header{
		BusinessPartner:         data.BusinessPartner,
		PurchaseRequisition:     data.PurchaseRequisition,
		PurchaseRequisitionType: data.PurchaseRequisitionType,
		SourceDetermination:     data.SourceDetermination,
		Language:                data.Language,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataPurchaseRequisition := sdc.PurchaseRequisition
	data := sdc.PurchaseRequisition.Item
	return &requests.Item{
		BusinessPartner:                dataPurchaseRequisition.BusinessPartner,
		PurchaseRequisition:            dataPurchaseRequisition.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		Product:                        data.Product,
		ProductGroup:                   data.ProductGroup,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		ProductGoodsReceiptDuration:    data.ProductGoodsReceiptDuration,
		PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		Organization:                   data.Organization,
		ShippingPlant:                  data.ShippingPlant,
		ShippingPlantStorageLocation:   data.ShippingPlantStorageLocation,
		ReceivingPlant:                 data.ReceivingPlant,
		ReceivingPlantStorageLocation:  data.ReceivingPlantStorageLocation,
		SourceOfSupplyIsAssigned:       data.SourceOfSupplyIsAssigned,
		OrderedQuantity:                data.OrderedQuantity,
		DeliveryDate:                   data.DeliveryDate,
		CreationDate:                   data.CreationDate,
		ProcessingStatus:               data.ProcessingStatus,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		Supplier:                       data.Supplier,
		RequisitionerName:              data.RequisitionerName,
		PurReqnItemCurrency:            data.PurReqnItemCurrency,
		ProductPlannedDeliveryDurn:     data.ProductPlannedDeliveryDurn,
		IsPurReqnBlocked:               data.IsPurReqnBlocked,
		GLAccount:                      data.GLAccount,
		BusinessArea:                   data.BusinessArea,
		Batch:                          data.Batch,
		MRPController:                  data.MRPController,
		Reservation:                    data.Reservation,
		LastChangeDateTime:             data.LastChangeDateTime,
		IsDeleted:                      data.IsDeleted,
	}
}

func (sdc *SDC) ConvertToItemDeliveryAddress() *requests.ItemDeliveryAddress {
	dataPurchaseRequisitionItem := sdc.PurchaseRequisition.Item
	data := dataPurchaseRequisitionItem.ItemDeliveryAddress
	return &requests.ItemDeliveryAddress{
		BusinessPartner:         dataPurchaseRequisitionItem.BusinessPartner,
		PurchaseRequisition:     dataPurchaseRequisitionItem.PurchaseRequisition,
		PurchaseRequisitionItem: dataPurchaseRequisitionItem.PurchaseRequisitionItem,
		AddressID:               data.AddressID,
		PostalCode:              data.PostalCode,
		LocalRegion:             data.LocalRegion,
		Country:                 data.Country,
		StreetName:              data.StreetName,
		CityName:                data.CityName,
	}
}
