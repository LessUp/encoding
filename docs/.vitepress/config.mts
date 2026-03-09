import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'zh-CN',
  title: 'Encoding',
  description: '编码算法集合 — 用 C++、Go、Rust 多语言实现经典压缩编码算法',

  // GitHub Pages 部署：base 需要与仓库名一致
  base: '/encoding/',

  cleanUrls: true,

  sitemap: {
    hostname: 'https://lessup.github.io/encoding/',
  },

  head: [
    ['meta', { property: 'og:type', content: 'website' }],
    ['meta', { property: 'og:title', content: 'Encoding — 编码算法集合' }],
    ['meta', { property: 'og:description', content: '用 C++、Go、Rust 多语言实现经典压缩编码算法，学习与对比' }],
    ['meta', { property: 'og:url', content: 'https://lessup.github.io/encoding/' }],
    ['meta', { name: 'twitter:card', content: 'summary' }],
    ['meta', { name: 'twitter:title', content: 'Encoding — 编码算法集合' }],
    ['meta', { name: 'twitter:description', content: '用 C++、Go、Rust 多语言实现经典压缩编码算法，学习与对比' }],
    ['link', { rel: 'icon', type: 'image/svg+xml', href: '/encoding/logo.svg' }],
  ],

  markdown: {
    lineNumbers: true,
    languageAlias: {
      cuda: 'cpp',
    },
  },

  lastUpdated: true,

  themeConfig: {
    nav: [
      { text: '指南', link: '/guide/getting-started' },
      { text: '算法详解', link: '/guide/algorithms' },
      { text: '项目结构', link: '/guide/project-structure' },
      {
        text: '相关链接',
        items: [
          { text: 'GitHub', link: 'https://github.com/LessUp/encoding' },
          { text: 'CHANGELOG', link: 'https://github.com/LessUp/encoding/blob/master/CHANGELOG.md' },
        ],
      },
    ],

    sidebar: [
      {
        text: '入门',
        items: [
          { text: '快速开始', link: '/guide/getting-started' },
          { text: '项目结构', link: '/guide/project-structure' },
        ],
      },
      {
        text: '算法',
        items: [
          { text: '算法总览与对比', link: '/guide/algorithms' },
        ],
      },
    ],

    editLink: {
      pattern: 'https://github.com/LessUp/encoding/edit/master/docs/:path',
      text: '在 GitHub 上编辑此页',
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/LessUp/encoding' },
    ],

    footer: {
      message: '基于 MIT 许可发布',
      copyright: 'Copyright © 2025-2026 LessUp',
    },

    search: {
      provider: 'local',
    },

    outline: {
      level: [2, 3],
      label: '目录',
    },

    lastUpdated: {
      text: '最后更新',
    },

    docFooter: {
      prev: '上一页',
      next: '下一页',
    },

    returnToTopLabel: '返回顶部',
    sidebarMenuLabel: '菜单',
    darkModeSwitchLabel: '主题',
    externalLinkIcon: true,
  },
})
