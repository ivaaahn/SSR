package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Profile struct {
	*BasePgRepo
}

func NewProfile(pg *postgres.Postgres, l logger.Interface) *Profile {
	return &Profile{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Profile) GetStProfile(email string) (*entity.StProfile, error) {
	const query = `
	select *
	from users u 
		join students s using (user_id)
	where email = $1
	`

	student := entity.StProfile{}

	err := repo.Conn.Get(&student, query, email)
	if err != nil {
		err := fmt.Errorf("Profile->GetStProfile->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &student, nil
}

func (repo *Profile) GetSvProfile(email string) (*entity.SvProfile, error) {
	const query = `
	select *
	from users u 
		join supervisors s using (user_id)
	where email = $1
	`

	supervisor := entity.SvProfile{}

	err := repo.Conn.Get(&supervisor, query, email)
	if err != nil {
		err := fmt.Errorf("Profile->GetSvProfile->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &supervisor, nil
}
