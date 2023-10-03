package kinesiscmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "getConfigData command gets kinesis config details",
	Long:  `getConfigData command gets kinesis config details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getConfigData started")
		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			streamName, _ := cmd.Flags().GetString("streamName")
			streamArn, _ := cmd.Flags().GetString("streamArn")

			if streamName == "" {
				log.Fatalln("stream name not provided. program exit")
				return
			} else if streamArn == "" {
				log.Fatalln("stream arn not provided. program exit")
				return
			}
			GetKinesisDetail(streamName, streamArn, *clientAuth)

		}
	},
}

func GetKinesisDetail(streamName string, streamArn string, auth client.Auth) (*kinesis.DescribeStreamOutput, error) {
	log.Println("Getting Kinesis  data")
	client := client.GetClient(auth, client.KINESIS_CLIENT).(*kinesis.Kinesis)

	input := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
		//StreamARN:  aws.String(streamArn),
	}

	if streamArn != "" {
		input.SetStreamARN(streamArn)
	}
	kinesisData, err := client.DescribeStream(input)

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

	GetConfigDataCmd.Flags().StringP("streamArn", "a", "", "kinesis  stream arn")
	if err := GetConfigDataCmd.MarkFlagRequired("streamArn"); err != nil {
		fmt.Println("--streamArn is required", err)
	}
}
