package pkg

import (
	//"api_gateway/basic/appconfig"
	//"context"
	//"github.com/minio/minio-go/v7"
	//"github.com/minio/minio-go/v7/pkg/credentials"
	//"log"
	"mime/multipart"
)

func MiNio(name string, open multipart.File) bool {
	//// MinIO 服务器配置
	//AppConf := appconfig.AppCong.Minio
	//endpoint := AppConf.Endpoint
	//accessKeyID := AppConf.Accesskeyid
	//secretAccessKey := AppConf.Accesskeysecret
	//
	//// 文件配置
	//fileName := name                 // 要上传的文件名
	//bucketName := AppConf.Bucketname // 目标桶名
	//objectName := fileName           // 上传后对象的名称（可包含路径前缀）
	//contentType := "text/plain"      // 文件的 MIME 类型
	//
	//// 初始化 MinIO 客户端
	//minioClient, err := minio.New(endpoint, &minio.Options{
	//	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	//	Secure: false,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// 创建上下文
	//ctx := context.Background()
	//
	//// 创建桶（如果不存在）
	//location := "test"
	//err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	//if err != nil {
	//	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	//	if errBucketExists != nil {
	//		log.Fatalln(errBucketExists)
	//	}
	//	if exists {
	//		log.Printf("Bucket %s already exists", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created bucket %s", bucketName)
	//}
	//
	//// 上传文件
	////filePath := filepath.Join("./tmp/", fileName) // 文件路径
	////fmt.Println(filePath)
	//_, err = minioClient.PutObject(ctx, bucketName, objectName, open, -1, minio.PutObjectOptions{ContentType: contentType})
	//log.Println(objectName)
	//if err != nil {
	//	log.Printf("Failed to upload %s: %v", objectName, err)
	//	return false
	//}
	//
	//log.Printf("Successfully uploaded %s%s", AppConf.Bucketurl, fileName)
	//return true

	return true
}
