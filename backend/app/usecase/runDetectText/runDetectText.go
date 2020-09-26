package runDetectText

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/lambda"
)

// import (
// 	entity/rawimage
// 	"fmt"
// 	"context"
// 	"github.com/aws/aws-lambda-go/lambda"
//	"os"
// 	import "github.com/aws/aws-sdk-go/service/rekognition"
// )

type RawImage struct {
	//stub;
	Location string
}

// HandleLambdaEvent takes RawImage, runs AWS Rekog store it temp in S3 bucket and return S3 bucket address in String
func HandleLambdaEvent(ctx context.Context, rawimage RawImage) (string, error) {
	return retrieveFront(rawimage), nil
}

// RunDetectText invokes lambda function to run DetectText with AWS Rekognition
func runDetectText() {
	lambda.Start(HandleLambdaEvent)
}

// helper function to retrieve image from dummy file
func retrieveFront(RawImage RawImage) string {
	location := RawImage.Location
	return rekognition(location)
}

//passes image into rekognition()
func rekognition(location string) string {

	// Read image from file that already exists
	existingImageFile, err := os.Open(location)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// create new rekognition client

	mySession := session.Must(session.NewSession())
	// Create a Rekognition client from just a session.
	//svc := rekognition.New(mySession)

	svc := rekognition.new(mySession, aws.NewConfig().WithRegion("us-east-1"))
	img := svc.DetectText(existingImageFile)

	// idk how to write the error handling or do the rest

	return img

}
