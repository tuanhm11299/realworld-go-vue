<script setup lang="ts">
// TODO(M7/M8): the home page carries the most state of any view — which feed
// tab is active, which tag, which page. Decide early where that state lives:
// the URL, the component, or the store. The e2e suite navigates by URL, so
// think hard about what has to survive a refresh.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Home"
    route="/"
    milestone="M7 + M8"
    :specs="['e2e/navigation.spec.ts', 'e2e/articles.spec.ts', 'e2e/health.spec.ts']"
    :requirements="[
      'Show the .banner hero (logged-out only)',
      'Offer the .feed-toggle tabs: Your Feed (auth only), Global Feed, and a tag tab once a tag is picked',
      'Render .article-preview cards from GET /api/articles or /api/articles/feed',
      'Show the .sidebar tag list from GET /api/tags and filter by the clicked tag',
      'Paginate with .pagination / .page-item / .page-link (limit 10 per page)',
      'Show .empty-feed-message when a feed has no articles',
      'Favourite / unfavourite from the card without a full reload',
    ]"
  />
</template>
