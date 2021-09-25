<template>
  <div class="l-col l-row l-v-padd">
    <div class="article-new l-row">
      <div class="article-new__form l-row">
        <!-- TODO: 入力フォームと共通化したい -->
        <form class="article-form l-row" name="article-form">
          <!-- TODO: TODO: エラーを表示したい -->
          <div class="article-form__title">
            <label class="article-form__label" for="form-title">タイトル</label>
            <input id="form-title" v-model=" article.title" class="article-form__input" type="text" name="title">
          </div>

          <div class="article-form__body">
            <label class="article-form__label article-form__label--body" for="form-body">本文</label>
            <textarea
              id="form-body"
              v-model="article.body"
              class="article-form__input article-form__input--body"
              rows="20"
              name="body"
            />
          </div>

          <div class="article-form__footer">
            <a href="javascript:void(0)" class="article-form__cancel btn btn-secondary">
              <nuxt-link to="/">
                キャンセル
              </nuxt-link>
            </a>
            <a id="article-form__save" href="javascript:void(0)" class="article-form__save btn btn-dark" @click="submit">
              保存
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, useRouter, useContext, reactive } from '@nuxtjs/composition-api'

export default defineComponent({
  name: 'Articles',
  setup () {
    const { $axios } = useContext()
    const article = reactive({ title: null, body: null })
    const router = useRouter()

    const submit = async (event: any) => {
      event.preventDefault()
      const params = { title: article.title, body: article.body }
      try {
        const { data } = await $axios.post('/api/articles', params)
        const id = data.Article.id
        router.push(`/articles/${id}`)
      } catch (err) {
        // console.log(err)
      }
    }

    return {
      article,
      submit
    }
  }
})
</script>

<style lang='scss' scoped>
.article-form {
  display: grid;
  grid-template-rows: max-content max-content 1fr max-content;
  grid-template-areas:
    'title'
    'body'
    'footer';
  grid-row-gap: 12px;
}

.article-form__label {
  display: block;
}

.article-form__input {
  display: block;
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 4px 8px;
  width: 100%;
}

.article-form__input:focus {
  outline: none;
}

.article-form__title {
  grid-area: title;
}

.article-form__body {
  grid-area: body;
  grid-template-columns: 1fr max-content;
  grid-template-rows: max-content 1fr;
  grid-template-areas:
    'label btn'
    'text text';
}

.article-form__body {
  display: grid;
}

.article-form__label--body {
  grid-area: label;
}

.article-form__input--body {
  resize: none;
  grid-area: text;
}

.article-form__footer {
  grid-area: footer;
  display: flex;
  justify-content: space-between;
}
</style>
