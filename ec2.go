package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func ck(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
}

var (
	name    = flag.String("name", "", "vpc name")
	cidr    = flag.String("cidr", "10.0.0.0/16", "cidr network block")
	reg     = flag.String("reg", "us-west-2", "region")
	subnets = flag.String("subnets", "", "comma seperate list of subnets in cidr notation")
)

func main() {
	flag.Parse()
	sess := session.Must(session.NewSession(aws.NewConfig().
		WithMaxRetries(3),
	))

	svc := ec2.New(sess, aws.NewConfig().
		WithRegion(*reg),
	)

	input := &ec2.CreateVpcInput{
		CidrBlock: aws.String(*cidr),
	}

	result, err := svc.CreateVpc(input)
	if err != nil {
		log.Println(err)
		log.Fatalln("https://console.aws.amazon.com/iam/home#/security_credential")
		log.Fatalln(`set AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
set AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`)
	}
	vpc := result.Vpc
	key := "Name"
	svc.CreateTags(&ec2.CreateTagsInput{
		Tags: []*ec2.Tag{&ec2.Tag{Key: &key, Value: name}},
		Resources: []*string{vpc.VpcId},
	})

	for _, sn := range strings.Split(*subnets, ",") {
		x := strings.IndexAny(sn, "@")
		var zone *string
		if x != -1 {
			z := sn[x+1:]
			zone = &z
			sn = sn[:x]
		}
		res, err := svc.CreateSubnet(&ec2.CreateSubnetInput{
			CidrBlock:        &sn,
			AvailabilityZone: zone,
			VpcId: vpc.VpcId,
		})
		no(err)
		fmt.Printf("subnet: %s: %s\n", res)
	}
	fmt.Printf("%#v\n", result)
}
func no(err error){
	if err != nil{
		log.Fatalln(err)
	}
}
