package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"kratos-sms/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSmsRepo)

// Data saved to database
type Data struct {
	client *mongo.Client
	log    *log.Helper
}

// NewData .
func NewData(bs *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	c := bs.GetData()
	l := log.NewHelper(logger)
	if c.GetDatabase().GetDriver() != "mongo" {
		l.Fatal("Only support mongodb driver.")
	}
	uri := c.GetDatabase().GetSource()
	mongof := c.GetDatabase().GetMongo()
	if mongof == nil {
		l.Fatal("Mongo configurations can't be empty.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), mongof.ConnectTimeout.AsDuration())
	defer cancel()

	option := options.Client()
	option.ApplyURI(uri)
	option.SetMinPoolSize(uint64(mongof.MinPoolSize))
	option.SetMaxPoolSize(uint64(mongof.MaxPoolSize))
	option.SetHeartbeatInterval(mongof.HeartbeatInterval.AsDuration())
	if mongof.ReadConcern != "" {
		option.ReadConcern = readconcern.New(readconcern.Level(mongof.ReadConcern))
	}
	if mongof.ReadPreferMode != conf.Mongo_NONE {
		readPref, err := readpref.New(readpref.Mode(mongof.ReadPreferMode))
		if err == nil {
			option.ReadPreference = readPref
		}
	}

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		l.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Connected to %s", uri)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		_ = client.Disconnect(ctx)
		defer cancel()
		l.Info("Closing the data resources")
	}

	return &Data{client: client, log: l}, cleanup, nil
}

func (d *Data) Collection(db, coll string) *mongo.Collection {
	return d.client.Database(db).Collection(coll)
}

func (d *Data) Db(db string) *mongo.Database {
	return d.client.Database(db)
}
