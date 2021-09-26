<template>
  <div v-if="article" class="l-col l-row l-v-padd">
    <div>
      <div class="px-3 pt-1 pb-3">
        <h1>
          {{ article.title }}
        </h1>
        <div class="d-flex">
          <div class="text-muted">
            更新: {{ article.updated }}
          </div>
          <div class="text-muted">
            公開: {{ article.created }}
          </div>
        </div>
        <nuxt-link v-if="isAuthenticated" class="text-muted" :to="`/articles/${articleId}/edit`">
          編集
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

export default defineComponent({
  name: 'Articles',
  components: {
    vueRemarkable
  },
  setup () {
    const route = useRoute()
    const { $axios } = useContext()

    const articleId = route.value.params.id
    const isAuthenticated = ref(false)
    const article = ref(null)

    onMounted(async () => {
      const { data: authData } = await $axios.get('/api/auth')
      isAuthenticated.value = authData.IsAuthenticated

      const { data: articleData } = await $axios.get(`/api/articles/${articleId}`)
      if (articleData) { article.value = articleData }
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
