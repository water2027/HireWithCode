// const baseUrl = import.meta.env.VITE_API_URL

interface Response<T> {
  code: number
  message: string
  data: T
}

class Request {
  private baseUrl: string
  constructor(url:string) {
    this.baseUrl = url || 'http://localhost:8080'
  }
  private async baseRequest<T>(url: string, init: RequestInit): Promise<T> {
    const resp = await fetch(`${this.baseUrl}${url}`, init)

    const jsonData = await resp.json() as Response<T>

    const { code, message, data } = jsonData

    // 约定小于100就失败吧, 反正场景不复杂
    if (code < 100) {
      console.error(message)
      return data
    }

    return data
  }

  public async get<T>(url: string, query: Record<string, string>) {
    const params = new URLSearchParams(query)
    const queryString = params.toString()
    const fullUrl = queryString ? `${url}?${queryString}` : url
    return this.baseRequest<T>(fullUrl, {
      method: 'GET',
    })
  }

  public async post<T>(url: string, query: Record<string, string>, body?: object) {
    const params = new URLSearchParams(query)
    const queryString = params.toString()
    const fullUrl = queryString ? `${url}?${queryString}` : url

    let finalBody
    try {
        finalBody = body ? JSON.stringify(body) : undefined
    } catch (error) {
        finalBody = undefined
    }
    return this.baseRequest<T>(fullUrl, {
        method: 'POST',
        headers: {
            'Content-Type':'application/json'
        },
        body: finalBody
    })
  }
}

export default new Request(import.meta.env.VITE_API_URL)