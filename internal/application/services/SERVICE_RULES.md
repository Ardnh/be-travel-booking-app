# Service Rules

## Clean Architecture Principles

### 1. Service Layer Responsibilities
The service layer orchestrates business logic and coordinates between handlers and repositories.

- **Business Logic**: Implement domain rules and workflows
- **Data Transformation**: Convert DTOs to entities and vice versa
- **Error Handling**: Return appropriate domain errors to handlers
- **Transaction Management**: Handle database transactions when needed

---

## Service Implementation Patterns

### Service Structure
Each service must implement the corresponding interface from `internal/domain/services`:

```go
package services

import (
    "context"
    "errors"

    "github.com/ardnh/be-travel-booking-app/internal/application/dto"
    "github.com/ardnh/be-travel-booking-app/internal/domain/entities"
    "github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
    errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
    "github.com/google/uuid"
    "github.com/sirupsen/logrus"
)

type YourEntityServiceImpl struct {
    yourEntityRepository repositories.YourEntityRepository
    log                  *logrus.Logger
}

func NewYourEntityServiceImpl(yourEntityRepository repositories.YourEntityRepository, log *logrus.Logger) *YourEntityServiceImpl {
    return &YourEntityServiceImpl{
        yourEntityRepository: yourEntityRepository,
        log:                  log,
    }
}
```

---

## Error Handling

### Return Domain Errors
Services must return domain errors, not HTTP errors. Let the handler translate to HTTP responses:

```go
func (s *PoolPointServiceImpl) GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error) {
    poolPoint, err := s.poolPointRepository.GetPoolPointByID(ctx, poolID)
    if err != nil {
        if errors.Is(err, errorConst.ErrNotFound) {
            s.log.WithFields(logrus.Fields{
                "pool_id": poolID,
                "error":   err,
            }).Error("pool point not found")
            return nil, errorConst.ErrNotFound
        }
        return nil, err
    }
    return poolPoint, nil
}
```

### Error Types Reference
- `errorConst.ErrNotFound` - Resource does not exist
- `errorConst.ErrConflict` - Resource already exists
- `errorConst.ErrBadRequest` - Invalid input
- `errorConst.ErrUnauthorized` - Authentication required
- `errorConst.ErrForbidden` - Access denied
- `errorConst.ErrInternalServer` - Server-side error

---

## Logging Best Practices

### Structured Logging
Use `logrus.Fields` for structured, searchable logs:

```go
s.log.WithFields(logrus.Fields{
    "vendor_id": vendorID,
    "error":     err,
}).Error("vendor not found")
```

### Log Levels
- `Error` - Unexpected errors, failures
- `Warn` - Suspicious activity (e.g., login attempt with wrong credentials)
- `Info` - Successful operations

---

## Entity Creation

### Generate UUIDs
Always generate new UUIDs for new entities:

```go
poolPoint := &entities.Pools{
    PoolID:   uuid.New(),
    VendorID: vendorID,
    Name:     req.Name,
    // ...
}
```

### Map DTO to Entity
Manually map DTO fields to entity fields:

```go
func (s *PoolPointServiceImpl) CreatePoolPoint(ctx context.Context, req dto.CreatePoolsDTO) (*entities.Pools, error) {
    vendorID, err := uuid.Parse(req.VendorID)
    if err != nil {
        s.log.WithFields(logrus.Fields{
            "vendor_id": req.VendorID,
            "error":     err,
        }).Error("failed to parse vendor id")
        return nil, errorConst.ErrBadRequest
    }

    poolPoint := &entities.Pools{
        PoolID:      uuid.New(),
        VendorID:    vendorID,
        Name:        req.Name,
        Slug:        req.Slug,
        Address:     req.Address,
        City:        req.City,
        Province:    req.Province,
        District:    req.District,
        Latitude:    req.Latitude,
        Longitude:   req.Longitude,
        OpenTime:    req.OpenTime,
        CloseTime:   req.CloseTime,
        Status:      req.Status,
        Description: req.Description,
        EmbedURL:    req.EmbedURL,
    }

    return s.poolPointRepository.CreatePoolPoint(ctx, *poolPoint)
}
```

---

## Entity Update

### Fetch Before Update
Always fetch existing entity before updating:

```go
func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendorID uuid.UUID, req dto.UpdateVendorDTO) error {
    vendor, err := s.vendorRepository.GetVendorByID(ctx, vendorID)
    if err != nil {
        if errors.Is(err, errorConst.ErrNotFound) {
            s.log.WithFields(logrus.Fields{
                "vendor_id": vendorID,
                "error":     err,
            }).Error("vendor not found")
            return errorConst.ErrNotFound
        }
        return err
    }

    // Update only provided fields
    if req.BusinessName != nil {
        vendor.BusinessName = *req.BusinessName
    }
    if req.PhoneNumber != nil {
        vendor.PhoneNumber = *req.PhoneNumber
    }
    // ...

    return s.vendorRepository.UpdateVendor(ctx, *vendor)
}
```

---

## Transaction Management

### Database Operations with External Systems
When combining database operations with external systems (e.g., Casbin), ensure atomicity:

```go
func (s *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequestDto) (*dto.RegisterResponseDto, error) {
    existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
    if err != nil && !errors.Is(err, errorConst.ErrNotFound) {
        s.log.WithFields(logrus.Fields{
            "email": req.Email,
            "error": err,
        }).Error("failed to check existing user")
        return nil, errorConst.ErrInternalServer
    }

    if existingUser != nil {
        s.log.WithField("email", req.Email).Warn("registration attempt with existing email")
        return nil, errorConst.ErrUserAlreadyExists
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        s.log.WithField("email", req.Email).Error("failed to hash password")
        return nil, errorConst.ErrInternalServer
    }

    user := entities.Users{
        UserID:       uuid.New(),
        Name:         req.Name,
        Email:        req.Email,
        PasswordHash: string(hashedPassword),
        Phone:        req.Phone,
    }

    err = s.userRepository.CreateUser(ctx, user)
    if err != nil {
        s.log.WithFields(logrus.Fields{
            "email": req.Email,
            "error": err,
        }).Error("failed to create user")
        return nil, errorConst.ErrInternalServer
    }

    // Add role grouping to Casbin after successful DB insert
    _, err = s.casbinEnforcer.AddGroupingPolicy(user.UserID.String(), constants.RoleDailyUser)
    if err != nil {
        s.log.WithFields(logrus.Fields{
            "email":  req.Email,
            "userID": user.UserID,
            "error":  err,
        }).Error("failed to add casbin grouping policy")
        return nil, errorConst.ErrInternalServer
    }

    s.casbinEnforcer.LoadPolicy()

    return nil, nil
}
```

---

## Security Considerations

### Password Handling
Never store plain-text passwords. Use bcrypt for hashing:

```go
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
if err != nil {
    return nil, errorConst.ErrInternalServer
}
```

### Protect User Enumeration
Return consistent errors to prevent information leakage:

```go
func (s *AuthServiceImpl) Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginResponseDto, error) {
    user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
    if err != nil {
        if errors.Is(err, errorConst.ErrNotFound) {
            s.log.WithField("email", req.Email).Warn("login attempt with unregistered email")
            return nil, errorConst.ErrUnauthorized  // Not ErrNotFound
        }
        return nil, errorConst.ErrInternalServer
    }

    if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        s.log.WithField("email", req.Email).Warn("invalid credentials provided")
        return nil, errorConst.ErrUnauthorized  // Same error as above
    }
    // ...
}
```

---

## Service Template

When creating a new service, follow this template:

```go
package services

import (
    "context"
    "errors"

    "github.com/ardnh/be-travel-booking-app/internal/application/dto"
    "github.com/ardnh/be-travel-booking-app/internal/domain/entities"
    "github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
    errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
    "github.com/google/uuid"
    "github.com/sirupsen/logrus"
)

type YourEntityServiceImpl struct {
    yourEntityRepository repositories.YourEntityRepository
    log                  *logrus.Logger
}

func NewYourEntityServiceImpl(yourEntityRepository repositories.YourEntityRepository, log *logrus.Logger) *YourEntityServiceImpl {
    return &YourEntityServiceImpl{
        yourEntityRepository: yourEntityRepository,
        log:                  log,
    }
}

func (s *YourEntityServiceImpl) GetYourEntityByID(ctx context.Context, id uuid.UUID) (*entities.YourEntity, error) {
    entity, err := s.yourEntityRepository.GetYourEntityByID(ctx, id)
    if err != nil {
        if errors.Is(err, errorConst.ErrNotFound) {
            s.log.WithFields(logrus.Fields{
                "id":    id,
                "error": err,
            }).Error("entity not found")
            return nil, errorConst.ErrNotFound
        }
        return nil, err
    }
    return entity, nil
}

func (s *YourEntityServiceImpl) CreateYourEntity(ctx context.Context, req dto.CreateYourEntityDTO) (*entities.YourEntity, error) {
    // Map DTO to entity
    entity := &entities.YourEntity{
        ID:   uuid.New(),
        Name: req.Name,
        // ...
    }

    // Call repository
    created, err := s.yourEntityRepository.CreateYourEntity(ctx, *entity)
    if err != nil {
        s.log.WithFields(logrus.Fields{
            "name":  req.Name,
            "error": err,
        }).Error("failed to create entity")
        return nil, err
    }

    return created, nil
}
```