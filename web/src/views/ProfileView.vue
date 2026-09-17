<script setup lang="ts">
// TODO(M8): both tabs are the same list with a different query parameter.
// One component, one composable, two call sites — and watch the route param so
// navigating from /profile/a to /profile/b actually refetches.
import TodoPanel from '@/components/TodoPanel.vue'
</script>

<template>
  <TodoPanel
    title="Profile"
    route="/profile/:username and /profile/:username/favorites"
    milestone="M8"
    :specs="['e2e/social.spec.ts', 'e2e/url-navigation.spec.ts', 'e2e/null-fields.spec.ts']"
    :requirements="[
      'Render .profile-page with .user-info: avatar, username, bio',
      'Show Follow/Unfollow for other users, Edit Profile Settings for yourself',
      'Two tabs: My Articles (author=:username) and Favorited Articles (favorited=:username)',
      'The favorites tab lives at its own URL and must survive a refresh',
      'Fall back to the default avatar when image is null',
      'Unknown username: show a not-found state, do not crash',
    ]"
  />
</template>
