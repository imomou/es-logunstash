package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	esc := &esClient{
		Host: "http://internal-prod-elas-clusterg-e3j9z3h937gx-2505122.ap-southeast-2.elb.amazonaws.com:9200",
	}
	is := &indicesService{
		ec: *esc,
	}

	is.Purge("logstash-prod-logstore-globalcloudtraillog-", 28)

	is.Purge("logstash-prod-logfarm-applicationlogs-", 14)
	is.Purge("logstash-elb-", 14)

	is.Purge("logstash-cf-access-", 7)
	is.Purge("logstash-s3-access-", 7)
	is.Purge("logstash-prod-snowplow-enrichedstream-1maltgsi3y3xd-", 7)
	is.Purge("logstash-prod-logstore-vpctrafficlog-ojme7t4aky29-", 7)
	is.Purge("logstash-prod-logstore-lxdockerlog-dhytnf0s89m1-", 7)

	fmt.Println("Operation completed...")

	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
	return fmt.Sprintf("Operation completed..."), nil
}
