<template>
  <div class="l-col">
    <div class="d-flex flex-column">
      <h1>
        記事一覧
      </h1>
      <article v-for="(article, index) in articles" :key="index">
        <div class="p-2">
          <nuxt-link :to="`/articles/${article.id}`">
            <p class="h3">
              {{ article.title }}
            </p>
          </nuxt-link>
          <div>
            <small class="text-muted">
              {{ article.created }}
            </small>
          </div>
        </div>
      </article>
      <button v-if="cursor !== 1" class="btn btn-dark" @click="load">
        もっとみる
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, useContext } from '@nuxtjs/composition-api'
import { Article } from '~/types/article'
export default defineComponent({
  name: 'TopPage',
  setup () {
    const { $axios } = useContext()
    const articles = ref<Article[]>([])
    const cursor = ref(0)
    type Data = {
      articles: Article[],
      cursor: number
    }
    const load = async () => {
      const { data }: { data: Data } = await $axios.get('/api/articles', { params: { cursor: cursor.value } })
      articles.value.push(...data.articles)
      cursor.value = data.cursor
      // console.log('cursor.value', cursor.value)
    }

    onMounted(async () => {
      await load()
    })

    return {
      articles,
      cursor,
      load
    }
  },
  head: {}
})
</script>
<style lang="scss" scoped>
</style>
