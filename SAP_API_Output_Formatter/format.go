package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-source-list-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToList(raw []byte, l *logger.Logger) (*PurchasingSourceList, error) {
	pm := &responses.PurchasingSourceList{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchasingSourceList. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	list := &List{
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

	return list, nil
}