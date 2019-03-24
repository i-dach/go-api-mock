package apm

import (
	"context"

	_ "github.com/aws/aws-xray-sdk-go/plugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func TraceSeg(c context.Context, service string) *context.Context {
	ctx, seg := xray.BeginSegment(c, service)
	seg.Close(nil)

	return &ctx
}

func TraceSubSeg(c context.Context, service string) *context.Context {
	ctx, subSeg := xray.BeginSubsegment(c, service)
	subSeg.Close(nil)

	return &ctx
}
