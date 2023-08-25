package routes

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type RouteRepo struct {
	ID        null.String `db:"id"`
	Name      null.String `db:"name"`
	SourceLat null.Float  `db:"source_lat"`
	SourceLng null.Float  `db:"source_lng"`
	DestLat   null.Float  `db:"dest_lat"`
	DestLng   null.Float  `db:"dest_lng"`
}

type RoutesRepositoryInterface interface {
	CreateRoute(data RouteRepo) (id null.Int, err error)
	GetRoute(id null.Int) (route RouteRepo, err error)
	FindAll() ([]*RouteRepo, error)
}

type routesRepository struct {
	db *sql.DB
}

func NewRoutesRepository(db *sql.DB) RoutesRepositoryInterface {
	return &routesRepository{
		db: db,
	}
}

func (r *routesRepository) CreateRoute(data RouteRepo) (id null.Int, err error) {

	sql := "INSERT INTO routes (name, source_lat, source_lng, dest_lat, dest_lng) VALUES (?,?,?,?,?)"

	result, err := r.db.Exec(sql, data.Name, data.SourceLat, data.SourceLng, data.DestLat, data.DestLng)
	if err != nil {
		logrus.Errorf("Error on create route: %v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		logrus.Errorf("Error on get last insert id: %v", err)
		return
	}

	id = null.IntFrom(int64(lastInsertID))
	return
}

func (r *routesRepository) GetRoute(id null.Int) (route RouteRepo, err error) {
	const query = `SELECT * FROM routes WHERE id = ?`

	err = r.db.QueryRow(query, id).Scan(&route.ID, &route.Name, &route.SourceLat, &route.SourceLng, &route.DestLat, &route.DestLng)
	if err != nil {
		logrus.Errorf("Error on get route: %v", err)
		return
	}

	return
}

func (r *routesRepository) FindAll() ([]*RouteRepo, error) {
	sqlSmt := "SELECT * FROM routes"

	rows, err := r.db.Query(sqlSmt)
	if err != nil {
		logrus.Errorf("Error on get routes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var routes []*RouteRepo
	for rows.Next() {
		var route RouteRepo

		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.SourceLat,
			&route.SourceLng,
			&route.DestLat,
			&route.DestLng,
		)
		if err != nil {
			logrus.Errorf("Error on get routes: %v", err)
			return nil, err
		}
		routes = append(routes, &route)
	}

	return routes, nil
}
