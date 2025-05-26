import request from '../request'

export interface SubmitRequest {
  email: string
  github_url: string
  online_url: string
  code: string
}

export function submitWork(req: SubmitRequest) {
  return request.post<string>('/work/submit', {
    body: req,
  })
}
