package main

import (
	"fmt"
	"log"
	"os"

	"breks/pkg/eks"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/sts"
)

func main() {

	// TOTO:
	// Elsewhere set environment variables - for development or testing only!
	// os.Setenv("AWS_ACCESS_KEY_ID", 		"your-access-key-id")
	// os.Setenv("AWS_SECRET_ACCESS_KEY", 	"your-secret-access-key")
	// // os.Setenv("AWS_SESSION_TOKEN", 	"your-session-token") 		// Optional, only if using temporary credentials

	// Retrieve credentials from environment variables
	region 				:= "us-west-2"
	accessKeyID 		:= os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey 	:= os.Getenv("AWS_SECRET_ACCESS_KEY")
	sessionToken 		:= os.Getenv("AWS_SESSION_TOKEN") // AWS_SESSION_TOKEN is optional


	// Create a new session in the us-west-1 region.
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(region)},
	// )
	// if err != nil {
	// 	log.Fatalf("ERROR: failed to create session, %v", err)
	// }

	// // Create a STS client from just a session.
	// svc := sts.New(sess)

	// // Assume an IAM role. Replace with the ARN of the role you want to assume.
	// roleToAssumeArn := os.Getenv("AWS_ROLE")
	// sessionName 	:= "MySsession1"
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	// // Call the AssumeRole API
	// resp, err := svc.AssumeRole(&sts.AssumeRoleInput{
	// 	RoleArn:         &roleToAssumeArn,
	// 	RoleSessionName: &sessionName,
	// })
	// if err != nil {
	// 	log.Fatalf("unable to assume role, %v", err)
	// }

	// // Now you have temporary credentials to use in the form of resp.Credentials
	// fmt.Println("Access Key ID:     ", *resp.Credentials.AccessKeyId)
	// fmt.Println("Secret Access Key: ", *resp.Credentials.SecretAccessKey)
	// fmt.Println("Session Token:     ", *resp.Credentials.SessionToken)

	// Create a new EKS client
	eksClient, err := eks.NewEKSClient(
							accessKeyID, // *resp.Credentials.AccessKeyId, //
							secretAccessKey, //*resp.Credentials.SecretAccessKey, //
							sessionToken,		// *resp.Credentials.SessionToken, //
							region) //sessionToken
	if err != nil {
		log.Fatalf("Failed to create EKS client: %v", err)
	}

	// List EKS clusters
	clusters, err := eksClient.ListClusters()
	if err != nil {
		log.Fatalf("Failed to list EKS clusters: %v", err)
	}

	fmt.Println("Clusters:", clusters)
}
