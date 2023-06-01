package boardRepository

import (
	"log"
	"task-core/models/entities"
	boardEnitity "task-core/models/entities/board"
	"task-core/models/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "board"

func Insert(c *fiber.Ctx, rawBody interface{}) (interface{}, error) {
	result, err := repositories.SuperInsertOne(c, collectionName, rawBody)
	return result, err
}

func BuildRows(rows []boardEnitity.Board) []boardEnitity.Board {
	var record = make([]boardEnitity.Board, 0)
	for _, v := range rows {
		record = append(record, boardEnitity.Board{
			ID:                v.ID,
			BelongToWorkspace: v.BelongToWorkspace,
			BelongToUser:      v.BelongToUser,
			Visibility:        v.Visibility,
			List:              v.List,
			Permission:        v.Permission,
			Labels:            v.Labels,
			UserMember:        v.UserMember,
		})
	}

	return record
}

func Find(c *fiber.Ctx, filter primitive.M, pagination *entities.PaginationRequests, sort *primitive.M) ([]boardEnitity.Board, error) {
	var opts = options.Find()

	if pagination != nil {
		var page int64 = 1
		var perPage int64 = 10
		if pagination.Page != 0 {
			page = pagination.Page
		}

		if pagination.PerPage != 0 {
			perPage = pagination.PerPage
		}

		var skip = (page - 1) * perPage

		opts = options.Find().SetSkip(int64(skip)).SetLimit(int64(perPage))
	}

	if sort == nil {
		opts.SetSort(bson.M{"_id": 1})
	} else {
		opts.SetSort(&sort)
	}

	var results []boardEnitity.Board = make([]boardEnitity.Board, 0)

	filter["deleted_by"] = nil

	cursor, err := repositories.SuperFind(c, collectionName, filter, opts)
	if err != nil {
		log.Println(err)

		return results, err
	}

	cursor.All(c.Context(), &results)

	return results, err
}

func FindOne(c *fiber.Ctx, filter primitive.M) (boardEnitity.Board, error) {
	var entity boardEnitity.Board
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
