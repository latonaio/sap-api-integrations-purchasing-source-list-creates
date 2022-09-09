package sap_api_input_reader

import (
	"sap-api-integrations-purchasing-source-list-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToPurchasingSourceList() *requests.PurchasingSourceList {
	data := sdc.List
	return &requests.PurchasingSourceList{
		Material:                   data.Material,
		Plant:                      data.Plant,
		SourceListRecord:           data.SourceListRecord,
		ValidityStartDate:          data.ValidityStartDate,
		ValidityEndDate:            data.ValidityEndDate,
		Supplier:                   data.Supplier,
		PurchasingOrganization:     data.PurchasingOrganization,
		SupplyingPlant:             data.SupplyingPlant,
		OrderQuantityUnit:          data.OrderQuantityUnit,
		PurchaseOutlineAgreement:   data.PurchaseOutlineAgreement,
		SupplierIsFixed:            data.SupplierIsFixed,
		SourceOfSupplyIsBlocked:    data.SourceOfSupplyIsBlocked,
		MRPSourcingControl:         data.MRPSourcingControl,
		LastChangeDateTime:         data.LastChangeDateTime,
		IssgPlantIsFixed:           data.IssgPlantIsFixed,
		PurOutlineAgreementIsFixed: data.PurOutlineAgreementIsFixed,
		SourceOfSupplyIsFixed:      data.SourceOfSupplyIsFixed,
	}
}