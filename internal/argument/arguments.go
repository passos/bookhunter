package argument

import (
	"github.com/bookstairs/bookhunter/internal/client"
	"github.com/bookstairs/bookhunter/internal/fetcher"
)

var (
	// Flags for talebook registering.

	Username = ""
	Password = ""
	Email    = ""

	// Common flags.

	Website    = ""
	UserAgent  = client.DefaultUserAgent
	Proxy      = ""
	ConfigRoot = ""

	// Common download flags.

	Formats       = fetcher.NormalizeFormats(fetcher.EPUB, fetcher.AZW3, fetcher.MOBI, fetcher.PDF, fetcher.ZIP)
	Extract       = false
	DownloadPath  = ""
	InitialBookID = 1
	Rename        = false
	Thread        = 1

	// Drive ISP configurations.

	RefreshToken = ""

	// Telegram configurations.

	ChannelID = ""
	Mobile    = ""
	ReLogin   = false
	AppID     = int64(0)
	AppHash   = ""
)

// NewFetcher will create the fetcher by the command line arguments.
func NewFetcher(category fetcher.Category, properties map[string]string) (fetcher.Fetcher, error) {
	cc, err := client.NewConfig(Website, UserAgent, Proxy, ConfigRoot)
	if err != nil {
		return nil, err
	}

	fs, err := fetcher.ParseFormats(Formats)
	if err != nil {
		return nil, err
	}

	return fetcher.New(&fetcher.Config{
		Config:        cc,
		Category:      category,
		Formats:       fs,
		Extract:       Extract,
		DownloadPath:  DownloadPath,
		InitialBookID: InitialBookID,
		Rename:        Rename,
		Thread:        Thread,
		Properties:    properties,
	})
}
