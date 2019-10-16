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
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
)

func main() {
	appConfig := config.Get()

	tracer, closer := setupJaeger(appConfig)
	defer closer.Close()

	router := gin.Default()
	router.Use(ginhttp.Middleware(tracer))

	_ = newApp(appConfig).RegisterRoute(router)
	_ = router.Run()
}

func setupJaeger(appConfig *config.Config) (opentracing.Tracer, io.Closer) {
	transport, err := zipkin.NewHTTPTransport(
		appConfig.JaegerEndpoint,
		zipkin.HTTPBatchSize(10),
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
