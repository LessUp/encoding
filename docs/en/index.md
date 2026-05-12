---
layout: home
---

<script setup>
import { onBeforeMount } from 'vue'

onBeforeMount(() => {
  if (typeof window !== 'undefined') {
    localStorage.setItem('docs-lang-preference', `${import.meta.env.BASE_URL}en/`)
  }
})
</script>

<div class="home-header">
  <div class="home-header-left">
    <div class="home-logo">CK</div>
    <div>
      <span class="home-title">CompressKit</span>
      <span class="home-subtitle">Lossless Compression Library</span>
    </div>
  </div>
  <div class="home-nav">
    <a href="./guide/getting-started">Get Started</a>
    <a href="https://github.com/LessUp/compress-kit">GitHub</a>
    <a href="../zh/">中文</a>
  </div>
</div>

<div class="home-intro-row">
  <div class="home-intro">
    CompressKit provides classic lossless compression algorithms with cross-language compatibility. Encode in C++, decode in Go. Encode in Rust, decode in C++. All implementations produce identical binary output.
  </div>
  <div class="home-stats">
    <span><strong>C++17</strong></span>
    <span><strong>Go</strong></span>
    <span><strong>Rust</strong></span>
  </div>
</div>

## Algorithms

<div class="feature-map">
  <div class="feature-card">
    <div class="feature-card-title">🌳 Huffman Coding</div>
    <div class="feature-card-desc">
      Optimal prefix codes based on symbol frequency. The classic approach to lossless compression.
    </div>
    <div class="feature-tags">
      <a href="./algorithms/huffman" class="feature-tag">Learn More</a>
      <span class="feature-tag">Fast Speed</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🧮 Arithmetic Coding</div>
    <div class="feature-card-desc">
      Entire message encoded as a single number. Achieves entropy limit for maximum compression.
    </div>
    <div class="feature-tags">
      <a href="./algorithms/arithmetic" class="feature-tag">Learn More</a>
      <span class="feature-tag">High Compression</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🎯 Range Coder</div>
    <div class="feature-card-desc">
      Integer-based arithmetic coding. Production-ready balance of speed and compression.
    </div>
    <div class="feature-tags">
      <a href="./algorithms/range" class="feature-tag">Learn More</a>
      <span class="feature-tag">Fast + High</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">📏 Run-Length Encoding</div>
    <div class="feature-card-desc">
      Simple and fast compression for repetitive data. Often used as preprocessing.
    </div>
    <div class="feature-tags">
      <a href="./algorithms/rle" class="feature-tag">Learn More</a>
      <span class="feature-tag">Very Fast</span>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">🔄 Cross-Language</div>
    <div class="feature-card-desc">
      Encode in one language, decode in another. All implementations produce identical binary output.
    </div>
    <div class="feature-tags">
      <a href="./guide/getting-started" class="feature-tag">Get Started</a>
      <a href="./testing/cross-language" class="feature-tag">Testing</a>
    </div>
  </div>

  <div class="feature-card">
    <div class="feature-card-title">📊 Benchmarks</div>
    <div class="feature-card-desc">
      Performance benchmarks across all algorithms and languages. Compare speed and compression.
    </div>
    <div class="feature-tags">
      <a href="./benchmarks/results" class="feature-tag">View Results</a>
      <a href="./benchmarks/how-to-run" class="feature-tag">Run Tests</a>
    </div>
  </div>
</div>

<div class="quick-start">
  <div class="quick-start-title">Quick Start</div>
  <div class="quick-start-content">
    <div class="command-block">
      <code>git clone https://github.com/LessUp/compress-kit.git && cd compress-kit && make build && make test</code>
    </div>
  </div>
</div>
