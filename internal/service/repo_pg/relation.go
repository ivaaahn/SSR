package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Relation struct {
	*BasePgRepo
}

func NewRelation(pg *postgres.Postgres, l logger.Interface) *Relation {
	return &Relation{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Relation) Get(id int) (*entity.Relation, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status,
		ssr.created_at,

		sv.department_id as "sv.department_id",
		sv.birthdate as "sv.birthdate",
		sv.about as "sv.about",
		
		svu.user_id as "sv.user.user_id",
		svu.email as "sv.user.email",
		svu.last_name as "sv.user.last_name",
		svu.first_name as "sv.user.first_name",
		
		st.year as "st.year",
		st.department_id as "st.department_id",

		stu.user_id as "st.user.user_id",
		stu.email as "st.user.email",
		stu.last_name as "st.user.last_name",
		stu.first_name as "st.user.first_name",

		w.semester as "work.semester",
		w.description as "work.description",
		w.work_id as "work.work_id",
		wk.name as "work.work_kind.name",
		wk.work_kind_id as "work.work_kind.work_kind_id",
		
		subj.subject_id as "work.subject.subject_id",
		subj.name as "work.subject.name",
		subj.department_id as "work.subject.department_id"
	from ssr
		join supervisors sv on ssr.supervisor_id = sv.user_id
		join students st on ssr.student_id = st.user_id
		join users stu on st.user_id = stu.user_id
		join users svu on sv.user_id = svu.user_id
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.ssr_id = $1;
	`

	relation := entity.Relation{}

	err := repo.Conn.Get(&relation, query, id)
	if err != nil {
		err := fmt.Errorf("Relation->Get->repo.Conn.Get: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &relation, nil
}

func (repo *Relation) GetRelationsBySupervisorID(supervisorID int) ([]*entity.RelationShort, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status,
		
		sv.department_id as "sv.department_id",
		svu.user_id as "sv.user.user_id",
		svu.last_name as "sv.user.last_name",
		svu.first_name as "sv.user.first_name",
		
		st.year as "st.year",
		st.department_id as "st.department_id",
		stu.user_id as "st.user.user_id",
		stu.last_name as "st.user.last_name",
		stu.first_name as "st.user.first_name",

		w.semester as "work.semester",
		w.description as "work.description",
		w.work_id as "work.work_id",
		wk.name as "work.work_kind.name",
		wk.work_kind_id as "work.work_kind.work_kind_id",
		
		subj.subject_id as "work.subject.subject_id",
		subj.name as "work.subject.name",
		subj.department_id as "work.subject.department_id"
	from ssr 
		join supervisors sv on ssr.supervisor_id = sv.user_id
		join students st on ssr.student_id = st.user_id
		join users stu on st.user_id = stu.user_id
		join users svu on sv.user_id = svu.user_id
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.supervisor_id = $1;
	`

	var relations []*entity.RelationShort

	err := repo.Conn.Select(&relations, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("Relation->GetRelations->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return relations, nil
}

func (repo *Relation) GetRelationsByStudentID(studentID int) ([]*entity.RelationShort, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status,
		
		sv.department_id as "sv.department_id",
		svu.user_id as "sv.user.user_id",
		svu.last_name as "sv.user.last_name",
		svu.first_name as "sv.user.first_name",
		
		st.year as "st.year",
		st.department_id as "st.department_id",
		stu.user_id as "st.user.user_id",
		stu.last_name as "st.user.last_name",
		stu.first_name as "st.user.first_name",

		w.semester as "work.semester",
		w.description as "work.description",
		w.work_id as "work.work_id",
		wk.name as "work.work_kind.name",
		wk.work_kind_id as "work.work_kind.work_kind_id",
		
		subj.subject_id as "work.subject.subject_id",
		subj.name as "work.subject.name",
		subj.department_id as "work.subject.department_id"
	from ssr 
		join supervisors sv on ssr.supervisor_id = sv.user_id
		join students st on ssr.student_id = st.user_id
		join users stu on st.user_id = stu.user_id
		join users svu on sv.user_id = svu.user_id
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.student_id = $1;
	`

	var relations []*entity.RelationShort

	err := repo.Conn.Select(&relations, query, studentID)
	if err != nil {
		err := fmt.Errorf("Relation->GetRelations->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return relations, nil
}

func (repo *Relation) Create(studentID, supervisorID, workID int) (int, error) {
	query := `
	insert into ssr (student_id, supervisor_id, work_id) 
	values ($1, $2, $3)
	returning ssr_id;
	`

	var relationID int
	err := repo.Conn.QueryRowx(query, studentID, supervisorID, workID).Scan(&relationID)
	if err != nil {
		err := fmt.Errorf("RelationRepo->Create->repo.Conn.QueryRowx: %w", err)
		repo.l.Error(err)
		return 0, err
	}

	return relationID, nil
}

func (repo *Relation) Update(id int, status entity.StatusSSR) (int, error) {
	query := `
	update ssr set status = $1
	where ssr_id = $2
	returning ssr_id;
	`

	var relationID int
	err := repo.Conn.QueryRowx(query, status, id).Scan(&relationID)
	if err != nil {
		err := fmt.Errorf("Relation->Update->repo.Conn.QueryRowx: %w", err)
		repo.l.Error(err)
		return 0, err
	}

	return relationID, nil
}

func (repo *Relation) GetRelationStatus(studentID, workID int) (entity.StatusSSR, error) {
	const query = `
	select ssr.status
	from ssr where ssr.student_id = $1 and ssr.work_id = $2;
	`

	var status entity.StatusSSR

	err := repo.Conn.Get(&status, query, studentID, workID)
	if err != nil {
		err := fmt.Errorf("Relation->GetRelationStatus->r.Conn.Select: %w", err)
		repo.l.Error(err)
		return "", err
	}

	return status, nil
}
