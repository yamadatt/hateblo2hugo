package service

import (
	"fmt"
	"path/filepath"
	"regexp"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"github.com/yamadatt/hateblo2hugo/helper"
	"github.com/yamadatt/hateblo2hugo/hugo"
	"github.com/yamadatt/movabletype"
)

type Migration interface {
	Execute() error
	OutputFilePath() string
}

type MigrationImpl struct {
	entry           *movabletype.Entry
	outputDirRoot   string
	outputEntryRoot string
}

func NewMigration(entry *movabletype.Entry, outputDirRoot string) Migration {

	return &MigrationImpl{
		entry:           entry,
		outputDirRoot:   outputDirRoot,
		outputEntryRoot: fmt.Sprintf("%s/content/post/entry", outputDirRoot),
	}
}

func (s *MigrationImpl) Execute() error {
	outpath := filepath.Join(s.outputEntryRoot, s.OutputFilePath())

	page := hugo.CreateHugoPage(s.entry)
	content, err := page.Render()
	if err != nil {
		return errors.Wrapf(err, "render hugo markdown is failed. [%s]", s.entry.Basename)
	}

	if err := helper.WriteFileWithDirectory(outpath, content, 0644); err != nil {
		return errors.Wrapf(err, "failed to write data file. path=%s", outpath)
	}

	return nil
}

func (s *MigrationImpl) OutputFilePath() string {
	// return fmt.Sprintf("%s.md", s.entry.Basename)

	du, _ := time.ParseDuration("-9h")
	d := s.entry.Date.Add(du)
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	dJST := d.In(jst)

	p := bluemonday.StrictPolicy()
	s.entry.Title = p.Sanitize(s.entry.Title)
	s.entry.Basename = p.Sanitize(s.entry.Basename)

	// fmt.Println("s.entry.Basename", s.entry.Basename)

	//＜と＞のコードがタイトルに入っているので削除する
	s.entry.Title = regexp.MustCompile(`&lt;.+/&gt;`).ReplaceAllString(s.entry.Title, "")
	s.entry.Basename = regexp.MustCompile(`&lt;.+/&gt;`).ReplaceAllString(s.entry.Basename, "")

	fmt.Println(dJST.Format("2006/01/02/150405"))
	if s.entry.Basename == "" {
		s.entry.Basename = s.entry.Title
	}

	fmt.Println("s.entry.Title", s.entry.Title)
	fmt.Println("s.entry.Basename", s.entry.Basename)

	s.entry.Basename = fmt.Sprintf("%s/%s", dJST.Format("2006/01/02/150405"), s.entry.Title)
	return fmt.Sprintf("%s/index.md", s.entry.Basename)

}
