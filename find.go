package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneWithFilter[model Model](filter bson.M, options ...*options.FindOneOptions) (result *model, err error) {
	var m model
	err = defaultDatabase.Collection(m.Collection()).FindOne(defaultInstance.config.contextFn(), filter, options...).Decode(&result)
	return
}

func FindWithFilter[model Model](filter bson.M, options ...*options.FindOptions) (result []model, err error) {
	var (
		m      model
		cursor *mongo.Cursor
	)
	cursor, err = defaultDatabase.Collection(m.Collection()).Find(defaultInstance.config.contextFn(), filter, options...)
	if err != nil {
		return
	}

	err = cursor.All(defaultInstance.config.contextFn(), &result)
	if err != nil {
		return
	}

	err = cursor.Err()
	return
}
