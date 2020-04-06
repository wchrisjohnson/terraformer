// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Input data for a TaskCompleted response to a decision task.
type RespondDecisionTaskCompletedInput struct {
	_ struct{} `type:"structure"`

	// The list of decisions (possibly empty) made by the decider while processing
	// this decision task. See the docs for the Decision structure for details.
	Decisions []Decision `locationName:"decisions" type:"list"`

	// User defined context to add to workflow execution.
	ExecutionContext *string `locationName:"executionContext" type:"string"`

	// The taskToken from the DecisionTask.
	//
	// taskToken is generated by the service and should be treated as an opaque
	// value. If the task is passed to another process, its taskToken must also
	// be passed. This enables it to provide its progress and respond with results.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RespondDecisionTaskCompletedInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RespondDecisionTaskCompletedInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RespondDecisionTaskCompletedInput"}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}
	if s.Decisions != nil {
		for i, v := range s.Decisions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Decisions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondDecisionTaskCompletedOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RespondDecisionTaskCompletedOutput) String() string {
	return awsutil.Prettify(s)
}

const opRespondDecisionTaskCompleted = "RespondDecisionTaskCompleted"

// RespondDecisionTaskCompletedRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Used by deciders to tell the service that the DecisionTask identified by
// the taskToken has successfully completed. The decisions argument specifies
// the list of decisions made while processing the task.
//
// A DecisionTaskCompleted event is added to the workflow history. The executionContext
// specified is attached to the event in the workflow execution history.
//
// Access Control
//
// If an IAM policy grants permission to use RespondDecisionTaskCompleted, it
// can express permissions for the list of decisions in the decisions parameter.
// Each of the decisions has one or more parameters, much like a regular API
// call. To allow for policies to be as readable as possible, you can express
// permissions on decisions as if they were actual API calls, including applying
// conditions to some parameters. For more information, see Using IAM to Manage
// Access to Amazon SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using RespondDecisionTaskCompletedRequest.
//    req := client.RespondDecisionTaskCompletedRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RespondDecisionTaskCompletedRequest(input *RespondDecisionTaskCompletedInput) RespondDecisionTaskCompletedRequest {
	op := &aws.Operation{
		Name:       opRespondDecisionTaskCompleted,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RespondDecisionTaskCompletedInput{}
	}

	req := c.newRequest(op, input, &RespondDecisionTaskCompletedOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RespondDecisionTaskCompletedRequest{Request: req, Input: input, Copy: c.RespondDecisionTaskCompletedRequest}
}

// RespondDecisionTaskCompletedRequest is the request type for the
// RespondDecisionTaskCompleted API operation.
type RespondDecisionTaskCompletedRequest struct {
	*aws.Request
	Input *RespondDecisionTaskCompletedInput
	Copy  func(*RespondDecisionTaskCompletedInput) RespondDecisionTaskCompletedRequest
}

// Send marshals and sends the RespondDecisionTaskCompleted API request.
func (r RespondDecisionTaskCompletedRequest) Send(ctx context.Context) (*RespondDecisionTaskCompletedResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RespondDecisionTaskCompletedResponse{
		RespondDecisionTaskCompletedOutput: r.Request.Data.(*RespondDecisionTaskCompletedOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RespondDecisionTaskCompletedResponse is the response type for the
// RespondDecisionTaskCompleted API operation.
type RespondDecisionTaskCompletedResponse struct {
	*RespondDecisionTaskCompletedOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RespondDecisionTaskCompleted request.
func (r *RespondDecisionTaskCompletedResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}