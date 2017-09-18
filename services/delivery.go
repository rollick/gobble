package services

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/rollick/decimal"
)

// Delivery is a delivery object
// http://apidocs.bubblepost.eu/get_delivery/#example
type Delivery struct {
	Type             string          `json:"type"`
	Commnt           string          `json:"comment"`
	TrackingNumber   string          `json:"tracking_number"`
	Reference        string          `json:"reference"`
	Status           int             `json:"status"`
	StatusText       string          `json:"status_text"`
	AbortReason      string          `json:"abort_reason"`
	TripCode         string          `json:"trip_code"`
	TimesRescheduled int             `json:"times_rescheduled"`
	ETA              *time.Time      `json:"eta"`
	DueDates         []*DueDateOther `json:"due_dates"`
	Items            []*Item         `json:"items"`
	Recipient        struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Phone           string `json:"phone"`
		Company         string `json:"company"`
		ShippingAddress struct {
			Address   string           `json:"address"`
			Bus       string           `json:"bus"`
			Latitude  *decimal.Decimal `json:latitude`
			Longitude *decimal.Decimal `json:longitude`
		} `json:"shipping_address"`
	} `json:"recipient"`
	Signature struct {
		Base64   string     `json:"base64"`
		SignedBy string     `json:"signed_by"`
		SignedAt *time.Time `json:"signed_at"`
	} `json:"signature"`
	History struct {
		Created    *time.Time `json:"created"`
		Status     int        `json:"status"`
		StatusText string     `json:"status_text"`
	} `json:"history"`
}

// DeliveryRequest is a delivery request object
// http://apidocs.bubblepost.eu/post_delivery/#example
type DeliveryRequest struct {
	Type           string     `json:"type"`
	Commnt         string     `json:"comment"`
	TrackingNumber string     `json:"tracking_number"`
	Reference      string     `json:"reference"`
	DueDates       []*DueDate `json:"due_dates"`
	Items          []*Item    `json:"items"`
	Recipient      *Recipient `json:"recipient"`
}

type Recipient struct {
	Name       string `json:"name"`
	Company    string `json:"company"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Bus        string `json:"bus"`
	City       string `json:"city"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PhoneExtra string `json:"phone_extra"`
	Email      string `json:"email"`
}

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Size  string `json:"size"` // small, medium, large, xlarge, xxlarge
}

type DueDate struct {
	Date     string `json:"date"`
	TimeSlot int    `json:"time_slot"`
}

type DueDateOther struct {
	Date     string `json:"date"`
	TimeSlot []int  `json:"time_slot"`
}

// DeliveryService provides methods for creating and reading deliveries.
type DeliveryService struct {
	sling *sling.Sling
}

type CancelResponse struct {
	Message        string `json:"msg"`
	TrackingNumber string `json:"tracking_number"`
}

// NewDeliveryService returns a new DeliveryService.
func NewDeliveryService(accessToken string) *DeliveryService {
	// Create Bubble Post API client
	client := NewClient(accessToken)

	return &DeliveryService{
		sling: client,
	}
}

// List returns the deliveries based on a list of tracking numbers
func (s *DeliveryService) List(params []*TrackingParam) ([]*Delivery, *http.Response, error) {
	dr := []*Delivery{}
	bpError := new(BubblePostError)
	resp, err := s.sling.New().Path("deliveries").QueryStruct(params).Receive(&dr, bpError)
	if err == nil {
		err = bpError
	}

	return dr, resp, err
}

// Fetch returns a delivery based on a tracking number
func (s *DeliveryService) Fetch(param *TrackingParam) (Delivery, *http.Response, error) {
	dr := new(Delivery)
	bpError := new(BubblePostError)
	resp, err := s.sling.New().Path("delivery").QueryStruct(param).Receive(dr, bpError)
	if err == nil {
		err = bpError
	}
	return *dr, resp, err
}

// Create will create a single new delivery
func (s *DeliveryService) Create(d *DeliveryRequest) (DeliveryRequest, *http.Response, error) {
	dr := new(DeliveryRequest)
	bpError := new(BubblePostError)
	resp, err := s.sling.New().Post("delivery").BodyJSON(d).Receive(dr, bpError)
	if err == nil {
		err = bpError
	}
	return *dr, resp, err
}

// Cancel will cancel a delivery based on a tracking number
func (s *DeliveryService) Cancel(d *DeliveryRequest) (CancelResponse, *http.Response, error) {
	cr := new(CancelResponse)
	bpError := new(BubblePostError)
	resp, err := s.sling.New().Post("delivery").BodyJSON(d).Receive(cr, bpError)
	if err == nil {
		err = bpError
	}
	return *cr, resp, err
}
