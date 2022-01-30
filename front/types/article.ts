import { Category } from './category'

export type Article = {
  id: number,
  title: string,
  body: string,
  categories: Category[],
  created: string,
  updated: string
}
