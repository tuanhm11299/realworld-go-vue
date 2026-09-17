import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/ — imported from vitest/config so the `test` block
// below is typed. Vite reads it exactly the same way.
export default defineConfig({
  plugins: [vue(), vueDevTools()],
  resolve: {
    alias: { '@': fileURLToPath(new URL('./src', import.meta.url)) },
  },
  server: {
    port: 5173,
    strictPort: true,
    proxy: {
      // Same-origin in dev, so you can ignore CORS while learning. Deployed,
      // the SPA and the API usually sit on different origins and you WILL need
      // the CORS middleware from internal/httpx — see docs/milestones/M7.md.
      '/api': {
        target: process.env.VITE_API_PROXY_TARGET ?? 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  test: {
    environment: 'jsdom',
    include: ['src/**/*.{test,spec}.ts'],
  },
})
