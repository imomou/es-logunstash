package main

import (
	"fmt"
)

func main() {
	esc := &esClient{
		Host: "http://internal-prod-elas-clusterg-e3j9z3h937gx-2505122.ap-southeast-2.elb.amazonaws.com:9200",
	}
	is := &indicesService{
		ec: *esc,
	}

	// prefixes := []string{
	// 	"logstash-elb-",
	// 	"logstash-billing-",
	// 	"logstash-cf-access-",
	// 	"logstash-prod-logfarm-applicationlogs-",
	// 	"logstash-prod-logstore-globalcloudtraillog-",
	// 	"logstash-prod-logstore-lxauthlog-mvqrjxk9bzkd-",
	// 	"logstash-prod-logstore-lxdockerlog-dhytnf0s89m1-",
	// 	"logstash-prod-logstore-lxsyslog-1uhlj0yn1qf1c-",
	// 	"logstash-prod-logstore-vpctrafficlog-ojme7t4aky29-",
	// 	"logstash-prod-snowplow-enrichedstream-1maltgsi3y3xd-",
	// 	"logstash-s3-access-",
	// }

	// is.ReportByIndexGrouping(prefixes)

	is.Purge("logstash-prod-logstore-globalcloudtraillog-", 28)

	is.Purge("logstash-prod-logfarm-applicationlogs-", 14)
	is.Purge("logstash-elb-", 14)

	is.Purge("logstash-cf-access-", 7)
	is.Purge("logstash-s3-access-", 7)
	is.Purge("logstash-prod-snowplow-enrichedstream-1maltgsi3y3xd-", 7)
	is.Purge("logstash-prod-logstore-vpctrafficlog-ojme7t4aky29-", 7)
	is.Purge("logstash-prod-logstore-lxdockerlog-dhytnf0s89m1-", 7)

	fmt.Println("Operation completed...")
}
