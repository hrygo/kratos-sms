package data

import (
  "context"

  "github.com/go-kratos/kratos/v2/log"

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

func (s smsRepo) SaveJournal(ctx context.Context, jo *biz.SmsJournal) (*biz.SmsJournal, error) {
  // TODO implement me
  panic("implement me")
}

func (s smsRepo) QueryJournal(ctx context.Context, queryId uint64) (*biz.SmsJournal, error) {
  // TODO implement me
  panic("implement me")
}

func (s smsRepo) FindTemplate(ctx context.Context, tempId string) (*biz.SmsTemplate, error) {
  // TODO implement me
  panic("implement me")
}
