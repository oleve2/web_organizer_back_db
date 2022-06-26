package backendServ

import (
	"database/sql"
	"fmt"
	"log"

	models "webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// ----------------------------
// knowbase
func (s *Service) GetAllPosts() ([]*models.PostDTO, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("select id, title, text, theme, part from posts order by theme, title")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	posts := make([]*models.PostDTO, 0)

	for rows.Next() {
		p := &models.PostDTO{}
		err := rows.Scan(&p.Id, &p.Title, &p.Text, &p.Theme, &p.Part)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	//
	return posts, nil
}

// ----------------------------
func (s *Service) GetOnePostByID(id string) (*models.PostDTO, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer db.Close()
	query := fmt.Sprintf(`select id, title, text, theme, part from posts where id=%s`, id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	posts := make([]*models.PostDTO, 0)

	for rows.Next() {
		p := &models.PostDTO{}
		err := rows.Scan(&p.Id, &p.Title, &p.Text, &p.Theme, &p.Part)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	//
	return posts[0], nil
}

// ----------------------------
func (s *Service) InsertNewPost(post *models.PostDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into posts (title, text, theme, part) values (?, ?, ?, ?)",
		post.Title, post.Text, post.Theme, post.Part,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ----------------------------
func (s *Service) UpdatePostById(post *models.PostDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("update posts set title=?, text=?, theme=?, part=? where id=?",
		post.Title, post.Text, post.Theme, post.Part, post.Id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ----------------------------
func (s *Service) DeletePostById(id string) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from posts where id=?", id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
