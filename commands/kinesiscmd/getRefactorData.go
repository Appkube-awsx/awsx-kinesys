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

// getRefactoringDataCmd represents the getRefactorData command
var GetRefactorDataCmd = &cobra.Command{
	Use:   "getRefactorData",
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
			streamArn, _ := cmd.Flags().GetString("streamArn")
			if streamName != "" {
				getKinesisRefactor(region, crossAccountRoleArn, acKey, secKey, streamName, streamArn, externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}
			if streamArn != "" {
				getKinesisRefactor(region, crossAccountRoleArn, acKey, secKey, streamName, streamArn, externalId)

			} else {
				log.Fatalln("stream not provided.Program exit")
			}

		}
	},
}

func getKinesisRefactor(region string, crossAccountRoleArn string, accessKey string, secretKey string, streamName string, streamArn string, externalId string) (*kinesis.DescribeStreamOutput, error) {
	kinesisClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	

	//Refactor code

	describeStreamInput := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
		StreamARN:  aws.String(streamArn),
	}

	describeStreamOutput, err := kinesisClient.DescribeStream(describeStreamInput)
	if err != nil {
		fmt.Println("Error describing stream:", err)
	}
	log.Println(describeStreamOutput)
	return describeStreamOutput, err

	shards := describeStreamOutput.StreamDescription.Shards
	for _, shard := range shards {
		shardIteratorType := "LATEST"

		getShardIteratorInput := &kinesis.GetShardIteratorInput{
			ShardId:           shard.ShardId,
			ShardIteratorType: aws.String(shardIteratorType),
			StreamName:        aws.String(streamName),
		}
		getShardIteratorOutput, err := kinesisClient.GetShardIterator(getShardIteratorInput)
		if err != nil {
			fmt.Println("Error getting shard iterator:", err)
			log.Println(getShardIteratorOutput)
			

		}

		shardIterator := getShardIteratorOutput.ShardIterator

		getRecordsInput := &kinesis.GetRecordsInput{
			ShardIterator: shardIterator,
		}
		getRecordsOutput, err := kinesisClient.GetRecords(getRecordsInput)
		if err != nil {
			fmt.Println("Error getting records:',err")
			log.Println(getRecordsOutput)

		}

		for _, record := range getRecordsOutput.Records {
			fmt.Println("Record:", string(record.Data))

		}
	}

}

func init() {
	GetRefactorDataCmd.Flags().StringP("streamName", "s", "", "kinesis  stream name")

	if err := GetRefactorDataCmd.MarkFlagRequired("streamName"); err != nil {
		fmt.Println("--streamName is required", err)
	}
	GetRefactorDataCmd.Flags().StringP("streamArn", "t", "", "kinesis  stream arn")

	if err := GetRefactorDataCmd.MarkFlagRequired("streamArn"); err != nil {
		fmt.Println("--streamArn is required", err)
	}
}
