package app

import (
	"context"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ClipContent struct {
	Id string
	Time string
	Content string
}

func NewClipContent(content string) *ClipContent {
	return &ClipContent{
		Id: cryptor.Md5String(content),
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Content: content,
	}
}


type Clipboard struct {
	ctx context.Context
	cfg *Config
	storeDir string
	contents map[string]*ClipContent
}

func NewClipboard(storeDir string, cfg *Config) *Clipboard {
	return &Clipboard{
		storeDir: storeDir,
		cfg: cfg,
	}
}

func (c *Clipboard) StartUp(ctx context.Context) {
	c.ctx = ctx
	c.read()
	c.listen(time.Millisecond * 500)
}

// 初始化读取记录
func (c *Clipboard) read() {
	contents := map[string]*ClipContent{}
	c.contents = contents
	f := c.storeDir + "history.yaml"
	vip := viper.New()
	vip.SetConfigFile(f)
	vip.SetConfigType("yaml")
	if err := vip.ReadInConfig(); err != nil {
		runtime.LogPrintf(c.ctx, "history文件读取错误，%s", err)
		return 
	}
	

	err := vip.Unmarshal(&contents)
	if err != nil {
		runtime.LogErrorf(c.ctx, "history文件解析错误，%s", err)
		return
	}
	c.contents = contents
}

// 监听剪切板
func (c *Clipboard) listen(interval time.Duration) {
	go func() {
		for {
			ct, err := runtime.ClipboardGetText(c.ctx)
			if err != nil {
				time.Sleep(interval)
				continue
			}
			content := NewClipContent(ct)
			if _, yes := c.contents[content.Id]; !yes {
				c.AddContent(content)
			}
		}
	}()
}

// 添加新内容
func (c *Clipboard) AddContent(content *ClipContent) {
	runtime.LogPrintf(c.ctx, "添加内容：%s", content.Content)
	c.contents[content.Id] = content
	runtime.EventsEmit(c.ctx, "NewClip", c.contents[content.Id])
}

// 获取所有内容
func (c *Clipboard) GetContents() map[string]*ClipContent {
	return c.contents
}

