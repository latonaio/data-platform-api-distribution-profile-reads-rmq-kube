package requests

type DistributionProfile struct {
	DistributionProfile	string	`json:"DistributionProfile"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
