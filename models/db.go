package models

import (
	"os"
	"gopkg.in/mgo.v2"
)

type MongoConnector struct {
	username string
	password string
	host string
	database string
}

// var MongoConnectorVar MongoConnector

func getEnv(environmentVariableName string, defaultValue string) string {
	environmentVariableValue := os.Getenv(environmentVariableName)
	if environmentVariableValue == "" {
		environmentVariableValue = defaultValue
	}
	return environmentVariableValue
}

func (mongoConnector MongoConnector) InitMongoValues() {
	mongoConnector.username = getEnv("MONGO_INITDB_ROOT_USERNAME", "admin")
	mongoConnector.password = getEnv("MONGO_INITDB_ROOT_PASSWORD", "admin")
	mongoConnector.host = getEnv("MONGO_HOST", "127.0.0.1")
	mongoConnector.database = getEnv("MONGO_DATABASE", "testdb")
}

func (mongoConnector MongoConnector) getSession() *mgo.Session {

	cred := &mgo.Credential{
		Username:  "admin",
		Password:  "admin",
		Mechanism: "SCRAM-SHA-1",
		Source:    "admin",
	}
	
	session, err := mgo.Dial(mongoConnector.host)

	if err != nil {
		panic(err)
	}

	err = session.Login(cred)

	if err != nil {
		panic(err)
	}

	return session
}
