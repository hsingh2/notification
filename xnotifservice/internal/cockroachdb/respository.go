package cockroachdb

import (
	"context"
	"database/sql"
	"errors"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"github.com/cockroachdb/cockroach-go/crdb"
	log "github.com/sirupsen/logrus"
)

//Error
var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db  *sql.DB
	log *log.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *sql.DB, logger *log.Logger) (template.Repository, error) {
	// return  repository
	return &repository{
		db:  db,
		log: logger,
	}, nil
}

// CreateOrder inserts a new order and its order items into db
func (repo repository) CreateNotificationTemplate(ctx context.Context, template template.NotificationTemplate) error {
	// Run a transaction to sync the query model.
	err := crdb.ExecuteTx(ctx, repo.db, nil, func(tx *sql.Tx) error {
		return createNotificationTemplate(tx, template)
	})
	if err != nil {
		return err
	}
	return nil
}

func createNotificationTemplate(tx *sql.Tx, template template.NotificationTemplate) error {
	return nil
}

// CountNotificationTemplate ...
func (repo repository) CountNotificationTemplate(ctx context.Context) (int, error) {
	count := 0
	return count, nil
}

// GetNotificationTemplateByID ...
func (repo repository) GetNotificationTemplateByID(ctx context.Context, id string) (template.NotificationTemplate, error) {
	notificationTemplate := template.NotificationTemplate{}
	return notificationTemplate, nil
}

//GetByPage returns all the Notification Template by page id
func (repo repository) GetByPage(context.Context) ([]template.NotificationTemplate, error) {
	return []template.NotificationTemplate{}, nil
}

//Update Notification Template
func (repo repository) UpdateNotificationTemplate(context.Context, template.NotificationTemplate) (template.NotificationTemplate, error) {
	return template.NotificationTemplate{}, nil
}

//Delete Notificaiton Template
func (repo repository) DeleteNotificationTemplate(context.Context, string) error {
	return nil
}

// Close implements DB.Close
func (repo repository) Close() error {
	return repo.db.Close()
}
