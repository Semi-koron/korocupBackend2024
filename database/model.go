package database

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	UserID    string `bun:",pk,notnull"`
	Name      string `bun:",notnull"`
	Email     string `bun:",unique,notnull"`
	Password  string `bun:",notnull"`
	Icon      int
	Profile   int
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",nullzero"`
}

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID        int       `bun:",pk,autoincrement"`
	UserID    string    `bun:",notnull"`
	User      *User     `bun:"rel:belongs-to,join:user_name=user_name"`
	Image     []byte    `bun:",notnull"`
	Reply     int       `bun:",nullzero"`
	Likes     int       `bun:",nullzero"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",nullzero"`
}

type Like struct {
	bun.BaseModel `bun:"table:likes,alias:l"`
	ID            int       `bun:",pk,autoincrement"`
	UserName      string    `bun:",notnull,unique:u_l"`
	User          *User     `bun:"rel:belongs-to,join:user_id=user_id"`
	PostID        int       `bun:",notnull,unique:u_l"`
	Post          *Post     `bun:"rel:belongs-to,join:post_id=id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
