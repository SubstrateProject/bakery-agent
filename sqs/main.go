package sqs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SendInput struct {
	QueueName string
	Body string
	Completed bool
}

func Send(params *SendInput) error {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	qm := NewQM(sess, params.QueueName)
	qm.sendMessage(params.Body, 1, "payload")
	return nil
}

type QueueMinder struct {
	AwsSession *session.Session
	SVC *sqs.SQS
	URL *string
}


func NewQM(awssess *session.Session, queueName string) (QueueMinder, error) {
	svc := sqs.New(awssess)
	params := &sqs.GetQueueUrlInput{
		QueueName:              aws.String(queueName), // Required
	}
	resp, err := svc.GetQueueUrl(params)

	if err != nil {
		log.Fatal(err)
	}

	qm := QueueMinder{
		AwsSession: aws,
		SVC: svc,
		URL: resp.QueueUrl,
	}
	return qm, nil
}

func (pqm *QueueMinder) sendMessage (body string, delay int, payload string) (*sqs.SendMessageOutput, error) {
	params := &sqs.SendMessageInput{
    MessageBody:  aws.String(body), // Required
    QueueUrl:     aws.String(&pqm.URL), // Required
    DelaySeconds: aws.Int64(delay),
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
    log.Fatal(err)
	}

	return resp, nil

}

func (pqm *QueueMinder) waitForMessage (signifier string, timeout int, pollrate int, totalpoll int) error {
	return nil
}
