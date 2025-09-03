package cf

import (
	"context"
	"encoding/json"
	"errors"
	"log"
)

type D1HostMapping struct {
	db *CFD1
}

type HostMapping struct {
	Host        string `json:"host"`
	ClusterIP   string `json:"cluster_ip"`
	ClusterCode string `json:"cluster_code"`
}

type hostMappingResult struct {
	Results []HostMapping `json:"results"`
	Success bool          `json:"success"`
}

const (
	initTableSQL = `
	CREATE TABLE IF NOT EXISTS host_mappings (
		host text NOT NULL PRIMARY KEY,
		cluster_ip text NOT NULL,
		cluster_code text
	)`
	initIndexSQL = `CREATE INDEX IF NOT EXISTS idx_host_mappings ON host_mappings(host)`
	insertSQL    = `INSERT INTO host_mappings (host, cluster_ip, cluster_code) VALUES (?, ?, ?)`
	selectSQL    = `SELECT host, cluster_ip, cluster_code FROM host_mappings WHERE host = ?`
	deleteSQL    = `DELETE FROM host_mappings WHERE host = ?`
)

func NewD1HostMapping(ctx context.Context, apiToken, dbID, accID string) (*D1HostMapping, error) {
	c := NewClient(apiToken)
	db := NewD1(c, dbID, accID)
	ins := &D1HostMapping{
		db: db,
	}
	_, err := ins.db.Raw(ctx, initTableSQL)
	if err != nil {
		log.Printf("[cf: d1 host mapping] failed to init table: %v", err)
		return nil, err
	}
	_, err = ins.db.Raw(ctx, initIndexSQL)
	if err != nil {
		log.Printf("[cf: d1 host mapping] failed to init index: %v", err)
		return nil, err
	}

	return ins, nil
}

func (h *D1HostMapping) GetHostMapping(ctx context.Context, host string) (*HostMapping, error) {
	res, err := h.db.Query(ctx, selectSQL, host)
	if err != nil {
		return nil, err
	}

	if len(res.Result) == 0 {
		return nil, nil
	}
	data := res.Result[0]
	var r hostMappingResult
	err = json.Unmarshal([]byte(data.JSON.RawJSON()), &r)
	if err != nil {
		return nil, err
	}

	if !r.Success || len(r.Results) == 0 {
		return nil, nil
	}
	record := r.Results[0]
	return &HostMapping{
		Host:        record.Host,
		ClusterIP:   record.ClusterIP,
		ClusterCode: record.ClusterCode,
	}, nil
}

func (h *D1HostMapping) InsertHostMapping(ctx context.Context, record HostMapping) error {
	if record.Host == "" || record.ClusterIP == "" {
		return errors.New("cf: d1 host mapping: host or cluster ip is empty")
	}
	_, err := h.db.Raw(ctx, insertSQL, record.Host, record.ClusterIP, record.ClusterCode)
	if err != nil {
		return err
	}
	return nil
}

func (h *D1HostMapping) DeleteHostMapping(ctx context.Context, host string) error {
	_, err := h.db.Raw(ctx, deleteSQL, host)
	if err != nil {
		return err
	}
	return nil
}
