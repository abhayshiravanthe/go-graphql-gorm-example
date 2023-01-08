package model

type Job struct {
	ID               string  `json:"id"`
	Text             string  `json:"text"`
	Description      string  `json:"description"`
	IsArchived       bool    `json:"isArchived"`
	IsDeleted        bool    `json:"isDeleted"`
	Tags             *string `json:"tags"`
	Skills           *string `json:"skills"`
	DeadlineDate     *string `json:"deadlineDate"`
	IsPermanent      *bool   `json:"isPermanent"`
	IsContract       *bool   `json:"isContract"`
	ContractDuration *string `json:"contractDuration"`
	Salary           *string `json:"salary"`
	Experience       *string `json:"experience"`
	Domain           *string `json:"domain"`
	Education        *string `json:"education"`
	Gender           *string `json:"gender"`
	User             *User   `json:"user"`
}
