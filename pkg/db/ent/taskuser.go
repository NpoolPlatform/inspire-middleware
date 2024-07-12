// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/taskuser"
	"github.com/google/uuid"
)

// TaskUser is the model entity for the TaskUser schema.
type TaskUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID uuid.UUID `json:"task_id,omitempty"`
	// EventID holds the value of the "event_id" field.
	EventID uuid.UUID `json:"event_id,omitempty"`
	// TaskState holds the value of the "task_state" field.
	TaskState string `json:"task_state,omitempty"`
	// RewardInfo holds the value of the "reward_info" field.
	RewardInfo string `json:"reward_info,omitempty"`
	// RewardState holds the value of the "reward_state" field.
	RewardState string `json:"reward_state,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case taskuser.FieldID, taskuser.FieldCreatedAt, taskuser.FieldUpdatedAt, taskuser.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case taskuser.FieldTaskState, taskuser.FieldRewardInfo, taskuser.FieldRewardState:
			values[i] = new(sql.NullString)
		case taskuser.FieldEntID, taskuser.FieldAppID, taskuser.FieldUserID, taskuser.FieldTaskID, taskuser.FieldEventID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskUser fields.
func (tu *TaskUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taskuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tu.ID = uint32(value.Int64)
		case taskuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tu.CreatedAt = uint32(value.Int64)
			}
		case taskuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tu.UpdatedAt = uint32(value.Int64)
			}
		case taskuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tu.DeletedAt = uint32(value.Int64)
			}
		case taskuser.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				tu.EntID = *value
			}
		case taskuser.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				tu.AppID = *value
			}
		case taskuser.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				tu.UserID = *value
			}
		case taskuser.FieldTaskID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				tu.TaskID = *value
			}
		case taskuser.FieldEventID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value != nil {
				tu.EventID = *value
			}
		case taskuser.FieldTaskState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field task_state", values[i])
			} else if value.Valid {
				tu.TaskState = value.String
			}
		case taskuser.FieldRewardInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reward_info", values[i])
			} else if value.Valid {
				tu.RewardInfo = value.String
			}
		case taskuser.FieldRewardState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reward_state", values[i])
			} else if value.Valid {
				tu.RewardState = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TaskUser.
// Note that you need to call TaskUser.Unwrap() before calling this method if this TaskUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (tu *TaskUser) Update() *TaskUserUpdateOne {
	return (&TaskUserClient{config: tu.config}).UpdateOne(tu)
}

// Unwrap unwraps the TaskUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tu *TaskUser) Unwrap() *TaskUser {
	_tx, ok := tu.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskUser is not a transactional entity")
	}
	tu.config.driver = _tx.drv
	return tu
}

// String implements the fmt.Stringer.
func (tu *TaskUser) String() string {
	var builder strings.Builder
	builder.WriteString("TaskUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tu.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.UserID))
	builder.WriteString(", ")
	builder.WriteString("task_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.TaskID))
	builder.WriteString(", ")
	builder.WriteString("event_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.EventID))
	builder.WriteString(", ")
	builder.WriteString("task_state=")
	builder.WriteString(tu.TaskState)
	builder.WriteString(", ")
	builder.WriteString("reward_info=")
	builder.WriteString(tu.RewardInfo)
	builder.WriteString(", ")
	builder.WriteString("reward_state=")
	builder.WriteString(tu.RewardState)
	builder.WriteByte(')')
	return builder.String()
}

// TaskUsers is a parsable slice of TaskUser.
type TaskUsers []*TaskUser

func (tu TaskUsers) config(cfg config) {
	for _i := range tu {
		tu[_i].config = cfg
	}
}
