import { createRouter, createWebHistory } from 'vue-router'

// The eight routes the RealWorld frontend spec requires.
// Source: .specs/realworld/docs/.../frontend/routing.md
//
// They are eagerly imported here for simplicity. Stage C of docs/roadmap.md asks
// you to switch to lazy `() => import(...)` chunks and measure what that buys.
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import SettingsView from '@/views/SettingsView.vue'
import EditorView from '@/views/EditorView.vue'
import ArticleView from '@/views/ArticleView.vue'
import ProfileView from '@/views/ProfileView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: HomeView },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/register', name: 'register', component: RegisterView },
    { path: '/settings', name: 'settings', component: SettingsView },
    { path: '/editor', name: 'editor', component: EditorView },
    { path: '/editor/:slug', name: 'editor-edit', component: EditorView, props: true },
    { path: '/article/:slug', name: 'article', component: ArticleView, props: true },
    { path: '/profile/:username', name: 'profile', component: ProfileView, props: true },
    {
      path: '/profile/:username/favorites',
      name: 'profile-favorites',
      component: ProfileView,
      props: true,
    },
    { path: '/:pathMatch(.*)*', name: 'not-found', component: NotFoundView },
  ],
})

// TODO(M7): guard /settings, /editor and /editor/:slug — an anonymous visitor
// must be redirected to /login. url-navigation.spec.ts tests this by typing the
// URL directly, so a guard in the component's onMounted is too late.
//
//   router.beforeEach((to) => { ... })
//
// Then ask yourself: a guard protects the UI, not the data. What stops someone
// calling your API directly? (Nothing, unless the backend checks too.)

export default router
