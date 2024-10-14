package entity

import "time"

func (AuthUser) TableName() string {
	return "auth_user"
}

type AuthUser struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName       string    `gorm:"size:255;not null" json:"first_name"`
	LastName        string    `gorm:"size:255;not null" json:"last_name"`
	Email           string    `gorm:"size:250;not null;unique" json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Username        string    `gorm:"size:255;not null" json:"username"`
	Password        string    `gorm:"size:255;not null" json:"password"`
	IsSuperuser     bool      `gorm:"not null" json:"is_superuser"`
	IsStaff         bool      `gorm:"not null" json:"is_staff"`
	IsActive        bool      `gorm:"not null" json:"is_active"`
	Photo           string    `gorm:"size:500;not null" json:"photo"`
	DateJoined      time.Time `gorm:"not null" json:"date_joined"`
	RememberToken   string    `gorm:"size:100" json:"remember_token"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `gorm:"size:100" json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       string    `gorm:"size:100" json:"updated_by"`
	DeletedAt       time.Time `json:"deleted_at"`
	DeletedBy       string    `gorm:"size:100" json:"deleted_by"`
}

func (AuthAudit) TableName() string {
	return "auth_audit"
}

type AuthAudit struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy        string    `gorm:"size:50" json:"created_by"`
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `gorm:"size:100" json:"updated_by"`
	Username         string    `gorm:"size:100" json:"username"`
	UserID           uint64    `gorm:"not null" json:"user_id"`
	IPAddress        string    `gorm:"size:50" json:"ip_address"`
	ActivityCategory string    `gorm:"size:100;not null" json:"activity_category"`
	ActivityDate     time.Time `gorm:"not null;default:'0000-00-00 00:00:00'" json:"activity_date"`
	ActivityName     string    `gorm:"size:100;not null" json:"activity_name"`
	ActivityRefCode  string    `gorm:"size:100;not null" json:"activity_ref_code"`
	ActivityStatus   string    `gorm:"size:10;not null" json:"activity_status"`
	DeviceInfo       string    `gorm:"size:100" json:"device_info"`
	DeviceID         string    `gorm:"size:100" json:"device_id"`
	ErrorCode        string    `gorm:"size:10;not null" json:"error_code"`
	ErrorDesc        string    `gorm:"size:1024" json:"error_desc"`
}

func (AuthCompany) TableName() string {
	return "auth_company"
}

type AuthCompany struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:100" json:"name"`
	Description string    `gorm:"size:1000" json:"description"`
	Address     string    `gorm:"size:500" json:"address"`
	Phone       string    `gorm:"size:10" json:"phone"`
	Email       string    `gorm:"size:100" json:"email"`
	Domain      string    `gorm:"size:100" json:"domain"`
	Village     string    `gorm:"size:100" json:"village"`
	District    string    `gorm:"size:100" json:"district"`
	City        string    `gorm:"size:100" json:"city"`
	Province    string    `gorm:"size:100" json:"province"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	PostalCode  string    `gorm:"size:100" json:"postal_code"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `gorm:"size:100" json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   string    `gorm:"size:100" json:"deleted_by"`
}

type AuthUserCompanyPrivilege struct {
	UserID      uint64 `gorm:"not null" json:"user_id"`
	CompanyID   uint64 `gorm:"not null" json:"company_id"`
	PrivilegeID uint64 `gorm:"not null" json:"privilege_id"`
}
func (AuthUserCompanyRole) TableName() string {
	return "auth_user_company_role"
}
type AuthUserCompanyRole struct {
	UserID    uint64 `gorm:"not null" json:"user_id"`
	CompanyID uint64 `gorm:"not null" json:"company_id"`
	RoleID    uint64 `gorm:"not null" json:"role_id"`
}

type AuthPrivilege struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:100" json:"name"`
	Description string    `gorm:"size:1000" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `gorm:"size:100" json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   string    `gorm:"size:100" json:"deleted_by"`
}

type AuthRole struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:100" json:"name"`
	Description string    `gorm:"size:1000" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `gorm:"size:100" json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   string    `gorm:"size:100" json:"deleted_by"`
}
