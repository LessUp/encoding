---
layout: home

hero:
  name: CompressKit
  text: ' '
  actions:
    - theme: brand
      text: English
      link: /en/
    - theme: alt
      text: 中文
      link: /zh/
---

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vitepress'

onMounted(() => {
  const router = useRouter()
  const lang = navigator.language || navigator.userLanguage
  if (lang.startsWith('zh')) {
    router.go('/zh/')
  } else {
    router.go('/en/')
  }
})
</script>
