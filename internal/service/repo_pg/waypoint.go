package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Waypoint struct {
	*BasePgRepo
}

func NewWaypointRepo(pg *postgres.Postgres, l logger.Interface) *Waypoint {
	return &Waypoint{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Waypoint) GetPlenty(workID int) ([]*entity.Waypoint, error) {
	query := `
	select * from waypoints w 
	where w.work_id = $1;
	`

	var waypoints []*entity.Waypoint

	err := repo.Conn.Select(&waypoints, query, workID)
	if err != nil {
		err := fmt.Errorf("Waypoint->getPlenty->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return waypoints, nil
}
