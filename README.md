- [What is awsx-kinesis](#awsx-kinesis)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-cloudelements
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture.svg)


# How to write plugin subcommand 
Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug  / publish and integrate into the main commmand.

# How to build / Test
            go run main.go
                - Program will print Calling aws-kinesis on console 

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-cloudelements) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-cloudelements --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2

# what it does 
This subcommand implement the following functionalities -
   getConfigData - It  will get the resource count summary for a given AWS account id and region.

# command input
  --valutURL = URL location of vault - that stores credentials to call API
  --acountId = The AWS account id.
  --zone = AWS region
  --accessKey = Access key for the AWS account
  --secretKey = Secret Key for the Aws Account
  --crossAccountRoleArn = Cross Acount Rols Arn for the account.
  --external Id = The AWS External id.
  --streamName= Insert the stream name 
#  command output
{
        

# How to run 
  From main awsx command , it is called as follows:
 awsx-kinesis  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
  can simply run it as standalone command as
  go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
  
 #awsx-kinesis
kinesis extension

AWSX Commands for AWSX-kinesis Cli's :
 1)CMD used to get list of kinesis stream :
   ./awsx-kinesis --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
 2)CMD used to get Config data (metadata) of AWS kinesis stream :
   ./awsx-kinesis --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --streamName=<>
   

