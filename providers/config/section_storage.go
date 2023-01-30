package config

type Storage struct {
	Driver    string
	AliYunOSS AliYunOSS
	AwsS3     AwsS3
}

type AliYunOSS struct {
	Bucket          string
	Region          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BaseURL         string
	Path            string
}

type AwsS3 struct {
	Bucket           string
	Region           string
	Endpoint         string
	DisableSSL       bool
	SecretID         string
	SecretKey        string
	BaseURL          string
	Path             string
	S3ForcePathStyle bool
}
