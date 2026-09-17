<script setup lang="ts">
// TODO(M8): remember that the server may change the slug when the title
// changes, so navigate to the slug in the RESPONSE, never the one you had.
// Also: what should happen if someone opens /editor/:slug for an article they
// do not own?
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Editor"
    route="/editor and /editor/:slug"
    milestone="M8"
    :specs="['e2e/articles.spec.ts', 'e2e/url-navigation.spec.ts', 'e2e/error-handling.spec.ts']"
    :requirements="[
      'Require auth',
      'Fields: input[name=title], input[name=description], textarea[name=body], input[placeholder=&quot;Enter tags&quot;]',
      'Enter adds a tag pill; the pill has an ion-close-round button that removes it',
      'With no :slug, POST /api/articles; with a :slug, GET then PUT it',
      'Redirect to /article/:slug on success, using the slug the server returned',
      'Show validation errors in ul.error-messages',
    ]"
  />
</template>
