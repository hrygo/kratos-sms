package data

import (
  "context"
  "time"

  "github.com/go-kratos/kratos/v2/log"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"

  "kratos-sms/internal/biz"
)

var _ biz.SmsRepo = (*smsRepo)(nil)

type smsRepo struct {
  data *Data
  log  *log.Helper
}

// NewSmsRepo .
func NewSmsRepo(data *Data, logger log.Logger) biz.SmsRepo {
  return &smsRepo{
    data: data,
    log:  log.NewHelper(logger),
  }
}

func (s *smsRepo) SaveJournal(ctx context.Context, jo *biz.SmsJournal) (*biz.SmsJournal, error) {
  ictx, cancel := context.WithTimeout(ctx, time.Second)
  defer cancel()

  result, err := s.journal().InsertOne(ictx, jo)
  if err != nil {
    s.log.WithContext(ictx).Error(err)
    return nil, err
  }
  id, ok := result.InsertedID.(primitive.ObjectID)
  if ok {
    jo.Id = id.Hex()
  }
  return jo, nil
}

func (s *smsRepo) QueryJournal(ctx context.Context, queryId uint64) (*biz.SmsJournal, error) {
  // TODO implement me
  panic("implement me")
}

func (s *smsRepo) FindTemplate(ctx context.Context, tempId string) (*biz.SmsTemplate, error) {
  // TODO implement me
  panic("implement me")
}

func (s *smsRepo) journal() *mongo.Collection {
  return s.data.Collection("sms", "journal")
}

func (s *smsRepo) template() *mongo.Collection {
  return s.data.Collection("sms", "template")
}

func (s *smsRepo) auth() *mongo.Collection {
  return s.data.Collection("sms", "auth")
}
