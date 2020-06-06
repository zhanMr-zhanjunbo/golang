package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
var (
	db orm.Ormer
)
type MovieInfo struct {
	Id                int
	MovieId           int
	MovieName         string
	MoviePic          string
	MovieDirector     string
	MovieWriter       string
	MovieStar         string
	MovieType         string
	MovieCountry      string
	MovieLanguage     string
	MovieOnTime       string
	MovieFilmLength   string
	MovieScore        string
	MovieScoreNumber  int
	CreateTime        string

} 
func init()  {
	orm.Debug=true
	orm.RegisterDataBase("default","mysql","root:microsys@tcp(localhost:3306)/test?charset=utf8",30)
	orm.RegisterModel(new(MovieInfo))
	db=orm.NewOrm()
}
func AddMovieInfo(m *MovieInfo)(int64,error) {
	n,err:=db.Insert(m)
	return n,err
}
func CheckSameName(name string)(int64,error){
	var m []MovieInfo
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("movie_info").Where("movie_name=?")
	sql:=qb.String()
	id,err:=db.Raw(sql,name).QueryRows(&m)
	return id,err
}