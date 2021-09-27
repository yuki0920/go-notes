<template>
  <div class="l-col pt-2">
    <div class="d-flex flex-column">
      <h1>
        記事一覧
      </h1>
      <template v-if="articles.length > 1">
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
        <button v-if="!finished" class="btn btn-dark" @click="load">
          もっとみる
        </button>
      </template>
      <div v-else class="d-flex justify-content-center align-items-center">
        <div class="spinner-border mr-2" style="width: 2rem; height: 2rem;" role="status" />
        <span>Loading...(Please wait 30 secs for heroku's server to start)</span>
        <div class="spinner-border ml-2" style="width: 2rem; height: 2rem;" role="status" />
      </div>
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
    const finished = ref(false)

    type Data = {
      articles: Article[],
      cursor: number
    }
    const load = async () => {
      const { data }: { data: Data } = await $axios.get('/api/articles', { params: { cursor: cursor.value } })
      if (cursor.value === data.cursor) {
        finished.value = true
        return
      }

      cursor.value = data.cursor
      articles.value.push(...data.articles)
    }

    onMounted(async () => {
      await load()
    })

    return {
      articles,
      cursor,
      load,
      finished
    }
  },
  head: {}
})
</script>
<style lang="scss" scoped>
</style>
