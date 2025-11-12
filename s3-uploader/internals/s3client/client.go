package s3client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Client struct {
	S3 *s3.Client
	Bucket string
	Region string
}


type ObjectInfo struct {
	Key string `json:"key"`
	Size int64 `json:"size"`
	LastModified time.Time `json:"lastModified"`
 	URL string `json:"url"`
}




func New(ctx context.Context, accessKey, secretKey, region, bucket string )(*Client, error) {

	// validering för våra parametrar

	if accessKey == "" || secretKey == "" || region == "" || bucket == "" {
		return nil, errors.New("saknar credentials ")
	}



	// bygga aws config med region + statiska credentials från .env filen


	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey,secretKey, ""),
		),
	)



	if err != nil {
		return nil, fmt.Errorf("kunde inte ladda aws konfigurationerna: %v", err)
	}



	client := s3.NewFromConfig(cfg)


	return &Client{
		S3: client,
		Bucket: bucket,
		Region: region,
	}, nil

}


func (c *Client) EnsureBucket(ctx context.Context) error {
	_, err := c.S3.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(c.Bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(c.Region),
		},
	})
	if err != nil {
		return fmt.Errorf("kunde inte ansluta till AWS S3 bucket")
	}

	return nil
}



func (c *Client) PutString(ctx context.Context, key, content string) error {
	if key == "" || content == "" {
		return errors.New("saknar key eller content")
	}

	_, err := c.S3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.Bucket),
		Key: aws.String(key),
		Body: strings.NewReader(content),
		ContentType: aws.String("text/plain; charset=utf-8"),
		ContentDisposition: aws.String("inline"),
	})

	if err != nil {
		return fmt.Errorf("kunde inte ladda upp objektet: %v", err)
	}

	return nil

}



func (c *Client) DeleteObject(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("saknar key")
	}

	_, err := c.S3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.Bucket),
		Key: aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("kunde itne radera objektet: %v", err)
	}
	return nil
}




func (c *Client) ListObjects(ctx context.Context, prefix string, urlTTL time.Duration) ([]ObjectInfo, error) {
    
	if c.Bucket == "" {
        return nil, errors.New("bucket saknas i S3-klienten")
    }

    presigner := s3.NewPresignClient(c.S3)

    var out []ObjectInfo
    var token *string

    for {
        input := &s3.ListObjectsV2Input{
            Bucket:            aws.String(c.Bucket),
            Prefix:            nil,
            ContinuationToken: token,
            MaxKeys:           aws.Int32(1000),
        }
        if prefix != "" {
            input.Prefix = aws.String(prefix)
        }

        resp, err := c.S3.ListObjectsV2(ctx, input)
        if err != nil {
            return nil, fmt.Errorf("lista objekt: %w", err)
        }

        for _, it := range resp.Contents {
            getOut, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
                Bucket: aws.String(c.Bucket),
                Key:    it.Key,
            }, s3.WithPresignExpires(urlTTL))
            if err != nil {
                return nil, fmt.Errorf("presign get för %s: %w", aws.ToString(it.Key), err)
            }

            out = append(out, ObjectInfo{
                Key:          aws.ToString(it.Key),
                Size:         aws.ToInt64(it.Size),
                LastModified: aws.ToTime(it.LastModified),
                URL:          getOut.URL,
            })
        }

        if *resp.IsTruncated {
            token = resp.NextContinuationToken
        } else {
            break
        }
    }

    return out, nil
}


