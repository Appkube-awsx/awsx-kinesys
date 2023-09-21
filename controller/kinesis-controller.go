package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-kinesys/command"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
)

func GetKinesisByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*kinesis.ListStreamsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetKinesisByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetKinesisByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*kinesis.ListStreamsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetKinesisByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetKinesisByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*kinesis.ListStreamsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetKinesisList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetKinesisStreams(clientAuth *client.Auth) (*kinesis.ListStreamsOutput, error) {
	response, err := command.GetKinesisList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
