package types

import "time"

type Item struct {
	AutoPay bool

	BestOfferDetails BestOfferDetails `xml:"BestOfferDetails"`

	BuyerRequirementDetails        *BuyerRequirementDetails `xml:"BuyerRequirementDetails,omitempty"`
	CategoryBasedAttributesPrefill bool                     `xml:"CategoryBasedAttributesPrefill"`
	CategoryMappingAllowed         bool                     `xml:"CategoryMappingAllowed"`
	Charity                        *Charity                 `xml:"Charity,omitempty"`
	ConditionDescription           string                   `xml:"ConditionDescription,omitempty"`
	ConditionID                    int                      `xml:"ConditionID,omitempty"`
	Country                        string                   `xml:"Country,omitempty"`
	CrossBorderTrade               string                   `xml:"CrossBorderTrade,omitempty"`
	Currency                       string                   `xml:"Currency,omitempty"`
	Description                    string                   `xml:"Description,omitempty"`

	DigitalGoodInfo struct {
		DigitalDelivery bool `xml:"DigitalDelivery"`
	} `xml:"DigitalGoodInfo"`

	DisableBuyerRequirements bool               `xml:"DisableBuyerRequirements"`
	DiscountPriceInfo        *DiscountPriceInfo `xml:"DiscountPriceInfo"`

	DispatchTimeMax         int    `xml:"DispatchTimeMax"`
	EBayNowEligible         bool   `xml:"eBayNowEligible"`
	EBayPlus                bool   `xml:"eBayPlus"`
	HitCounter              string `xml:"HitCounter,omitempty"`
	IncludeRecommendations  bool   `xml:"IncludeRecommendations"`
	InventoryTrackingMethod string `xml:"InventoryTrackingMethod,omitempty"`

	ItemCompatibilityList *struct {
		Compatibility struct {
			CompatibilityNotes string `xml:"CompatibilityNotes"`
			NameValueList      []struct {
				Name  string `xml:"Name"`
				Value string `xml:"Value"`
			} `xml:"NameValueList"`
		} `xml:"Compatibility"`
	} `xml:"ItemCompatibilityList,omitempty"`

	ItemSpecifics ItemSpecifics `xml:"ItemSpecifics"`

	ListingDesigner *struct {
		LayoutID           int  `xml:"LayoutID"`
		OptimalPictureSize bool `xml:"OptimalPictureSize"`
		ThemeID            int  `xml:"ThemeID"`
	} `xml:"ListingDesigner,omitempty"`

	ListingDetails *struct {
		BestOfferAutoAcceptPrice float32 `xml:"BestOfferAutoAcceptPrice"`
		LocalListingDistance     string  `xml:"LocalListingDistance"`
		MinimumBestOfferPrice    float32 `xml:"MinimumBestOfferPrice"`
	} `xml:"ListingDetails,omitempty"`

	ListingDuration    string `xml:"ListingDuration,omitempty"`
	ListingEnhancement string `xml:"ListingEnhancement,omitempty"`
	ListingType        string `xml:"ListingType,omitempty"`
	Location           string `xml:"Location,omitempty"`

	PaymentMethods       string `xml:"PaymentMethods,omitempty"`
	PayPalEmailAddress   string `xml:"PayPalEmailAddress"`
	PickupInStoreDetails struct {
		EligibleForPickupInStore bool `xml:"EligibleForPickupInStore"`
	} `xml:"PickupInStoreDetails"`

	PictureDetails PictureDetails `xml:"PictureDetails"`

	PostalCode            string                `xml:"PostalCode"`
	PrimaryCategory       CategoryID            `xml:"PrimaryCategory"`
	PrivateNotes          string                `xml:"PrivateNotes,omitempty"`
	ProductListingDetails ProductListingDetails `xml:"ProductListingDetails"`
	Quantity              int                   `xml:"Quantity"`

	QuantityInfo *struct {
		MinimumRemnantSet int `xml:"MinimumRemnantSet"`
	} `xml:"QuantityInfo,omitempty"`

	QuantityRestrictionPerBuyer *struct {
		MaximumQuantity int `xml:"MaximumQuantity"`
	} `xml:"QuantityRestrictionPerBuyer,omitempty"`

	ReturnPolicy *ReturnPolicy `xml:"ReturnPolicy,omitempty"`

	ScheduleTime      *time.Time  `xml:"ScheduleTime,omitempty"`
	SecondaryCategory *CategoryID `xml:"SecondaryCategory,omitempty"`
	SellerProfiles    *struct {
		SellerPaymentProfile struct {
			PaymentProfileID   int64  `xml:"PaymentProfileID"`
			PaymentProfileName string `xml:"PaymentProfileName"`
		} `xml:"SellerPaymentProfile"`
		SellerReturnProfile struct {
			ReturnProfileID   int64  `xml:"ReturnProfileID"`
			ReturnProfileName string `xml:"ReturnProfileName"`
		} `xml:"SellerReturnProfile"`
		SellerShippingProfile struct {
			ShippingProfileID   int64  `xml:"ShippingProfileID"`
			ShippingProfileName string `xml:"ShippingProfileName"`
		} `xml:"SellerShippingProfile"`
	} `xml:"SellerProfiles,omitempty"`

	SellerProvidedTitle             string                           `xml:"SellerProvidedTitle,omitempty"`
	ShippingDetails                 *ShippingDetails                 `xml:"ShippingDetails"`
	ShippingPackageDetails          *ShippingPackageDetails          `xml:"ShippingPackageDetails,omitempty"`
	ShippingServiceCostOverrideList *ShippingServiceCostOverrideList `xml:"ShippingServiceCostOverrideList,omitempty"`
	ShippingTermsInDescription      bool                             `xml:"ShippingTermsInDescription"`
	ShipToLocations                 string                           `xml:"ShipToLocations"`
	Site                            string                           `xml:"Site"` // https://developer.ebay.com/devzone/xml/docs/reference/ebay/types/SiteCodeType.html
	SKU                             string                           `xml:"SKU"`
	StartPrice                      float32                          `xml:"StartPrice,omitempty"`
	Storefront                      *StoreFront                      `xml:"Storefront,omitempty"`
	SubTitle                        string                           `xml:"SubTitle,omitempty"`
	TaxCategory                     string                           `xml:"TaxCategory,omitempty"`
	Title                           string                           `xml:"Title"`
	UseTaxTable                     bool                             `xml:"UseTaxTable,omitempty"`
	UUID                            string                           `xml:"UUID,omitempty"`
	Variations                      *Variations                      `xml:"Variations,omitempty"`
	VATDetails                      *VATDetails                      `xml:"VATDetails,omitempty"`
	VIN                             string                           `xml:"VIN,omitempty"`
	VRM                             string                           `xml:"VRM,omitempty"`
}

type ShippingDetails struct {
	CalculatedShippingRate *struct {
		InternationalPackagingHandlingCosts float32 `xml:"InternationalPackagingHandlingCosts,omitempty"`
		MeasurementUnit                     string  `xml:"MeasurementUnit,omitempty"` // English/Metric MeasurementSystemCodeType
		OriginatingPostalCode               string  `xml:"OriginatingPostalCode"`
		PackagingHandlingCosts              float32 `xml:"PackagingHandlingCosts"`
		ShippingIrregular                   bool    `xml:"ShippingIrregular"`
	} `xml:"CalculatedShippingRate,omitempty"`

	CODCost               float32 `xml:"CODCost"`
	ExcludeShipToLocation string  `xml:"ExcludeShipToLocation,omitempty"`

	GlobalShipping                           bool   `xml:"GlobalShipping"`
	InternationalPromotionalShippingDiscount bool   `xml:"InternationalPromotionalShippingDiscount,omitempty"`
	InternationalShippingDiscountProfileID   string `xml:"InternationalShippingDiscountProfileID,omitempty"`
	InternationalShippingServiceOption       *struct {
		ShippingService               string  `xml:"ShippingService"`
		ShippingServiceAdditionalCost float32 `xml:"ShippingServiceAdditionalCost"`
		ShippingServiceCost           float32 `xml:"ShippingServiceCost"`
		ShippingServicePriority       int     `xml:"ShippingServicePriority"`
		ShipToLocation                string  `xml:"ShipToLocation"`
	} `xml:"InternationalShippingServiceOption,omitempty"`
	PaymentInstructions         string `xml:"PaymentInstructions"`
	PromotionalShippingDiscount bool   `xml:"PromotionalShippingDiscount"`
	RateTableDetails            *struct {
		DomesticRateTable        string `xml:"DomesticRateTable"`
		DomesticRateTableID      string `xml:"DomesticRateTableId"`
		InternationalRateTable   string `xml:"InternationalRateTable"`
		InternationalRateTableID string `xml:"InternationalRateTableId"`
	} `xml:"RateTableDetails,omitempty"`
	SalesTax *struct {
		SalesTaxPercent       float32 `xml:"SalesTaxPercent"`
		SalesTaxState         string  `xml:"SalesTaxState"`
		ShippingIncludedInTax bool    `xml:"ShippingIncludedInTax"`
	} `xml:"SalesTax,omitempty"`
	ShippingDiscountProfileID string                   `xml:"ShippingDiscountProfileID,omitempty"`
	ShippingServiceOptions    []ShippingServiceOptions `xml:"ShippingServiceOptions,omitempty"`
	ShippingType              string                   `xml:"ShippingType"`
}

type CategoryID struct {
	CategoryID string `xml:"CategoryID"`
}

type ProductListingDetails struct {
	BrandMPN                       BrandMPN `xml:"BrandMPN"`
	EAN                            string   `xml:"EAN,omitempty"`
	IncludeeBayProductDetails      bool     `xml:"IncludeeBayProductDetails"`
	IncludeStockPhotoURL           bool     `xml:"IncludeStockPhotoURL"`
	ISBN                           string   `xml:"ISBN,omitempty"`
	ProductReferenceID             string   `xml:"ProductReferenceID,omitempty"`
	ReturnSearchResultOnDuplicates bool     `xml:"ReturnSearchResultOnDuplicates"`
	TicketListingDetails           *struct {
		EventTitle  string `xml:"EventTitle"`
		PrintedDate string `xml:"PrintedDate"`
		PrintedTime string `xml:"PrintedTime"`
		Venue       string `xml:"Venue"`
	} `xml:"TicketListingDetails,omitempty"`
	UPC                       string `xml:"UPC,omitempty"`
	UseFirstProduct           bool   `xml:"UseFirstProduct"`
	UseStockPhotoURLAsGallery bool   `xml:"UseStockPhotoURLAsGallery"`
}

type BrandMPN struct {
	Brand string `xml:"Brand"`
	MPN   string `xml:"MPN"`
}

type BuyerRequirementDetails struct {
	MaximumItemRequirements      MaximumItemRequirements      `xml:"MaximumItemRequirements"`
	MaximumUnpaidItemStrikesInfo MaximumUnpaidItemStrikesInfo `xml:"MaximumUnpaidItemStrikesInfo"`
	ShipToRegistrationCountry    bool                         `xml:"ShipToRegistrationCountry"`
	ZeroFeedbackScore            bool                         `xml:"ZeroFeedbackScore"`
}

type MaximumItemRequirements struct {
	MaximumItemCount     int `xml:"MaximumItemCount"`
	MinimumFeedbackScore int `xml:"MinimumFeedbackScore"`
}

type MaximumUnpaidItemStrikesInfo struct {
	Count  int    `xml:"Count"`
	Period string `xml:"Period"`
}

type DiscountPriceInfo struct {
	MadeForOutletComparisonPrice   float32 `xml:"MadeForOutletComparisonPrice"`
	MinimumAdvertisedPrice         float32 `xml:"MinimumAdvertisedPrice"`
	MinimumAdvertisedPriceExposure string  `xml:"MinimumAdvertisedPriceExposure"`
	OriginalRetailPrice            float32 `xml:"OriginalRetailPrice"`
	SoldOffeBay                    bool    `xml:"SoldOffeBay"`
	SoldOneBay                     bool    `xml:"SoldOneBay"`
}

type ShippingServiceOptions struct {
	FreeShipping                  bool    `xml:"FreeShipping"`
	ShippingService               string  `xml:"ShippingService,omitempty"`
	ShippingServiceAdditionalCost float32 `xml:"ShippingServiceAdditionalCost,omitempty"`
	ShippingServiceCost           float32 `xml:"ShippingServiceCost,omitempty"`
	ShippingServicePriority       int     `xml:"ShippingServicePriority,omitempty"`
	ShippingSurcharge             float32 `xml:"ShippingSurcharge,omitempty"`
}

type ShippingPackageDetails struct {
	MeasurementUnit   string  `xml:"MeasurementUnit"` // English/Metric MeasurementSystemCodeType
	PackageDepth      float32 `xml:"PackageDepth"`
	PackageLength     float32 `xml:"PackageLength"`
	PackageWidth      float32 `xml:"PackageWidth"`
	ShippingIrregular bool    `xml:"ShippingIrregular"`
	ShippingPackage   string  `xml:"ShippingPackage"` // https://developer.ebay.com/devzone/xml/docs/reference/ebay/types/ShippingPackageCodeType.html
	WeightMajor       float32 `xml:"WeightMajor"`
	WeightMinor       float32 `xml:"WeightMinor"`
}

type ShippingServiceCostOverrideList struct {
	ShippingServiceCostOverride struct {
		ShippingServiceAdditionalCost struct {
		} `xml:"ShippingServiceAdditionalCost"`
		ShippingServiceCost struct {
		} `xml:"ShippingServiceCost"`
		ShippingServicePriority struct {
		} `xml:"ShippingServicePriority"`
		ShippingServiceType struct {
		} `xml:"ShippingServiceType"`
		ShippingSurcharge struct {
		} `xml:"ShippingSurcharge"`
	} `xml:"ShippingServiceCostOverride"`
}

type StoreFront struct {
	StoreCategory2ID   int64  `xml:"StoreCategory2ID"`
	StoreCategory2Name string `xml:"StoreCategory2Name"`
	StoreCategoryID    int64  `xml:"StoreCategoryID"`
	StoreCategoryName  string `xml:"StoreCategoryName"`
}

type Variations struct {
	Pictures struct {
		VariationSpecificName struct {
		} `xml:"VariationSpecificName"`
		VariationSpecificPictureSet struct {
			PictureURL struct {
			} `xml:"PictureURL"`
			VariationSpecificValue struct {
			} `xml:"VariationSpecificValue"`
		} `xml:"VariationSpecificPictureSet"`
	} `xml:"Pictures"`
	Variation struct {
		DiscountPriceInfo struct {
			MadeForOutletComparisonPrice struct {
			} `xml:"MadeForOutletComparisonPrice"`
			MinimumAdvertisedPrice struct {
			} `xml:"MinimumAdvertisedPrice"`
			MinimumAdvertisedPriceExposure struct {
			} `xml:"MinimumAdvertisedPriceExposure"`
			OriginalRetailPrice struct {
			} `xml:"OriginalRetailPrice"`
			SoldOffeBay struct {
			} `xml:"SoldOffeBay"`
			SoldOneBay struct {
			} `xml:"SoldOneBay"`
		} `xml:"DiscountPriceInfo"`
		Quantity struct {
		} `xml:"Quantity"`
		SKU struct {
		} `xml:"SKU"`
		StartPrice struct {
		} `xml:"StartPrice"`
		VariationProductListingDetails struct {
			EAN struct {
			} `xml:"EAN"`
			ISBN struct {
			} `xml:"ISBN"`
			ProductReferenceID struct {
			} `xml:"ProductReferenceID"`
			UPC struct {
			} `xml:"UPC"`
		} `xml:"VariationProductListingDetails"`
		VariationSpecifics struct {
			NameValueList struct {
				Name struct {
				} `xml:"Name"`
				Value struct {
				} `xml:"Value"`
			} `xml:"NameValueList"`
		} `xml:"VariationSpecifics"`
	} `xml:"Variation"`
	VariationSpecificsSet struct {
		NameValueList struct {
			Name struct {
			} `xml:"Name"`
			Value struct {
			} `xml:"Value"`
		} `xml:"NameValueList"`
	} `xml:"VariationSpecificsSet"`
}

type VATDetails struct {
	BusinessSeller       bool    `xml:"BusinessSeller"`
	RestrictedToBusiness bool    `xml:"RestrictedToBusiness"`
	VATPercent           float32 `xml:"VATPercent"`
}

type Charity struct {
	CharityID       string  `xml:"CharityID"`
	CharityNumber   int     `xml:"CharityNumber"`
	DonationPercent float32 `xml:"DonationPercent"`
}

type ReturnPolicy struct {
	Description                           string `xml:"Description"`
	InternationalRefundOption             string `xml:"InternationalRefundOption"`
	InternationalReturnsAcceptedOption    string `xml:"InternationalReturnsAcceptedOption"`
	InternationalReturnsWithinOption      string `xml:"InternationalReturnsWithinOption"`
	InternationalShippingCostPaidByOption string `xml:"InternationalShippingCostPaidByOption"`
	RefundOption                          string `xml:"RefundOption"`
	ReturnsAcceptedOption                 string `xml:"ReturnsAcceptedOption"`
	ReturnsWithinOption                   string `xml:"ReturnsWithinOption"`
	ShippingCostPaidByOption              string `xml:"ShippingCostPaidByOption"`
}

type PictureDetails struct {
	GalleryDuration string   `xml:"GalleryDuration,omitempty"`
	GalleryType     string   `xml:"GalleryType,omitempty"`
	PhotoDisplay    string   `xml:"PhotoDisplay,omitempty"`
	PictureSource   string   `xml:"PictureSource,omitempty"`
	PictureURL      []string `xml:"PictureURL"`
}

type NameValueList struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type ItemSpecifics struct {
	NameValueList []NameValueList `xml:"NameValueList"`
}

type BestOfferDetails struct {
	BestOfferEnabled bool `xml:"BestOfferEnabled"`
}
