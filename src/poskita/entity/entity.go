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
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"size:100" json:"name"`
	Description string     `gorm:"size:1000" json:"description"`
	Address     string     `gorm:"size:500" json:"address"`
	Phone       string     `gorm:"size:10" json:"phone"`
	Email       string     `gorm:"size:100" json:"email"`
	Domain      string     `gorm:"size:100" json:"domain"`
	Village     string     `gorm:"size:100" json:"village"`
	District    string     `gorm:"size:100" json:"district"`
	City        string     `gorm:"size:100" json:"city"`
	Province    string     `gorm:"size:100" json:"province"`
	UserID      uint64     `gorm:"not null" json:"user_id"`
	PostalCode  string     `gorm:"size:100" json:"postal_code"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   string     `gorm:"size:100" json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   string     `gorm:"size:100" json:"deleted_by"`
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

func (AuthPrivilege) TableName() string {
	return "auth_privilage"
}

type AuthPrivilege struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"size:100" json:"name"`
	Description string     `gorm:"size:1000" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   string     `gorm:"size:100" json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   string     `gorm:"size:100" json:"deleted_by"`
}

func (AuthRole) TableName() string {
	return "auth_role"
}

type AuthRole struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"size:100" json:"name"`
	Description string     `gorm:"size:1000" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `gorm:"size:100" json:"created_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   string     `gorm:"size:100" json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   string     `gorm:"size:100" json:"deleted_by"`
}

// master data
func (Product) TableName() string {
	return "m_product"
}

type Product struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedBy    string    `json:"deleted_by"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	CategoryID   int64     `json:"category_id"`
	IsPreOrder   bool      `json:"is_pre_order"`
	IsSecondhand bool      `json:"is_secondhand"`
	Rating       float32   `json:"rating"`
	MinimumOrder int32     `json:"minimum_order"`
	PreOrderDay  int32     `json:"pre_order_day"`
	SupplierID   int64     `json:"supplier_id"`
	CompanyID    int64     `json:"company_id"`
}

func (ProductCategory) TableName() string {
	return "m_product_category"
}

type ProductCategory struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CompanyID   uint64     `json:"company_id"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy   string     `json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedBy   string     `gorm:"null" json:"deleted_by"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (ProductStock) TableName() string {
	return "m_product_stock"
}

type ProductStock struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedBy  string    `json:"deleted_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	TotalStock int32     `json:"total_stock"`
	KeepStock  int32     `json:"keep_stock"`
	ShowStock  string    `json:"show_stock"`
	UnitPrice  float64   `json:"unit_price"`
}

func (ProductSupplier) TableName() string {
	return "m_product_supplier"
}

type ProductSupplier struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	CompanyID   uint64     `json:"company_id"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy   string     `json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedBy   string     `json:"deleted_by"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (ProductVariation) TableName() string {
	return "m_product_variation"
}

type ProductVariation struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedBy  string    `json:"deleted_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	ProductID  uint64    `json:"product_id"`
	VarianName string    `json:"varian_name"`
}

func (ProductVariationOption) TableName() string {
	return "m_product_variation_option"
}

type ProductVariationOption struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy        string    `json:"created_by"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy        string    `json:"updated_by"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	DeletedBy        string    `json:"deleted_by"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
	ProductVarianID  uint64    `json:"product_varian_id"`
	VarianOptionName string    `json:"varian_option_name"`
	SKU              string    `json:"sku"`
	Price            float64   `json:"price"`
	ProductStockID   uint64    `json:"product_stock_id"`
}

func (ProductVariationImage) TableName() string {
	return "m_product_variation_image"
}

type ProductVariationImage struct {
	ID                       uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductVariationOptionID int64     `json:"product_varian_option_id"`
	VarianImage              string    `json:"varian_image"`
	CreatedBy                string    `json:"created_by"`
	CreatedAt                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy                string    `json:"updated_by"`
	UpdatedAt                time.Time `json:"updated_at,omitempty"`
	DeletedBy                string    `json:"deleted_by"`
	DeletedAt                time.Time `json:"deleted_at,omitempty"`
}

func (MAppConfig) TableName() string {
	return "m_app_config"
}

type MAppConfig struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy   string     `json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" gorm:"autoUpdateTime" json:"updated_at"`
	DeletedBy   string     `json:"deleted_by"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	Code        string     `json:"code"`
	Module      string     `json:"module"`
	CompanyID   string     `json:"company_id"`
	Value       string     `json:"value"`
	Description string     `gorm:"type:text" json:"description"`
}
