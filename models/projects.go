package models

import (
	"database/sql"
	// "fmt"
)

type Project struct {
	ID                int            `json:"id"`
	Primary_marker_id int            `json:"primary_marker_id"`
	Contractor_id     int            `json:"contractor_id"`
	Pm_id             int            `json:"pm_id"`
	Title             string         `json:"title"`
	Start_date        sql.NullString `json:"start_date"`
	End_date          sql.NullString `json:"end_date"`
}

type GetProject struct {
	ID                    int            `json:"id"`
	Primary_marker_id     int            `json:"primary_marker_id`
	Contractor_id         int            `json:"contractor_id"`
	Pm_id                 int            `json:"pm_id"`
	Title                 string         `json:"title"`
	Start_date            sql.NullString `json:"start_date"`
	End_date              sql.NullString `json:"end_date"`
	Contractor            string         `json:"contractor"`
	Contractor_phone      string         `json:"contractor_phone"`
	Contractor_cell       string         `json:"contractor_cell"`
	Contractor_email      string         `json:"contractor_email"`
	Project_manager       string         `json:"project_manager"`
	Project_manager_cell  string         `json:"project_manager_cell"`
	Project_manager_phone string         `json:"project_manager_phone"`
	Project_manager_email string         `json:"project_manager_email"`
	Color                 string         `json:"color"`
}

// TaskCollection is collection of Tasks
type ProjectCollection struct {
	Projects []Project `json:"projects"`
}

type GetProjectCollection struct {
	GetProjects []GetProject `json:"projects"`
}

func GetProjects(db *sql.DB) GetProjectCollection {

	sql := `SELECT projects.id,
    primary_marker_id,
    contractor_id,
    pm_id, title,
    start_date,
    end_date,
    CONCAT(B.first_name, B.last_name) as project_manager,
    B.phone as project_manager_phone,
    B.cell as project_manager_cell,
    B.email as project_manager_email,
    CONCAT(A.first_name, ' ', A.last_name) as contractor,
    A.phone as contractor_phone,
    A.cell as contractor_cell,
    A.email as contractor_email,
    contractors.color
    FROM projects 
    LEFT JOIN contractors on contractors.user_id = contractor_id
    LEFT JOIN project_managers on project_managers.user_id = pm_id
    LEFT JOIN users A on contractors.user_id = A.id
    LEFT JOIN users B on project_managers.user_id = B.id;`

	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := GetProjectCollection{}

	for rows.Next() {
		project := GetProject{}
		err2 := rows.Scan(
			&project.ID,
			&project.Primary_marker_id,
			&project.Contractor_id,
			&project.Pm_id,
			&project.Title,
			&project.Start_date,
			&project.End_date,
			&project.Project_manager,
			&project.Project_manager_phone,
			&project.Project_manager_cell,
			&project.Project_manager_email,
			&project.Contractor,
			&project.Contractor_phone,
			&project.Contractor_cell,
			&project.Contractor_email,
			&project.Color)

		if err2 != nil {
			panic(err2)
		}

		result.GetProjects = append(result.GetProjects, project)
	}
	// fmt.Printf("%+v", result)
	return result
}

func PostProject(db *sql.DB, project Project) (int64, error) {
	sql := `
    INSERT INTO projects(primary_marker_id, contractor_id, pm_id, title, start_date, end_date) 
    VALUES(?, ?, ?, ?, ?, ?)
    `
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(project.Primary_marker_id, project.Contractor_id, project.Pm_id, project.Title, project.Start_date, project.End_date)
	// Exit if we get an error
	checkErr(err2)

	return result.LastInsertId()
}

func PutProject(db *sql.DB, project Project, id int) (int64, error) {
	sql := `
    UPDATE projects SET primary_marker_id = ?, contractor_id = ?, pm_id = ?, title = ?, start_date = ?, end_date = ? where id = ?
    `
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(project.Primary_marker_id, project.Contractor_id, project.Pm_id, project.Title, project.Start_date, project.End_date, id)
	// Exit if we get an error
	checkErr(err2)

	return result.LastInsertId()
}

func DeleteProject(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM projects WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	checkErr(err2)

	return result.RowsAffected()
}
