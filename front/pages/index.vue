<template>
  <div class="l-col">
    <div class="page">
      <h1 class="page__title">
        記事一覧
      </h1>
      <nuxt-link class="page__new btn--primary" to="/articles/new">
        新規
      </nuxt-link>
      <div class="page__articles">
        <div class="articles">
          <article v-for="(article, index) in articles" :key="index">
            <nuxt-link class="articles__item" :to="`/articles/${article.id}`">
              <div class="articles__item-title">
                {{ article.title }}
              </div>
              <div class="articles__item-date">
                {{ article.created }}
              </div>
              <button class="articles__item-delete">
                <i class="fas fa-trash-alt" />
              </button>
            </nuxt-link>
          </article>
        </div>
      </div>
      <div class="articles__item-tmpl" style="display: none">
        <article>
          <a class="articles__item" href="">
            <div class="articles__item-title" />
            <div class="articles__item-date" />
            <button class="articles__item-delete"><i class="fas fa-trash-alt" /></button>
          </a>
        </article>
      </div>
      <div v-if="cursor !== 1" class="page__more" @click="load">
        もっとみる
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, useContext } from '@nuxtjs/composition-api'

export default defineComponent({
  name: 'TopPage',
  setup () {
    const articles = ref([])
    const { $axios } = useContext()
    const cursor = ref(0)
    const load = async () => {
      const { data } = await $axios.get('/api/articles', { params: { cursor: cursor.value } })
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
  background: #2e82e8;
  color: #fff;
  text-align: center;
  border-radius: 10px;
  padding: 8px 0;
  font-size: 20px;
  cursor: pointer;
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
    'title delete'
    'date delete';
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

.articles__item-delete {
  grid-area: delete;
  align-self: center;
  background: #fe4d08;
  color: #fff;
  height: 28px;
  width: 28px;
  font-size: 16px;
  border-radius: 4px;
}
</style>
