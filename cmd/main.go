package main

import (
	"github.com/DanYesmagulov/go-video-streaming/pkg/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	streaming "github.com/DanYesmagulov/go-video-streaming"
	"github.com/DanYesmagulov/go-video-streaming/pkg/handler"
	"github.com/DanYesmagulov/go-video-streaming/pkg/repository"
	"github.com/DanYesmagulov/go-video-streaming/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error during configs init: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error during env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error occured during db init: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error during config init: %s", err.Error())
	}

	//=======================================================

	// Init client storage connection
	storageProvider, err := newStorageProvider(&storage.Configuration{
		Host:            viper.GetString("minio.host"),
		AccessKey:       viper.GetString("minio.access_key"),
		SecretAccessKey: viper.GetString("minio.secret_key"),
		Bucket:          viper.GetString("minio.bucket"),
	})
	if err != nil {
		logrus.Fatalf("Error during storage init: %s", err.Error())
	}

	//========================================================

	repos := repository.NewRepository(db)
	services := service.NewService(service.Deps{
		Repos:           repos,
		StorageProvider: storageProvider,
	})
	handlers := handler.NewHandler(services)

	/*client := storage.NewFileStorage(minioClient, viper.GetString("minio.bucket"), viper.GetString("minio.host"),)
	bucketName := "video-file"
	objectName := "profile.jpeg"
	filepath := "./assets/images/profile.jpeg"
	err = client.UploadFile(bucketName, objectName, filepath)
	if err != nil {
		logrus.Fatalf("Error during file upload: %s", err.Error())
	}*/

	srv := new(streaming.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured during http server run: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func newStorageProvider(config *storage.Configuration) (storage.Provider, error) {
	client, err := minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	provider := storage.NewFileStorage(client, config.Bucket, config.Host)

	return provider, nil
}
