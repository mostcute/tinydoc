# Tinydoc 简介

golang开发的 静态文件管理器 + [Docute](https://github.com/egoist/docute) 组成的文档网站。



看了多种文档系统，如gitbook，HEXO,Vue/VitePress等



相比于在本地和项目一个目录中维护一大推散乱的文档，文档网站的存在使项目开发者更有动力去编写文档，有效的组织文件使项目介绍更加条理清晰。可以帮助项目被感兴趣的人快速了解。



但是额外搭建一套文档系统如gitbook，wiki，又极大的占据的维护者的精力，部分文档系统还有用户系统，团队协作等功能，更接近于一套复杂的研发管理工具。



而作为一个项目维护者，维护文档的初心就是能有效组织文档关联性，并且优雅的展现在其他人面前。



## Tinydoc 做了什么？

tinydoc就是简单的将 [Docute](https://github.com/egoist/docute) 生成的文件托管到http静态文件服务中，提供文档展示

如果你喜欢docsify，可以将docsify替换掉Docute

如果你是使用git仓库来维护文档,tinydoc将会支持从git仓库来自动更新你的文档

如果你需要一点点权限保护，仅允许部分人阅读，tinydoc将提供basicauth来保护你的文档

想要不通过nginx直接支持https? tinydoc也将支持



以上就是tinydoc的全部功能，相信已经可以满足大部分个人项目的文档展示需求，接下来就动力满满的去把时间花在项目本身上吧！



## 帮助信息

[Docute 官方文档]([Docute (egoist.dev)](https://docute.egoist.dev/zh/))