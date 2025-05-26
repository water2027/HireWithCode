import request from '../request'

export interface SendCodeRequest {
  email: string
}

export function sendCode(req: SendCodeRequest) {
  return request.post<string>('/work/send_code', {
    body: req,
  })
}
