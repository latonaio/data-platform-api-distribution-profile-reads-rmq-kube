package dpfm_api_output_formatter

import (
	"data-platform-api-distribution-profile-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToDistributionProfile(rows *sql.Rows) (*[]DistributionProfile, error) {
	defer rows.Close()
	distributionProfile := make([]DistributionProfile, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DistributionProfile{}

		err := rows.Scan(
			&pm.DistributionProfile,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &distributionProfile, nil
		}

		data := pm
		distributionProfile = append(distributionProfile, DistributionProfile{
			DistributionProfile:	data.DistributionProfile,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &distributionProfile, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.DistributionProfile,
			&pm.Language,
			&pm.DistributionProfileName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			DistributionProfile:		data.DistributionProfile,
			Language:          			data.Language,
			DistributionProfileName:	data.DistributionProfileName,
			CreationDate:				data.CreationDate,
			LastChangeDate:				data.LastChangeDate,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
