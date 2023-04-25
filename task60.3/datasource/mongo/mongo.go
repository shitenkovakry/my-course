package mongo

import (
	"context"
	"curse/task60.3/logger"
	"curse/task60.3/models"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ordersCollection = "orders"
)

type OrdersManager struct {
	db *mongo.Database

	log logger.Logger
}

func (orders *OrdersManager) obtainNextID() (int, error) {
	nextID := 0
	collectionOrders := orders.db.Collection(ordersCollection)

	// Определение этапов агрегации.
	pipeline := []bson.M{
		{"$sort": bson.M{"id": -1}},
		{"$limit": 1},
		{"$project": bson.M{"_id": 0, "id": 1}},
	}

	// Создание объекта Aggregation.
	agg, err := collectionOrders.Aggregate(context.Background(), pipeline)
	if err != nil {
		panic(err)
	}

	// Получение результатов агрегации.
	var result models.Orders

	if err := agg.All(context.Background(), &result); err != nil {
		panic(err)
	}

	// Вывод результата.
	if len(result) > 0 {
		nextID = result[0].ID
	}

	return nextID + 1, nil
}

func (orders *OrdersManager) findOrderByFilter(filter *bson.M) (*models.Order, error) {
	collectionOrders := orders.db.Collection(ordersCollection)
	result := collectionOrders.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, errors.Wrap(err, "can not find by ID")
	}

	var foundOrder *models.Order
	if err := result.Decode(&foundOrder); err != nil {
		return nil, errors.Wrap(err, "can not decode found order")
	}

	return foundOrder, nil
}

func (orders *OrdersManager) Insert(order *models.Order) (*models.Order, error) {
	collectionOrders := orders.db.Collection(ordersCollection)

	nextID, err := orders.obtainNextID()
	if err != nil {
		panic(err)
	}

	order.ID = nextID
	opts := options.InsertOne()

	result, err := collectionOrders.InsertOne(context.Background(), order, opts)
	if err != nil {
		return nil, errors.Wrap(err, "can not insert order")
	}

	filter := &bson.M{
		"_id": result.InsertedID,
	}

	insertedOrder, err := orders.findOrderByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find inserted order")
	}

	return insertedOrder, nil
}

func (mongo *OrdersManager) UpdateAddress(ID int, address string) (*models.Order, error) {
	collectionOrders := mongo.db.Collection(ordersCollection)

	filter := &bson.M{
		"id": ID,
	}

	upd := &bson.M{
		"$set": &bson.M{
			"address": address,
		},
	}

	_, err := collectionOrders.UpdateOne(context.Background(), filter, upd)
	if err != nil {
		return nil, errors.Wrap(err, "can not update order address")
	}

	updatedOrder, err := mongo.findOrderByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find updated by address order")
	}

	return updatedOrder, nil
}

func (mongo *OrdersManager) UpdateTelephone(ID int, telephone string) (*models.Order, error) {
	collectionOrders := mongo.db.Collection(ordersCollection)

	filter := &bson.M{
		"id": ID,
	}

	upd := &bson.M{
		"$set": &bson.M{
			"telephone": telephone,
		},
	}

	_, err := collectionOrders.UpdateOne(context.Background(), filter, upd)
	if err != nil {
		return nil, errors.Wrap(err, "can not update order telephone")
	}

	updatedOrder, err := mongo.findOrderByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find updated by telephone order")
	}

	return updatedOrder, nil
}

func (mongo *OrdersManager) UpdateStatus(ID int, status string) (*models.Order, error) {
	collectionOrders := mongo.db.Collection(ordersCollection)

	filter := &bson.M{
		"id": ID,
	}

	upd := &bson.M{
		"$set": &bson.M{
			"status": status,
		},
	}

	_, err := collectionOrders.UpdateOne(context.Background(), filter, upd)
	if err != nil {
		return nil, errors.Wrap(err, "can not update order status")
	}

	updatedOrder, err := mongo.findOrderByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find updated by status order")
	}

	return updatedOrder, nil
}

func (mongo *OrdersManager) ReadAll() (models.Orders, error) {
	collectionOrders := mongo.db.Collection(ordersCollection)
	filter := &bson.M{}

	cursor, err := collectionOrders.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.Wrapf(err, "can not read orders")
	}

	var orders models.Orders

	if err := cursor.All(context.Background(), &orders); err != nil {
		return nil, errors.Wrap(err, "can not read cursor")
	}

	return orders, nil
}

func New(log logger.Logger, username, password string, dbHosts []string, database string) *OrdersManager {
	client, err := connect(dbHosts, username, password)
	if err != nil {
		panic(err)
	}

	db := client.Database(database)

	return &OrdersManager{
		log: log,
		db:  db,
	}
}
