package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"seriesTracker/src/database/models"
	"sync"
)

// Database manejada con SQLITE (ya no postgres porque me rendí intentando hacerla funcionar en el server)

// Singleton para no tener que estar abriendo nuevas conexiones a la DB
// en el server
var (
	db   *sql.DB
	once sync.Once
	err  error
)

func GetDB() (*sql.DB, error) {
	once.Do(func() {
		db, err = sql.Open("sqlite", "file:src/database/sql/series.db")
		if err != nil {
			return
		}
		err = db.Ping()
	})

	return db, err
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func Index() (*[]models.Serie, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM series")
	if err != nil {
		return nil, err
	}

	return scanSeriesRows(rows)
}

func SearchByName(query string) (*[]models.Serie, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM series WHERE name LIKE ? COLLATE NOCASE", "%"+query+"%")
	if err != nil {
		return nil, err
	}

	return scanSeriesRows(rows)
}

func scanSeriesRows(rows *sql.Rows) (*[]models.Serie, error) {
	defer rows.Close()

	var series []models.Serie
	for rows.Next() {
		var s models.Serie
		err := rows.Scan(
			&s.Id_serie,
			&s.Name,
			&s.Description,
			&s.Current_episode,
			&s.Total_episodes,
			&s.Img_src,
		)
		if err != nil {
			return nil, err
		}
		series = append(series, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &series, nil
}

func Show(id int) (*models.Serie, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	var serie models.Serie
	err = db.QueryRow("SELECT * FROM series WHERE id_serie=?", id).Scan(
		&serie.Id_serie,
		&serie.Name,
		&serie.Description,
		&serie.Current_episode,
		&serie.Total_episodes,
		&serie.Img_src,
	)
	if err != nil {
		return nil, err
	}

	return &serie, nil
}

func Store(s models.Serie) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	_, err2 := db.Exec(`INSERT INTO series 
                (name,
                description,
                current_episode,
                total_episodes,
                img_src) 
            VALUES (?, ?, ?, ?, ?)`,
		s.Name,
		s.Description,
		s.Current_episode,
		s.Total_episodes,
		s.Img_src)

	if err2 != nil {
		return err2
	}

	return nil
}

func Update(id int, s models.Serie) error {
	db, err = GetDB()
	if err != nil {
		return err
	}

	res, err2 := db.Exec(`
        UPDATE series
        SET name = ?,
            description = ?,
            current_episode = ?,
            total_episodes = ?,
            img_src = ?
        WHERE id_serie = ?
    `, s.Name, s.Description, s.Current_episode, s.Total_episodes, s.Img_src, id)

	if err2 != nil {
		return err2
	}

	rowsAffected, err3 := res.RowsAffected()
	if err3 != nil {
		return err3
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func Destroy(id int) error {
	db, err = GetDB()
	if err != nil {
		return err
	}

	res, err2 := db.Exec("DELETE FROM series WHERE id_serie = ?", id)
	if err2 != nil {
		return err2
	}

	rowsAffected, err3 := res.RowsAffected()
	if err3 != nil {
		return err3
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
