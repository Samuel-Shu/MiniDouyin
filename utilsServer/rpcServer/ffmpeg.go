package rpcServer

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	pb "utilsServer/rpcPb"
)

type UtilsServerServer struct {}

func (u *UtilsServerServer)ParseCover(ctx context.Context,ffmpegReq *pb.FfmpegReq) (ffmpegRes *pb.FfmpegRes,err error) {
	// Returns specified frame as []byte
	fmt.Println("\n ParseCover start!")
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(ffmpegReq.VideoUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", ffmpegReq.FrameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Panicln(err)
		//return nil, err
	}
	coverByte := buf.Bytes()
	ffmpegRes=&pb.FfmpegRes{CoverByte: coverByte,Error: ""}
	return
}

func (u *UtilsServerServer)Md5(ctx context.Context,md5Req *pb.Md5Req )(md5Res *pb.Md5Res,err error)  {
	fmt.Println("\n md5 server start!")
	h := md5.New()
	_, err = io.WriteString(h, md5Req.Password)
	md5Password := string(h.Sum([]byte(nil)))
	md5Password = fmt.Sprintf("%x",md5Password)
	md5Res = &pb.Md5Res{Md5Password: md5Password}
	return
}

