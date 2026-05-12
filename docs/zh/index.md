---
layout: home
---

<script setup>
import { onBeforeMount } from 'vue'

onBeforeMount(() => {
  if (typeof window !== 'undefined') {
    localStorage.setItem('docs-lang-preference', `${import.meta.env.BASE_URL}zh/`)
  }
})
</script>

<div class="home-header">
  <div class="home-header-left">
    <div class="home-logo">CK</div>
    <div>
      <span class="home-title">CompressKit</span>
      <span class="home-subtitle">无损压缩算法库</span>
    </div>
  </div>
  <div class="home-nav">
    <a href="./guide/getting-started">快速开始</a>
    <a href="https://github.com/LessUp/compress-kit">GitHub</a>
    <a href="../en/">English</a>
  </div>
</div>

<div class="home-intro-row">
  <div class="home-intro">
    CompressKit 提供经典的无损压缩算法，支持跨语言兼容。C++ 编码，Go 解码。Rust 编码，C++ 解码。所有实现产生完全相同的二进制输出。
  </div>
  <div class="home-stats">
    <span><strong>C++17</strong></span>
    <span><strong>Go</strong></span>
    <span><strong>Rust</strong></span>
  </div>
</div>

## 算法

<div class="feature-map">
  <div class="feature-card">
    <div class="feature-card-title">🌳 霍夫曼编码</div>
    <div class="feature-card-desc">
      基于符号频率的最优前缀码。经典的无损压缩方法。
    </div>
    <div class="feature-tags">
      <a href="./algorithms/huffman" class="feature-tag">了解更多</a>
      <span class="feature-tag">速度快</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🧮 算术编码</div>
    <div class="feature-card-desc">
      整个消息编码为单个数字。达到熵极限，实现最大压缩率。
    </div>
    <div class="feature-tags">
      <a href="./algorithms/arithmetic" class="feature-tag">了解更多</a>
      <span class="feature-tag">高压缩率</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🎯 区间编码</div>
    <div class="feature-card-desc">
      基于整数的算术编码。生产级的速度与压缩率平衡。
    </div>
    <div class="feature-tags">
      <a href="./algorithms/range" class="feature-tag">了解更多</a>
      <span class="feature-tag">快 + 高压缩</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">📏 行程编码</div>
    <div class="feature-card-desc">
      针对重复数据的简单快速压缩。常作为预处理步骤使用。
    </div>
    <div class="feature-tags">
      <a href="./algorithms/rle" class="feature-tag">了解更多</a>
      <span class="feature-tag">极快</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🔄 跨语言兼容</div>
    <div class="feature-card-desc">
      一种语言编码，另一种语言解码。所有实现产生完全相同的二进制输出。
    </div>
    <div class="feature-tags">
      <a href="./guide/getting-started" class="feature-tag">快速开始</a>
      <a href="./testing/cross-language" class="feature-tag">测试</a>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">📊 性能基准</div>
    <div class="feature-card-desc">
      跨所有算法和语言的性能基准测试。比较速度和压缩率。
    </div>
    <div class="feature-tags">
      <a href="./benchmarks/results" class="feature-tag">查看结果</a>
      <a href="./benchmarks/how-to-run" class="feature-tag">运行测试</a>
    </div>
  </div>
</div>

<div class="quick-start">
  <div class="quick-start-title">快速开始</div>
  <div class="quick-start-content">
    <div class="command-block">
      <code>git clone https://github.com/LessUp/compress-kit.git && cd compress-kit && make build && make test</code>
    </div>
  </div>
</div>
