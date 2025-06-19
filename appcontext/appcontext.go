package appcontext

import (
	"context"
	"database/sql"

	"github.com/Doraverse-Workspace/go-utilities/language"
	"github.com/Doraverse-Workspace/go-utilities/logger"
	"github.com/Doraverse-Workspace/go-utilities/timezone"
	"github.com/segmentio/ksuid"
)

type contextKey int

const (
	userContextKey contextKey = iota
	platformContextKey
	provinceContextKey
	ipContextKey
	langContextKey
	timezoneContextKey
	isMobileKey
	userAgentKey
	dbTransactionKey
)

type AppContext struct {
	requestID string
	traceID   string
	logger    *logger.Logger
	context   context.Context
}

type Fields = logger.Fields

func newWithSource(ctx context.Context, source string) *AppContext {
	var (
		requestID = generateID()
		traceID   = generateID()
	)

	return &AppContext{
		requestID: requestID,
		traceID:   traceID,
		logger:    logger.NewLogger(logger.Fields{"requestId": requestID, "traceId": traceID, "source": source}),
		context:   ctx,
	}
}

func NewRest(ctx context.Context) *AppContext {
	return newWithSource(ctx, "rest")
}

func NewGRPC(ctx context.Context) *AppContext {
	return newWithSource(ctx, "grpc")
}

func NewWorker(ctx context.Context) *AppContext {
	return newWithSource(ctx, "worker")
}

func (appCtx *AppContext) SetTraceID(traceID string) {
	appCtx.traceID = traceID
	appCtx.logger.AddData(logger.Fields{"traceId": traceID})
}

func (appCtx *AppContext) GetTraceID() string {
	return appCtx.traceID
}

func (appCtx *AppContext) AddLogData(fields Fields) {
	appCtx.logger.AddData(fields)
}

func (appCtx *AppContext) Logger() *logger.Logger {
	return appCtx.logger

}

func (appCtx *AppContext) Context() context.Context {
	return appCtx.context
}

func (appCtx *AppContext) SetContext(ctx context.Context) {
	appCtx.context = ctx
}

func (appCtx *AppContext) SetUserID(id string) {
	appCtx.context = context.WithValue(appCtx.context, userContextKey, id)
}

func (appCtx *AppContext) GetUserID() string {
	id, ok := appCtx.context.Value(userContextKey).(string)
	if !ok {
		return ""
	}
	return id
}

func (appCtx *AppContext) SetPlatformID(id string) {
	appCtx.context = context.WithValue(appCtx.context, platformContextKey, id)
}

func (appCtx *AppContext) GetPlatformID() string {
	id, ok := appCtx.context.Value(platformContextKey).(string)
	if !ok {
		return ""
	}
	return id
}

func (appCtx *AppContext) SetProvince(code int) {
	appCtx.context = context.WithValue(appCtx.context, provinceContextKey, code)
}

func (appCtx *AppContext) GetProvince() int {
	id, ok := appCtx.context.Value(provinceContextKey).(int)
	if !ok {
		return -1
	}
	return id
}

func (appCtx *AppContext) SetIP(ip string) {
	appCtx.context = context.WithValue(appCtx.context, ipContextKey, ip)
}

func (appCtx *AppContext) GetIP() string {
	ip, ok := appCtx.context.Value(ipContextKey).(string)
	if !ok {
		return ""
	}
	return ip
}

func (appCtx *AppContext) SetLang(lang string) {
	appCtx.context = context.WithValue(appCtx.context, langContextKey, lang)
}

func (appCtx *AppContext) GetLang() language.Language {
	lang, ok := appCtx.context.Value(langContextKey).(string)
	if !ok {
		return language.Vietnamese
	}

	dLang := language.ToLanguage(lang)
	if !dLang.IsValid() {
		return language.Vietnamese
	}

	return dLang
}

func (appCtx *AppContext) SetTimezone(tz string) {
	appCtx.context = context.WithValue(appCtx.context, timezoneContextKey, tz)
}

func (appCtx *AppContext) GetTimezone() timezone.Timezone {
	tz, ok := appCtx.context.Value(timezoneContextKey).(string)
	if !ok {
		return *timezone.UTC
	}

	utz, err := timezone.GetTimezoneData(tz)
	if err != nil {
		appCtx.logger.Error("error when getting user timezone", err, Fields{"timezone": tz})
	}
	return *utz
}

func (appCtx *AppContext) SetIsMobile(isMobile bool) {
	appCtx.context = context.WithValue(appCtx.context, isMobileKey, isMobile)
}

func (appCtx *AppContext) GetIsMobile() bool {
	isMobile, ok := appCtx.context.Value(isMobileKey).(bool)
	if !ok {
		return true
	}
	return isMobile
}

func (appCtx *AppContext) SetDBTransaction(tx *sql.Tx) {
	appCtx.context = context.WithValue(appCtx.context, dbTransactionKey, tx)
}

func (appCtx *AppContext) GetDBTransaction() *sql.Tx {
	tx, ok := appCtx.context.Value(dbTransactionKey).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}

func (appCtx *AppContext) SetUserAgent(ua string) {
	appCtx.context = context.WithValue(appCtx.context, userAgentKey, ua)
}

func (appCtx *AppContext) GetUserAgent() string {
	ip, ok := appCtx.context.Value(userAgentKey).(string)
	if !ok {
		return ""
	}
	return ip
}

func generateID() string {
	return ksuid.New().String()
}
