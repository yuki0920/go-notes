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
            <i class="fas fa-eye article-form__open-preview" />
            <textarea
              id="form-body"
              v-model="article.body"
              class="article-form__input article-form__input--body"
              name="body"
            />
          </div>

          <!-- TODO: プレビュー機能を復活させたい -->
          <!-- <div class="article-form__preview">
            <div class="form__label article-form____label--preview">
              プレビュー
            </div>
            <i class="fas fa-eye-slash article-form__close-preview" />
            <div class="article-form__preview-body">
              <div id="article-body" class="article-form__preview-body-contents" />
            </div>
          </div> -->

          <div class="article-form__footer">
            <button class="article-form__cancel btn--info">
              <nuxt-link to="/articles">
                キャンセル
              </nuxt-link>
            </button>
            <button id="article-form__save" class="article-form__save btn--primary" @click="submit">
              保存
            </button>
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
        await $axios.post('/api/articles', params)
        router.push('/articles')
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
    'errors'
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

.article-form__errors {
  grid-area: errors;
  background: #facac8;
  border-radius: 4px;
}

.article-form__error {
  list-style: none;
  padding: 8px 8px 0;
}

.article-form__error:last-child {
  padding-bottom: 8px;
}

.article-form__title {
  grid-area: title;
}

.article-form__body,
.article-form__preview {
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

.article-form__preview {
  display: none;
}

.article-form__label--body,
.article-form__label--preview {
  grid-area: label;
}

.article-form__open-preview,
.article-form__close-preview {
  grid-area: btn;
  justify-self: center;
  align-self: center;
}

.article-form__input--body {
  resize: none;
  grid-area: text;
}

.article-form__preview-body {
  grid-area: text;
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 0.4px;
  padding: 4px 8px;
  width: 100%;
  word-break: break-all;
  white-space: normal;
  position: relative;
}

.article-form__preview-body-contents {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  overflow: scroll;
  padding: 24px 0;
}

.article-form__preview-body-contents > *:first-child {
  margin-top: 24px;
}

.article-form__footer {
  grid-area: footer;
  display: flex;
  justify-content: space-between;
}

</style>
