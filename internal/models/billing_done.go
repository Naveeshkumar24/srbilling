package models

import "net/http"

type BillingPo struct {
	ID              int     `json:"id"`
	Timestamp       string  `json:"timestamp"`
	EnggName        string  `json:"engg_name"`
	Supplier        string  `json:"supplier"`
	BillNo          string  `json:"bill_no"`
	BillDate        string  `json:"bill_date"`
	CustomerName    string  `json:"customer_name"`
	CustomerPoNo    string  `json:"customer_po_no"`
	CustomerPoDate  string  `json:"customer_po_date"`
	ItemDescription string  `json:"item_description"`
	BilledQty       int     `json:"billed_qty"`
	Unit            string  `json:"unit"`
	NetValue        float64 `json:"net_value"`
	CGST            float64 `json:"cgst"`
	IGST            float64 `json:"igst"`
	Totaltax        float64 `json:"total_tax"`
	Gross           float64 `json:"gross"`
	DispatchThrough string  `json:"dispatch_through"`
}
type BillingPoDropDown struct {
	EnggName     string `json:"engg_name"`
	Supplier     string `json:"supplier_name"`
	CustomerName string `json:"customer_name"`
	Unit         string `json:"unit_name"`
}

type BillingPoInterface interface {
	FetchDropDown() ([]BillingPoDropDown, error)
	SubmitFormBillingPoData(billingPo BillingPo) error
	FetchBillingPoData(r *http.Request) ([]BillingPo, error)
	UpdateBillingPoData(customerPo BillingPo) error
	DeleteBillingPoData(id int) error
}
