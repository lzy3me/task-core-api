package workspaceRepository

import (
	"log"
	"task-core/models/entities"
	workspaceEntitiy "task-core/models/entities/workspace"
	"task-core/models/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "workspace"

func Insert(c *fiber.Ctx, rawBody interface{}) (interface{}, error) {
	result, err := repositories.SuperInsertOne(c, collectionName, rawBody)
	return result, err
}

func BuildRows(rows []workspaceEntitiy.Workspace) []workspaceEntitiy.Workspace {
	var record = make([]workspaceEntitiy.Workspace, 0)
	for _, v := range rows {
		record = append(record, workspaceEntitiy.Workspace{
			ID:          v.ID,
			Type:        v.Type,
			Name:        v.Name,
			Description: v.Description,
			Visibility:  v.Visibility,
		})
	}

	return record
}

func Find(c *fiber.Ctx, filter primitive.M, pagination *entities.PaginationRequests, sort *primitive.M) ([]workspaceEntitiy.Workspace, error) {
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

	var results []workspaceEntitiy.Workspace = make([]workspaceEntitiy.Workspace, 0)

	filter["deleted_by"] = nil

	cursor, err := repositories.SuperFind(c, collectionName, filter, opts)
	if err != nil {
		log.Println(err)

		return results, err
	}

	cursor.All(c.Context(), &results)

	return results, err
}

func FindOne(c *fiber.Ctx, filter primitive.M) (workspaceEntitiy.Workspace, error) {
	var entity workspaceEntitiy.Workspace
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

func Aggregate(c *fiber.Ctx, pipeline mongo.Pipeline) ([]workspaceEntitiy.Workspace, error) {
	cursor, err := repositories.SuperAggregate(c, collectionName, pipeline)
	var results []workspaceEntitiy.Workspace = make([]workspaceEntitiy.Workspace, 0)
	if err != nil {
		log.Println(err)

		return results, err
	}

	errC := cursor.All(c.Context(), &results)
	if errC != nil {
		log.Println("errC", errC)
	}

	// record := BuildRows(results)

	return results, err
}

func CountAggregate(c *fiber.Ctx, pipeline mongo.Pipeline) (int64, error) {
	cursor, err := repositories.SuperAggregate(c, collectionName, pipeline)
	var countLoaded = []entities.CountAggregate{}
	var cnt int64 = 0

	if err != nil {
		log.Println(err)

		return cnt, err
	}

	cursor.All(c.Context(), &countLoaded)

	if len(countLoaded) > 0 {
		cnt = int64(countLoaded[0].Count)
	}

	return cnt, err
}
