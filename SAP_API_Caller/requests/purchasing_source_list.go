package requests

type PurchasingSourceList struct {
	Material                   string `json:"Material"`
	Plant                      string `json:"Plant"`
	SourceListRecord           *string `json:"SourceListRecord"`
	ValidityStartDate          *string `json:"ValidityStartDate"`
	ValidityEndDate            *string `json:"ValidityEndDate"`
	Supplier                   *string `json:"Supplier"`
	PurchasingOrganization     *string `json:"PurchasingOrganization"`
	SupplyingPlant             *string `json:"SupplyingPlant"`
	OrderQuantityUnit          *string `json:"OrderQuantityUnit"`
	PurchaseOutlineAgreement   *string `json:"PurchaseOutlineAgreement"`
	SupplierIsFixed            *bool   `json:"SupplierIsFixed"`
	SourceOfSupplyIsBlocked    *bool   `json:"SourceOfSupplyIsBlocked"`
	MRPSourcingControl         *string `json:"MRPSourcingControl"`
	LastChangeDateTime         *string `json:"LastChangeDateTime"`
	IssgPlantIsFixed           *bool   `json:"IssgPlantIsFixed"`
	PurOutlineAgreementIsFixed *bool   `json:"PurOutlineAgreementIsFixed"`
	SourceOfSupplyIsFixed      *bool   `json:"SourceOfSupplyIsFixed"`
}
