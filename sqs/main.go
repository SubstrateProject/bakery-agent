package sqs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type QueueMinder struct {
	AwsSession *session.Session
	SVC *sqs.SQS
	URL *string
}


func NewQM(aws *session.Session, queue_name string) (ProvisionQueueMinder, error) {
	sqs := sqs.New(aws)
	params := &sqs.GetQueueUrlInput{
    QueueName:              aws.String(queue_name), // Required
	}
	resp, err := svc.GetQueueUrl(params)

	if err != nil {
    log.Error(err)
		return
	}

	return QueueMinder{
		AwsSession: aws,
		SVC: sqs,
		URL: resp.QueueUrl,
	}
}

func (pqm *ProvisionQueueMinder) sendMessage (body string, delay int, payload string, ) (*sqs.SendMessageOutput, error) {
	params := &sqs.SendMessageInput{
    MessageBody:  aws.String(body), // Required
    QueueUrl:     aws.String(pqm.URL), // Required
    DelaySeconds: aws.Int64(ds),
    MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Key": { // Required
				DataType: aws.String("String"), // Required
				BinaryListValues: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				BinaryValue: []byte("PAYLOAD"),
				StringListValues: []*string{
					aws.String("String"), // Required
					// More values...
				},
				StringValue: aws.String("String"),
			},
			// More values...
    },
	}

	resp, err := pqm.SVC.SendMessage(params)
	if err != nil {
    log.Error(err)
	}

	return resp, nil

}

func (pqm *ProvisionQueueMinder) waitForMessage (signifier string, timeout int, pollrate int, totalpoll int) error {

}
