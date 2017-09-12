package commands

type PaginationRequest struct {
	EntriesPerPage       string `xml:",omitempty"`
	OrdersPerPage        string `xml:",omitempty"`
	PageNumber           string `xml:",omitempty"`
	TotalNumberOfEntries int    `xml:",omitempty"`
	TotalNumberOfPages   int    `xml:",omitempty"`
}
