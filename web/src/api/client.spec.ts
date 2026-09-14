import { describe, it, expect } from 'vitest'
import { ApiError } from './client'

// A worked example of the shape your tests should take. Vitest is already wired
// up (`make test-web`), so there is no excuse to put off writing the next one.
describe('ApiError', () => {
  it('flattens the spec error envelope into renderable lines', () => {
    const err = new ApiError(422, {
      email: ['has already been taken'],
      password: ['is too short', "can't be blank"],
    })

    expect(err.messages).toEqual([
      'email has already been taken',
      'password is too short',
      "password can't be blank",
    ])
  })

  it('tolerates a response with no error payload', () => {
    expect(new ApiError(500).messages).toEqual([])
  })
})

// TODO(M7): once request() exists, test it against a stubbed fetch:
//   - attaches `Authorization: Token ...` only when auth is requested
//   - throws ApiError with the parsed body on 422
//   - does not explode when the body is HTML rather than JSON
