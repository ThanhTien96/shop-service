package api

import (
	"net/http"
	"shop-test/errRes"
	"shop-test/model"
	"shop-test/pkg/db"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ItemColName = "items"
)

func ApiListItem(svc model.Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		svc.Logger().Infof("INF: ApiListItem is called")

		paging := svc.NewPagingRequest(1, 100)

		pageParam := c.QueryParam("page")
		if pageParam != "" {
			p, err := strconv.Atoi(pageParam)
			if err == nil {
				paging.Page = p
			}
		}

		pageSizeParam := c.QueryParam("page_size")
		if pageSizeParam != "" {
			ps, err := strconv.Atoi(pageSizeParam)
			if err == nil {
				paging.PageSize = ps
			}
		}

		var items []*model.Item
		coll := db.GetDBCollection(ItemColName)

		skip := (paging.Page - 1) * paging.PageSize
		totalItems, err := coll.CountDocuments(c.Request().Context(), bson.D{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, JsonError(http.StatusInternalServerError, err.Error()))
		}

		cursor, err := coll.Find(c.Request().Context(), bson.D{}, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(paging.PageSize)))
		if err != nil {
			return c.JSON(http.StatusNotFound, JsonError(http.StatusBadRequest, err.Error()))
		}
		defer cursor.Close(c.Request().Context())

		for cursor.Next(c.Request().Context()) {
			var item *model.Item
			if err := cursor.Decode(&item); err != nil {
				return c.JSON(http.StatusInternalServerError, JsonError(http.StatusInternalServerError, err.Error()))
			}
			items = append(items, item)
		}

		if err := cursor.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, JsonError(http.StatusInternalServerError, err.Error()))
		}

		pagingResult := &model.Paging{
			TotalItem: int(totalItems),
			Page:      int(paging.Page),
			PageSize:  int(paging.PageSize),
		}

		return c.JSON(http.StatusOK, JsonSuccessPaging("Get item list successfully", items, pagingResult))
	})
}

func ApiGetItem(svc model.Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		svc.Logger().Infof("INF: ApiGetItem is called")

		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, JsonError(http.StatusBadRequest, errRes.ErrBindingParam.Error()))
		}
		var item *model.Item
		var itemId primitive.ObjectID
		coll := db.GetDBCollection(ItemColName)
		itemId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, JsonError(http.StatusBadRequest, errRes.ErrBindingParam.Error()))
		}

		filter := bson.D{{"_id", itemId}}

		result := coll.FindOne(c.Request().Context(), filter)
		if result.Err() != nil {
			if result.Err() == mongo.ErrNoDocuments {
				return c.JSON(http.StatusNotFound, JsonError(http.StatusNotFound, errRes.ErrItemNotFound.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, JsonError(http.StatusInternalServerError, errRes.ErrFailQuery.Error()))
			}
		}

		if err := result.Decode(&item); err != nil {
			return c.JSON(http.StatusInternalServerError, JsonError(http.StatusInternalServerError, errRes.ErrFailDecode.Error()))
		}

		return c.JSON(http.StatusOK, JsonSuccess("Get item successfully", item))
	})

}

func ApiCreateItem(svc model.Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		svc.Logger().Infof("INF: apiCreateItem is called")

		var itemRequest model.CreateItemRequest
		if err := c.Bind(&itemRequest); err != nil {
			svc.Logger().Errorf("INF: apiCreateItem is error", err)
			return c.JSON(http.StatusBadRequest, JsonError(http.StatusBadRequest, errRes.ErrCreatePayloadFail.Error()))
		}

		if err := itemRequest.ValidateItem(); err != nil {
			svc.Logger().Errorf("ERR: ", err.Error())
			return c.JSON(http.StatusBadRequest, JsonError(http.StatusBadRequest, err.Error()))
		}

		coll := db.GetDBCollection(ItemColName)

		item := &model.Item{
			Name:     itemRequest.Name,
			Currency: itemRequest.Currency,
			Stock:    itemRequest.Stock,
			BaseModel: model.BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		result, err := coll.InsertOne(c.Request().Context(), item)
		if err != nil {
			svc.Logger().Errorf("INF: apiCreateItem is error", err)
			return c.JSON(http.StatusBadRequest, JsonError(http.StatusBadRequest, errRes.ErrCreateItemFail.Error()))
		}

		item.ID = result.InsertedID.(primitive.ObjectID)

		return c.JSON(http.StatusOK, JsonSuccess("Create item successfully.", item))
	})
}

