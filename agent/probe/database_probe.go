package probe

import (
	"context"
	"database/sql"
	"github.com/your-project/models"
)

type DatabaseProbe struct {
	proxyPort int
	dbConns   map[string]*sql.DB
}

func (p *DatabaseProbe) InjectFault(exp *models.Experiment) error {
	switch exp.Action {
	case models.ActionDBTimeout:
		return p.injectTimeout(exp)
	case models.ActionDBError:
		return p.injectError(exp)
	case models.ActionDBSlow:
		return p.injectLatency(exp)
	default:
		return ErrUnsupportedAction
	}
} 