<script setup lang="ts">
// TODO(M7): the error list is not decoration — error-handling.spec.ts asserts
// on it. Write the {errors:{field:[msg]}} -> flat string[] mapping once, in a
// helper, and reuse it on every form.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Sign in"
    route="/login"
    milestone="M7"
    :specs="['e2e/auth.spec.ts', 'e2e/error-handling.spec.ts', 'e2e/health.spec.ts']"
    :requirements="[
      'Render an h1 containing &quot;Sign in&quot; and a link to /register',
      'Fields: input[name=email], input[name=password]',
      'POST /api/users/login, store the token, redirect to /',
      'Render server errors in ul.error-messages (the 422 {errors:{field:[msg]}} shape)',
      'Disable the submit button while the request is in flight',
    ]"
  />
</template>
