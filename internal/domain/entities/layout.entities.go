package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type CellType string
type CellStatus string
type WindowPosition string

const (
	CellTypeSeat   CellType = "seat"
	CellTypeDriver CellType = "driver"
	CellTypeAisle  CellType = "aisle"
)

const (
	StatusAvailable CellStatus = "available"
	StatusBooked    CellStatus = "booked"
	StatusBlocked   CellStatus = "blocked"
	StatusSelected  CellStatus = "selected"
)

type Cell struct {
	Type           CellType       `json:"type"`
	Row            int            `json:"row"`
	Col            int            `json:"col"`
	ID             string         `json:"id"`
	Status         CellStatus     `json:"status"`
	IsWindow       bool           `json:"isWindow"`
	WindowPosition WindowPosition `json:"windowPosition"`
}

type LayoutConfig [][]Cell

// ============================================================
// Layout
// ============================================================
type Layouts struct {
	LayoutID     uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"layout_id"`
	Name         string       `gorm:"type:varchar(255);not null" json:"name"`
	GridSizeX    int          `gorm:"type:int;not null" json:"grid_size_x"`
	GridSizeY    int          `gorm:"type:int;not null" json:"grid_size_y"`
	SeatCount    int          `gorm:"type:int;not null" json:"seat_count"`
	LayoutConfig LayoutConfig `gorm:"type:jsonb;not null;default:'[]'" json:"layout_config"`
	CreatedBy    *uuid.UUID   `gorm:"type:uuid;column:created_by" json:"created_by,omitempty"`

	// Relations
	BaseModel
}

func (Layouts) TableName() string { return "layouts" }

func (lc LayoutConfig) Value() (driver.Value, error) {
	if lc == nil {
		return "[]", nil
	}

	b, err := json.Marshal(lc)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (lc *LayoutConfig) Scan(value any) error {
	var bytes []byte

	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	case nil:
		*lc = LayoutConfig{} // nil dari DB → empty slice
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return json.Unmarshal(bytes, lc)
}
