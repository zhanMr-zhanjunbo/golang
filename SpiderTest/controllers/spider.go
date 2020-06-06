package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"SpiderTest/models"
	"time"
	"fmt"
)

type MovieSpiderController struct {
	beego.Controller
}
func CheckErr(err error){
	if err!=nil{
		//fmt.Println(err.Error())
		panic(err)
	}
}
func (c *MovieSpiderController)GetSpider()  {
	  c.EnableRender=false
	  var movieInfo models.MovieInfo
	  models.CreateConnect("127.0.0.1:6379")
      url:="https://movie.douban.com/subject/30318116/"
      models.PushInQueue(url)
      for{
      	length:=models.GetQueueLength()
      	if length==0{
      		break
		}
		movieHtml:=models.PopOutQueue()
		if models.IsVisit(movieHtml){
			  continue
		}
		html,err:=httplib.Get(movieHtml).String()
		CheckErr(err)
		movieInfo=models.MovieInfo{}
		movieName:=models.GetMoveName(html)
		if movieName!=""{
			id,_:=models.CheckSameName(movieName)
			if id==0{
				movieInfo.MovieId          =  models.GetMoveId(html)
				movieInfo.MovieName        =  movieName
				movieInfo.MoviePic         =  models.GetMoviePic(html)
				movieInfo.MovieDirector    =  models.GetMovieDirector(html)
				movieInfo.MovieWriter      =  models.GetMovieWriter(html)
				movieInfo.MovieStar        =  models.GetMovieStar(html)
				movieInfo.MovieType        =  models.GetMovieType(html)
				movieInfo.MovieCountry     =  models.GetMovieCountry(html)
				movieInfo.MovieLanguage    =  models.GetMovieLanguage(html)
				movieInfo.MovieOnTime      =  models.GetMovieOnTime(html)
				movieInfo.MovieFilmLength  =  models.GetMovieFilmLength(html)
				movieInfo.MovieScore       =  models.GetMovieScore(html)
				movieInfo.MovieScoreNumber =  models.GetMovieScoreNumber(html)
				movieInfo.CreateTime       =  time.Now().Format("2006-01-02 15:04:05")
				_,err:=models.AddMovieInfo(&movieInfo)
				CheckErr(err)
			}else{
				continue
			}
		}
		  urls:=models.GetMovieURL(html)
		  for _,urlList:=range urls{
			  models.PushInQueue(urlList)
		  }
		  c.Ctx.WriteString(fmt.Sprintf("%v",urls))
		  models.AddToSet(movieHtml)
		  time.Sleep(time.Second)
	  }
	  c.Ctx.WriteString("end of spider!")

}