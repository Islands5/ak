package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	var path, tagKey, tagValue string
	flag.StringVar(&path, "P", "empty", "input path of public key file")
	flag.StringVar(&path, "path", "empty", "input path of public key file")

	flag.StringVar(&tagKey, "tk", "empty", "input tag key")
	flag.StringVar(&tagKey, "tag-key", "empty", "input tag key")

	flag.StringVar(&tagValue, "tv", "empty", "input tag value")
	flag.StringVar(&tagValue, "tag-value", "empty", "input tag value")

	flag.Parse()

	fmt.Printf("path: %v\n", path)
	fmt.Printf("tag-key: %v\n", tagKey)
	fmt.Printf("tag-value: %v\n", tagValue)

	for _, ele := range getPublicIPAddresses(tagKey, tagValue) {
		fmt.Printf("IPAddresses: %v\n", ele)
	}

}

func getPublicIPAddresses(tagKey string, tagValue string) []string {
	var ipAddresses []string
	awsTag := "tag:" + tagKey
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err.Error())
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(awsTag),
				Values: []*string{
					aws.String(tagValue),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println(err.Error())
	}

	// reservation sets
	for idx := range resp.Reservations {
		// instances
		for _, inst := range resp.Reservations[idx].Instances {
			ipAddresses = append(ipAddresses, *inst.PublicIpAddress)
		}
	}

	return ipAddresses
}
