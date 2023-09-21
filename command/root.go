package command

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"log"

	"github.com/Appkube-awsx/awsx-kinesis/command/kinesiscmd"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

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
		log.Fatalln("Error: in getting kinesis streams", err)
	}
	log.Println(streamList)
	return streamList, err
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
