// Code generated by goctl. DO NOT EDIT!
// Source: fileud.proto

package fileudclient

import (
	"context"

	"gcloud/fileud/rpc/fileud"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CosObject                      = fileud.CosObject
	FileDownloadReply              = fileud.FileDownloadReply
	FileDownloadRequest            = fileud.FileDownloadRequest
	FileUploadChunkCompleteReply   = fileud.FileUploadChunkCompleteReply
	FileUploadChunkCompleteRequest = fileud.FileUploadChunkCompleteRequest
	FileUploadChunkReply           = fileud.FileUploadChunkReply
	FileUploadChunkRequest         = fileud.FileUploadChunkRequest
	FileUploadPrepareReply         = fileud.FileUploadPrepareReply
	FileUploadPrepareRequest       = fileud.FileUploadPrepareRequest
	FileUploadReply                = fileud.FileUploadReply
	FileUploadRequest              = fileud.FileUploadRequest

	Fileud interface {
		FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadReply, error)
		FileDownload(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (*FileDownloadReply, error)
		FileUploadChunk(ctx context.Context, in *FileUploadChunkRequest, opts ...grpc.CallOption) (*FileUploadChunkReply, error)
		FileUploadChunkComplete(ctx context.Context, in *FileUploadChunkCompleteRequest, opts ...grpc.CallOption) (*FileUploadChunkCompleteReply, error)
		FileUploadPrepare(ctx context.Context, in *FileUploadPrepareRequest, opts ...grpc.CallOption) (*FileUploadPrepareReply, error)
	}

	defaultFileud struct {
		cli zrpc.Client
	}
)

func NewFileud(cli zrpc.Client) Fileud {
	return &defaultFileud{
		cli: cli,
	}
}

func (m *defaultFileud) FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadReply, error) {
	client := fileud.NewFileudClient(m.cli.Conn())
	return client.FileUpload(ctx, in, opts...)
}

func (m *defaultFileud) FileDownload(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (*FileDownloadReply, error) {
	client := fileud.NewFileudClient(m.cli.Conn())
	return client.FileDownload(ctx, in, opts...)
}

func (m *defaultFileud) FileUploadChunk(ctx context.Context, in *FileUploadChunkRequest, opts ...grpc.CallOption) (*FileUploadChunkReply, error) {
	client := fileud.NewFileudClient(m.cli.Conn())
	return client.FileUploadChunk(ctx, in, opts...)
}

func (m *defaultFileud) FileUploadChunkComplete(ctx context.Context, in *FileUploadChunkCompleteRequest, opts ...grpc.CallOption) (*FileUploadChunkCompleteReply, error) {
	client := fileud.NewFileudClient(m.cli.Conn())
	return client.FileUploadChunkComplete(ctx, in, opts...)
}

func (m *defaultFileud) FileUploadPrepare(ctx context.Context, in *FileUploadPrepareRequest, opts ...grpc.CallOption) (*FileUploadPrepareReply, error) {
	client := fileud.NewFileudClient(m.cli.Conn())
	return client.FileUploadPrepare(ctx, in, opts...)
}
