/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-kinesis/authenticater"
	"github.com/Appkube-awsx/awsx-kinesis/client"
	"github.com/Appkube-awsx/awsx-kinesis/commands/kinesiscmd"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxKinesisCmd = &cobra.Command{
	Use:   "Kinesis",
	Short: "get Kinesis Details command gets resource counts",
	Long:  `get Kinesis Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command kinesis started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getKinesisList(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func getKinesisList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*kinesis.ListStreamsOutput, error) {
	log.Println("Getting kinesis list summary")
	kinesisClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &kinesis.ListStreamsInput{}
	streamList, err := kinesisClient.ListStreams(input)
	if err != nil {
		log.Fatalln("Error: in getting kinesis list", err)
	}
	log.Println(streamList)
	return streamList, err
}

//func GetConfig(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
//	return getLambdaList(region, crossAccountRoleArn, accessKey, secretKey)
//}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxKinesisCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxKinesisCmd.AddCommand(kinesiscmd.GetConfigDataCmd)
	AwsxKinesisCmd.AddCommand(kinesiscmd.GetCostDataCmd)
	AwsxKinesisCmd.AddCommand(kinesiscmd.GetCostSpikeCmd)
	AwsxKinesisCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKinesisCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxKinesisCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKinesisCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKinesisCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKinesisCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKinesisCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
