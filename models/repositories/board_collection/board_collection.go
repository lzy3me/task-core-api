package boardcollectionRepository

import (
	"log"
	boardcollectionEntity "task-core/models/entities/board_collection"
	"task-core/models/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "board_collection"

func Insert(c *fiber.Ctx, rawBody interface{}) (interface{}, error) {
	result, err := repositories.SuperInsertOne(c, collectionName, rawBody)
	return result, err
}

func BuildRows(rows []boardcollectionEntity.BoardCollection) []boardcollectionEntity.BoardCollection {
	var record = make([]boardcollectionEntity.BoardCollection, 0)
	for _, v := range rows {
		record = append(record, boardcollectionEntity.BoardCollection{
			ID:            v.ID,
			BelongToBoard: v.BelongToBoard,
			BelongToList:  v.BelongToList,
			Name:          v.Name,
			Description:   v.Description,
			WatchUsers:    v.WatchUsers,
			AssignUsers:   v.AssignUsers,
			DueDate:       v.DueDate,
			Labels:        v.Labels,
			Collection:    v.Collection,
		})
	}

	return record
}

func Find(c *fiber.Ctx, pipeline mongo.Pipeline) ([]boardcollectionEntity.BoardCollection, error) {
	cursor, err := repositories.SuperAggregate(c, collectionName, pipeline)
	var results []boardcollectionEntity.BoardCollection = make([]boardcollectionEntity.BoardCollection, 0)
	if err != nil {
		log.Fatal(err)
		return results, err
	}

	errC := cursor.All(c.Context(), &results)
	if errC != nil {
		log.Fatal("errC", errC)
	}

	return results, err
}

func FindOne(c *fiber.Ctx, filter primitive.M) (boardcollectionEntity.BoardCollection, error) {
	var entity boardcollectionEntity.BoardCollection
	errFind := repositories.SuperFindOne(c, collectionName, filter, &entity)
	return entity, errFind
}

func Update(c *fiber.Ctx, filter primitive.M, body primitive.M, upsert bool) (interface{}, error) {
	result, errFind := repositories.SuperUpdate(c, collectionName, filter, body, upsert)
	return result, errFind
}

func SoftDelete(c *fiber.Ctx, filter primitive.M, by *primitive.ObjectID) (interface{}, error) {
	result, errFind := repositories.SuperSoftDelete(c, collectionName, filter, by)
	return result, errFind
}
