package models

type Orders struct {
	Order_uid          string   `json:"order_uid" db:"order_uid"`
	Track_number       string   `json:"track_number" db:"track_number"`
	Entry              string   `json:"entry" db:"entry"`
	Delivery           Delivery `json:"delivery" db:"delivery"`
	Payment            Payment  `json:"payment" db:"payment"`
	Items              []Item   `json:"items" db:"items"`
	Locale             string   `json:"locale" db:"locale"`
	Internal_signature string   `json:"internal_signature" db:"internal_signature"`
	Customer_id        string   `json:"customer_id" db:"customer_id"`
	Delivery_service   string   `json:"delivery_service" db:"delivery_service"`
	Shardkey           string   `json:"shardkey" db:"shardkey"`
	Sm_id              int      `json:"sm_id" db:"sm_id"`
	Date_created       string   `json:"date_created" db:"date_created"`
	Oof_shard          string   `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	Transaction   string `json:"transaction" db:"transaction"`
	Request_id    string `json:"request_id" db:"request_id"`
	Currency      string `json:"currency" db:"currency"`
	Provider      string `json:"provider" db:"provider"`
	Amount        int    `json:"amount" db:"amount"`
	Payment_dt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank          string `json:"bank" db:"bank"`
	Delivery_cost int    `json:"delivery_cost" db:"delivery_cost"`
	Goods_total   int    `json:"goods_total" db:"goods_total"`
	Custom_fee    int    `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id" db:"chrt_id"`
	Track_number string `json:"rack_number" db:"rack_number"`
	Price        int    `json:"price" db:"price"`
	Rid          string `json:"rid" db:"rid"`
	Name         string `json:"name" db:"name"`
	Sale         int    `json:"sale" db:"sale"`
	Size         string `json:"size" db:"size"`
	Total_price  int    `json:"total_price" db:"total_price"`
	Nm_id        int    `json:"nm_id" db:"nm_id"`
	Brand        string `json:"brand" db:"brand"`
	Status       int    `json:"status" db:"status"`
}
