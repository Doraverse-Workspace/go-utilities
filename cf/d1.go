package cf

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/d1"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

type CFD1 struct {
	dbID   string
	accID  string
	client *cloudflare.Client
}

func NewD1(client *cloudflare.Client, dbID, accID string) *CFD1 {
	return &CFD1{
		client: client,
		dbID:   dbID,
		accID:  accID,
	}
}

func (d *CFD1) Query(ctx context.Context, sql string, params ...string) (*pagination.SinglePage[d1.QueryResult], error) {
	return d.client.D1.Database.Query(ctx, d.dbID, d1.DatabaseQueryParams{
		AccountID: cloudflare.F(d.accID),
		Sql:       cloudflare.F(sql),
		Params:    cloudflare.F(params),
	})
}

func (d *CFD1) Raw(ctx context.Context, sql string, params ...string) (*pagination.SinglePage[d1.DatabaseRawResponse], error) {
	return d.client.D1.Database.Raw(ctx, d.dbID, d1.DatabaseRawParams{
		AccountID: cloudflare.F(d.accID),
		Sql:       cloudflare.F(sql),
		Params:    cloudflare.F(params),
	})
}
