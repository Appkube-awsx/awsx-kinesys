package kinesiscmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-kinesis/authenticater"
	"github.com/Appkube-awsx/awsx-kinesis/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of getConfigData command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
                log.Println("Command getKinesisConfigData started")
		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			streamName, _ := cmd.Flags().GetString("streamName")
			streamArn, _ := cmd.Flags().GetString("streamArn")
			
			if streamName != "" {
				getKinesisDetail(region, crossAccountRoleArn, acKey, secKey, streamName, streamArn,  externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}
			if streamArn != "" {
				getKinesisDetail(region, crossAccountRoleArn, acKey, secKey, streamName, streamArn, externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}

		}
	},
}

func getKinesisDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, streamName string, streamArn string, externalId string) (*kinesis.DescribeStreamOutput, error) {
	log.Println("Getting Kinesis  data")
	kinesisClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &kinesis.DescribeStreamInput{
		StreamName:       aws.String(streamName),
		StreamARN: aws.String(streamArn),
	}

	kinesisData, err := kinesisClient.DescribeStream(input)

	if err != nil {
		log.Fatalln("Error: in getting kinesis data", err)
	}
	log.Println(kinesisData)
	return kinesisData, err

}

func init() {
	GetConfigDataCmd.Flags().StringP("streamName", "s", "", "kinesis  stream name")

	if err := GetConfigDataCmd.MarkFlagRequired("streamName"); err != nil {
		fmt.Println("--streamName is required", err)
	}
	GetConfigDataCmd.Flags().StringP("streamArn", "t", "", "kinesis  stream arn")

	if err := GetConfigDataCmd.MarkFlagRequired("streamArn"); err != nil {
		fmt.Println("--streamArn is required", err)
	}
}
