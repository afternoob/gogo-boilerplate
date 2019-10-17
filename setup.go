package main

import (
	"io"
	"log"

	"github.com/afternoob/gogo-boilerplate/app"
	"github.com/afternoob/gogo-boilerplate/config"
	companyRepo "github.com/afternoob/gogo-boilerplate/repository/company/store"
	staffRepo "github.com/afternoob/gogo-boilerplate/repository/staff/store"
	companyService "github.com/afternoob/gogo-boilerplate/service/company"
	staffService "github.com/afternoob/gogo-boilerplate/service/staff"
	"github.com/devit-tel/goxid"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
)

func setupJaeger(appConfig *config.Config) (opentracing.Tracer, io.Closer) {
	transport, err := zipkin.NewHTTPTransport(
		appConfig.JaegerEndpoint,
		zipkin.HTTPBatchSize(50),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	if err != nil {
		log.Fatalf("Cannot initialize HTTP transport: %v", err)
	}

	return jaeger.NewTracer(
		"GoGoBoilerplate",
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(transport),
	)
}

func newApp(appConfig *config.Config) *app.App {
	xid := goxid.New()

	companyStore := companyRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	company := companyService.New(xid, companyStore)

	staffStore := staffRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	staff := staffService.New(xid, staffStore, companyStore)

	return app.New(staff, company)
}
