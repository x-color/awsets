package lister

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/trek10inc/awsets/context"
	"github.com/trek10inc/awsets/resource"
)

type AWSGlueConnection struct {
}

func init() {
	i := AWSGlueConnection{}
	listers = append(listers, i)
}

func (l AWSGlueConnection) Types() []resource.ResourceType {
	return []resource.ResourceType{
		resource.GlueConnection,
	}
}

func (l AWSGlueConnection) List(ctx context.AWSetsCtx) (*resource.Group, error) {
	svc := glue.NewFromConfig(ctx.AWSCfg)
	res, err := svc.GetConnections(ctx.Context, &glue.GetConnectionsInput{
		HidePassword: aws.Bool(true),
		MaxResults:   aws.Int32(100),
	})

	rg := resource.NewGroup()
	paginator := glue.NewGetConnectionsPaginator(req)
	for paginator.Next(ctx.Context) {
		page := paginator.CurrentPage()
		for _, v := range page.ConnectionList {
			r := resource.New(ctx, resource.GlueConnection, v.Name, v.Name, v)
			rg.AddResource(r)
		}
	}

	err := paginator.Err()
	return rg, err
}
