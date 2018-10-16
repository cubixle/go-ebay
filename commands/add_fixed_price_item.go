package commands

import (
	"encoding/xml"

	"github.com/cubixle2/go-ebay/config"
	"github.com/cubixle2/go-ebay/types"
)

func NewAddFixedPriceItemRequest(cfg *config.Config, item *types.Item) *AddFixedPriceItemRequest {
	// condition is always new (for now).
	item.ConditionID = 1000
	item.Country = "GB"
	item.Currency = "GBP"
	item.Site = "UK"
	item.PaymentMethods = types.BuyerPaymentMethodCodeTypePayPal
	item.HitCounter = "NoHitCounter"
	item.ListingType = types.ListingTypeCodeTypeFixedPriceItem

	return &AddFixedPriceItemRequest{
		Item: *item,
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		Xmlns:       "urn:ebay:apis:eBLBaseComponents",
		DetailLevel: "ItemReturnDescription",
		Version:     "837",
	}
}

type AddFixedPriceItemRequest struct {
	Item                 types.Item           `xml:"Item"`
	Xmlns                string               `xml:"xmlns,attr"`
	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`
	DetailLevel          string
	Version              string
}

func (c AddFixedPriceItemRequest) CallName() string {
	return "AddFixedPriceItem"
}

// GetRequestBody returns bytes
func (c AddFixedPriceItemRequest) GetRequestBody() []byte {
	body, _ := xml.Marshal(c)
	return body
}

func (c AddFixedPriceItemRequest) SetToken(token string) {
	c.RequesterCredentials.EBayAuthToken = token
}

// AddFixedPriceItemResponse was generated 2018-09-11 14:04:26 by lukerodham on Lukes-MacBook-Pro.local.
type AddFixedPriceItemResponse struct {
	XMLName     xml.Name `xml:"AddFixedPriceItemResponse"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	Category2ID struct {
		Text string `xml:",chardata"`
	} `xml:"Category2ID"`
	CategoryID struct {
		Text string `xml:",chardata"`
	} `xml:"CategoryID"`
	DiscountReason struct {
		Text string `xml:",chardata"`
	} `xml:"DiscountReason"`
	EndTime struct {
		Text string `xml:",chardata"`
	} `xml:"EndTime"`
	Fees struct {
		Text string `xml:",chardata"`
		Fee  struct {
			Text string `xml:",chardata"`
			Fee  struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"Fee"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"Name"`
			PromotionalDiscount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"PromotionalDiscount"`
		} `xml:"Fee"`
	} `xml:"Fees"`
	ItemID struct {
		Text string `xml:",chardata"`
	} `xml:"ItemID"`
	ListingRecommendations struct {
		Text           string `xml:",chardata"`
		Recommendation struct {
			Text string `xml:",chardata"`
			Code struct {
				Text string `xml:",chardata"`
			} `xml:"Code"`
			FieldName struct {
				Text string `xml:",chardata"`
			} `xml:"FieldName"`
			Group struct {
				Text string `xml:",chardata"`
			} `xml:"Group"`
			Message struct {
				Text string `xml:",chardata"`
			} `xml:"Message"`
			Metadata struct {
				Text string `xml:",chardata"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"Name"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"Value"`
			} `xml:"Metadata"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"Type"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"Value"`
		} `xml:"Recommendation"`
	} `xml:"ListingRecommendations"`
	ProductSuggestions struct {
		Text              string `xml:",chardata"`
		ProductSuggestion struct {
			Text string `xml:",chardata"`
			EPID struct {
				Text string `xml:",chardata"`
			} `xml:"EPID"`
			Recommended struct {
				Text bool `xml:",chardata"`
			} `xml:"Recommended"`
			StockPhoto struct {
				Text string `xml:",chardata"`
			} `xml:"StockPhoto"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"Title"`
		} `xml:"ProductSuggestion"`
	} `xml:"ProductSuggestions"`
	SKU struct {
		Text string `xml:",chardata"`
	} `xml:"SKU"`
	StartTime struct {
		Text string `xml:",chardata"`
	} `xml:"StartTime"`
	Ack struct {
		Text string `xml:",chardata"`
	} `xml:"Ack"`
	Build struct {
		Text string `xml:",chardata"`
	} `xml:"Build"`
	CorrelationID struct {
		Text string `xml:",chardata"`
	} `xml:"CorrelationID"`
	DuplicateInvocationDetails struct {
		Text   string `xml:",chardata"`
		Status struct {
			Text string `xml:",chardata"`
		} `xml:"Status"`
	} `xml:"DuplicateInvocationDetails"`
	Errors struct {
		Text                string `xml:",chardata"`
		ErrorClassification struct {
			Text string `xml:",chardata"`
		} `xml:"ErrorClassification"`
		ErrorCode struct {
			Text string `xml:",chardata"`
		} `xml:"ErrorCode"`
		ErrorParameters struct {
			Text    string `xml:",chardata"`
			ParamID string `xml:"ParamID,attr"`
			Value   struct {
				Text string `xml:",chardata"`
			} `xml:"Value"`
		} `xml:"ErrorParameters"`
		LongMessage struct {
			Text string `xml:",chardata"`
		} `xml:"LongMessage"`
		SeverityCode struct {
			Text string `xml:",chardata"`
		} `xml:"SeverityCode"`
		ShortMessage struct {
			Text string `xml:",chardata"`
		} `xml:"ShortMessage"`
	} `xml:"Errors"`
	HardExpirationWarning struct {
		Text string `xml:",chardata"`
	} `xml:"HardExpirationWarning"`
	Message struct {
		Text string `xml:",chardata"`
	} `xml:"Message"`
	Timestamp struct {
		Text string `xml:",chardata"`
	} `xml:"Timestamp"`
	Version struct {
		Text string `xml:",chardata"`
	} `xml:"Version"`
}
