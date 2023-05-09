package kinesiscmd

import (
	"log"

	"github.com/Appkube-awsx/awsx-kinesis/authenticater"
	"github.com/Appkube-awsx/awsx-kinesis/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

// getKinesisCostDataCmd represents the getConfigData command
var GetKinesisCostDataCmd = &cobra.Command{
	Use:   "getKinesisCostData",
	Short: "A brief description of getKinesisCost command",
	Long:  `getKinesisCostData`,
	Run: func(cmd *cobra.Command, args []string) {
                 log.Println("Command getKinesis CostData started")
		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {

			getKinesisCostDetail(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func getKinesisCostDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*costexplorer.GetCostAndUsageOutput, error) {
	log.Println("Getting kinesis cost data")
	costClient := client.GetCostClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2022-06-01"),
			End:   aws.String("2022-07-01"),
		},
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
			aws.String("BLENDED_COST"),
			aws.String("AMORTIZED_COST"),
			aws.String("NET_AMORTIZED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("REGION"),
			},
		},
		Granularity: aws.String("MONTHLY"),
		Filter: &costexplorer.Expression{
			And: []*costexplorer.Expression{
				{
					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("SERVICE"),
						Values: []*string{
							aws.String("Amazon kinesis Service"),
						},
					},
				},
				{

					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("RECORD_TYPE"),
						Values: []*string{
							aws.String("Credit"),
						},
					},
				},
			},
		},
	}

	costData, err := costClient.GetCostAndUsage(input)
	if err != nil {
		log.Fatalln("Error: in getting cost data", err)
	}

	log.Println(costData)
	return costData, err
}

func init() {

}
