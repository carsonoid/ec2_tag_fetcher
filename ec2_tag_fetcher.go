package main // import "github.com/carsonoid/ec2_tag_fetcher"

import (
    "fmt"
    "os"
    "encoding/json"
    "flag"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/aws/ec2metadata"

)

func main() {
    // Flags
    humanPtr := flag.Bool("H", false, "Human-readable output")
    flag.Parse()

    // Create Metadata Client
    md := ec2metadata.New(session.New(), aws.NewConfig())

    // Sanity Check
    if is_ec2 := md.Available(); ! is_ec2 {
        fmt.Println("Not an ec2 instance")
        os.Exit(255)
    }

    // Get our current instance id
    instance_id, err := md.GetMetadata("instance-id")
    if err != nil {
        // TODO smart error
        panic(err)
    }

    // Get rour current instance region 
    region, err      := md.Region()
    if err != nil {
        // TODO smart error
        panic(err)
    }

    // Create an EC2 service object in the same region as current instance
    svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

    // Filter tags to just this instance
    params := &ec2.DescribeTagsInput{
        Filters: []*ec2.Filter{
            &ec2.Filter{
                Name: aws.String("resource-id"),
                Values: []*string{ aws.String(instance_id), },
            },
            &ec2.Filter{
                Name: aws.String("resource-type"),
                Values: []*string{ aws.String("instance"), },
            },
        },
    }

    // Call the DescribeTags Operation
    resp, err := svc.DescribeTags(params)
    if err != nil {
        // TODO smart error
        panic(err)
    }

    // Output appropriately
    if *humanPtr {
        fmt.Println("> Number of Tags : ", len(resp.Tags))
        for _, tag := range resp.Tags {
            fmt.Println(" - Key: ", *tag.Key)
            fmt.Println("   Value: ", *tag.Value)
        }
    } else {
        // Print out the json encoded list of tags
        if res, err:= json.Marshal(resp.Tags); err == nil {
            fmt.Printf("%s", res)
        } else {
            panic(err)
        }
    }
}
