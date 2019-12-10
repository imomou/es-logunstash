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

	is.Purge("logstash-prod-logstore-globalcloudtraillog-", 60)

	is.Purge("logstash-prod-logfarm-applicationlogs-", 7)
	is.Purge("logstash-elb-", 7)

	is.Purge("logstash-cf-access-", 3)
	is.Purge("logstash-s3-access-", 5)
	is.Purge("logstash-prod-logstore-vpctrafficlog-", 3)
	is.Purge("logstash-prod-logstore-lxsyslog-", 7)
	is.Purge("logstash-prod-snowplow-enrichedstream-1maltgsi3y3xd-", 3)
	is.Purge("logstash-prod-logstore-vpctrafficlog-ojme7t4aky29-", 3)
	is.Purge("logstash-prod-logstore-lxdockerlog-dhytnf0s89m1-", 3)

	fmt.Println("Operation completed...")

	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
	return fmt.Sprintf("Operation completed..."), nil
}
