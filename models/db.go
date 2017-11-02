package models

import (
	"os"
	"gopkg.in/mgo.v2"
)

type MongoConnector struct {
	mongodbRootUsername string
	mongodbRootPassword string
	mongodbHost string
}

var mongoConnector MongoConnector

func getEnv(environmentVariableName string, defaultValue string) string {
	environmentVariableValue := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	if environmentVariableValue == "" {
		environmentVariableValue = defaultValue
	}
	return environmentVariableValue
}

func (mongoConnector MongoConnector) initMongoValues() {
	mongoConnector.mongodbRootUsername = getEnv("MONGO_INITDB_ROOT_USERNAME", "admin")
	mongoConnector.mongodbRootPassword = getEnv("MONGO_INITDB_ROOT_PASSWORD", "admin")
	mongoConnector.mongodbHost = getEnv("MONGO_HOST", "127.0.0.1")
}

func (mongoConnector MongoConnector) getSession() (*mgo.Session, error) {
	session, err := mgo.Dial(mongoConnector.mongodbHost)
	return session, err
}

func init() {
	mongoConnector.initMongoValues()
}
