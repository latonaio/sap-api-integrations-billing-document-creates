package sap_api_input_reader

import (
	"sap-api-integrations-billing-document-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.BillingDocument
	return &requests.Header{
			BillingDocument:            data.BillingDocument,
			BillingDocumentType:        data.BillingDocumentType,
			SDDocumentCategory:         data.SDDocumentCategory,
			BillingDocumentCategory:    data.BillingDocumentCategory,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			SalesOrganization:          data.SalesOrganization,
			DistributionChannel:        data.DistributionChannel,
			Division:                   data.Division,
			BillingDocumentDate:        data.BillingDocumentDate,
			BillingDocumentIsCancelled: data.BillingDocumentIsCancelled,
			CancelledBillingDocument:   data.CancelledBillingDocument,
			IsExportDelivery:           data.IsExportDelivery,
			TotalNetAmount:             data.TotalNetAmount,
			TransactionCurrency:        data.TransactionCurrency,
			TaxAmount:                  data.TaxAmount,
			TotalGrossAmount:           data.TotalGrossAmount,
			CustomerPriceGroup:         data.CustomerPriceGroup,
			IncotermsClassification:    data.IncotermsClassification,
			CustomerPaymentTerms:       data.CustomerPaymentTerms,
			PaymentMethod:              data.PaymentMethod,
			CompanyCode:                data.CompanyCode,
			AccountingDocument:         data.AccountingDocument,
			ExchangeRateDate:           data.ExchangeRateDate,
			ExchangeRateType:           data.ExchangeRateType,
			DocumentReferenceID:        data.DocumentReferenceID,
			SoldToParty:                data.SoldToParty,
			PartnerCompany:             data.PartnerCompany,
			PurchaseOrderByCustomer:    data.PurchaseOrderByCustomer,
			CustomerGroup:              data.CustomerGroup,
			Country:                    data.Country,
			CityCode:                   data.CityCode,
			Region:                     data.Region,
			CreditControlArea:          data.CreditControlArea,
			OverallBillingStatus:       data.OverallBillingStatus,
			AccountingPostingStatus:    data.AccountingPostingStatus,
			AccountingTransferStatus:   data.AccountingTransferStatus,
			InvoiceListStatus:          data.InvoiceListStatus,
			BillingDocumentListType:    data.BillingDocumentListType,
			BillingDocumentListDate:    data.BillingDocumentListDate,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataBillingDocument := sdc.BillingDocument
	data := sdc.BillingDocument.BillingDocumentItem
	return &requests.Item{
		    BillingDocument:              dataBillingDocument.BillingDocument,
			BillingDocumentItem:          data.BillingDocumentItem,
			SalesDocumentItemCategory:    data.SalesDocumentItemCategory,
			ReturnItemProcessingType:     data.ReturnItemProcessingType,
			CreationDate:                 data.CreationDate,
			Division:                     data.Division,
			Material:                     data.Material,
			InternationalArticleNumber:   data.InternationalArticleNumber,
			Batch:                        data.Batch,
			MaterialGroup:                data.MaterialGroup,
			Plant:                        data.Plant,
			StorageLocation:              data.StorageLocation,
			BillingDocumentItemText:      data.BillingDocumentItemText,
			BillingQuantity:              data.BillingQuantity,
			BillingQuantityUnit:          data.BillingQuantityUnit,
			NetAmount:                    data.NetAmount,
			TransactionCurrency:          data.TransactionCurrency,
			GrossAmount:                  data.GrossAmount,
			PricingDate:                  data.PricingDate,
			TaxAmount:                    data.TaxAmount,
			MaterialPricingGroup:         data.MaterialPricingGroup,
			MainItemMaterialPricingGroup: data.MainItemMaterialPricingGroup,
			BusinessArea:                 data.BusinessArea,
			ProfitCenter:                 data.ProfitCenter,
			WBSElement:                   data.WBSElement,
			ControllingArea:              data.ControllingArea,
			ProfitabilitySegment:         data.ProfitabilitySegment,
			OrderID:                      data.OrderID,
			CostCenter:                   data.CostCenter,
			ReferenceSDDocument:          data.ReferenceSDDocument,
			ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
			MatlAccountAssignmentGroup:   data.MatlAccountAssignmentGroup,
			SalesDocument:                data.SalesDocument,
			SalesDocumentItem:            data.SalesDocumentItem,
			SDDocumentReason:             data.SDDocumentReason,
			ShippingPoint:                data.ShippingPoint,
	}
}
