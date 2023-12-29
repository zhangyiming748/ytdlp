package ytdlp

import (
	"fmt"
	"github.com/zhangyiming748/ytdlp/replace"
	"github.com/zhangyiming748/ytdlp/sql"
	"github.com/zhangyiming748/ytdlp/util"
	"log/slog"
	"os/exec"
	"strconv"
	"sync"
)

func Ytdlp(fp string) {
	lines := util.ReadByLine(fp)
	var wg sync.WaitGroup
	concurrency := util.GetVal("yt-dlp", "concurrency")
	concurrencyNum, err := strconv.Atoi(concurrency)
	if err != nil {
		panic("yt-dlp下载并发数设置错误")
	}
	if concurrencyNum < 1 {
		panic("下载并发数需要大于1")
	}
	ch := make(chan struct{}, concurrencyNum)
	for _, line := range lines {
		ch <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			ytdlp(line)
			<-ch
		}()
		slog.Debug("协程任务运行完成")
	}
	wg.Wait()
}
func ytdlp(line string) {
	slog.Debug("协程任务运行")
	prefix := util.GetVal("yt-dlp", "saveTo")
	proxy := util.GetVal("yt-dlp", "proxy")
	cmd := exec.Command("yt-dlp", "--proxy", proxy, "--no-playlist", "-P", prefix, "-o", "%(title)s.%(ext)s", line)
	slog.Info("yt-dlp命令开始执行", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("文件名", "%(title)s.%(ext)s"))
	y := new(sql.Ytdlp)
	y.URL = line
	y.Status = "失败"
	if name, err := getName(line); err != nil {
		slog.Warn("获取文件名失败")
	} else {
		y.Name = name
	}
	y.SetOne()

	err := util.ExecCommand(cmd)
	if err != nil {
		slog.Warn("yt-dlp命令下载失败", slog.String("命令原文", fmt.Sprint(cmd)), slog.Any("错误原文", err))
		return
	}
	y.Status = "下载成功"
	y.UpdateStatusById()
	slog.Info("yt-dlp命令执行完成")
}

func getName(link string) (string, error) {
	// cmd := exec.Command("yt-dlp", "--print", "filename", "-o", "%(title)s.%(ext)s", link)
	cmd := exec.Command("yt-dlp", "--print", "filename", "-o", "%(title)s", link)
	fname, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return replace.ForFileName(string(fname)), err
}
