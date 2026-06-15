package repository

import (
    "database/sql"
    "back-end/internal/models"
)

// FetchMenus mengambil semua data dari tabel menus
func FetchMenus(db *sql.DB) ([]models.Menu, error) {
    rows, err := db.Query("SELECT id, name, price,pic FROM menu")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var menus []models.Menu
    for rows.Next() {
        var m models.Menu
        if err := rows.Scan(&m.ID, &m.Name, &m.Price ,&m.Pic); err != nil {
            return nil, err
        }
        menus = append(menus, m)
    }
    return menus, nil
}