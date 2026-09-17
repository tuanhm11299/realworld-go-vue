import { defineConfig } from '@playwright/test'
import { baseConfig } from './.specs/e2e/playwright.base'

/**
 * Runs the UPSTREAM RealWorld e2e suite against this implementation.
 *
 * The specs are not ours and are not in git — `make specs-sync` fetches them
 * into .specs/. They are the frontend half of your to-do list: every failure
 * names something the app does not do yet.
 *
 * TEST_MODE=fullstack
 * -------------------
 * We own both halves of the stack, so there is no hosted demo API to lean on and
 * no seeded `johndoe` to borrow. In fullstack mode the suite drives everything
 * through the UI and creates its own users. See .specs/e2e/helpers/config.ts.
 *
 * Once you have a dev seeder (M8 stretch), switching to
 *   TEST_MODE=spa API_BASE=http://localhost:8080/api
 * unlocks the extra tests that intercept browser API traffic.
 *
 * Prerequisite: Postgres must be up (`make up && make migrate-up`). Playwright
 * starts the API and the Vite dev server itself.
 *
 * @playwright/test is pinned to an exact version in package.json on purpose:
 * Playwright ships version-locked browser binaries, so a floating range would
 * silently change which Chromium CI downloads.
 */
export default defineConfig({
  ...baseConfig,
  testDir: './.specs/e2e',

  use: {
    ...baseConfig.use,
    baseURL: process.env.BASE_URL ?? 'http://localhost:5173',
  },

  webServer: [
    {
      command: 'go run ./cmd/api',
      cwd: 'api',
      url: 'http://localhost:8080/healthz',
      reuseExistingServer: !process.env.CI,
      timeout: 120_000,
      stdout: 'pipe',
      stderr: 'pipe',
    },
    {
      command: 'npm run dev',
      cwd: 'web',
      url: 'http://localhost:5173',
      reuseExistingServer: !process.env.CI,
      timeout: 120_000,
    },
  ],
})
