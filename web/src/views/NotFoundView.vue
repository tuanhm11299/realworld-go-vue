<script setup lang="ts">
// Not required by the RealWorld spec, but url-navigation.spec.ts visits unknown
// URLs and expects the app to stay alive rather than render a blank page.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Not found"
    route="anything unmatched"
    milestone="M8"
    :specs="['e2e/url-navigation.spec.ts']"
    :requirements="[
      'Keep the navbar and footer visible',
      'Offer a way back to the home page',
      'Do not leave the previous page’s content on screen',
    ]"
  />
</template>
