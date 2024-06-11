import type { Flag } from '@/entities/Flag'
import fetcher from './fetcher'

const flags = {
  getAll: () => fetcher.get<Flag[]>('/flags')
}

export default flags
