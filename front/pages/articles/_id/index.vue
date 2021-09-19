<template>
  <div v-if="article" class="l-col l-row l-v-padd">
    <div class="article">
      <div class="article__header">
        <h1 class="article__title">
          {{ article.title }}
        </h1>
        <nuxt-link class="article__edit btn--info" :to="`/articles/${article.id}/edit`">
          編集
        </nuxt-link>
        <div class="article__date">
          <div class="article__updated">
            更新: {{ article.updated }}
          </div>
          <div class="article__published">
            公開: {{ article.created }}
          </div>
        </div>
      </div>
      <div class="article-body">
        <vue-remarkable>
          {{ article.body }}
        </vue-remarkable>
      </div>
    </div>
    <div class="article-footer" />
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
    const id = route.value.params.id
    const { $axios } = useContext()
    const article = ref(null)

    onMounted(async () => {
      const { data } = await $axios.get(`/api/articles/${id}`)
      if (data) { article.value = data }
    })

    return {
      article
    }
  }
})
</script>

<style lang='scss' scoped>
.article {
  background: #fff;
  overflow: hidden;
  width: 100%;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: repeat(2, max-content);
  grid-template-areas:
    'header'
    'body';
  grid-row-gap: 36px;
}

.article__header {
  grid-area: header;
  padding: 16px 4% 16px;
  display: grid;
  align-items: center;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr max-content;
  grid-column-gap: 12px;
  grid-row-gap: 12px;
  grid-template-areas:
    'title edit'
    'date date';
}

.article__title {
  grid-area: title;
  font-size: 28px;
}

.article__edit {
  grid-area: edit;
}

.article__date {
  grid-area: date;
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: max-content max-content;
  grid-column-gap: 12px;
  color: #777;
  font-size: 12px;
}

.article__body {
  grid-area: body;
}
</style>
