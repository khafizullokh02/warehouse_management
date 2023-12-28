// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	zero "gopkg.in/guregu/null.v4/zero"
)

type ActionType string

const (
	ActionTypeSell   ActionType = "sell"
	ActionTypeBuy    ActionType = "buy"
	ActionTypeRefund ActionType = "refund"
	ActionTypeDefect ActionType = "defect"
	ActionTypeNone   ActionType = "none"
)

func (e *ActionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ActionType(s)
	case string:
		*e = ActionType(s)
	default:
		return fmt.Errorf("unsupported scan type for ActionType: %T", src)
	}
	return nil
}

type NullActionType struct {
	ActionType ActionType `json:"action_type"`
	Valid      bool       `json:"valid"` // Valid is true if ActionType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullActionType) Scan(value interface{}) error {
	if value == nil {
		ns.ActionType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ActionType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullActionType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ActionType), nil
}

type AgreementFormsStatus string

const (
	AgreementFormsStatusAccepted   AgreementFormsStatus = "accepted"
	AgreementFormsStatusPending    AgreementFormsStatus = "pending"
	AgreementFormsStatusInProgress AgreementFormsStatus = "in_progress"
	AgreementFormsStatusFailed     AgreementFormsStatus = "failed"
	AgreementFormsStatusRejected   AgreementFormsStatus = "rejected"
	AgreementFormsStatusDone       AgreementFormsStatus = "done"
	AgreementFormsStatusNone       AgreementFormsStatus = "none"
)

func (e *AgreementFormsStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AgreementFormsStatus(s)
	case string:
		*e = AgreementFormsStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AgreementFormsStatus: %T", src)
	}
	return nil
}

type NullAgreementFormsStatus struct {
	AgreementFormsStatus AgreementFormsStatus `json:"agreement_forms_status"`
	Valid                bool                 `json:"valid"` // Valid is true if AgreementFormsStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAgreementFormsStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AgreementFormsStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AgreementFormsStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAgreementFormsStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AgreementFormsStatus), nil
}

type CurrencyCode string

const (
	CurrencyCodeUsd  CurrencyCode = "usd"
	CurrencyCodeUzs  CurrencyCode = "uzs"
	CurrencyCodeUe   CurrencyCode = "u.e"
	CurrencyCodeNone CurrencyCode = "none"
)

func (e *CurrencyCode) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CurrencyCode(s)
	case string:
		*e = CurrencyCode(s)
	default:
		return fmt.Errorf("unsupported scan type for CurrencyCode: %T", src)
	}
	return nil
}

type NullCurrencyCode struct {
	CurrencyCode CurrencyCode `json:"currency_code"`
	Valid        bool         `json:"valid"` // Valid is true if CurrencyCode is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrencyCode) Scan(value interface{}) error {
	if value == nil {
		ns.CurrencyCode, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CurrencyCode.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrencyCode) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CurrencyCode), nil
}

type EntryGroupsStatus string

const (
	EntryGroupsStatusInitial    EntryGroupsStatus = "initial"
	EntryGroupsStatusInProgress EntryGroupsStatus = "in_progress"
	EntryGroupsStatusInDelivery EntryGroupsStatus = "in_delivery"
	EntryGroupsStatusDelivered  EntryGroupsStatus = "delivered"
	EntryGroupsStatusAccepted   EntryGroupsStatus = "accepted"
	EntryGroupsStatusNone       EntryGroupsStatus = "none"
)

func (e *EntryGroupsStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EntryGroupsStatus(s)
	case string:
		*e = EntryGroupsStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EntryGroupsStatus: %T", src)
	}
	return nil
}

type NullEntryGroupsStatus struct {
	EntryGroupsStatus EntryGroupsStatus `json:"entry_groups_status"`
	Valid             bool              `json:"valid"` // Valid is true if EntryGroupsStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEntryGroupsStatus) Scan(value interface{}) error {
	if value == nil {
		ns.EntryGroupsStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EntryGroupsStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEntryGroupsStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EntryGroupsStatus), nil
}

type PricingType string

const (
	PricingTypeRetail    PricingType = "retail"
	PricingTypeWholesale PricingType = "wholesale"
	PricingTypeNone      PricingType = "none"
)

func (e *PricingType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PricingType(s)
	case string:
		*e = PricingType(s)
	default:
		return fmt.Errorf("unsupported scan type for PricingType: %T", src)
	}
	return nil
}

type NullPricingType struct {
	PricingType PricingType `json:"pricing_type"`
	Valid       bool        `json:"valid"` // Valid is true if PricingType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPricingType) Scan(value interface{}) error {
	if value == nil {
		ns.PricingType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PricingType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPricingType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PricingType), nil
}

type Account struct {
	ID        int32            `json:"id"`
	UserID    int32            `json:"user_id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type AgreementForm struct {
	ID                     int32      `json:"id"`
	FromAccount            int32      `json:"from_account"`
	ToAccount              int32      `json:"to_account"`
	ProductIds             []int32    `json:"product_ids"`
	ActionTypeForAgreement ActionType `json:"action_type_for_agreement"`
	WholesalePrice         float64    `json:"wholesale_price"`
	RetailPrice            float64    `json:"retail_price"`
}

type EntryGroup struct {
	ID                int32             `json:"id"`
	Quantity          int32             `json:"quantity"`
	ActionType        ActionType        `json:"action_type"`
	PricingType       PricingType       `json:"pricing_type"`
	Price             float64           `json:"price"`
	Currency          CurrencyCode      `json:"currency"`
	EntryGroupsStatus EntryGroupsStatus `json:"entry_groups_status"`
	CreatedAt         pgtype.Timestamp  `json:"created_at"`
	UpdatedAt         pgtype.Timestamp  `json:"updated_at"`
	DeletedAt         zero.Time         `json:"deleted_at"`
}

type EntryItem struct {
	ID           int32  `json:"id"`
	ProductID    int32  `json:"product_id"`
	EntryGroupID int32  `json:"entry_group_id"`
	SupCode      string `json:"sup_code"`
}

type Product struct {
	ID             int32            `json:"id"`
	Name           string           `json:"name"`
	SupCode        string           `json:"sup_code"`
	BarCode        string           `json:"bar_code"`
	Image          string           `json:"image"`
	Brand          string           `json:"brand"`
	WholesalePrice float64          `json:"wholesale_price"`
	RetailPrice    float64          `json:"retail_price"`
	Discount       float64          `json:"discount"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
}

type User struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt zero.Time        `json:"deleted_at"`
}
