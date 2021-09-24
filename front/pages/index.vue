<template>
  <div class="l-col">
    <div class="page">
      <h1 class="page__title">
        記事一覧
      </h1>
      <nuxt-link class="page__new btn btn-dark" to="/articles/new">
        新規
      </nuxt-link>
      <div class="page__articles">
        <div class="articles">
          <article v-for="(article, index) in articles" :key="index">
            <div class="articles__item">
              <nuxt-link :to="`/articles/${article.id}`">
                <div class="articles__item-title">
                  {{ article.title }}
                </div>
              </nuxt-link>
              <div class="articles__item-date">
                {{ article.created }}
              </div>
            </div>
          </article>
        </div>
      </div>
      <button v-if="cursor !== 1" class="page__more btn btn-dark" @click="load">
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
.page {
  display: grid;
  align-items: center;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr max-content;
  grid-column-gap: 16px;
  grid-row-gap: 16px;
  grid-template-areas:
    'title new'
    'list list'
    'more more';
  padding: 16px 0;
}

.page__title {
  grid-area: title;
}

.page__new {
  grid-area: new;
}

.page__articles {
  grid-area: list;
}

.page__more {
  grid-area: more;
}

.articles {
  display: grid;
  grid-auto-rows: 1fr;
  grid-row-gap: 16px;
}

.articles__item {
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr max-content;
  grid-column-gap: 16px;
  grid-row-gap: 4px;
  grid-template-areas:
    'title'
    'date';
  background: #fff;
  padding: 12px;
}

.articles__item-title {
  grid-area: title;
  font-size: 20px;
}

.articles__item-date {
  grid-area: date;
  font-size: 12px;
  color: #999;
}
</style>
