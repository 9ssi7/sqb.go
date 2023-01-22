package sqb_go

type Order string

const (
	asc  Order = "ASC"
	desc Order = "DESC"
)

type order struct {
	ASC  Order
	DESC Order
}

var Orders = order{
	ASC:  asc,
	DESC: desc,
}
