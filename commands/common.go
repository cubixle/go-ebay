package commands

import "time"

type RequesterCredentials struct {
	EBayAuthToken string `xml:"eBayAuthToken"`
}

type Response interface {
	ResponseErrors() ebayErrors
	GetResponse()
}

type ebayErrors []ebayResponseError

type ebayResponseError struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int
	SeverityCode        string
	ErrorClassification string
}

type ShippingServiceOption struct {
	ShippingService               string
	ShippingServiceCost           float64
	ShippingServiceAdditionalCost float64
	FreeShipping                  bool
}

type ReturnPolicy struct {
	ReturnsAccepted, ReturnsAcceptedOption, ReturnsWithinOption, RefundOption string
}

type Storefront struct {
	StoreCategoryID string
}

type ProductListingDetails struct {
	UPC      string
	BrandMPN BrandMPN
}

type PrimaryCategory struct {
	CategoryID string
}

type ItemSpecifics struct {
	NameValueList []NameValueList
}

type NameValueList struct {
	Name  string
	Value []string
}

type BrandMPN struct {
	Brand, MPN string
}

type BestOfferDetails struct {
	BestOfferEnabled bool
}

type ebayResponse struct {
	Timestamp time.Time
	Ack       string
	Errors    []ebayResponseError
}
