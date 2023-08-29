package models

import(
    "fmt"
    "sync"
    "time"
)



// 数据库的表项
type Video struct{
    VideoId   int64`gorm:"primaryKey"`
    UserId    int64 
    playUrl    string
    CreateAt    time.Time
    UpdateAt    time.Time
    DeleteAt     time.Time
}



type VideoDao struct{}

var videoDao *VideoDao

var videoOnce sync.Once




func (*VideoDao)VideoDaoInstance()*VideoDao {
    videoOnce.Do(
        func(){
            videoDao = &VideoDao{}
        })
    return videoDao
}

// 创建 video

func (*VideoDao) CreateVideo(video *Video) (*Video, error) {
	/*Video := Video{Name: Videoname, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := SqlSession.Create(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}




// 根据UserId，查出Video列表
func (d *VideoDao) FindVideoById(id int64) (*Video, error) {

	video := Video{VideoId: id}



	result := SqlSession.Where("Video_id = ?", id).First(&video)

	err := result.Error

	if err != nil {

		return nil, err

	}

	return &video, err

}



// 根据UserId，查出Video列表

func (*VideoDao) QueryVideoByUserId(userId int64) ([]*Video, error) {

	var videos []*Video

	err := SqlSession.Where("user_id = ?", userId).Find(&videos).Error

	if err != nil {

		fmt.Println("查询Video列表失败")

		return nil, err

	}

	return videos, nil

}



// 根据时间和需要查询的条数，获取video列表

func (*VideoDao) QueryVideo(date *string, limit int) []*Video {

	fmt.Println(*date)

	var VideoList []*Video

	SqlSession.Where("create_at < ?", *date).Order("create_at desc").Find(&VideoList)

	if len(VideoList) <= limit {

		fmt.Println(VideoList)

		return VideoList

	}

	fmt.Println(VideoList)

	return VideoList[0:limit]

}

