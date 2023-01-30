package logic

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"lol/internal/svc"
	"lol/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	request *http.Request
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadLogic {
	return &UploadLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		request: r,
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResponse, err error) {
	// go-zero从jwt token解析后会将用户生成token时传入的kv原封不动的放在http.Request的Context中，因此我们可以通过Context就可以拿到你想要的值
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	file, fileHeader, err := l.request.FormFile(l.svcCtx.Config.Upload.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tmpFile, err := os.Create(path.Join(l.svcCtx.Config.Upload.Dir, fileHeader.Filename))
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()
	io.Copy(tmpFile, file)

	return &types.UploadResponse{
		FilePath: path.Join(l.svcCtx.Config.Upload.Dir, fileHeader.Filename),
		Code:     0,
	}, nil
}
