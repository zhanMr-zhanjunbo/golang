主键ID
电影ID             movie_id
电影名称           movie_name
电影图片           movie_pic
导演姓名           movie_director
编剧               movie_writer
主演               movie_star
类型               movie_type
地点               movie_country
语言               movie_language
上映时间           movie_show_time
片长               movie_film_length
豆瓣评分           movie_score
评分人数           movie_score_number
创建时间           create_time
use test;
CREATE table movie_info(
   id                   int(11) UNSIGNED auto_increment,
   movie_id             int(11) UNSIGNED NOT NULL COMMENT '电影id',
   movie_name           varchar(255) COMMENT '电影名称',
   movie_pic            varchar(255) COMMENT '电影图片',
   movie_director       varchar(255) COMMENT '电影导演',
   movie_writer         varchar(255) COMMENT '电影编剧',
   movie_star           varchar(500) COMMENT '电影主演',                                        //演员过多,255个字节不够
   movie_type           varchar(255) COMMENT '电影类型',
   movie_country        varchar(255) COMMENT '电影地点',
   movie_language       varchar(255) COMMENT '电影语言',
   movie_on_time        timestamp DEFAULT '1970-12-01 01:01:00' COMMENT '电影上映时间',         //mysql版本问题不能设置时间'0000-00-00 00:00:00'
   movie_film_length    varchar(255) COMMENT '电影片长',
   movie_score          varchar(255) COMMENT '电影评分',
   movie_score_number   int(11) UNSIGNED NOT NULL COMMENT '评分人数',
   remark               varchar(255) default '' COMMENT '备注',
   create_time          timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   modify_time          timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
   PRIMARY KEY(id),
   KEY(movie_id),
   key(create_time),
   KEY(modify_time)
);