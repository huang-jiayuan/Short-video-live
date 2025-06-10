package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"user-server/basic/global"
	"user-server/models"
	__ "user-server/proto"
)

func (s *Server) VideoWorksList(ctx context.Context, in *__.VideoWorksListRequest) (*__.VideoWorksListResponse, error) {
	w := &models.VideoWorks{}
	page := in.Page
	if page < 1 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize < 10 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	works, err := w.FindWorks(offset, pageSize)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var Item []*__.VideoWorksList
	for _, m := range works {
		Item = append(Item, &__.VideoWorksList{
			Title:       m.Title,
			Desc:        m.Desc,
			MusicId:     m.MusicId,
			IpAddress:   m.IpAddress,
			BrowseCount: m.BrowseCount,
		})
	}
	return &__.VideoWorksListResponse{List: Item}, nil
}
func (s *Server) AddVideoWorks(_ context.Context, in *__.AddVideoWorksRequest) (*__.AddVideoWorksResponse, error) {
	if strings.TrimSpace(in.Title) == "" {
		return nil, errors.New("视频标题不能为空")
	}
	if len(in.Title) > 100 {
		return nil, errors.New("视频标题不能超过100个字符")
	}
	if len(in.Desc) > 255 {
		return nil, errors.New("视频描述不能超过255个字符")
	}
	w := models.VideoWorks{}

	err := w.AddVideoWorks(in)
	if err != nil {
		return nil, errors.New("作品发布失败")
	}
	return &__.AddVideoWorksResponse{Greet: "发布成功，作品正在审核中"}, nil
}
func (s *Server) CreateVideoWorksComment(_ context.Context, in *__.CreateVideoWorksCommentRequest) (*__.CreateVideoWorksCommentResponse, error) {

	c := &models.VideoWorkComment{}
	w := &models.VideoWorks{}
	works, err := w.FindVideoWorks(in.WorkId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if works.Id == 0 {
		return nil, errors.New("该视频不存在")
	}
	tx := *global.DB.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tag := in.Tag
	if works.CommentCount == 0 {
		tag = 1
	}
	err = c.CreateComment(in.WorkId, in.UserId, in.Pid, tag, in.Content)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("评论发布失败")
	}
	count := works.CommentCount + 1
	err = w.UpdateCommentCount(in.WorkId, int64(count))
	if err != nil {
		tx.Rollback()
		return nil, errors.New("修改成功")
	}
	tx.Commit()
	return &__.CreateVideoWorksCommentResponse{Greet: "评价发布成功"}, nil
}
func (s *Server) WorksItem() {
	
}
