package dao

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDb    *sql.DB
	MysqlDbErr error
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
)

func InitDB() {
	MysqlDb, MysqlDbErr = sql.Open("mysql", Username+":"+Password+"@tcp("+Host+":"+Port+")/"+Database+"?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if MysqlDbErr != nil {
		panic(MysqlDbErr.Error())
	} else {
		fmt.Println("connect success!!")
	}
}

func InsertNews(news News) (bool, error) {
	res, err := MysqlDb.Exec("INSERT IGNORE INTO news(类别,标题,采集时间,网址)VALUES(?,?,?,?)", news.Category, news.Title, news.Time, news.Url)
	if err != nil {
		return false, err
	} else {
		i, _ := res.RowsAffected()
		if i == 0 {
			return false, nil
		}
		return true, nil
	}
}

func GetAllNews() ([]News, error) {
	newslist := make([]News, 0)
	rows, err := MysqlDb.Query("SELECT * FROM news  ORDER BY 采集时间 DESC")
	if err != nil {
		return nil, err
	}
	var news News
	for rows.Next() {
		rows.Scan(&news.Category, &news.Title, &news.Time, &news.Url)
		newslist = append(newslist, news)
	}
	return newslist, nil
}

func GetSearchNews(query string) ([]News, error) {
	newslist := make([]News, 0)
	rows, err := MysqlDb.Query("SELECT 类别,标题,采集时间,网址 FROM news where 标题 like ? ORDER BY 采集时间 DESC ", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	var news News
	for rows.Next() {
		rows.Scan(&news.Category, &news.Title, &news.Time, &news.Url)
		newslist = append(newslist, news)
	}
	return newslist, nil
}

func GetCategoryNews(query string) ([]News, error) {
	newslist := make([]News, 0)
	rows, err := MysqlDb.Query("SELECT 类别,标题,采集时间,网址 FROM news where 类别= ? ORDER BY 采集时间 DESC ", query)
	if err != nil {
		return nil, err
	}
	var news News
	for rows.Next() {
		rows.Scan(&news.Category, &news.Title, &news.Time, &news.Url)
		newslist = append(newslist, news)
	}
	return newslist, nil
}

func DeleteNews(query string) (bool, error) {
	_, err := MysqlDb.Exec("DELETE FROM news where 标题=?", query)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func AlterNews(query string, change string) (bool, error) {
	_, err := MysqlDb.Exec("UPDATE news set 标题=? where 标题=?", change, query)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func DeleteOverdueNews(query time.Time) (int64, error) {
	res, err := MysqlDb.Exec("DELETE FROM news where 采集时间<?", query)
	if err != nil {
		return 0, err
	} else {
		rows, _ := res.RowsAffected()
		return rows, nil
	}
}
