type HttpMethod = 'GET'

interface FetcherArgs {
  body?: Record<string, unknown>
  endpoint: string
  method: HttpMethod
}

async function _fetcher<Response extends object>({
  body,
  endpoint,
  method
}: FetcherArgs): Promise<Response> {
  const url = `http://localhost:8080${endpoint}`
  const headers = new Headers()
  const fetchArgs: RequestInit = {}

  if (body) {
    headers.append('Content-Type', 'application/json')
    fetchArgs.body = JSON.stringify(body)
  }

  const res = await fetch(url, {
    method
  })

  if (!res.ok) {
    throw new Error('request error')
  }

  return (await res.json()) as Response
}

const get = <Response extends object>(endpoint: string) =>
  _fetcher<Response>({ endpoint, method: 'GET' })

const fetcher = { get }

export default fetcher
