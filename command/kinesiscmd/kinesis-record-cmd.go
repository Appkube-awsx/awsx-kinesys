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

// GetRecordCmd represents the GetRecordCmd command
var GetRecordCmd = &cobra.Command{
	Use:   "getRecord",
	Short: "getRecord command gets kinesis records",
	Long:  `getRecord command gets kinesis records of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getRecord started")
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
			GetKinesisRecords(streamName, streamArn, *clientAuth)

		}
	},
}

func GetKinesisRecords(streamName string, streamArn string, auth client.Auth) ([]*kinesis.Record, error) {
	describeStreamOutput, err := GetKinesisDetail(streamName, streamArn, auth)
	if err != nil {
		log.Fatalln("Error in getting kinesis detail", err)
		return nil, err
	}
	allRecords := []*kinesis.Record{}

	shards := describeStreamOutput.StreamDescription.Shards
	client := client.GetClient(auth, client.KINESIS_CLIENT).(*kinesis.Kinesis)
	shardIteratorType := "LATEST"

	for _, shard := range shards {

		shardIteratorInput := &kinesis.GetShardIteratorInput{
			ShardId:           shard.ShardId,
			ShardIteratorType: aws.String(shardIteratorType),
			StreamName:        aws.String(streamName),
		}

		shardIteratorOutput, err := client.GetShardIterator(shardIteratorInput)
		if err != nil {
			fmt.Println("Error in getting kinesis shard", err)
			return nil, err
		}
		shardIterator := shardIteratorOutput.ShardIterator

		recordsInput := &kinesis.GetRecordsInput{
			ShardIterator: shardIterator,
		}
		recordsOutput, err := client.GetRecords(recordsInput)
		if err != nil {
			fmt.Println("Error getting kinesis shard records', err")
			return nil, err
		}

		for _, record := range recordsOutput.Records {
			fmt.Println("Record:", string(record.Data))
			allRecords = append(allRecords, record)
		}
	}
	return allRecords, nil
}

func init() {
	GetRecordCmd.Flags().StringP("streamName", "s", "", "kinesis  stream name")
	if err := GetRecordCmd.MarkFlagRequired("streamName"); err != nil {
		fmt.Println("--streamName is required", err)
	}

	GetRecordCmd.Flags().StringP("streamArn", "a", "", "kinesis  stream arn")
	if err := GetRecordCmd.MarkFlagRequired("streamArn"); err != nil {
		fmt.Println("--streamArn is required", err)
	}
}
