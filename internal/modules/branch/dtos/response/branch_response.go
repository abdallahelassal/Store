package response


type BranchResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type BranchListResponse struct {
	Data  []*BranchResponse 	`json:"branches"`
	Total int64 				`json:"total"`
}