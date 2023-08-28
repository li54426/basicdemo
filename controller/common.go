package controller


// 类型设计

type User struct{
  Id int64 `json:"id, omitempty"`
  Name string `json:"name, omitempty"`
  
}


type Video struct {
  Id int64 `json:"id, omitempty"`
  Author User `json:"author"`
  PlayUrl string `json:"paly_url" json:"play_url, omitempty"`
}


// type File Video



type Response struct{
    StatusCode int32 `json:"status_code"`
    StatusMsg   string `json:"status_msg,omitempty"`
}