package mongo

import (
	"context"
	"curse/practice31/logger"
	"curse/practice31/models"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	usersCollection = "users"
)

type UsersManager struct {
	db *mongo.Database

	log logger.Logger
}

func (users *UsersManager) obtainNextID() (int, error) {
	nextID := 0
	collectionUsers := users.db.Collection(usersCollection)

	// Определение этапов агрегации.
	pipeline := []bson.M{
		{"$sort": bson.M{"id": -1}},
		{"$limit": 1},
		{"$project": bson.M{"_id": 0, "id": 1}},
	}

	// Создание объекта Aggregation.
	agg, err := collectionUsers.Aggregate(context.Background(), pipeline)
	if err != nil {
		panic(err)
	}

	// Получение результатов агрегации.
	var result models.Users

	if err := agg.All(context.Background(), &result); err != nil {
		panic(err)
	}

	// Вывод результата.
	if len(result) > 0 {
		nextID = result[0].ID
	}

	return nextID + 1, nil
}

func (users *UsersManager) findUserByFilter(filter *bson.M) (*models.User, error) {
	collectionUsers := users.db.Collection(usersCollection)
	result := collectionUsers.FindOne(context.Background(), filter)

	err := result.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, models.ErrNoDocuments
	}

	if err != nil {
		return nil, errors.Wrap(err, "can not find by ID")
	}

	var foundUser *models.User
	if err := result.Decode(&foundUser); err != nil {
		return nil, errors.Wrap(err, "can not decode found user")
	}

	return foundUser, nil
}

func (users *UsersManager) Insert(user *models.User) (*models.User, error) {
	collectionUsers := users.db.Collection(usersCollection)

	nextID, err := users.obtainNextID()
	if err != nil {
		panic(err)
	}

	user.ID = nextID
	opts := options.InsertOne()

	result, err := collectionUsers.InsertOne(context.Background(), user, opts)
	if err != nil {
		return nil, errors.Wrap(err, "can not insert user")
	}

	filter := &bson.M{
		"_id": result.InsertedID,
	}

	insertedUser, err := users.findUserByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find inserted user")
	}

	return insertedUser, nil
}

func (mongo *UsersManager) UpdateAge(id int, age int) (*models.User, error) {
	collectionUsers := mongo.db.Collection(usersCollection)

	filter := &bson.M{
		"id": id,
	}

	upd := &bson.M{
		"$set": &bson.M{
			"age": age,
		},
	}

	_, err := collectionUsers.UpdateOne(context.Background(), filter, upd)
	if err != nil {
		return nil, errors.Wrap(err, "can not update age by user")
	}

	updatedUser, err := mongo.findUserByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find updated by age user")
	}

	return updatedUser, nil
}

func (mongo *UsersManager) DeleteUser(id int) (*models.User, error) {
	collectionUsers := mongo.db.Collection(usersCollection)

	filter := &bson.M{
		"id": id,
	}

	deletedUser, err := mongo.findUserByFilter(filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find deleted user")
	}

	if _, err := collectionUsers.DeleteOne(context.Background(), filter); err != nil {
		return nil, errors.Wrap(err, "can not delete user")
	}

	return deletedUser, nil
}

func New(log logger.Logger, username, password string, dbHosts []string, database string) *UsersManager {
	client, err := connect(dbHosts, username, password)
	if err != nil {
		panic(err)
	}

	db := client.Database(database)

	return &UsersManager{
		log: log,
		db:  db,
	}
}
