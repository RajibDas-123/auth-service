package repo

import (
	"context"

	"os"
	"time"

	"github.com/RajibDas-123/ms-grpc-auth/auth/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connection *mongo.Client
var err error
var TemplateMap = map[string]string{}

func Initialize() {

	TemplateMap = map[string]string{"1": "InvoiceFolder", "2": "FormsFolder", "3": "DocumentsFolder", "4": "BillsFolder", "5": "ChequeFolder"}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if Connection, err = mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("DB_DSN"),
	)); err != nil {
		logging.DBLogger.Error("Failed to connect to the database", err)
	} else {
		logging.DBLogger.Info("Successfully connect to the database")
	}
}
