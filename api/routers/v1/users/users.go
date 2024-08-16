package users

import (
    "context"
    "net/http"

    "github.com/danielgtaylor/huma/v2"
)

type IntNot3 int

type GreetingOutput struct {
    Body struct {
        Message string `json:"message" example:"Example: Hello, world!" doc:"Doc: Greeting message"`
    }
}

type ReviewInput struct {
    Body struct {
        Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
        Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
        Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
    }
}

type MyInput struct {
    QueryCount IntNot3 `query:"count" example:"2" minimum:"1" maximum:"10"`
    Value      string  `json:"value,omitempty" dependentRequired:"dependent1,dependent2"`
    Dependent1 string  `json:"dependent1,omitempty"`
    Dependent2 string  `json:"dependent2,omitempty"`
}

func AddRouters(api *huma.API) {

    prefix := "/users"
    huma.Get(*api, prefix, func(ctx context.Context, input *MyInput) (*GreetingOutput, error) {
        output := &GreetingOutput{}
        output.Body.Message = "Welcome to Get!"
        return output, nil
    })

    huma.Get(*api, prefix+"/{name}", func(ctx context.Context, input *struct {
        Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
    }) (*GreetingOutput, error) {
        output := &GreetingOutput{}
        output.Body.Message = input.Name
        return output, nil
    })

    huma.Register(*api, huma.Operation{
        OperationID: "ID-get-greeting",
        Method:      http.MethodGet,
        Path:        "/users/{name}",
        Summary:     huma.GenerateSummary(http.MethodGet, "/users/{name}", struct{}{}),
        Description: "Description of Get a greeting",
        Tags:        []string{"Users", "Greeting", "Get"},
        Errors:      []int{http.StatusInternalServerError, http.StatusBadRequest},
    }, func(ctx context.Context, input *struct {
        Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
    }) (*GreetingOutput, error) {
        output := &GreetingOutput{}
        output.Body.Message = input.Name
        return output, nil
    })

    minRaiting := 1.0
    maxRaiting := 5.0
    responses := map[string]*huma.Response{
        "200": &huma.Response{},
        // "201": &huma.Response{},
    }
    huma.Register(*api, huma.Operation{
        OperationID:   "ID-post-review",
        Method:        http.MethodPost,
        Path:          "/users/{username}/reviews",
        Summary:       huma.GenerateSummary(http.MethodPost, "/users/{username}/reviews", ReviewInput{}),
        Description:   "Description of Post a review",
        Tags:          []string{"Users", "Reviews", "Post"},
        DefaultStatus: http.StatusCreated,
        Responses:     responses,
        Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest},
        RequestBody: &huma.RequestBody{
            Required: true,
            Content: map[string]*huma.MediaType{
                "application/json": {
                    Schema: &huma.Schema{
                        Type: "object",
                        Properties: map[string]*huma.Schema{
                            "author": {
                                Type: "string",
                            },
                            "rating": {
                                Type:    "integer",
                                Minimum: &minRaiting,
                                Maximum: &maxRaiting,
                            },
                            "message": {
                                Type: "string",
                            },
                        },
                    },
                },
            },
        },
    }, func(ctx context.Context, i *struct {
        Name string `path:"username" maxLength:"30" example:"Niki" doc:"Name Author"`
    }) (*struct{}, error) {
        // save data in database
        status := http.StatusServiceUnavailable
        return nil, huma.NewError(status, "Database error")
    })

    huma.Post(*api, prefix, func(ctx context.Context, input *struct {
        Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
        Body struct {
            User string `body:"username" maxLength:"30" example:"niki" doc:"Username"`
        }
    }) (*struct{}, error) {
        // save data in database
        return nil, nil
    })
}
