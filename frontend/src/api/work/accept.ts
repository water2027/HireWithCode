import request from '../request'

export interface AcceptChallengeRequest {
  email: string
  github_id: string
  code: string
}

export function acceptChallenge(req: AcceptChallengeRequest) {
  return request.post<string>('/work/accept', {
    body: req,
  })
}
