package enums

import "errors"

type TransactionType uint

const (
	NormalPurchase TransactionType = iota + 1
	PurchaseWithInstallments
	Withdrawal
	CreditVoucher
)

var typeToString = map[TransactionType]string{
	NormalPurchase:           "Normal Purchase",
	PurchaseWithInstallments: "Purchase With Installments",
	Withdrawal:               "Withdrawal",
	CreditVoucher:            "Credit Voucher",
}

var typeToSign = map[TransactionType]int64{
	NormalPurchase:           -1,
	PurchaseWithInstallments: -1,
	Withdrawal:               -1,
	CreditVoucher:            1,
}

func GetAllTransactionTypes() []TransactionType {
	return []TransactionType{
		NormalPurchase,
		PurchaseWithInstallments,
		Withdrawal,
		CreditVoucher,
	}
}

func IsValidTransactionType(transactionType int64) error {
	for _, v := range GetAllTransactionTypes() {
		if v == TransactionType(transactionType) {
			return nil
		}
	}
	return errors.New("invalid transaction type")
}

func GetTransactionTypeFromInt(transactionType int64) (*TransactionType, error) {
	for k, v := range typeToString {
		if v == typeToString[TransactionType(transactionType)] {
			return &k, nil
		}

	}
	return nil, errors.New("Transaction type not found")
}

func (t TransactionType) String() string {
	return typeToString[t]
}

func (t TransactionType) Sign() int64 {
	return typeToSign[t]
}

func (t TransactionType) Int() *int64 {
	val := int64(t)
	return &val
}
