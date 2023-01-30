package logic

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"lol/internal/svc"
	"lol/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: w,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadRequest) error {
	// go-zero从jwt token解析后会将用户生成token时传入的kv原封不动的放在http.Request的Context中，因此我们可以通过Context就可以拿到你想要的值
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	fileByte, err := ioutil.ReadFile(path.Join(l.svcCtx.Config.Upload.Dir, req.File))
	if err != nil {
		return err
	}
	n, err := l.writer.Write(fileByte)
	if err != nil {
		return err
	}
	if n < len(fileByte) {
		return io.ErrClosedPipe
	}
	return nil
}
