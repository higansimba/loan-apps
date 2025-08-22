package entity

type Collection string

const (
	CollectionPayments        Collection = "payments"
	CollectionLoans           Collection = "loans"
	CollectionUsers           Collection = "users"
	CollectionLoansSubmission Collection = "loans_submission"
	CollectionLoansApproval   Collection = "loans_approval"
)

func (c Collection) String() string {
	return string(c)
}
