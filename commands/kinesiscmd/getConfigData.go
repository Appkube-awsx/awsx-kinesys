/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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
			scalingType, _ := cmd.Flags().GetString("scalingType")
			targetShardCount, _ := cmd.Flags().GetInt64("targetShardCount")
			if streamName != "" {
				getKinesisDetail(region, crossAccountRoleArn, acKey, secKey, streamName, scalingType, targetShardCount, externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}
			if scalingType != "" {
				getKinesisDetail(region, crossAccountRoleArn, acKey, secKey, streamName, scalingType, targetShardCount, externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}

		}
	},
}

func getKinesisDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, streamName string, scalingType string, targetShardCount int64, externalId string) (*kinesis.UpdateShardCountOutput, error) {
	log.Println("Getting Kinesis  data")
	kinesisClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &kinesis.UpdateShardCountInput{
		StreamName:       aws.String(streamName),
		ScalingType:      aws.String(scalingType),
		TargetShardCount: aws.Int64(targetShardCount),
	}

	kinesisData, err := kinesisClient.UpdateShardCount(input)

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
}
