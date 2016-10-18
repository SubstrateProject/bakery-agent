package main

import (
	"github.com/SubstrateProject/bakery-agent/sqs"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(
		"bakery-agent",
		"bakery-agent: an ami baking helper",
	).Version("0.1")
)

var (
	dlCommand   = app.Command("grab", "Download and unpack a tarball")
	sendCommand = app.Command("send", "send a message to a queue")
	recvCommand = app.Command("recv", "recv a message from a queue")
	runCommand  = app.Command("run", "Do a provision run")
)

var (
	sendQueueName = sendCommand.Flag("queueName", "name of queue to send to").Default("substrate-provisioner-queue")
	recvQueueName = recvCommand.Flag("queueName", "name of queue to recv to").Default("substrate-provisioner-queue")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case sendCommand.FullCommand():
		err := sqs.Send()
		app.FatalIfError(err, "send")
	case recvCommand.FullCommand():
		err := sqs.Recv()
		app.FatalIfError(err, "recv")
	}
}
