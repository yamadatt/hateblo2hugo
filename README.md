hateblo2hugo
=======

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


