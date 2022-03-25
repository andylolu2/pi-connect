package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostIp(hostname, ip string) {
	client := CreateConnection()
	defer CloseConnection(client)

	coll := client.Database("pi_connect").Collection("ip_addr")

	hostNameItem := primitive.E{Key: "hostname", Value: hostname}
	filter := bson.D{hostNameItem}
	replacement := bson.D{hostNameItem, {Key: "ip", Value: ip}}
	opts := options.Replace().SetUpsert(true)
	_, err := coll.ReplaceOne(context.TODO(), filter, replacement, opts)
	if err != nil {
		log.Fatal(err)
	}
}
