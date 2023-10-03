package command

import (
	"encoding/json"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"log"

	"github.com/Appkube-awsx/awsx-kinesys/command/kinesiscmd"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

type KinesysObj struct {
	Stream interface{} `json:"stream"`
	Tags   interface{} `json:"tags"`
}

// AwsxKinesisCmd represents the base command when called without any subcommands
var AwsxKinesisCmd = &cobra.Command{
	Use:   "getKinesis",
	Short: "getKinesis command gets kinesis details",
	Long:  `getKinesis command gets kinesis details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command kinesis started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetKinesisList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

func GetKinesisList(auth client.Auth) (*kinesis.ListStreamsOutput, error) {
	log.Println("Getting kinesis streams")
	client := client.GetClient(auth, client.KINESIS_CLIENT).(*kinesis.Kinesis)
	input := &kinesis.ListStreamsInput{}
	streamList, err := client.ListStreams(input)
	if err != nil {
		log.Println("Error: in getting kinesis streams", err)
		return nil, err
	}
	log.Println(streamList)
	return streamList, err
}

func GetKinesisDetailsWithTags(auth client.Auth) (string, error) {
	log.Println("Getting kinesis details with tags")
	streamList, err := GetKinesisList(auth)
	if err != nil {
		log.Println("Error: in getting kinesis streams", err)
		return "", err
	}
	client := client.GetClient(auth, client.KINESIS_CLIENT).(*kinesis.Kinesis)
	allKinesysDetailWithTag := []KinesysObj{}
	for _, name := range streamList.StreamNames {
		kinesysDetail, err := kinesiscmd.GetKinesisDetail(*name, "", auth)
		if err != nil {
			log.Println("Error: in getting kinesys detail", err)
			continue
		}

		tagInput := &kinesis.ListTagsForStreamInput{
			StreamName: name,
		}
		tagOutput, err := client.ListTagsForStream(tagInput)
		if err != nil {
			log.Println("Error: in getting kinesys tag", err)
			continue
		}
		kinesysObj := KinesysObj{
			Stream: kinesysDetail,
			Tags:   tagOutput,
		}
		allKinesysDetailWithTag = append(allKinesysDetailWithTag, kinesysObj)
	}
	jsonData, err := json.Marshal(allKinesysDetailWithTag)
	log.Println(string(jsonData))
	return string(jsonData), err

}

// Execute adds all child command to the root command and sets flags appropriately.
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
	AwsxKinesisCmd.AddCommand(kinesiscmd.GetRecordCmd)

	AwsxKinesisCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKinesisCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKinesisCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxKinesisCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKinesisCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKinesisCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKinesisCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKinesisCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
