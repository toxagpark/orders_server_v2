package dto

type Order struct {
	OrderUID          string `gorm:"primaryKey"`
	TrackNumber       string
	Entry             string
	Delivery          Delivery `gorm:"embedded;embeddedPrefix:delivery_"`
	Payment           Payment  `gorm:"embedded;embeddedPrefix:payment_"`
	Items             []Item   `gorm:"foreignKey:OrderUID;references:OrderUID"`
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string
	SMID              int
	DateCreated       string
	OOFShard          string
}

type Delivery struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Payment struct {
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDT    int64
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

type Item struct {
	ID          uint `gorm:"primaryKey"`
	ChrtID      int
	TrackNumber string
	Price       int
	RID         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  int
	NmID        int
	Brand       string
	Status      int
	OrderUID    string
}
