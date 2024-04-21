hateblo2hugo
=======

docker build . -t goapp

ダウンロード込みで実行

```bash
docker run -v ${PWD}:/app --name golang-app --rm -t goapp migrate -i 7butk7ed2atx.txt -u
```

記事を移動。

```bash
cp  ~/git/hateblo2hugo/content/post/entry/* ~/git/hugobali/content/english/blog -r
```



```bash

rsync -av  ~/git/hateblo2hugo/content/post/entry/* ~/git/hugobali/content/english/blog

```


```bash

rsync -ahvnr  ~/git/hateblo2hugo/content/post/entry/* ~/git/hugobali/content/english/blog --existing

```


* -a を指定することでなるべくコピー元のファイルと同一条件でコピーする。
* -h でファイルサイズの視認性をよくする。
* -v で詳細を出力する。
* -n または --dry-run で対象ファイルを確認する。
* --existing 更新分のみ(追加は無視される)

netlify cliの場合

```bash

sudo hugo -b "https://hugobali.netlify.app"

netlify deploy --dir=public --site=5361fff5-bcef-4fbc-9e2f-0902b48ddb03 --prod
```


hateblo2hugo is a tool to migrate blog data of [hatenablog](http://hatenablog.com/) to markdown data for Hugo.

### Install

```bash
$ go get github.com/yamadatt/hateblo2hugo
```

### Preparation

Before use this tool, you must export blog data from hatenablog. Data format of hatenablog is Movable Type.


### Usage

```bash
$ hateblo2hugo migrate -i ~/your_path/your_hatenablog.export.txt -o ~/your_path/your_hugo_blog/blog/ -u
```

### Migration Features

* Remove Hatena Keyword link
* Download images from hatena photo life, and locate files to `{blog_dir}/static/images` directory.
* Embed contents
    * Tweet
    * Speakerdeck
    * General links
    * GitHub Repository (use [lepture/github-cards](https://github.com/lepture/github-cards) )
* Code syntax


