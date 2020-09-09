package usecase

// import (
// 	entity/rawimage
// 	"fmt"
// 	"context"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// HandleLambdaEvent takes RawImage, runs AWS Rekog store it temp in S3 bucket and return S3 bucket address in String
func HandleLambdaEvent(event RawImage) (string, error) {
	//stub
}

// RunDetectText invokes lambda function to run DetectText with AWS Rekognition
func runDetectText() {
	lambda.Start(HandleLambdaEvent)
}
