package modelDatabase

import (
	"time"

	PTGUdatabase "github.com/parinyapt/StreamySnap_AuthService/utils/database"
)

type Client struct {
	UUID      string    `gorm:"column:client_uuid;primary_key"`
	Name      string    `gorm:"column:client_name"`
	CreatedAt time.Time `gorm:"column:client_created_at"`
	UpdatedAt time.Time `gorm:"column:client_updated_at"`
}

func (Client) TableName() string {
	return PTGUdatabase.GenerateTableName("client")
}

type ClientService struct {
	UUID        string    `gorm:"column:client_service_uuid;primary_key"`
	SecretKey   string    `gorm:"column:client_service_secret_key"`
	ClientUUID  string    `gorm:"column:client_service_client_uuid;foreignkey:ClientUUID"`
	Name        string    `gorm:"column:client_service_name"`
	Status      string    `gorm:"column:client_service_status;type:enum('active', 'inactive');default:active"`
	CallBackURL string    `gorm:"column:client_service_callback_url"`
	CreatedAt   time.Time `gorm:"column:client_service_created_at"`
	UpdatedAt   time.Time `gorm:"column:client_service_updated_at"`
}

func (ClientService) TableName() string {
	return PTGUdatabase.GenerateTableName("client_service")
}

type Account struct {
	UUID      string    `gorm:"column:account_uuid;primary_key"`
	Name      string    `gorm:"column:account_name"`
	Image     string    `gorm:"column:account_image"`
	Email     string    `gorm:"column:account_email"`
	Password  string    `gorm:"column:account_password"`
	Status    string    `gorm:"column:account_status;type:enum('active', 'inactive');default:active"`
	CreatedAt time.Time `gorm:"column:account_created_at"`
	UpdatedAt time.Time `gorm:"column:account_updated_at"`
}

func (Account) TableName() string {
	return PTGUdatabase.GenerateTableName("account")
}

type AccountOAuth struct {
	ID          string    `gorm:"column:account_oauth_id;primary_key;auto_increment"`
	AccountUUID string    `gorm:"column:account_oauth_account_uuid"`
	Token       string    `gorm:"column:account_oauth_token"`
	Provider    string    `gorm:"column:account_oauth_provider"`
	CreatedAt   time.Time `gorm:"column:account_oauth_created_at"`
	UpdatedAt   time.Time `gorm:"column:account_oauth_updated_at"`
}

func (AccountOAuth) TableName() string {
	return PTGUdatabase.GenerateTableName("account_oauth")
}

type OAuthProvider struct {
	ID   string `gorm:"column:oauth_provider_id;primary_key;auto_increment"`
	Name string `gorm:"column:oauth_provider_name"`
}

func (OAuthProvider) TableName() string {
	return PTGUdatabase.GenerateTableName("oauth_provider")
}

type AuthPageSession struct {
	UUID              string    `gorm:"column:auth_page_session_uuid;primary_key"`
	ClientServiceUUID string    `gorm:"column:auth_page_session_client_service_uuid"`
	ExpiredAt         time.Time `gorm:"column:auth_page_session_expired_at"`
	CreatedAt         time.Time `gorm:"column:auth_page_session_created_at"`
}

func (AuthPageSession) TableName() string {
	return PTGUdatabase.GenerateTableName("auth_page_session")
}

type AuthTemporaryToken struct {
	UUID              string    `gorm:"column:auth_temporary_token_uuid;primary_key"`
	AccountUUID       string    `gorm:"column:auth_temporary_token_account_uuid"`
	ClientServiceUUID string    `gorm:"column:auth_temporary_token_client_service_uuid"`
	ExpiredAt         time.Time `gorm:"column:auth_temporary_token_expired_at"`
	CreatedAt         time.Time `gorm:"column:auth_temporary_token_created_at"`
}

func (AuthTemporaryToken) TableName() string {
	return PTGUdatabase.GenerateTableName("auth_temporary_token")
}

type AuthHistory struct {
	UUID               string    `gorm:"column:auth_history_uuid;primary_key"`
	AccountUUID        string    `gorm:"column:auth_history_account_uuid"`
	ClientServiceUUID  string    `gorm:"column:auth_history_client_service_uuid"`
	RefreshToken       string    `gorm:"column:auth_history_refresh_token"`
	RefreshTokenExpire time.Time `gorm:"column:auth_history_refresh_token_expire"`
	ExpiredAt          time.Time `gorm:"column:auth_history_expired_at"`
	CreatedAt          time.Time `gorm:"column:auth_history_created_at"`
}

func (AuthHistory) TableName() string {
	return PTGUdatabase.GenerateTableName("auth_history")
}
