package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/admin"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/cockroachdb"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/config"
	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/implementation"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/middleware"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport"
	httptransport "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"

	_ "github.com/lib/pq"
)

var (
	logger        = logrus.New()
	serviceConfig config.Config
	errs          chan error
	db            *sql.DB
	repository    template.Repository
	consul        admin.ServiceDiscovery
)

const (
	appName    = "notification"
	configFile = "bootstrap.yml"
)

func init() {

	// Log as JSON instead of the default ASCII formatter.
	// Log as JSON instead of the default ASCII formatter, but wrapped with the runtime Formatter.
	formatter := runtime.Formatter{ChildFormatter: &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}}
	// Enable line number logging as well
	formatter.Line = true

	logger.SetFormatter(&formatter)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	//initialize the configs
	configs, err := ioutil.ReadFile(configFile)
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err, "filename": configFile}).Error("invalid service configuration provided file supplied")
		//level.Error(logger).Log("error", "invalid service configuration provided file supplied", "Filename", configFile)
		os.Exit(1)
	}
	//unmarshall service configs
	if err := yaml.Unmarshal(configs, &serviceConfig); err != nil {
		logger.WithFields(logrus.Fields{"error": err, "filename": configFile}).Error("unable to unmarshal the config file")
		os.Exit(1)
	}

	// Connect to the "ordersdb" database
	db, err = sql.Open("postgres", "postgresql://root@localhost:26257/notification?sslmode=disable")
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err, "dbconn": "localhost:26257/notification?sslmode=disable"}).Error("failed to open database connection")
		os.Exit(1)
	}

	//create a repository client
	repository, err = cockroachdb.New(db, logger)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// Create service discovery registry client
	consul, err = admin.NewSDClient("192.168.1.20", "8500", "192.168.1.20", "9213", logger)
	if err != nil {
		//level.Error(logger).Log("exit", err)
		os.Exit(1)
	}
}

func main() {

	logger.WithField("msg", "xnotification service started...").Info()
	defer logger.WithField("msg", "xnotification service ended.").Info()

	//add authentication
	//authClient := server.NewAuthClient(cfg, logger)
	//authBefore := []kithttp.RequestFunc{kitjwt.HTTPToContext()}

	//serveroptions
	serverOptions := []kithttp.ServerOption{}

	//setup template service
	templateService := implementation.NewTemplateService(repository, logger)
	templateService = middleware.LoggingMiddleware(logger)(templateService)
	templateEndpoints := transport.MakeEndpoints(templateService)

	//setup admin service
	adminSVC := admin.NewAdminService(logger)
	adminSVC = admin.LoggingMiddleware(logger)(adminSVC)
	adminEndpoints := admin.MakeEndpoints(adminSVC)

	//setup router
	router := mux.NewRouter()

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	router.Handle("/debug/vars", http.DefaultServeMux)

	//add NotificationTemplate routes to the router
	httptransport.NewNotificationTemplateService(router, templateEndpoints, serverOptions)
	admin.AddAdminServiceRoutes(router, adminEndpoints, serverOptions)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", serviceConfig.Server.Port),
			Handler: router,
		}
		consul.Register()
		errs <- server.ListenAndServe()
	}()

	logger.WithField("error", <-errs).Error("listenserve return error")
	consul.Deregister()
}
