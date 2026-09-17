<script setup lang="ts">
// TODO(M8): this view is where null-handling bugs surface. A user with
// bio: null must not see "null" in the textarea, and submitting an untouched
// blank password must not set the password to the empty string.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Settings"
    route="/settings"
    milestone="M8"
    :specs="['e2e/settings.spec.ts', 'e2e/null-fields.spec.ts', 'e2e/url-navigation.spec.ts']"
    :requirements="[
      'Require auth: an anonymous visit redirects to /login',
      'Fields: input[name=image], input[name=username], textarea[name=bio], input[name=email], input[name=password]',
      'Prefill from the current user; an empty bio/image must render as an empty field, not the string &quot;null&quot;',
      'PUT /api/user with ONLY the fields that changed',
      'Leave the password blank to keep it unchanged',
      'Offer a logout button that clears the token and redirects',
    ]"
  />
</template>
