package requests

type Text struct {
	DistributionProfile     string  `json:"DistributionProfile"`
	Language          		string  `json:"Language"`
	DistributionProfileName	string  `json:"DistributionProfileName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
