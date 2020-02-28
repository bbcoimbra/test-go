package main

type OperationType struct {
	OperationTypeID int64  `json:"account_id,omitempty"`
	Description     string `json:"description"`
}

var operation_types = []OperationType{
	{
		OperationTypeID: 1,
		Description:     "COMPRA A VISTA",
	},
	{
		OperationTypeID: 2,
		Description:     "COMPRA PARCELADA",
	},
	{
		OperationTypeID: 3,
		Description:     "SAQUE",
	},
	{
		OperationTypeID: 4,
		Description:     "PAGAMENTO",
	},
}
