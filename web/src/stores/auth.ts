import { defineStore } from 'pinia'

/**
 * Authentication state: the current user and their JWT.
 *
 * EXERCISE (M7)
 * -------------
 * The frontend spec says to keep the token in localStorage. Build:
 *
 *   state:    user (User | null), token (string | null), status
 *   getters:  isAuthenticated, username
 *   actions:  login, register, restore, updateUser, logout
 *
 * Details that will bite you if you skip them:
 *   - restore() runs before the first render (see main.ts). A stored token may
 *     be expired or tampered with, so verify it with GET /api/user rather than
 *     trusting what is in localStorage.
 *   - logout() must clear both the store and localStorage, then redirect.
 *   - localStorage throws in some privacy modes. Wrap access in try/catch.
 *
 * And a question to answer properly rather than hand-wave: localStorage is
 * readable by any script on the page, so a single XSS steals the token.
 * httpOnly cookies fix that and bring CSRF instead. The spec chose localStorage;
 * Stage H of docs/roadmap.md asks you to implement the alternative and write up
 * the trade-off.
 */
export const useAuthStore = defineStore('auth', {
  state: () => ({
    // TODO(M7): type this with the User interface from @/types/api
    user: null as unknown | null,
    token: null as string | null,
  }),

  getters: {
    isAuthenticated: (state) => state.token !== null && state.user !== null,
  },

  actions: {
    async restore(): Promise<void> {
      // TODO(M7): read the token, verify it with GET /api/user, populate state.
    },

    async login(_email: string, _password: string): Promise<void> {
      throw new Error('auth store: login() is not implemented yet — see LEARNING.md M7')
    },

    async register(_username: string, _email: string, _password: string): Promise<void> {
      throw new Error('auth store: register() is not implemented yet — see LEARNING.md M7')
    },

    logout(): void {
      // TODO(M7): clear state and localStorage, then redirect to '/'.
    },
  },
})
