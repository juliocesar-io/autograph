// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package gamelift

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// Welcome to the Amazon GameLift API Reference. Amazon GameLift is a managed
// Amazon Web Services (AWS) service for developers who need a scalable, server-based
// solution for multiplayer games. Amazon GameLift provides setup and deployment
// of game servers, and handles infrastructure scaling and session management.
//
// This reference describes the low-level service API for GameLift. You can
// call this API directly or use the AWS SDK (https://aws.amazon.com/tools/#sdk)
// for your preferred language. The AWS SDK includes a set of high-level GameLift
// actions multiplayer game sessions. Alternatively, you can use the AWS command-line
// interface (https://aws.amazon.com/cli/) (CLI) tool, which includes commands
// for GameLift. For administrative actions, you can also use the Amazon GameLift
// console.
//
// More Resources
//
//   Amazon GameLift Developer Guide (http://docs.aws.amazon.com/gamelift/latest/developerguide/):
// Learn more about GameLift features and how to use them   Lumberyard and GameLift
// Tutorials (https://gamedev.amazon.com/forums/tutorials): Get started fast
// with walkthroughs and sample projects  GameDev Blog (https://aws.amazon.com/blogs/gamedev/):
// Stay up to date with new features and techniques  GameDev Forums (https://gamedev.amazon.com/forums/spaces/123/gamelift-discussion.html):
// Connect with the GameDev community  Manage Games and Players Through GameLift
//
// Call these actions from your game clients and/or services to create and
// manage multiplayer game sessions and player sessions.
//
//   Game sessions  CreateGameSession DescribeGameSessions DescribeGameSessionDetails
// UpdateGameSession SearchGameSessions    Player sessions  CreatePlayerSession
// CreatePlayerSessions DescribePlayerSessions    Other actions:  GetGameSessionLogUrl
//    Set Up and Manage Game Servers
//
// Use these administrative actions to configure GameLift to host your game
// servers. When setting up GameLift, you'll need to (1) configure a build for
// your game and upload build files, and (2) set up one or more fleets to host
// game sessions. Once you've created and activated a fleet, you can assign
// aliases to it, scale capacity, track performance and utilization, etc.
//
//   Manage your builds:  ListBuilds CreateBuild DescribeBuild UpdateBuild
// DeleteBuild RequestUploadCredentials    Manage your fleets:  ListFleets CreateFleet
// Describe fleets:  DescribeFleetAttributes DescribeFleetCapacity DescribeFleetPortSettings
// DescribeFleetUtilization DescribeEC2InstanceLimits DescribeFleetEvents DescribeRuntimeConfiguration
//   Update fleets:  UpdateFleetAttributes UpdateFleetCapacity UpdateFleetPortSettings
// UpdateRuntimeConfiguration   DeleteFleet    Manage fleet aliases:  ListAliases
// CreateAlias DescribeAlias UpdateAlias DeleteAlias ResolveAlias    Manage
// autoscaling:  PutScalingPolicy DescribeScalingPolicies DeleteScalingPolicy
//    To view changes to the API, see the GameLift Document History (http://docs.aws.amazon.com/gamelift/latest/developerguide/doc-history.html)
// page.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type GameLift struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "gamelift"

// New creates a new instance of the GameLift client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a GameLift client from just a session.
//     svc := gamelift.New(mySession)
//
//     // Create a GameLift client with additional configuration
//     svc := gamelift.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *GameLift {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *GameLift {
	svc := &GameLift{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-10-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "GameLift",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a GameLift operation and runs any
// custom request initialization.
func (c *GameLift) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
