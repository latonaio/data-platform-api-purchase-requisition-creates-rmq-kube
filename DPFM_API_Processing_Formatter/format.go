package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-purchase-requisition-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		PurchaseRequisition:		data.PurchaseRequisition,
		PurchaseRequisitionDate:	data.PurchaseRequisitionDate,
		Project:					data.Project,
		WBSElement:					data.WBSElement,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item
	return &ItemUpdates{
			PurchaseRequisition:						data.PurchaseRequisition,
			PurchaseRequisitionItem:					data.PurchaseRequisitionItem,
			SupplyChainRelationshipID:					data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:			data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:		data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:	data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:	data.SupplyChainRelationshipProductionPlantID,
			Seller:										data.Seller,
			DeliverToParty:								data.DeliverToParty,
			DeliverFromParty:							data.DeliverFromParty,
			DeliverToPlant:								data.DeliverToPlant,
			DeliverToPlantStorageLocation:				data.DeliverToPlantStorageLocation,
			DeliverFromPlant:							data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:			data.DeliverFromPlantStorageLocation,
			Product:									data.Product,
			RequestedQuantity:							data.RequestedQuantity,
			Project:									data.Project,
			WBSElement:									data.WBSElement,
			PurchaseRequisitionItemText:				data.PurchaseRequisitionItemText,
			PurchaseRequisitionItemTextByBuyer:			data.PurchaseRequisitionItemTextByBuyer,
			PurchaseRequisitionItemTextBySeller:		data.PurchaseRequisitionItemTextBySeller,
			PurchaseRequisitionItemPrice:				data.PurchaseRequisitionItemPrice,
			PurchaseRequisitionItemPriceQuantity:		data.PurchaseRequisitionItemPriceQuantity,
			ProductPlannedDeliveryDuration:				data.ProductPlannedDeliveryDuration,
			ProductPlannedDeliveryDurationUnit:			data.ProductPlannedDeliveryDurationUnit,
			TransactionCurrency:						data.TransactionCurrency,
			StockConfirmationBusinessPartner:			data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:						data.StockConfirmationPlant,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		PurchaseRequisition:     dataHeader.PurchaseRequisition,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		PurchaseRequisition:	dataHeader.PurchaseRequisition,
		AddressID:   			data.AddressID,
		PostalCode:  			data.PostalCode,
		LocalRegion: 			data.LocalRegion,
		Country:     			data.Country,
		District:    			data.District,
		StreetName:  			data.StreetName,
		CityName:    			data.CityName,
		Building:    			data.Building,
		Floor:       			data.Floor,
		Room:        			data.Room,
	}
}
