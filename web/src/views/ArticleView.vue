<script setup lang="ts">
// TODO(M8/M9): rendering markdown means rendering user input as HTML.
// xss-security.spec.ts will try to inject a script through the article body,
// the title, and the bio. Sanitise (DOMPurify) or use a renderer that escapes
// by default — and never reach for v-html without thinking first.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Article"
    route="/article/:slug"
    milestone="M8 + M9"
    :specs="['e2e/articles.spec.ts', 'e2e/comments.spec.ts', 'e2e/xss-security.spec.ts']"
    :requirements="[
      'Render .article-page with .banner and .article-content',
      'Render the markdown body client-side',
      'Show .article-meta twice (banner and footer) with follow and favourite buttons',
      'Show Edit / Delete buttons to the author only',
      'List comments in .card elements; show .mod-options delete only on your own',
      'Post comments from textarea[placeholder=&quot;Write a comment...&quot;]',
      'Prompt anonymous visitors to sign in instead of showing the comment form',
    ]"
  />
</template>
