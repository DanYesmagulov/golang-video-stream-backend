## Бэкенд для платформы на GO и postgresql, где пользователь регистрируется, создает курс с видео-файлами и может оценивать и комментировать чужие курсы. 
# С дальнейшей конвертацией, сжатием, генерацией превью и т.д. для видео-роликов с использованием ffmpeg

Используются библиотеки: 
 - gin для роутинга, валидации полей и респонса
 - sqlx для работы с бд, в данном случае postgresql
 - viper, godotenv для чтения окружения переменных
 - minio sdk для go, которое является "API compatible with Amazon S3", что делает миграцию на AWS S3 и digital ocean spaces очень простой
 
 Директории устроены в сответсвии чистой архитектуре Роберта Мартина
 - cmd/ Для файла main.go, как точки входа приложения
 - pkg/
    - repository - работа с бд
    - handler - контроллер
    - service - бизнес-логика приложения
    - store (domain, entity) - для бизнес-сущностей в приложении

![Сущности](entities.png)
