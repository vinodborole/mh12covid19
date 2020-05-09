package main

import (
	"context"
	"covid19/config"
	"covid19/config/domain"
	"covid19/controllers"
	"covid19/infra"
	"covid19/infra/auth"
	"covid19/infra/database"
	u "covid19/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"net/http"
)

var db *gorm.DB
var err error

func main() {
	err := config.ReadYamlConfigFile()
	if err != nil {
		u.Errorln(err)
		return
	}
	fmt.Println("Initialise logging")
	logConfig()
	user := config.GetConfig().MH12Config.Database.Username
	password := config.GetConfig().MH12Config.Database.Password
	dbName := config.GetConfig().MH12Config.Database.DBName
	location := config.GetConfig().MH12Config.Database.Location
	port := config.GetConfig().MH12Config.Database.Port
	fmt.Println("Setup Database")
	database.Setup(dbName, location, user, password, port)
	database.GetWorkingInstance().MigrateSchema()
	fmt.Println("Create default user")
	createDefaultUser()
	generateCaseJSON()
	//createWards()
	fmt.Println("Starting HTTP server..")
	_, err = StartHTTPServer()
	if err != nil {
		fmt.Println("Error Starting MH12 Http Server : ", err.Error())
		u.Errorln(err.Error())
		return
	}
}

func logConfig() {
	logFile := config.GetConfig().MH12Config.Log.Location + config.GetConfig().MH12Config.Application.Name + ".log"
	LogWriter := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes
		MaxAge:     config.GetConfig().MH12Config.Log.MaxAge,
		MaxBackups: config.GetConfig().MH12Config.Log.MaxBackups,
		Compress:   true,
	}
	writerMap := lfshook.WriterMap{
		log.DebugLevel: &LogWriter,
		log.InfoLevel:  &LogWriter,
		log.WarnLevel:  &LogWriter,
		log.ErrorLevel: &LogWriter,
	}
	log.AddHook(lfshook.NewHook(
		writerMap,
		&log.JSONFormatter{},
	))
	log.SetOutput(ioutil.Discard)
	level := config.GetConfig().MH12Config.Log.Level
	if len(level) > 0 {
		switch level {
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warning":
			log.SetLevel(log.WarnLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		case "debug":
			log.SetLevel(log.DebugLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

//StartHTTPServer start http server
func StartHTTPServer() (*http.Server, error) {
	port := config.GetConfig().MH12Config.Application.Port
	if port == "" {
		port = "8080"
	}
	u.Infoln("MH12 port : ", port)
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/login", controllers.Authenticate).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/wards", controllers.GetWards).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/news", controllers.GetNews).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/ward-details", controllers.GetWardDetails).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/district-summary", controllers.GetDistrictSummary).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/district-summary-latest", controllers.GetDistrictSummaryLatest).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/patient-summary", controllers.GetPatientSummary).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/patient-summary-delta", controllers.GetPatientSummaryDelta).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/ward", controllers.AddWardCase).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/add-bulk-ward", controllers.AddBulkWardCase).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/district", controllers.AddDistrictSummary).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/patient", controllers.AddPatientSummary).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/add-news", controllers.AddNews).Methods(http.MethodPost)

	router.Use(auth.JwtAuthentication) //attach JWT auth middleware
	server := &http.Server{Addr: ":" + port, Handler: router}
	server.ListenAndServe()
	return server, nil
}

func createDefaultUser() {
	configObj := config.GetConfig()
	// make account entry
	account := domain.User{
		ID:       1,
		Name:     configObj.MH12Config.Application.User.Name,
		Phone:    configObj.MH12Config.Application.User.Phone,
		Email:    configObj.MH12Config.Application.User.Email,
		Password: configObj.MH12Config.Application.User.Password,
	}
	_, _ = infra.GetUseCaseInteractor().CreateUser(&account)
}

func generateCaseJSON() {
	infra.GetUseCaseInteractor().GenerateDistrictSummaryLatestJSON()
	infra.GetUseCaseInteractor().GenerateDistrictSummaryJSON()
	infra.GetUseCaseInteractor().GenerateDistrictDetailJSON()
	infra.GetUseCaseInteractor().GenerateWardCasesJSON()
	infra.GetUseCaseInteractor().GenerateNewsJSON()
	infra.GetUseCaseInteractor().GeneratePatientDeltaSummaryJSON()
	infra.GetUseCaseInteractor().GeneratePatientSummaryJSON()
	infra.GetUseCaseInteractor().GenerateMetaJSON()
}
func createWards() {
	aundhWard := domain.Ward{Name: "Aundh Banner", Code: "AB"}
	bhawaniPethWard := domain.Ward{Name: "Bhawani Peth", Code: "BP"}
	bibwewadiWard := domain.Ward{Name: "Bibwewadi", Code: "BW"}
	dhankawadiWard := domain.Ward{Name: "Dhankawadi - Sahakarnagar", Code: "DS"}
	dholePatilWard := domain.Ward{Name: "Dhole Patil", Code: "DP"}
	hadapsarMundhwaWard := domain.Ward{Name: "Hadapsar Mundhwa", Code: "HM"}
	kasbaWard := domain.Ward{Name: "Kasba - Visharambaghwada", Code: "KV"}
	kondhwaWard := domain.Ward{Name: "Kondhwa - Yewalewadi", Code: "KY"}
	kothrudWard := domain.Ward{Name: "Kothrud - Bavdhan", Code: "KB"}
	wadgoansheriWard := domain.Ward{Name: "Nagar Road - Wadgoansheri", Code: "NW"}
	shivajiNagarWard := domain.Ward{Name: "Shivaji Nagar - Ghole Road", Code: "SN"}
	sinhagadRoadWard := domain.Ward{Name: "Singhagad Road", Code: "SR"}
	wanawadiWard := domain.Ward{Name: "Wanawadi - Ramtekdi", Code: "WR"}
	warjeWard := domain.Ward{Name: "Warje - Karve Nagar", Code: "WK"}
	yerwadaWard := domain.Ward{Name: "Yerwada - Kalas - Dhanori", Code: "YKD"}
	puneRuralWard := domain.Ward{Name: "Pune Rural", Code: "PR"}

	_, _ = infra.GetUseCaseInteractor().CreateWard(&aundhWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&bhawaniPethWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&bibwewadiWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&dhankawadiWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&dholePatilWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&hadapsarMundhwaWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&kasbaWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&kondhwaWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&kothrudWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&wadgoansheriWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&shivajiNagarWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&sinhagadRoadWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&wanawadiWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&warjeWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&yerwadaWard)
	_, _ = infra.GetUseCaseInteractor().CreateWard(&puneRuralWard)

}

//ShutdownHTTPServer shutdown http server
func ShutdownHTTPServer(server *http.Server) error {
	u.Infoln("Shutting down mh12 server..")
	if err := server.Shutdown(context.TODO()); err != nil {
		msg := fmt.Sprintf("Stopping x-node http server error(): %s", err)
		u.Errorln(msg)
		panic(err) // failure/timeout shutting down the server gracefully
	}
	return nil
}

//CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

//ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, HEAD, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, X-Requested-With, Access-Control-Allow-Headers, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}
