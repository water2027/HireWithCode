<script lang="ts" setup>
import DOMPurify from 'dompurify'
import MarkdownIt from 'markdown-it'
import { computed } from 'vue'
import 'github-markdown-css/github-markdown-light.css'

const { content } = defineProps<{
  content: string
}>()

const md = new MarkdownIt({
  html: true,
  breaks: true,
  linkify: true,
  typographer: true,
})

const markdown = computed(() => {
  const rendered = md.render(content)
  const finalHtml = DOMPurify.sanitize(rendered, {
    FORBID_TAGS: ['style', 'script', 'iframe'],
    FORBID_ATTR: ['style', 'onerror', 'onclick', 'onload'],
    ALLOWED_TAGS: ['p', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'ul', 'ol', 'li', 'a', 'strong', 'em', 'blockquote', 'code', 'pre'],
    ALLOWED_ATTR: ['href', 'target', 'rel', 'title'],
  })
  return finalHtml
})
</script>

<template>
  <div class="markdown-body" v-html="markdown" />
</template>
