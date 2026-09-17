/**
 * Thin wrapper around fetch for the Conduit API.
 *
 * EXERCISE (M7)
 * -------------
 * Everything below is a stub. Build it so that no component ever calls fetch
 * directly — one place to attach the token, one place to parse errors, one place
 * to change when the API moves.
 *
 * What it needs to do:
 *   - Prefix every path with BASE_URL.
 *   - Attach `Authorization: Token <jwt>` when logged in. Note "Token", not
 *     "Bearer" — the RealWorld spec is non-standard here.
 *   - Parse the response before deciding success: a 422 has a useful body.
 *   - Throw a typed ApiError carrying status and the {errors:{field:[msg]}}
 *     payload, so forms can render it and error-handling.spec.ts passes.
 *   - Survive a response that is not JSON at all (a 502 from a proxy is HTML) —
 *     user-fetch-errors.spec.ts checks you do not explode on that.
 *
 * Design question worth sitting with: should this module read the token from the
 * Pinia store (store -> client -> store, a cycle) or should the store inject it
 * (client knows nothing about Pinia)? Pick one and write down why.
 */

/** Vite proxies this to the Go API in dev; see vite.config.ts. */
export const BASE_URL = '/api'

/** Error thrown for any non-2xx response. */
export class ApiError extends Error {
  constructor(
    readonly status: number,
    /** The spec's 422 shape: { "errors": { "email": ["has already been taken"] } } */
    readonly errors: Record<string, string[]> = {},
    message = `API request failed with ${status}`,
  ) {
    super(message)
    this.name = 'ApiError'
  }

  /** Flatten to the "field message" lines that ul.error-messages renders. */
  get messages(): string[] {
    return Object.entries(this.errors).flatMap(([field, msgs]) => msgs.map((m) => `${field} ${m}`))
  }
}

export interface RequestOptions {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
  body?: unknown
  /** Send the auth header. Some endpoints change behaviour when a token is present. */
  auth?: boolean
  query?: Record<string, string | number | undefined>
  signal?: AbortSignal
}

export async function request<T>(_path: string, _options: RequestOptions = {}): Promise<T> {
  // TODO(M7): implement. Suggested order of work:
  //   1. Build the URL (BASE_URL + path + serialised, undefined-skipping query).
  //   2. Build headers (Content-Type when there is a body; auth when asked).
  //   3. await fetch, then read the body ONCE (res.json() twice throws).
  //   4. if (!res.ok) throw new ApiError(res.status, body?.errors ?? {}).
  //   5. Return the parsed body, typed as T.
  throw new Error('api/client.ts: request() is not implemented yet — see LEARNING.md M7')
}

// Convenience wrappers to build on top of request() once it works.
// TODO(M7): add typed endpoint functions, e.g.
//
//   export const login = (email: string, password: string) =>
//     request<{ user: User }>('/users/login', { method: 'POST', body: { user: { email, password } } })
//
// Keep them in src/api/ grouped by resource (auth.ts, articles.ts, profiles.ts,
// comments.ts, tags.ts) rather than one 400-line file.
