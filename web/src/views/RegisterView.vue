<script setup lang="ts">
// TODO(M7): registration and login differ by one field and one URL. Extract
// the shared form shell only once you have written both — premature extraction
// is how component hierarchies get ugly.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Sign up"
    route="/register"
    milestone="M7"
    :specs="['e2e/auth.spec.ts', 'e2e/error-handling.spec.ts', 'e2e/health.spec.ts']"
    :requirements="[
      'Render an h1 containing &quot;Sign up&quot; and a link to /login',
      'Fields: input[name=username], input[name=email], input[name=password]',
      'POST /api/users, store the token, redirect to /',
      'Surface duplicate-email and duplicate-username errors in ul.error-messages',
    ]"
  />
</template>
