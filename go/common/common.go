package common

import (
	"os"
	"time"

	"github.com/developerforce/pub-sub-api/go/proto"
)

var (
	// topic and subscription-related variables
	TopicName           = os.Getenv("TOPIC")
	ReplayPreset        = proto.ReplayPreset_EARLIEST
	ReplayId     []byte = nil
	Appetite     int32  = 5

	// gRPC server variables
	GRPCEndpoint    = "api.pubsub.salesforce.com:" + os.Getenv("PORT")
	GRPCDialTimeout = 5 * time.Second
	GRPCCallTimeout = 5 * time.Second

	// OAuth header variables
	GrantType    = os.Getenv("GRANT_TYPE")
	ClientId     = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	Username     = os.Getenv("USERNAME")
	Password     = os.Getenv("PASSWORD")

	// OAuth server variables
	OAuthEndpoint    = os.Getenv("LOGIN_URL")
	OAuthDialTimeout = 5 * time.Second
)
