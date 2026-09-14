<script setup lang="ts">
// TODO(M7): this is where the persistent chrome lives.
//
// The e2e suite expects, on every page:
//   <nav class="navbar">  with  <a class="navbar-brand">conduit</a>
//   <a class="nav-link">  items, switching between the logged-out set
//   (Home / Sign in / Sign up) and the logged-in set
//   (Home / New Article / Settings / <username>)
// plus the footer.
//
// Exact contract: .specs/e2e/SELECTORS.md
// Markup to copy:  .specs/realworld/docs/.../frontend/templates.md
//
// Until you build it, health.spec.ts fails on the very first assertion — which
// is the right place to start.
</script>

<template>
  <RouterView />
</template>
