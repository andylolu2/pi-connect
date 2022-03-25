package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetIp(hostname string) string {
	client := CreateConnection()
	defer CloseConnection(client)

	coll := client.Database("pi_connect").Collection("ip_addr")

	var result bson.D
	filter := bson.D{{Key: "hostname", Value: hostname}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Fatalf("No such hostname: %s\n", hostname)
	} else if err != nil {
		panic(err)
	}

	ip := result.Map()["ip"].(string)
	return ip
}
