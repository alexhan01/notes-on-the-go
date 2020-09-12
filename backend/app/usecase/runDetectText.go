package usecase

// import (
// 	entity/rawimage
// 	"fmt"
// 	"context"
// 	"github.com/aws/aws-lambda-go/lambda"
//	"os"
// 	import "github.com/aws/aws-sdk-go/service/rekognition"
// )

// HandleLambdaEvent takes RawImage, runs AWS Rekog store it temp in S3 bucket and return S3 bucket address in String
func HandleLambdaEvent(event RawImage) (string, error) {
	retrieveFront(RawImage)
}

// RunDetectText invokes lambda function to run DetectText with AWS Rekognition
func runDetectText() {
	lambda.Start(HandleLambdaEvent)
	

}

// helper function to retrieve image from dummy file
func retrieveFront() {
	location := RawImage.location
	rekognition(location)
}

//passes image into rekognition()
func rekognition(string location) {

	// Read image from file that already exists
	existingImageFile, err := os.Open(location)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	//do the acutal shit
	rekognition := new AWS.Rekognition()
	return rekognition.DetectText(existingImageFile)// idk how to write the error handling in go


}
