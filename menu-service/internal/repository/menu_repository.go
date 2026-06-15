package repository

import (
    "database/sql"
)

type Menu struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Pic   string  `json:"pic"`
}

func FetchMenus(db *sql.DB) ([]Menu, error) {
    rows, err := db.Query("SELECT id, name, price, pic FROM menu")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var menus []Menu
    for rows.Next() {
        var m Menu
        // Scan data kolom 'pic' ke dalam struct properti m.Pic
        if err := rows.Scan(&m.ID, &m.Name, &m.Price, &m.Pic); err != nil {
            return nil, err
        }
        menus = append(menus, m)
    }
    return menus, nil
}