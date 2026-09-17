import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// TODO(M7): restore the session before the first render, so a reload does not
// briefly show the logged-out navbar. Something like:
//
//   const auth = useAuthStore()
//   await auth.restore()   // read the token from localStorage, GET /api/user
//
// Think about what should happen when the stored token is expired or forged.

app.mount('#app')
