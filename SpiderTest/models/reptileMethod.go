package models

import (
	"regexp"
	"strconv"
)
func GetMoveId(movieHtml string)int{
	reg:=regexp.MustCompile(`https://movie.douban.com/subject/(.*)/`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return 0
	}
	id,_:=strconv.Atoi(result[0][1])
	return id
}
func GetMoveName(movieHtml string)string {
	reg:=regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	return string(result[0][1])
}
func GetMoviePic(movieHtml string) string {
	reg:=regexp.MustCompile(`<img\s*src="(.*?)"`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var picSets []string
	for _,v:=range result{
		picSets=append(picSets,v[1])
	}
	return picSets[1]
}
func GetMovieDirector(movieHtml string) string {
	reg:=regexp.MustCompile(`<a\s*href="/celebrity/(.*)/"\s*rel="v:directedBy">(.*)</a>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var directorSets string=""
	for _,v:=range result{
		directorSets+=v[2]
	}
	return directorSets
}
func GetMovieWriter(movieHtml string) string {
	reg:=regexp.MustCompile(`<a\s*href="/celebrity/(.*?)">(.*?)</a>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var writerSets []string
	for _,v:=range result{
		writerSets=append(writerSets,v[2])
	}
	return writerSets[0]
}
func GetMovieStar(movieHtml string) string {
	reg:=regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var starSets string=""
	for _,v:=range result{
		starSets+=v[1]+"/"
	}
	return starSets
}
func GetMovieType (movieHtml string) string {
	reg:=regexp.MustCompile(`<span\s*property="v:genre">(.*?)</span>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var typeSets string=""
	for _,v:=range result{
		typeSets+=v[1]+"/"
	}
	return typeSets
}
func GetMovieCountry  (movieHtml string) string {
	reg:=regexp.MustCompile(`<span\s*class="pl">制片国家/地区:</span>(.*)<br/>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	return result[0][1]
}
func GetMovieLanguage (movieHtml string) string {
	reg:=regexp.MustCompile(`<span\s*class="pl">语言:</span>(.*)<br/>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	return result[0][1]
}
func GetMovieOnTime (movieHtml string) string {
	reg:=regexp.MustCompile(`<span\s*property="v:initialReleaseDate" content="(.*?)"`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var timeSets string=""
	for _,v:=range result{
		timeSets+=v[1]+"/"
	}
	return timeSets
}
func GetMovieFilmLength (movieHtml string) string {
	reg:=regexp.MustCompile(`<span\s*property="v:runtime"\s*content="(.*)">(.*)</span>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	var filmLength string=""
	for _,v:=range result{
		filmLength+=v[2]
	}
	return filmLength
}
func GetMovieScore (movieHtml string) string {
	reg:=regexp.MustCompile(`<strong\s*class="ll rating_num"\s*property="v:average">(.*)</strong>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return ""
	}
	return result[0][1]
}
func GetMovieScoreNumber(movieHtml string)int{
	reg:=regexp.MustCompile(`<span\s*property="v:votes">(.*)</span>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result)==0{
		return 0
	}
	scoreNumber,_:=strconv.Atoi(result[0][1])
	return scoreNumber
}
func GetMovieURL(movieHtml string) []string {
	reg:=regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)
	var urlSets []string
	for _,v:=range result{
		urlSets=append(urlSets,v[1])
	}
	return urlSets
}