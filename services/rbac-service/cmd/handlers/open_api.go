package swagger

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"services.core-service/app_error"

	"github.com/getkin/kin-openapi/openapi3"
)

func NewOpenAPI3() openapi3.T {

	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "ToDo API",
			Description: "REST APIs used for interacting with the ToDo Service",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/MarioCarrion/todo-api-microservice-example",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://0.0.0.0:9234",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Priority": openapi3.NewSchemaRef("",
			openapi3.NewStringSchema().
				WithEnum("none", "low", "medium", "high").
				WithDefault("none")),
		"Dates": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("start", openapi3.NewStringSchema().
					WithFormat("date-time").
					WithNullable()).
				WithProperty("due", openapi3.NewStringSchema().
					WithFormat("date-time").
					WithNullable())),
		"Task": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("description", openapi3.NewStringSchema()).
				WithProperty("is_done", openapi3.NewBoolSchema()).
				WithPropertyRef("priority", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Priority",
				}).
				WithPropertyRef("dates", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Dates",
				})),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateTasksRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a task.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("description", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).
					WithPropertyRef("dates", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Dates",
					})),
		},
		"UpdateTasksRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a task.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("description", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("is_done", openapi3.NewBoolSchema().
						WithDefault(false)).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).
					WithPropertyRef("dates", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Dates",
					})),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating tasks.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Task",
					}))),
		},
		"ReadTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one task.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Task",
					}))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/tasks": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateTask",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateTasksRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateTasksResponse",
					},
				},
			},
		},
		"/tasks/{taskId}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "ReadTask",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("taskId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadTasksResponse",
					},
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateTask",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("taskId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateTasksRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task was updated"),
					},
				},
			},
		},
	}

	return swagger
}

func RegisterOpenAPI() func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		swagger := NewOpenAPI3()

		swaggerRoute := engine.Group("/swagger")
		{
			swaggerRoute.GET("/openapi3.json", func(c *gin.Context) {
				content, err := json.Marshal(&swagger)
				app_error.MustError(err)
				c.JSON(http.StatusOK, content)
			})

			swaggerRoute.GET("/openapi3.yaml", func(c *gin.Context) {
				content, err := json.Marshal(&swagger)
				app_error.MustError(err)
				c.JSON(http.StatusOK, content)
			})
		}
	}
}
