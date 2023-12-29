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
	y := new(sql.Ytdlp)
	y.URL = line
	y.Status = "下载失败"
	if name, err := getName(line); err != nil {
		slog.Warn("获取文件名失败")
	} else {
		yt := new(sql.Ytdlp)
		yt.Name = name
		y.Name = name
		if s, _, _ := yt.FindDupByName(); len(s) > 0 {
			slog.Warn("有可能已经下载过了", slog.String("文件名", name), slog.String("URL", y.URL))
		}
	}
	y.SetOne()
	slog.Info("yt-dlp命令开始执行", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("文件名", y.Name))
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
	proxy := "127.0.0.1:7890"
	//proxy := util.GetVal("yt-dlp", "proxy")
	cmd := exec.Command("yt-dlp", "--proxy", proxy, "--print", "filename", "-o", "%(title)s", link)
	fname, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err1")
		return "", err
	} else {
		slog.Debug("从命令获取的文件名", slog.String("文件名", string(fname)))
	}
	return replace.ForFileName(string(fname)), err
}
