<script setup lang="ts">
/**
 * Scaffold-only component: renders "here is what this route must become".
 *
 * It is NOT part of the RealWorld spec and uses no spec CSS classes, so it can
 * never accidentally satisfy an e2e selector. Delete each usage as you build the
 * real view; when the last one is gone, delete this file too.
 */
defineProps<{
  title: string
  route: string
  milestone: string
  specs: string[]
  requirements: string[]
}>()

const routes: { to: string; label: string }[] = [
  { to: '/', label: 'Home' },
  { to: '/login', label: 'Sign in' },
  { to: '/register', label: 'Sign up' },
  { to: '/settings', label: 'Settings' },
  { to: '/editor', label: 'New article' },
  { to: '/article/slug-here', label: 'Article' },
  { to: '/profile/username-here', label: 'Profile' },
]
</script>

<template>
  <div class="scaffold">
    <p class="scaffold__kicker">Not implemented &middot; {{ milestone }}</p>
    <h1 class="scaffold__title">{{ title }}</h1>
    <p class="scaffold__route">
      <code>{{ route }}</code>
    </p>

    <h2>This view must</h2>
    <ul>
      <li v-for="r in requirements" :key="r">{{ r }}</li>
    </ul>

    <h2>Proven by</h2>
    <ul>
      <li v-for="s in specs" :key="s">
        <code>{{ s }}</code>
      </li>
    </ul>

    <p class="scaffold__hint">
      DOM contract: <code>.specs/e2e/SELECTORS.md</code> &middot; Markup:
      <code>.specs/realworld/docs/src/content/docs/specifications/frontend/templates.md</code>
    </p>

    <nav class="scaffold__nav">
      <RouterLink v-for="r in routes" :key="r.to" :to="r.to">{{ r.label }}</RouterLink>
    </nav>
  </div>
</template>

<style scoped>
.scaffold {
  max-width: 44rem;
  margin: 3rem auto;
  padding: 0 1.5rem;
  font-family: ui-sans-serif, system-ui, sans-serif;
  line-height: 1.6;
  color: #373a3c;
}
.scaffold__kicker {
  margin: 0;
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: #b85c5c;
}
.scaffold__title {
  margin: 0.25rem 0 0;
  font-size: 2rem;
  font-weight: 600;
}
.scaffold__route {
  margin: 0.25rem 0 2rem;
  color: #999;
}
.scaffold h2 {
  margin: 1.75rem 0 0.5rem;
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: #999;
}
.scaffold ul {
  margin: 0;
  padding-left: 1.25rem;
}
.scaffold code {
  padding: 0.1em 0.35em;
  border-radius: 3px;
  background: #f2f2f2;
  font-size: 0.875em;
}
.scaffold__hint {
  margin-top: 2rem;
  font-size: 0.875rem;
  color: #999;
}
.scaffold__nav {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid #e5e5e5;
  font-size: 0.875rem;
}
.scaffold__nav a {
  color: #5cb85c;
  text-decoration: none;
}
.scaffold__nav a:hover {
  text-decoration: underline;
}
</style>
