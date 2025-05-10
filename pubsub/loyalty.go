package pubsub

type PubsubLoyaltyTransactionMktRequest struct {
	JobID               string  `json:"job_id"`
	RequestTime         string  `json:"request_type"`
	BillingStore        string  `json:"billing_store"`
	Phone               string  `json:"phone"`
	NetFare             int64   `json:"net_fare"`
	Status              string  `json:"status"`
	Fare                int64   `json:"fare"`
	Discount            int64   `json:"discount"`
	Tax                 int64   `json:"tax"`
	ReferenceBillNumber string  `json:"reference_bill_number"`
	PaymentMethod       string  `json:"payment_method"`
	PromoCode           string  `json:"promo_code"`
	PromoDesc           string  `json:"promo_desc"`
	DiscountType        int     `json:"discount_type"`
	OrderID             string  `json:"order_id"`
	CreatedTime         string  `json:"created_time"`
	PickupTime          string  `json:"pickup_time"`
	DropoffTime         string  `json:"dropoff_time"`
	Identifier          string  `json:"identifier"`
	ServiceType         string  `json:"service_type"`
	OrderChannel        string  `json:"channel"`
	CarID               string  `json:"car_id"`
	CarType             string  `json:"car_type"`
	DriverID            string  `json:"driver_id"`
	PickupLocationName  string  `json:"pickup_location_name"`
	DropoffLocationName string  `json:"dropoff_location_name"`
	City                string  `json:"city"`
	BinCode             string  `json:"bin_code"`
	BookingType         string  `json:"booking_type"`
	FareType            string  `json:"fare_type"`
	PickupLandmark      string  `json:"pickup_landmark"`
	DropoffLandmark     string  `json:"dropoff_landmark"`
	MerchantID          string  `json:"merchant_id"`
	PickupLat           float64 `json:"pickup_lat"`
	DestinationLat      float64 `json:"destination_lat"`
	PickupLong          float64 `json:"pickup_long"`
	DestinationLong     float64 `json:"destination_long"`
	DriverName          string  `json:"driver_name"`
	Itop                string  `json:"itop"`
}

type PubsubLoyaltyMktResponse struct {
	Message string `json:"message"`
}
