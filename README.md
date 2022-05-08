# SSR (Student-Supervisor Relationship) - упрости взаимодействие с научным руководителем

## Идея

За время своего обучения в учебном заведении *стдент* зачастую сталкиваются с проблемой выбора научного руководителя и дальнейшего взаимодействия с ним для выполнения той или иной работы (НИР, Курсовая, Дипломная), а *научный руководитель* - с проблемой отслеживания прогресса студента. 

**Идея проекта** - сделать жизнь студентов и научных руководителей проще, предоставив удобную систему взаимодействия и поиска.

## Актуальность

**Актуальность проекта** проверена на личном опыте. Когда подходит время выбирать научного руководителя студент попросту не знает многих преподавателей, поэтому не имеет целостной картины для принятия решения. А кафедра обычно этому никак не содействует, назначая научного руководителя случайным образом. Студент оказывается вынужден взаимодействовать с человеком, который ему не по душе, откуда следует низкая мотивация студента выполнять работу. А научный руководитель зачастую даже или не знает своих подопечных, или ввиду своего возраста не помнит, не говоря уже о каком-то контроле прогресса.


## Аналоги

**Аналогичных проектов** найти не удалось. Единственный возможный вариант сегодя -- по кусочкам собирать информацию из различных источников о преподавателях. А научному руководителю для отслеживания прогресса самостоятельно вести на каждого студента таблицу успеваемости.


## Разбиение на компоненты

## **Слои**
1. Delivery - Controllers
2. Services - Business Logic
3. Repos - DAL


## **Компоненты**
1. Auth - регистрация, авторизация


## **Контроллеры**
* *AuthUser(login, pswd)*

* *GetSvWorksList()*
* *GetSvWork()*
* *GetStudentProgress()*
* *AcceptStudentWaypoint()*
* *GetWorksList()*
* *GetStudentWorksList()*
* *GetStudentWork()*
* *GetAvailableSupervisors()*
* *GetSupervisorInfo()*


* *GetSvBidsList()*
* *GetSvBid()*
* *HandleSvBid()*
* *CreateBid()*
* *GetStudentBidsList()*
* *GetStudentBid()*
* *CancelStudentBid()*


* *GetStudentProfile()*
* *GetSvProfile()*
* *PatchStudentProfile()*
* *PatchSvProfile()*


* *GetSupervisorsList()*
* *GetWorksList()*
* *GetWorkInfo()*
* *CreateWork()*
* *GetWorkStat()*
* *UpdateWorkInfo()*
* *AssignSupervisorOnWork()*
* *AssignSupervisorOnWorkBulk()*



    




## **Services**
1. **WorkService**
   * *GetSvWorksList(sv_id: int) -> list[SvWorksDTO]*
   * *GetSvWork(sv_id: int, work_id: int) -> SvWorkDTO*
   * *GetStudentProgress(sv_id: int, work_id: int, st_id: int) -> StudentProgressDTO*
   * *AcceptStudentWaypoint(id: int)*
   * *GetWorksList(subject_id: int) -> list[WorkInfoDTO]*
   * *GetStudentWorksList(sv_id: int) -> list[StudentWorkInfoDTO]*
   * *GetStudentWork(sv_id: int, work_id: int) -> StudentWorkInfoDTO*
   * *GetAvailableSupervisors(work_id: int) -> list[SupervisorInfoDTO]*
   * *GetSupervisorInfo(sv_id: int) -> SupervisorInfoDTO*


2. **BidService**
   * *GetSvBidsList(supervisor_id: int) -> list[SvBidInfoDTO]*
   * *GetSvBid(supervisor_id: int, bid_id: int) -> SvBidInfoDTO*
   * *HandleSvBid(bid_id: int, action: Enum)*
   * *CreateBid(bid: BidInfoDTO)*
   * *GetStudentBidsList(supervisor_id: int) -> list[StudentBidInfoDTO]*
   * *GetStudentBid(supervisor_id: int, bid_id: int) -> StudentBidInfoDTO*
   * *CancelStudentBid(bid_id: int)*


3. **ProfileService**
   * *GetStudentProfile(user_id: int) -> StudentProfileDTO* 
   * *GetSvProfile(user_id: int) -> SvProfileDTO*
   * *PatchStudentProfile(changes: StudentProfileDTO)*
   * *PatchSvProfile(changes: SvProfileDTO)*


4. **HeadSupervisorService**
   * *GetSupervisorsList() -> list[SupervisorInfoDTO]*
   * *GetWorksList() -> list[WorkInfoDTO]*
   * *GetWorkInfo(work_id: int) -> list[WorkInfoDTO]*
   * *CreateWork(data: WorkCreateDTO) -> int*
   * *GetWorkStat(work_id: int) -> WorkStatForHeadSvDTO*
   * *UpdateWorkInfo(data: WorkInfoDTO)*
   * *AssignSupervisorOnWork(work_id: int, sv_id: int)*
   * *AssignSupervisorOnWorkBulk(work_id: int, supervisors_ids: list[int])*
   

## Models
1. **CreateSSRModel:**
   * work_id
   * sv_id
   * student_id

2. **StudentModel**
   * email
   * first_name
   * last_name
   * photo
   * student_card
   * year
   * department: DepartmentModel

3. **SvModel**
   * email
   * first_name
   * last_name
   * photo
   * about
   * birthdate
   * department: DepartmentModel

4. **DepartmentModel:**
   * name

**Репозитории**
1. UsersRepo
   * Get(user_id: int)
   * GetPlenty()
2. SupervisorsRepo (use UsersRepo)
3. StudentsRepo (use UsersRepo)
   * Get(student_id: int, ext: bool)
   * 
4. SSRRepo
   * GetPlenty(req_role: Enum, req_id: int, sort, offset, order, ...)
   * Get(bid_id: int)
   * Create(model: CreateSSRModel)
5. WorksRepo
   * 



**Flows**
1. Список работ рук-ля: WorkController -> WorkService.GetSupervisorWorks(email) -> Join 