package main

import (
	_ "github.com/looplab/logspout-logstash"
	_ "github.com/gliderlabs/logspout/httpstream"
    _ "github.com/gliderlabs/logspout/routesapi"
	_ "github.com/remind101/logspout-kinesis"
	_ "github.com/gliderlabs/logspout/transports/tcp"
	_ "github.com/gliderlabs/logspout/transports/udp"
	_ "github.com/gliderlabs/logspout/transports/tls"
)
