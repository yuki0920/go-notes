<template>
  <div v-if="article" class="container">
    <div>
      <div class="px-3 pt-1 pb-3">
        <h1>
          {{ article.title }}
        </h1>
        <div>
          <span v-for="category in article.categories" :key="category.title" class="badge badge-dark inline-block mb-1 mr-1">
            {{ category.title }}
          </span>
        </div>
        <div class="d-flex">
          <small class="text-muted">
            Published: {{ article.created }} / Updated: {{ article.updated }}
          </small>
        </div>
        <nuxt-link v-if="isAuthenticated" :to="`/articles/${articleId}/edit`">
          Edit
        </nuxt-link>
      </div>
      <div class="px-3">
        <vue-remarkable>
          {{ article.body }}
        </vue-remarkable>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, useRoute, useContext, ref } from '@nuxtjs/composition-api'
// @ts-ignore # NOTE: 型定義ファイルがないため
import vueRemarkable from 'vue-remarkable'
import { Article } from '~/types/article'

export default defineComponent({
  name: 'ArticleIndex',
  components: {
    vueRemarkable
  },
  setup () {
    const route = useRoute()
    const { $axios } = useContext()

    const articleId = route.value.params.id
    const isAuthenticated = ref(false)
    const article = ref<Article | null>(null)

    onMounted(async () => {
      const { data: authData } = await $axios.get('/api/auth')
      isAuthenticated.value = authData.IsAuthenticated

      const { data }: { data: Article } = await $axios.get(`/api/articles/${articleId}`)
      if (data) { article.value = data }
    })

    return {
      articleId,
      isAuthenticated,
      article
    }
  }
})
</script>

<style lang='scss' scoped>
</style>
