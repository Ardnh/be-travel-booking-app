# Handler Rules

## 1. Parse the Request Body, Parameters, and Path

Handlers are responsible for extracting and parsing incoming request data.

### Request Body Parsing
Use `c.Bind().Body(&req)` to parse JSON request bodies into DTOs:

```go
func (h *PoolPointHandler) CreatePoolPoint(c fiber.Ctx) error {
    var req dto.CreatePoolsDTO
    if err := c.Bind().Body(&req); err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, err)
    }
    // ...
}
```

### Path Parameters Parsing
Use `c.Params()` to extract path parameters and validate them:

```go
func (h *PoolPointHandler) GetPoolPointByID(c fiber.Ctx) error {
    id := c.Params("id")
    poolID, err := uuid.Parse(id)
    if err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid pool point ID", err)
    }
    // ...
}
```

### Query Parameters Parsing
Use `c.Query()` to extract query parameters with default values where appropriate:

```go
func (h *PoolPointHandler) GetAllPoolPoints(c fiber.Ctx) error {
    page := c.Query("page", "1")
    pageSize := c.Query("page_size", "30")
    search := c.Query("search")
    sortBy := c.Query("sort_by", "created_at")
    sortOrder := c.Query("sort_order", "desc")
    
    pageInt, err := strconv.Atoi(page)
    if err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, err)
    }
    // ...
}
```

---

## 2. Validate Input

All input must be validated in the handler layer before reaching the service layer. Invalid input should return an error response immediately.

### Using Struct Validator
Use the validator instance to validate request DTOs:

```go
if err := h.validator.Struct(&req); err != nil {
    return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
}
```

### Manual Validation
Some validations require manual checks (e.g., UUID parsing, numeric bounds):

```go
// UUID validation
poolID, err := uuid.Parse(id)
if err != nil {
    return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid pool point ID", err)
}

// Pagination bounds
if pageInt <= 0 {
    pageInt = 1
}
if pageSizeInt <= 0 || pageSizeInt >= 1000 {
    pageSizeInt = 30
}
```

### Error-Specific Handling
Handle service-layer errors appropriately:

```go
if err != nil {
    if errors.Is(err, errorConst.ErrBadRequest) {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid location type", err)
    }
    return httpResponses.NewErrorResponse(c, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
}
```

---

## 3. Return Data to the Client

Always use the standardized response helpers to return data to the client.

### Success Response
```go
return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool point retrieved successfully", poolPoint)
```

### Success Response with Pagination
```go
pagination := dto.Pagination{
    CurrentPage: pageInt,
    PageSize:    pageSizeInt,
    TotalItems:  int(total),
    TotalPages:  (int(total) + pageSizeInt - 1) / pageSizeInt,
    HasNext:     pageInt*pageSizeInt < int(total),
    HasPrevious: pageInt > 1,
}

return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Pool points retrieved successfully", poolPoints, pagination)
```

### Error Response
```go
return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid vendor ID", err)
return httpResponses.NewErrorResponse(c, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
```

Or using the standard error constants:
```go
return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
```

### Alternative: Using HandleError
For service errors, you can use the centralized error handler:

```go
if err != nil {
    return httpResponses.HandleError(c, err)
}
```

### Response Status Codes
- `fiber.StatusCreated` for successful creation (201)
- `fiber.StatusOK` for successful retrieval/update/deletion (200)
- `fiber.StatusBadRequest` for validation errors (400)
- `fiber.StatusUnauthorized` for authorization failures (401)
- `fiber.StatusForbidden` for forbidden access (403)
- `fiber.StatusNotFound` for resource not found (404)
- `fiber.StatusConflict` for conflicts (409)
- `fiber.StatusInternalServerError` for server errors (500)

### Additional Notes

#### Logging
Log key operations for debugging and monitoring:

```go
h.log.Infof("Starting CreateVendor")
h.log.Errorf("Failed to parse vendor ID: %v", err)
h.log.Infof("Successfully created vendor")
```

#### Context Values
Retrieve authenticated user info from context locals:

```go
userID, ok := c.Locals("user_id").(string)
if !ok || userID == "" {
    return httpResponses.NewErrorResponse(c, fiber.StatusUnauthorized, fiber.ErrUnauthorized.Message, "User ID not found")
}
```

---

## Handler Template

When creating a new handler, follow this template:

```go
package handlers

import (
    "github.com/ardnh/be-travel-booking-app/internal/application/dto"
    "github.com/ardnh/be-travel-booking-app/internal/domain/services"
    httpResponses "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
    validator_utils "github.com/ardnh/be-travel-booking-app/internal/utils/validator"
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v3"
    "github.com/google/uuid"
    "github.com/sirupsen/logrus"
)

type YourEntityHandler struct {
    yourEntityService services.YourEntityService
    validator         *validator.Validate
    log               *logrus.Logger
}

func NewYourEntityHandler(yourEntityService services.YourEntityService, validator *validator.Validate, log *logrus.Logger) *YourEntityHandler {
    return &YourEntityHandler{
        yourEntityService: yourEntityService,
        validator:         validator,
        log:               log,
    }
}

func (h *YourEntityHandler) CreateYourEntity(c fiber.Ctx) error {
    // 1. Parse request body
    var req dto.CreateYourEntityDTO
    if err := c.Bind().Body(&req); err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, err)
    }

    // 2. Validate input
    if err := h.validator.Struct(&req); err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
    }

    // 3. Call service
    result, err := h.yourEntityService.CreateYourEntity(c.Context(), req)
    if err != nil {
        return httpResponses.NewErrorResponse(c, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, err)
    }

    // 4. Return response
    return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Your entity created successfully", result)
}
```