<template>
  <div class="container pt-2">
    <div class="d-flex flex-column">
      <h1>
        記事一覧
      </h1>
      <template v-if="articles.length > 0">
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
        <div class="overflow-auto">
          <b-pagination-nav :link-gen="linkGen" :number-of-pages="totalPage" use-router />
        </div>
      </template>
      <div v-else class="d-flex flex-column justify-content-center align-items-center mt-5">
        <div class="spinner-border ml-2" style="width: 10rem; height: 10rem;" role="status" />
        <h3 class="mt-3">
          Loading
        </h3>
        <p>(Please wait 30 secs for heroku's server to start)</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, useContext, useRouter } from '@nuxtjs/composition-api'
import { Article } from '~/types/article'
export default defineComponent({
  name: 'TopPage',
  setup () {
    const { $axios } = useContext()
    const articles = ref<Article[]>([])
    const totalPage = ref(0)

    type Data = {
      articles: Article[],
      totalPage: number
    }

    const load = async (pageNum = 1) => {
      const { data }: { data: Data } = await $axios.get(`/api/articles?page=${pageNum}`)
      articles.value = data.articles
      totalPage.value = data.totalPage
    }

    const linkGen = (pageNum: number) => {
      return pageNum === 1 ? '' : `?page=${pageNum}`
    }

    onMounted(async () => {
      await load(1)
    })

    const router = useRouter()
    router.beforeEach((to, from, next) => {
      if (to.name === from.name && (typeof to.query.page === 'string' || to.query.page === undefined)) {
        load(parseInt(to.query.page, 10) || 1)
        scrollTo(0, 0)
      }
      next()
    })

    return {
      articles,
      totalPage,
      load,
      linkGen
    }
  },
  head: {}
})
</script>
<style lang="scss" scoped>
</style>
