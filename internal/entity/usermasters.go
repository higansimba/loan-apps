package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NotificationStatusInvoice mewakili preferensi notifikasi untuk invoice
type NotificationStatusInvoice struct {
	Email  bool `json:"email" bson:"email"`
	SMS    bool `json:"sms" bson:"sms"`
	Call   bool `json:"call" bson:"call"`
	Others bool `json:"others" bson:"others"`
}

// SubscriptionData mewakili keseluruhan data langganan user
type Usermaster struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Enduser        primitive.ObjectID `bson:"enduser"`
	ParentID       primitive.ObjectID `bson:"parent_id"`
	PhoneNumber    string             `bson:"phone_number"`
	UserType       string             `bson:"user_type"`
	AccountStatus  string             `bson:"account_status"`
	Master         bool               `bson:"master"`
	Flag           string             `json:"flag"`
	TestingPurpose bool               `bson:"testing_purpose"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	DeleteAt       time.Time          `bson:"deleted_at,omitempty"`
	DeleteSource   string             `bson:"delete_source,omitempty"`
}

func (b *Usermaster) CollectionName() string {
	return "usermasters"
}

type Application struct {
	ID   primitive.ObjectID `bson:"_id"`
	User primitive.ObjectID `bson:"user"`
}

func (a *Application) CollectionName() string {
	return "applications"
}

type Billingmaster struct {
	ID                        primitive.ObjectID        `bson:"_id" json:"_id"`
	NotificationStatusInvoice NotificationStatusInvoice `bson:"notification_status_invoice" json:"notification_status_invoice"`
	Application               primitive.ObjectID        `bson:"application" json:"application"`
	PaymentPeriod             string                    `bson:"payment_period" json:"payment_period"`
	NextBillingDate           time.Time                 `bson:"next_billing_date" json:"next_billing_date"`
	PackageSubscription       string                    `bson:"package_subscription" json:"package_subscription"`
	PackageDetail             string                    `bson:"package_detail" json:"package_detail"`
	BillingType               string                    `bson:"billing_type" json:"billing_type"`
	AccountStatus             string                    `bson:"account_status" json:"account_status"`
	AccountCode               string                    `bson:"accountcode" json:"accountcode"`
	BillingDate               time.Time                 `bson:"billing_date" json:"billing_date"`
	CreatedAt                 time.Time                 `bson:"created_at" json:"created_at"`
	UpdatedAt                 time.Time                 `bson:"updated_at" json:"updated_at"`
	Version                   int                       `bson:"__v" json:"__v"`
	ReminderExpiredDate       time.Time                 `bson:"reminder_expired_date" json:"reminder_expired_date"`
	Error                     string                    `bson:"error" json:"error"`
	Processed                 bool                      `bson:"processed" json:"processed"`
	IsNew                     bool                      `bson:"is_new" json:"is_new"`
	SubscriptionType          string                    `bson:"subscription_type" json:"subscription_type"`
	UserType                  string                    `bson:"user_type" json:"user_type"`
	GracePeriodMonth          int                       `bson:"grace_period_month" json:"grace_period_month"`
	AIPlan                    string                    `bson:"ai_plan" json:"ai_plan"`
	AIPlanStatus              string                    `bson:"ai_plan_status" json:"ai_plan_status"`
	GracePeriodDate           time.Time                 `bson:"grace_period_date" json:"grace_period_date"`
	DeletedAt                 time.Time                 `bson:"ai_plus_account_deleted_at,omitempty"`
	TestingPurpose            bool                      `bson:"testing_purpose"`
	DeleteSource              string                    `bson:"ai_plus_account_source,omitempty"`
}

func (a *Billingmaster) CollectionName() string {
	return "billingmasters"
}

type UserAggregation struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Email         string             `bson:"email" json:"email"`
	EndUser       primitive.ObjectID `bson:"enduser" json:"enduser"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number"`
	ParentID      primitive.ObjectID `bson:"parent_id" json:"parent_id"`
	UserType      string             `bson:"user_type" json:"user_type"`
	AccountStatus string             `bson:"account_status" json:"account_status"`
	Master        bool               `bson:"master" json:"master"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
	Source        string             `bson:"delete_source"`
	DeletedAt     time.Time          `bson:"deleted_at"`
	V             int                `bson:"__v" json:"__v"`
	Application   struct {
		ID        primitive.ObjectID `bson:"_id" json:"_id"`
		APIKey    string             `bson:"api_key" json:"api_key"`
		User      primitive.ObjectID `bson:"user" json:"user"`
		Type      string             `bson:"type" json:"type"`
		CreatedAt time.Time          `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
		V         int                `bson:"__v" json:"__v"`
	} `bson:"application" json:"application"`
	Idx int `bson:"idx" json:"idx"`
	// BillingMaster (from $lookup)
	BillingMaster struct {
		ID                     primitive.ObjectID `bson:"_id" json:"_id"`
		Application            primitive.ObjectID `bson:"application" json:"application"`
		PaymentPeriod          string             `bson:"payment_period" json:"payment_period"`
		NextBillingDate        time.Time          `bson:"next_billing_date" json:"next_billing_date"`
		PackageSubscription    string             `bson:"package_subscription" json:"package_subscription"`
		PackageDetail          string             `bson:"package_detail" json:"package_detail"`
		BillingType            string             `bson:"billing_type" json:"billing_type"`
		AccountStatus          string             `bson:"account_status" json:"account_status"`
		AccountCode            string             `bson:"accountcode" json:"accountcode"`
		BillingDate            time.Time          `bson:"billing_date" json:"billing_date"`
		CreatedAt              time.Time          `bson:"created_at" json:"created_at"`
		UpdatedAt              time.Time          `bson:"updated_at" json:"updated_at"`
		V                      int                `bson:"__v" json:"__v"`
		ReminderExpiredDate    time.Time          `bson:"reminder_expired_date" json:"reminder_expired_date"`
		Error                  string             `bson:"error" json:"error"`
		Processed              bool               `bson:"processed" json:"processed"`
		IsNew                  bool               `bson:"is_new" json:"is_new"`
		SubscriptionType       string             `bson:"subscription_type" json:"subscription_type"`
		UserType               string             `bson:"user_type" json:"user_type"`
		GracePeriodMonth       int                `bson:"grace_period_month" json:"grace_period_month"`
		AIPlan                 string             `bson:"ai_plan" json:"ai_plan"`
		AIPlanStatus           string             `bson:"ai_plan_status" json:"ai_plan_status"`
		GracePeriodDate        time.Time          `bson:"grace_period_date" json:"grace_period_date"`
		NotificationStatusBill struct {
			Email  bool `bson:"email" json:"email"`
			SMS    bool `bson:"sms" json:"sms"`
			Call   bool `bson:"call" json:"call"`
			Others bool `bson:"others" json:"others"`
		} `bson:"notification_status_invoice" json:"notification_status_invoice"`
	} `bson:"billingmaster" json:"billingmaster"`
	Idxbmaster int `bson:"idxbmaster" json:"idxbmaster"`
}

type UserAggregationV2 struct {
	ID                  primitive.ObjectID `bson:"_id" json:"id"`
	Application         primitive.ObjectID `bson:"application" json:"application"`
	NextBillingDate     time.Time          `bson:"next_billing_date" json:"next_billing_date"`
	BillingType         string             `bson:"billing_type" json:"billing_type"`
	AccountStatus       string             `bson:"account_status" json:"account_status"`
	AccountCode         string             `bson:"accountcode" json:"accountcode"`
	BillingDate         time.Time          `bson:"billing_date" json:"billing_date"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updated_at"`
	ReminderExpiredDate time.Time          `bson:"reminder_expired_date" json:"reminder_expired_date"`
	IsNew               bool               `bson:"is_new" json:"is_new"`
	SubscriptionType    string             `bson:"subscription_type" json:"subscription_type"`
	GracePeriodMonth    int                `bson:"grace_period_month" json:"grace_period_month"`
	AIPlan              string             `bson:"ai_plan" json:"ai_plan"`
	AIPlanStatus        string             `bson:"ai_plan_status" json:"ai_plan_status"`
	GracePeriodDate     time.Time          `bson:"grace_period_date" json:"grace_period_date"`
	DeletedAt           time.Time          `bson:"ai_plus_account_deleted_at,omitempty"`
	DeleteSource        string             `bson:"ai_plus_account_source,omitempty"`

	UserMasters Umaster `json:"usermaster"`
}

type Umaster struct {
	ID           primitive.ObjectID `json:"usermaster_id"`
	Email        string             `json:"email"`
	DeletedAt    time.Time          `json:"deleted_at,omitempty"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	DeleteSource string             `json:"delete_source,omitempty"`
}
