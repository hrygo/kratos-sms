package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"kratos-sms/internal/biz"
	"kratos-sms/internal/conf"
	mylog "kratos-sms/internal/log"
)

var _ biz.SmsRepo = (*smsRepo)(nil)

type smsRepo struct {
	conf config.Config
	bs   *conf.Bootstrap
	data *Data
	log  *log.Helper
}

func (s *smsRepo) Replace(logger *log.Helper) {
	s.log = logger
}

// NewSmsRepo .
func NewSmsRepo(cc config.Config, bs *conf.Bootstrap, data *Data, logger log.Logger) biz.SmsRepo {
	if bs.Log.Filter == nil {
		log.Fatal("Configuration `bs.log.filter` has not set.")
	}

	logger = log.With(logger, "caller", log.Caller(6))
	repo := &smsRepo{
		conf: cc,
		bs:   bs,
		data: data,
		log: log.NewHelper(log.NewFilter(logger,
			mylog.KeyFilter,
			mylog.FilterLevel(bs.Log.Filter.DataLogLevel),
		)),
	}
	// 根据配置变更动态修改过滤器级别
	mylog.FilterChangeWatch(cc, repo, mylog.DataLogLvlConfKey, logger)

	return repo
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
